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

package logservice

import (
	"github.com/jinzhu/gorm"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda/modules/msp/instance/db"
	"github.com/erda-project/erda/modules/msp/resource/deploy/handlers"
)

type config struct {
}

// +provider
type provider struct {
	*handlers.DefaultDeployHandler
	Cfg                  *config
	Log                  logs.Logger
	DB                   *gorm.DB `autowired:"mysql-client"`
	LogServiceInstanceDB *db.LogServiceInstanceDB
	LogDeploymentDb      *db.LogDeploymentDB
	LogInstanceDb        *db.LogInstanceDB
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.DefaultDeployHandler = handlers.NewDefaultHandler(p.DB, p.Log)
	p.LogServiceInstanceDB = &db.LogServiceInstanceDB{DB: p.DB}
	p.LogDeploymentDb = &db.LogDeploymentDB{DB: p.DB}
	p.LogInstanceDb = &db.LogInstanceDB{DB: p.DB}
	return nil
}

func init() {
	servicehub.Register("erda.msp.resource.deploy.handlers.logservice", &servicehub.Spec{
		Services: []string{
			"erda.msp.resource.deploy.handlers.logservice",
		},
		Description: "erda.msp.resource.deploy.handlers.logservice",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
