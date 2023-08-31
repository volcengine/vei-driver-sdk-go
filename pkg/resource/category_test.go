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

	sdkmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
)

func TestGetCategory(t *testing.T) {
	type args struct {
		req sdkmodels.CommandRequest
	}
	tests := []struct {
		name string
		args args
		want Category
	}{
		{name: "nil attributes", args: args{req: sdkmodels.CommandRequest{}}, want: Property},
		{name: "empty attributes", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{}}}, want: Property},
		{name: "not string", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{CategoryKey: 123}}}, want: Property},
		{name: "property", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{CategoryKey: "property"}}}, want: Property},
		{name: "service", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{CategoryKey: "service"}}}, want: Service},
		{name: "event", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{CategoryKey: "event"}}}, want: Event},
		{name: "other", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{CategoryKey: "other"}}}, want: Property},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCategory(tt.args.req); got != tt.want {
				t.Errorf("GetCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetCategory(t *testing.T) {
	type args struct {
		attributes map[string]interface{}
		category   Category
	}
	tests := []struct {
		name  string
		args  args
		wantC Category
	}{
		{name: "nil attributes", args: args{}, wantC: Property},
		{name: "property", args: args{category: Property}, wantC: Property},
		{name: "service", args: args{category: Service}, wantC: Service},
		{name: "event", args: args{category: Event}, wantC: Event},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.attributes = SetCategory(tt.args.attributes, tt.args.category)
			c := GetCategory(sdkmodels.CommandRequest{Attributes: tt.args.attributes})
			if c != tt.wantC {
				t.Errorf("SetCategory() = %v, want %v", c, tt.wantC)
			}
		})
	}
}

func TestIsProperty(t *testing.T) {
	type args struct {
		req sdkmodels.CommandRequest
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "false", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{}}}, want: true},
		{name: "false", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Service.String()}}}, want: false},
		{name: "true", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Property.String()}}}, want: true},
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
		req sdkmodels.CommandRequest
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "false", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Property.String()}}}, want: false},
		{name: "true", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Service.String()}}}, want: true},
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
		req sdkmodels.CommandRequest
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "false", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Property.String()}}}, want: false},
		{name: "true", args: args{req: sdkmodels.CommandRequest{Attributes: map[string]interface{}{CategoryKey: Event.String()}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEvent(tt.args.req); got != tt.want {
				t.Errorf("IsEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
