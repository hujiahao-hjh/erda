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

package dbclient

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/spec"
)

// UpdatePipelineTaskSnippetDetail 更新 snippet task 的 snippet 信息
func (client *Client) UpdatePipelineTaskSnippetDetail(id uint64, snippetDetail apistructs.PipelineTaskSnippetDetail, ops ...SessionOption) error {
	session := client.NewSession(ops...)
	defer session.Close()

	_, err := session.ID(id).Cols("snippet_pipeline_detail").Update(&spec.PipelineTask{SnippetPipelineDetail: &snippetDetail})
	return err
}
