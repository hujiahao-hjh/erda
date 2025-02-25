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

package migrator

import "github.com/pingcap/parser/ast"

type StmtType string

const (
	DDL = "DDL"
	DML = "DML"
)

type Block struct {
	typ   StmtType
	nodes []ast.Node
}

func (block Block) Type() StmtType {
	return block.typ
}

func (block Block) Nodes() []ast.Node {
	return block.nodes
}

func AppendBlock(blocks []Block, node ast.Node, typ StmtType) []Block {
	if len(blocks) == 0 {
		return []Block{{typ: typ, nodes: []ast.Node{node}}}
	}
	last := len(blocks) - 1
	if blocks[last].Type() == typ {
		blocks[last].nodes = append(blocks[last].nodes, node)
		return blocks
	}
	return append(blocks, Block{
		typ:   typ,
		nodes: []ast.Node{node},
	})
}
