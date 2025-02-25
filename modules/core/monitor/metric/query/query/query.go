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

package query

import (
	"net/url"
	"time"

	"github.com/olivere/elastic"

	"github.com/erda-project/erda-infra/providers/i18n"
	tsql "github.com/erda-project/erda/modules/core/monitor/metric/query/es-tsql"
)

// metric keys
const (
	TimestampKey   = tsql.TimestampKey
	NameKey        = tsql.NameKey
	TagKey         = "tags"
	FieldKey       = "fields"
	ClusterNameKey = TagKey + ".cluster_name"
)

// ResultSet .
type ResultSet struct {
	*tsql.ResultSet
	Details interface{}
	Elapsed struct {
		Search time.Duration `json:"search"`
	} `json:"elapsed"`
}

// Queryer .
type Queryer interface {
	Query(tsql, statement string, params map[string]interface{}, options url.Values) (*ResultSet, error)
	QueryWithFormat(tsql, statement, format string, langCodes i18n.LanguageCodes, params map[string]interface{}, filters []*Filter, options url.Values) (*ResultSet, interface{}, error)
	Export(tsql, statement string, params map[string]interface{}, options url.Values, handle func(id string, source []byte) error) error // 暂时没用到, 待移除

	QueryRaw(metrics, clusters []string, start, end int64, searchSource *elastic.SearchSource) (*elastic.SearchResult, error)
	SearchRaw(indices []string, searchSource *elastic.SearchSource) (*elastic.SearchResult, error)
}
