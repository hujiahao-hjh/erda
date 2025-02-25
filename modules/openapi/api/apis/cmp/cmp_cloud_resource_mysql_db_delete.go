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

package cmp

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
	"github.com/erda-project/erda/modules/openapi/api/spec"
)

var CMP_CLOUD_RESOURCE_MYSQL_DB_DELETE = apis.ApiSpec{
	Path:         "/api/cloud-mysql/actions/delete-db",
	BackendPath:  "/api/cloud-mysql/actions/delete-db",
	Host:         "cmp.marathon.l4lb.thisdcos.directory:9027",
	Scheme:       "http",
	Method:       "DELETE",
	CheckLogin:   true,
	RequestType:  apistructs.DeleteCloudResourceMysqlDBRequest{},
	ResponseType: apistructs.CloudAddonResourceDeleteRespnse{},
	Doc:          "删除 mysql database",
	Audit: func(ctx *spec.AuditContext) error {
		var request apistructs.DeleteCloudResourceMysqlDBRequest
		if err := ctx.BindRequestData(&request); err != nil {
			return err
		}

		project, err := ctx.GetProject(request.ProjectID)
		if err != nil {
			return err
		}

		return ctx.CreateAudit(&apistructs.Audit{
			ScopeType:    apistructs.ProjectScope,
			ScopeID:      project.ID,
			ProjectID:    project.ID,
			TemplateName: apistructs.DeleteMysqlDbTemplate,
			Context: map[string]interface{}{
				"projectName": project.Name,
				"addonID":     request.AddonID,
			},
		})
	},
}
