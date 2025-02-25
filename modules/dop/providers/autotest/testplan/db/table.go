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

package db

import (
	"time"

	"github.com/erda-project/erda/pkg/database/dbengine"
)

// TestPlanV2 测试计划V2
type TestPlanV2 struct {
	dbengine.BaseModel
	Name        string
	Desc        string
	CreatorID   string
	UpdaterID   string
	ProjectID   uint64
	SpaceID     uint64
	PipelineID  uint64
	PassRate    float64
	ExecuteTime *time.Time
}

// TableName table name
func (TestPlanV2) TableName() string {
	return "dice_autotest_plan"
}
