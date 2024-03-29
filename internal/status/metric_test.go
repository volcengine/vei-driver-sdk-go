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

package status

import (
	"fmt"
	"testing"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/volcengine/vei-driver-sdk-go/extension/dtos"
	"github.com/volcengine/vei-driver-sdk-go/extension/interfaces"
	"github.com/volcengine/vei-driver-sdk-go/extension/interfaces/mocks"
	"github.com/volcengine/vei-driver-sdk-go/extension/responses"
	"github.com/volcengine/vei-driver-sdk-go/pkg/models"
)

func MockDeviceStatusClient() interfaces.DeviceStatusClient {
	mockClient := &mocks.DeviceStatusClient{}
	mockClient.On("DeviceStatusByName", mock.Anything, "device1").Return(
		responses.DeviceStatusResponse{}, errors.NewCommonEdgeXWrapper(fmt.Errorf("status not found")))
	mockClient.On("DeviceStatusByName", mock.Anything, mock.Anything).Return(
		responses.DeviceStatusResponse{Status: dtos.DeviceStatus{DeviceName: "any", OperatingState: string(models.UP)}}, nil)
	mockClient.On("Update", mock.Anything, mock.Anything).Return(
		common.BaseResponse{}, errors.NewCommonEdgeXWrapper(fmt.Errorf("update failed")))
	return mockClient
}

func TestNewManagedDevice(t *testing.T) {
	client = MockDeviceStatusClient()

	device := NewManagedDevice("device1")
	require.NotNil(t, device)
	require.Empty(t, device.Status)

	device = NewManagedDevice("any")
	require.NotNil(t, device)
	require.Equal(t, string(models.UP), device.Status)
}

func TestReport(t *testing.T) {
	client = MockDeviceStatusClient()
	interval = 3

	device := NewManagedDevice("any")

	go device.ReportPeriodically()

	// update device status to DOWN
	device.Status = string(models.DOWN)
	device.ReportImmediately()

	// update device status to UP
	device.Status = string(models.UP)
	device.ReportImmediately()

	// update device reason
	device.Reason = "PASSWORD ERROR"
	device.ReportImmediately()

	// collected inc
	device.DeltaCollected.Inc(10)
	device.LastReportedTime = time.Now().UnixMilli()
	device.ReportImmediately()

	// failures inc
	device.DeltaFailures.Inc(10)
	device.ReportImmediately()

	// wait for ticker trigger
	time.Sleep(time.Second * time.Duration(interval+1))
	device.Stop()
}
