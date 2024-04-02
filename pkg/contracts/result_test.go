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
	"reflect"
	"testing"
	"time"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/stretchr/testify/require"
)

var (
	_ Result = (*SimpleResult)(nil)
	_ Result = (*NativeResult)(nil)
)

func TestSimpleResult_Value(t *testing.T) {
	value := "string-value"
	result := NewSimpleResult(value)
	require.Equal(t, value, result.Value())
}

func TestSimpleResult_UnixNano(t *testing.T) {
	value := "string-value"
	result := NewSimpleResult(value)
	require.Equal(t, int64(0), result.UnixNano())
	now := time.Now()
	result = NewSimpleResult(value).WithTime(now)
	require.Equal(t, now.UnixNano(), result.UnixNano())
}

func TestSimpleResult_Tags(t *testing.T) {
	value := "string-value"
	tags := map[string]string{}
	result := NewSimpleResult(value).WithTags(tags)
	require.Equal(t, tags, result.Tags())
}

func TestSimpleResult_CommandValue(t *testing.T) {
	tests := []struct {
		name      string
		valueType string
		value     interface{}
		cast      bool
		wantValue interface{}
		wantErr   bool
	}{
		{name: "error type", valueType: common.ValueTypeInt32, value: "100", wantValue: 100, wantErr: true},
		{name: "cast type", valueType: common.ValueTypeInt32, value: "100", cast: true, wantValue: int32(100), wantErr: false},
		{name: "correct type", valueType: common.ValueTypeString, value: "100", wantValue: "100", wantErr: false},
		{name: "nil result", valueType: common.ValueTypeInt32, value: nil, wantValue: 100, wantErr: true},
		{name: "string arr", valueType: common.ValueTypeStringArray, value: []string{"1", "2", "3"}, wantValue: []string{"1", "2", "3"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSimpleResult(tt.value).WithCast(tt.cast).CommandValue(tt.name, tt.valueType)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !reflect.DeepEqual(got.Value, tt.wantValue) {
				t.Errorf("CommandValue() got = %v, want %v", got.Value, tt.wantValue)
			}
		})
	}
}

func TestNativeResult_Value(t *testing.T) {
	native, _ := models.NewCommandValue("", common.ValueTypeString, "string-value")
	result := NewNativeResult(native)
	require.Equal(t, native.Value, result.Value())

	result = NewNativeResult(nil)
	require.Equal(t, nil, result.Value())
}

func TestNativeResult_UnixNano(t *testing.T) {
	native, _ := models.NewCommandValue("", common.ValueTypeString, "string-value")
	result := NewNativeResult(native)
	require.Equal(t, native.Origin, result.UnixNano())

	result = NewNativeResult(nil)
	require.Equal(t, int64(0), result.UnixNano())
}

func TestNativeResult_Tags(t *testing.T) {
	native, _ := models.NewCommandValue("", common.ValueTypeString, "string-value")
	result := NewNativeResult(native)
	require.Equal(t, native.Tags, result.Tags())

	result = NewNativeResult(nil)
	require.Equal(t, map[string]string(nil), result.Tags())
}

func TestNativeResult_CommandValue(t *testing.T) {
	native, _ := models.NewCommandValue("", common.ValueTypeString, "string-value")
	result := NewNativeResult(native)
	cv, err := result.CommandValue("", "")
	require.NoError(t, err)
	require.Equal(t, native, cv)

	var null *models.CommandValue
	result = NewNativeResult(null)
	cv, err = result.CommandValue("", "")
	require.NoError(t, err)
	require.Equal(t, null, cv)
}
