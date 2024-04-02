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
	"fmt"
	"math"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/spf13/cast"
)

func CastCommandValue(resourceName string, valueType string, reading interface{}) (result *models.CommandValue, err error) {
	if !CheckValueRange(valueType, reading) {
		err = fmt.Errorf("parse reading failed, reading %v is out of the type(%v)'s range", reading, valueType)
		return nil, err
	}

	var val interface{}
	switch valueType {
	case common.ValueTypeBool:
		val, err = cast.ToBoolE(reading)
	case common.ValueTypeString:
		val, err = cast.ToStringE(reading)
	case common.ValueTypeUint8:
		val, err = cast.ToUint8E(reading)
	case common.ValueTypeUint16:
		val, err = cast.ToUint16E(reading)
	case common.ValueTypeUint32:
		val, err = cast.ToUint32E(reading)
	case common.ValueTypeUint64:
		val, err = cast.ToUint64E(reading)
	case common.ValueTypeInt8:
		val, err = cast.ToInt8E(reading)
	case common.ValueTypeInt16:
		val, err = cast.ToInt16E(reading)
	case common.ValueTypeInt32:
		val, err = cast.ToInt32E(reading)
	case common.ValueTypeInt64:
		val, err = cast.ToInt64E(reading)
	case common.ValueTypeFloat32:
		val, err = cast.ToFloat32E(reading)
	case common.ValueTypeFloat64:
		val, err = cast.ToFloat64E(reading)
	case common.ValueTypeObject:
		val = reading
	default:
		return nil, fmt.Errorf("cast result failed, unsupported value type: %s", valueType)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s' reading: %v", resourceName, err)
	}

	return models.NewCommandValue(resourceName, valueType, val)
}

func CheckValueRange(valueType string, reading interface{}) bool {
	switch valueType {
	case common.ValueTypeInt8, common.ValueTypeInt16, common.ValueTypeInt32, common.ValueTypeInt64:
		val := cast.ToInt64(reading)
		return checkIntValueRange(valueType, val)
	case common.ValueTypeUint8, common.ValueTypeUint16, common.ValueTypeUint32, common.ValueTypeUint64:
		val := cast.ToUint64(reading)
		return checkUintValueRange(valueType, val)
	case common.ValueTypeFloat32, common.ValueTypeFloat64:
		val := cast.ToFloat64(reading)
		return checkFloatValueRange(valueType, val)
	default:
		return true
	}
}

func checkIntValueRange(valueType string, val int64) bool {
	var isValid = false
	switch valueType {
	case common.ValueTypeInt8:
		isValid = Ternary(val >= math.MinInt8 && val <= math.MaxInt8, true, false)
	case common.ValueTypeInt16:
		isValid = Ternary(val >= math.MinInt16 && val <= math.MaxInt16, true, false)
	case common.ValueTypeInt32:
		isValid = Ternary(val >= math.MinInt32 && val <= math.MaxInt32, true, false)
	case common.ValueTypeInt64:
		isValid = true
	}
	return isValid
}

func checkUintValueRange(valueType string, val uint64) bool {
	var isValid = false
	switch valueType {
	case common.ValueTypeUint8:
		isValid = Ternary(val <= math.MaxUint8, true, false)
	case common.ValueTypeUint16:
		isValid = Ternary(val <= math.MaxUint16, true, false)
	case common.ValueTypeUint32:
		isValid = Ternary(val <= math.MaxUint32, true, false)
	case common.ValueTypeUint64:
		isValid = Ternary(val <= math.MaxUint64, true, false)
	}
	return isValid
}

func checkFloatValueRange(valueType string, val float64) bool {
	var isValid = false
	switch valueType {
	case common.ValueTypeFloat32:
		isValid = Ternary(!math.IsNaN(val) && math.Abs(val) <= math.MaxFloat32, true, false)
	case common.ValueTypeFloat64:
		isValid = Ternary(!math.IsNaN(val) && !math.IsInf(val, 0), true, false)
	}
	return isValid
}
