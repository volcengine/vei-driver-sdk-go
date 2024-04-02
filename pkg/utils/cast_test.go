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

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
)

func TestCastCommandValue(t *testing.T) {
	type args struct {
		resourceName string
		valueType    string
		reading      interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "unsupported type", args: args{valueType: "unknown"}, wantErr: true},
		{name: "bool", args: args{valueType: common.ValueTypeBool, reading: "bool"}, wantErr: true},
		{name: "bool", args: args{valueType: common.ValueTypeBool, reading: "true"}, wantErr: false},
		{name: "string", args: args{valueType: common.ValueTypeString, reading: "string"}, wantErr: false},
		{name: "uint8", args: args{valueType: common.ValueTypeUint8, reading: "255"}, wantErr: false},
		{name: "uint16", args: args{valueType: common.ValueTypeUint16, reading: "255"}, wantErr: false},
		{name: "uint32", args: args{valueType: common.ValueTypeUint32, reading: "255"}, wantErr: false},
		{name: "uint64", args: args{valueType: common.ValueTypeUint64, reading: "255"}, wantErr: false},
		{name: "int8", args: args{valueType: common.ValueTypeInt8, reading: "255"}, wantErr: true},
		{name: "int8", args: args{valueType: common.ValueTypeInt8, reading: "127"}, wantErr: false},
		{name: "int16", args: args{valueType: common.ValueTypeInt16, reading: "255"}, wantErr: false},
		{name: "int32", args: args{valueType: common.ValueTypeInt32, reading: "255"}, wantErr: false},
		{name: "int64", args: args{valueType: common.ValueTypeInt64, reading: "255"}, wantErr: false},
		{name: "float32", args: args{valueType: common.ValueTypeFloat32, reading: "255"}, wantErr: false},
		{name: "float64", args: args{valueType: common.ValueTypeFloat64, reading: "255"}, wantErr: false},
		{name: "object", args: args{valueType: common.ValueTypeObject, reading: map[string]string{}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CastCommandValue(tt.args.resourceName, tt.args.valueType, tt.args.reading)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCheckValueRange(t *testing.T) {
	type args struct {
		valueType string
		reading   interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "string", args: args{valueType: common.ValueTypeString, reading: "string"}, want: true},
		{name: "bool", args: args{valueType: common.ValueTypeBool, reading: true}, want: true},
		{name: "object", args: args{valueType: common.ValueTypeObject, reading: map[string]string{}}, want: true},
		{name: "int32", args: args{valueType: common.ValueTypeInt32, reading: 123}, want: true},
		{name: "uint32", args: args{valueType: common.ValueTypeUint32, reading: 123}, want: true},
		{name: "float32", args: args{valueType: common.ValueTypeFloat32, reading: 123}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValueRange(tt.args.valueType, tt.args.reading); got != tt.want {
				t.Errorf("CheckValueRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkIntValueRange(t *testing.T) {
	type args struct {
		valueType string
		val       int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{valueType: common.ValueTypeInt8, val: math.MaxInt64}, want: false},
		{args: args{valueType: common.ValueTypeInt16, val: math.MaxInt64}, want: false},
		{args: args{valueType: common.ValueTypeInt32, val: math.MaxInt64}, want: false},
		{args: args{valueType: common.ValueTypeInt64, val: math.MaxInt64}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIntValueRange(tt.args.valueType, tt.args.val); got != tt.want {
				t.Errorf("checkIntValueRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkUintValueRange(t *testing.T) {
	type args struct {
		valueType string
		val       uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{valueType: common.ValueTypeUint8, val: math.MaxInt64}, want: false},
		{args: args{valueType: common.ValueTypeUint16, val: math.MaxInt64}, want: false},
		{args: args{valueType: common.ValueTypeUint32, val: math.MaxInt64}, want: false},
		{args: args{valueType: common.ValueTypeUint64, val: math.MaxUint64}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkUintValueRange(tt.args.valueType, tt.args.val); got != tt.want {
				t.Errorf("checkUintValueRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkFloatValueRange(t *testing.T) {
	type args struct {
		valueType string
		val       float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{valueType: common.ValueTypeFloat32, val: math.NaN()}, want: false},
		{args: args{valueType: common.ValueTypeFloat64, val: math.NaN()}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkFloatValueRange(tt.args.valueType, tt.args.val); got != tt.want {
				t.Errorf("checkFloatValueRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
