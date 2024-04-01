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

package responses

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/volcengine/vei-driver-sdk-go/extension/dtos"
)

func TestNewDeviceStatusResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedDeviceStatus := dtos.DeviceStatus{DeviceName: "test device status"}
	actual := NewDeviceStatusResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedDeviceStatus)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedDeviceStatus, actual.Status)
}

func TestNewMultiDeviceStatusResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedDeviceStatus := []dtos.DeviceStatus{
		{DeviceName: "test device1"},
		{DeviceName: "test device2"},
	}
	expectedTotalCount := uint32(2)
	actual := NewMultiDeviceStatusResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedDeviceStatus)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedDeviceStatus, actual.Status)
}
