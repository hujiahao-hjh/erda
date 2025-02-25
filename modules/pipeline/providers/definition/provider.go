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

package definition

import (
	logs "github.com/erda-project/erda-infra/base/logs"
	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	"github.com/erda-project/erda-infra/providers/mysqlxorm"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/definition/pb"
	"github.com/erda-project/erda/modules/pipeline/providers/definition/db"
)

type config struct {
}

// +provider
type provider struct {
	Cfg      *config
	Log      logs.Logger
	Register transport.Register  `autowired:"service-register"`
	MySQL    mysqlxorm.Interface `autowired:"mysql-xorm"`

	definitionService *definitionService
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.definitionService = &definitionService{p: p, dbClient: &db.Client{Interface: p.MySQL}}
	if p.Register != nil {
		pb.RegisterDefinitionServiceImp(p.Register, p.definitionService)
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.core.pipeline.definition.DefinitionService" || ctx.Type() == pb.DefinitionServiceServerType() || ctx.Type() == pb.DefinitionServiceHandlerType():
		return p.definitionService
	}
	return p
}

func init() {
	servicehub.Register("erda.core.pipeline.definition", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
		OptionalDependencies: []string{"service-register"},
		Description:          "",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
