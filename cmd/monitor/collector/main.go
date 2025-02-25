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

package main

import (
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda/conf"
	"github.com/erda-project/erda/pkg/common"

	// modules
	_ "github.com/erda-project/erda-infra/providers/health"
	_ "github.com/erda-project/erda-infra/providers/kafka"
	_ "github.com/erda-project/erda-infra/providers/pprof"

	// providers
	_ "github.com/erda-project/erda-infra/providers"
	_ "github.com/erda-project/erda/modules/core/monitor/collector"
	_ "github.com/erda-project/erda/modules/oap/collector/authentication"
	_ "github.com/erda-project/erda/modules/oap/collector/receivers/common"
	_ "github.com/erda-project/erda/modules/oap/collector/receivers/jaeger"

	// grpc
	_ "github.com/erda-project/erda-infra/providers/grpcclient"
	_ "github.com/erda-project/erda-proto-go/core/services/authentication/credentials/accesskey/client"
)

//go:generate sh -c "cd ${PROJ_PATH} && go generate -v -x github.com/erda-project/erda/modules/monitor/core/collector"
func main() {
	common.Run(&servicehub.RunOptions{
		ConfigFile: conf.MonitorCollectorConfigFilePath,
		Content:    conf.MonitorCollectorDefaultConfig,
	})
}
