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

package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	err := Validate("plain text")
	require.Error(t, err)
}

func TestTernary(t *testing.T) {
	type args struct {
		x bool
		a any
		b any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{"int1", args{true, 1, 2}, 1},
		{"int2", args{false, 1, 2}, 2},
		{"string1", args{true, "1", "2"}, "1"},
		{"string2", args{false, "1", "2"}, "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ternary(tt.args.x, tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}
