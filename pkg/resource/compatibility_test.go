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

package resource

import (
	"testing"
)

func TestCompatibleName(t *testing.T) {
	type args struct {
		module   string
		function string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{module: "default", function: "p1"}, want: "p1"},
		{args: args{module: "m1", function: "p1"}, want: "m1:p1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompatibleName(tt.args.module, tt.args.function); got != tt.want {
				t.Errorf("CompatibleName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name         string
		args         args
		wantModule   string
		wantFunction string
	}{
		{args: args{name: "p1"}, wantModule: "default", wantFunction: "p1"},
		{args: args{name: "m1:p1"}, wantModule: "m1", wantFunction: "p1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotModule, gotFunction := SplitName(tt.args.name)
			if gotModule != tt.wantModule {
				t.Errorf("SplitName() gotModule = %v, want %v", gotModule, tt.wantModule)
			}
			if gotFunction != tt.wantFunction {
				t.Errorf("SplitName() gotFunction = %v, want %v", gotFunction, tt.wantFunction)
			}
		})
	}
}
