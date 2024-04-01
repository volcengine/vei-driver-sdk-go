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
	"time"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"

	"github.com/volcengine/vei-driver-sdk-go/pkg/common"
	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

// Result defines the interface for a result of the Request.
type Result interface {
	// Value returns the value read by driver instance.
	Value() interface{}
	// UnixNano returns the unix nano timestamp of the result read by driver instance.
	UnixNano() int64
	// Tags returns the custom information of the result.
	Tags() map[string]string
	// CommandValue constructs EdgeX CommandValue based on the content of the result.
	CommandValue(resourceName string, valueType string) (*models.CommandValue, error)
}

type SimpleResult struct {
	value  interface{}
	origin *time.Time
	tags   map[string]string
}

func NewSimpleResult(value interface{}) *SimpleResult {
	return &SimpleResult{value: value}
}

func (r *SimpleResult) WithTime(origin time.Time) *SimpleResult {
	r.origin = &origin
	return r
}

func (r *SimpleResult) WithTags(tags map[string]string) *SimpleResult {
	r.tags = tags
	return r
}

func (r *SimpleResult) Value() interface{} {
	return r.value
}

func (r *SimpleResult) UnixNano() int64 {
	if r.origin == nil {
		return 0
	}
	return r.origin.UnixNano()
}

func (r *SimpleResult) Tags() map[string]string {
	return r.tags
}

func (r *SimpleResult) CommandValue(resourceName string, valueType string) (*models.CommandValue, error) {
	valueType = utils.Ternary(valueType == string(common.StringArray), string(common.Object), valueType)
	return models.NewCommandValueWithOrigin(resourceName, valueType, r.value, r.UnixNano())
}

type NativeResult struct {
	native *models.CommandValue
}

func NewNativeResult(native *models.CommandValue) *NativeResult {
	return &NativeResult{native: native}
}

func (r *NativeResult) Value() interface{} {
	if r.native == nil {
		return nil
	}
	return r.native.Value
}

func (r *NativeResult) UnixNano() int64 {
	if r.native == nil {
		return 0
	}
	return r.native.Origin
}

func (r *NativeResult) Tags() map[string]string {
	if r.native == nil {
		return nil
	}
	return r.native.Tags
}

func (r *NativeResult) CommandValue(_ string, _ string) (*models.CommandValue, error) {
	return r.native, nil
}
