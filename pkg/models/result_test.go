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

package models

import (
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
	name := "resource_name"
	value := "string-value"
	result := NewSimpleResult(value)

	cv, err := result.CommandValue(name, common.ValueTypeInt32)
	require.Error(t, err)
	require.Nil(t, cv)

	cv, err = result.CommandValue(name, common.ValueTypeString)
	require.NoError(t, err)
	require.NotNil(t, cv)
	require.Equal(t, value, cv.Value)

	result = NewSimpleResult(nil)
	cv, err = result.CommandValue(name, common.ValueTypeInt32)
	require.Error(t, err)
	require.Nil(t, cv)

	stringArr := []string{"1", "2", "3"}
	result = NewSimpleResult(stringArr)
	cv, err = result.CommandValue(name, common.ValueTypeStringArray)
	require.NoError(t, err)
	require.NotNil(t, cv)
	require.Equal(t, stringArr, cv.Value)
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
