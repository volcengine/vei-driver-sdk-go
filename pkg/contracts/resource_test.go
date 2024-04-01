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
	"testing"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
)

func TestResourceCategory_String(t *testing.T) {
	tests := []struct {
		name string
		c    ResourceCategory
	}{
		{name: "property", c: Property},
		{name: "service", c: Service},
		{name: "event", c: Event},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.name {
				t.Errorf("String() = %v, want %v", got, tt.name)
			}
		})
	}
}

func TestGetResourceCategory(t *testing.T) {
	type args struct {
		req models.CommandRequest
	}
	tests := []struct {
		name string
		args args
		want ResourceCategory
	}{
		{name: "nil attributes", args: args{req: models.CommandRequest{}}, want: Property},
		{name: "empty attributes", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{}}}, want: Property},
		{name: "not string", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{CategoryKey: 123}}}, want: Property},
		{name: "property", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{CategoryKey: "property"}}}, want: Property},
		{name: "service", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{CategoryKey: "service"}}}, want: Service},
		{name: "event", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{CategoryKey: "event"}}}, want: Event},
		{name: "other", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{CategoryKey: "other"}}}, want: Property},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetResourceCategory(tt.args.req); got != tt.want {
				t.Errorf("GetResourceCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetResourceCategory(t *testing.T) {
	type args struct {
		attributes map[string]interface{}
		category   ResourceCategory
	}
	tests := []struct {
		name  string
		args  args
		wantC ResourceCategory
	}{
		{name: "nil attributes", args: args{}, wantC: Property},
		{name: "property", args: args{category: Property}, wantC: Property},
		{name: "service", args: args{category: Service}, wantC: Service},
		{name: "event", args: args{category: Event}, wantC: Event},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.attributes = SetResourceCategory(tt.args.attributes, tt.args.category)
			c := GetResourceCategory(models.CommandRequest{Attributes: tt.args.attributes})
			if c != tt.wantC {
				t.Errorf("SetResourceCategory() = %v, want %v", c, tt.wantC)
			}
		})
	}
}

func TestIsProperty(t *testing.T) {
	type args struct {
		req models.CommandRequest
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "false", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{}}}, want: true},
		{name: "false", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Service}}}, want: false},
		{name: "true", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Property}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsProperty(tt.args.req); got != tt.want {
				t.Errorf("IsProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsService(t *testing.T) {
	type args struct {
		req models.CommandRequest
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "false", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Property}}}, want: false},
		{name: "true", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Service}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsService(tt.args.req); got != tt.want {
				t.Errorf("IsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEvent(t *testing.T) {
	type args struct {
		req models.CommandRequest
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "false", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Property}}}, want: false},
		{name: "true", args: args{req: models.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Event}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEvent(tt.args.req); got != tt.want {
				t.Errorf("IsEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcatResourceName(t *testing.T) {
	type args struct {
		module   string
		resource string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{module: "default", resource: "p1"}, want: "p1"},
		{args: args{module: "m1", resource: "p1"}, want: "m1:p1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatResourceName(tt.args.module, tt.args.resource); got != tt.want {
				t.Errorf("ConcatResourceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitResourceName(t *testing.T) {
	type args struct {
		full string
	}
	tests := []struct {
		name         string
		args         args
		wantModule   string
		wantResource string
	}{
		{args: args{full: "p1"}, wantModule: "default", wantResource: "p1"},
		{args: args{full: "m1:p1"}, wantModule: "m1", wantResource: "p1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotModule, gotResource := SplitResourceName(tt.args.full)
			if gotModule != tt.wantModule {
				t.Errorf("SplitResourceName() gotModule = %v, want %v", gotModule, tt.wantModule)
			}
			if gotResource != tt.wantResource {
				t.Errorf("SplitResourceName() gotFunction = %v, want %v", gotResource, tt.wantResource)
			}
		})
	}
}
