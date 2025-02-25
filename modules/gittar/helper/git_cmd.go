// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helper

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda/modules/gittar/models"
	"github.com/erda-project/erda/modules/gittar/pkg/gitmodule"
	"github.com/erda-project/erda/modules/gittar/webcontext"
)

func recycleChildProcess() {
	//是init进程
	if os.Getpid() == 1 {
		sc := make(chan os.Signal)
		signal.Notify(sc, syscall.SIGCHLD)
		//回收被kill的进程的子进程
		go func() {
			var status syscall.WaitStatus
			for range sc {
				wpid, err := syscall.Wait4(-1, &status, syscall.WNOHANG, nil)
				if err == nil {
					logrus.Infof("wait child process success %s", strconv.Itoa(wpid))
				}
			}
		}()
	}
}

func gitCommand(version string, args ...string) *exec.Cmd {
	command := exec.Command("git", args...)
	if len(version) > 0 {
		command.Env = append(os.Environ(), fmt.Sprintf("GIT_PROTOCOL=%s", version))
	}
	return command
}

// Run Command extend
func runCommand2(w io.Writer, cmd *exec.Cmd, readers ...io.Reader) {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		logrus.Printf("[ERROR] get command stdin error %v", err)
		return
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Printf("[ERROR] get command stdout error %v", err)
		return
	}
	if err := cmd.Start(); err != nil {
		logrus.Printf("[ERROR] command start error %v", err)
		return
	}
	logrus.Infof("command %v processId:%s", cmd.Args, strconv.Itoa(cmd.Process.Pid))

	for _, v := range readers {
		_, err := io.Copy(stdin, v)
		if err != nil {
			logrus.Infof("[ERROR] stdin read error %v", err)
			logrus.Infof("kill process %s", strconv.Itoa(cmd.Process.Pid))
			cmd.Process.Kill()
			cmd.Wait()
			return
		}
	}

	stdin.Close()
	_, err = io.Copy(w, stdout)
	if err != nil {
		logrus.Infof("[ERROR] stdout write error %v", err)
		logrus.Infof("kill process %s", strconv.Itoa(cmd.Process.Pid))
		cmd.Process.Kill()
		cmd.Wait()
		return
	}
	err = cmd.Wait()
	if err != nil {
		logrus.Printf("[ERROR] command exec error %v", err)
	}

}

// RunAdvertisement command
func RunAdvertisement(service string, c *webcontext.Context) {
	version := c.MustGet("gitProtocol").(string)
	headerNoCache(c)
	c.Header("Content-type", fmt.Sprintf("application/x-git-%s-advertisement", service))

	c.Status(http.StatusOK)
	if len(version) == 0 {
		c.GetWriter().Write(packetWrite("# service=git-" + service + "\n"))
		c.GetWriter().Write(packetFlush())
	}

	runCommand2(c.GetWriter(), gitCommand(
		version,
		service,
		"--stateless-rpc",
		"--advertise-refs",
		c.Repository.DiskPath(),
	), c.GetRequestBody())
}

// RunProcess command
func RunProcess(service string, c *webcontext.Context) {
	version := c.MustGet("gitProtocol").(string)

	c.Header("Content-Type", fmt.Sprintf("application/x-git-%s-result", service))
	headerNoCache(c)

	var (
		reqBody = c.GetRequestBody()
		err     error
	)

	// Handle GZIP
	if c.GetHeader("Content-Encoding") == "gzip" {
		reqBody, err = gzip.NewReader(reqBody)
		if err != nil {
			logrus.Errorf("HTTP.Get: fail to create gzip reader: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	defer reqBody.Close()

	// Grab pushEvent
	var pushEvents []*models.PayloadPushEvent
	if service == "receive-pack" {
		header, err := ReadGitSendPackHeader(reqBody)
		if err != nil {
			logrus.Errorf("receive-pack error %v", err)
			c.AbortWithStatus(500)
			return
		}
		logrus.Infof("push header:" + string(header))
		re := regexp.MustCompile(
			`(?mi)(?P<before>[0-9a-fA-F]{40}) (?P<after>[0-9a-fA-F]{40}) (?P<ref>refs\/(heads|tags)\/.*?)\0`,
		)

		for _, matches := range re.FindAllSubmatch(header, -1) {
			pushEvent := &models.PayloadPushEvent{
				Before:            string(matches[1]),
				After:             string(matches[2]),
				Ref:               string(bytes.Trim(matches[3], "\x00")),
				IsTag:             string(matches[4]) == "tags",
				Pusher:            c.MustGet("user").(*models.User),
				TotalCommitsCount: 0,
			}
			pushEvent.IsDelete = pushEvent.After == gitmodule.INIT_COMMIT_ID
			pushEvents = append(pushEvents, pushEvent)
		}

		repository := c.MustGet("repository").(*gitmodule.Repository)
		if preReceiveHook(pushEvents, c) {

			runCommand2(c.GetWriter(), gitCommand(
				version,
				service,
				"--stateless-rpc",
				repository.DiskPath(),
			), bytes.NewReader(header), reqBody)
			go PostReceiveHook(pushEvents, c)
		}
	} else {
		runCommand2(c.GetWriter(), gitCommand(
			version,
			service,
			"--stateless-rpc",
			c.MustGet("repository").(*gitmodule.Repository).DiskPath(),
		), reqBody)
	}
}

func RunArchive(c *webcontext.Context, ref string, format string) {
	c.EchoContext.Response().Header().Add("Content-Disposition", "attachment; filename="+
		strings.Replace(ref, "/", "-", -1)+"."+format)

	fullPath, _ := filepath.Abs(c.MustGet("repository").(*gitmodule.Repository).DiskPath())
	runCommand2(c.GetWriter(), gitCommand(
		"",
		"archive",
		"--format="+format,
		ref,
		"--remote=file:///"+fullPath))
}

// OutPutArchive 创建打包文件
func OutPutArchive(c *webcontext.Context, ref string, format string) string {

	fullPath, _ := filepath.Abs(c.MustGet("repository").(*gitmodule.Repository).DiskPath())
	filename := fullPath + "/" + strings.Replace(ref, "/", "-", -1) + "." + format

	runCommand2(c.GetWriter(), gitCommand(
		"",
		"archive",
		"--format="+format,
		ref,
		"--remote=file:///"+fullPath,
		"--output="+filename))
	return filename
}

// OutPutArchiveDelete 删除打包文件  （文件打包上传完必须删除）
func OutPutArchiveDelete(c *webcontext.Context, path string) {

	runCommand2(c.GetWriter(), exec.Command(
		"rm",
		"-r",
		"-f",
		path))
}
