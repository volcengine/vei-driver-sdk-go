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
	"math"
	"testing"

	sdkmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

type TestObject struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func TestParametersFromURLRawQuery(t *testing.T) {
	type args struct {
		req sdkmodels.CommandRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "invalid raw query", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{URLRawQuery: ";;;"}}}, wantErr: true},
		{name: "empty raw query", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{URLRawQuery: ""}}}, wantErr: false},
		{name: "not base64", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{URLRawQuery: ServiceParams + "=123"}}}, wantErr: true},
		{name: "parse success", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{URLRawQuery: ServiceParams + "=eyJ4IjoxMCwieSI6MjB9"}}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParametersFromURLRawQuery(tt.args.req)
			if err != nil {
				t.Log(err.Error())
				if !tt.wantErr {
					t.Errorf("ParametersFromURLRawQuery() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

func TestObjectToQueryParam(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "nil obj", args: args{obj: nil}, wantErr: true},
		{name: "unsupported type", args: args{obj: make(chan int)}, wantErr: true},
		{name: "unsupported value", args: args{obj: math.Inf(0)}, wantErr: true},
		{name: "int", args: args{obj: 123}, wantErr: false},
		{name: "string", args: args{obj: `{"x":10,"y":20}`}, wantErr: false},
		{name: "struct", args: args{obj: &TestObject{X: 10, Y: 20}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ObjectToQueryParam(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObjectToQueryParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				t.Log(got)
			}
		})
	}
}

func TestExtractPassage(t *testing.T) {
	type args struct {
		protocols map[string]models.ProtocolProperties
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "key not found", args: args{protocols: nil}, wantErr: true},
		{name: "found", args: args{protocols: map[string]models.ProtocolProperties{Passage: nil}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := PassageFromProtocols(tt.args.protocols)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractPassage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
