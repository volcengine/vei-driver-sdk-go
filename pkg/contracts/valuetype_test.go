/*
 * Copyright 2023 Beijing Volcano Engine Technology Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package contracts

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"strings"
	"testing"
)

const filename = "valuetype.go"

func TestValueType(t *testing.T) {
	fileset := token.NewFileSet()
	f, err := parser.ParseFile(fileset, filename, nil, parser.ParseComments)
	if err != nil {
		t.Errorf("failed to parse the value type file: %v", err)
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check(".", fileset, []*ast.File{f}, nil)
	if err != nil {
		t.Error(err)
	}

	for _, decl := range f.Scope.Objects {
		if decl.Kind == ast.Con {
			name := decl.Name
			value := pkg.Scope().Lookup(decl.Name).(*types.Const).Val().String()
			value = strings.Trim(value, "\"")
			if name != value {
				t.Errorf("the constant [%s] maybe inconsistent with edgex", name)
			}
		}
	}
}

func TestValueType_String(t *testing.T) {
	tests := []struct {
		name string
		t    ValueType
	}{
		{name: "Bool", t: Bool},
		{name: "String", t: String},
		{name: "Int32", t: Int32},
		{name: "Float32", t: Float32},
		{name: "Float64", t: Float64},
		{name: "Object", t: Object},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.name {
				t.Errorf("String() = %v, want %v", got, tt.name)
			}
		})
	}
}
