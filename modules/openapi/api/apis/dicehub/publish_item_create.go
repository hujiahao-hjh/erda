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

package dicehub

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
	"github.com/erda-project/erda/modules/openapi/api/spec"
)

var PUBLISH_ITEM_CREATE = apis.ApiSpec{
	Path:         "/api/publish-items",
	BackendPath:  "/api/publish-items",
	Host:         "dicehub.marathon.l4lb.thisdcos.directory:10000",
	Scheme:       "http",
	Method:       "POST",
	CheckLogin:   true,
	CheckToken:   true,
	RequestType:  apistructs.CreatePublishItemRequest{},
	ResponseType: apistructs.CreatePublishItemResponse{},
	Doc:          "summary: 创建发布",
	Audit: func(ctx *spec.AuditContext) error {

		var resp apistructs.CreatePublishItemResponse
		err := ctx.BindResponseData(&resp)
		if err != nil {
			return err
		}

		if resp.Success {
			return ctx.CreateAudit(&apistructs.Audit{
				ScopeType:    apistructs.OrgScope,
				ScopeID:      uint64(ctx.OrgID),
				TemplateName: apistructs.CreatePublishItemsTemplate,
				Context:      map[string]interface{}{"type": resp.Data.Type, "publishItemContent": resp.Data.Name},
			})
		} else {
			return nil
		}
	},
}
