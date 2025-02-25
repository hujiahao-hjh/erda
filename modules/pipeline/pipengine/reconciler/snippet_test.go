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

package reconciler

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestParsePipelineOutputRef(t *testing.T) {
	reffedTask, reffedKey, err := parsePipelineOutputRef("${Dice 文档:OUTPUT:status}")
	spew.Dump(reffedTask, reffedKey)
	assert.NoError(t, err)
	assert.Equal(t, "Dice 文档", reffedTask)
	assert.Equal(t, "status", reffedKey)
}

func TestParsePipelineOutputRefV2(t *testing.T) {
	reffedTask, reffedKey, err := parsePipelineOutputRefV2("${{ outputs.a.b }}")
	spew.Dump(reffedTask, reffedKey)
	assert.NoError(t, err)
	assert.Equal(t, "a", reffedTask)
	assert.Equal(t, "b", reffedKey)
}
