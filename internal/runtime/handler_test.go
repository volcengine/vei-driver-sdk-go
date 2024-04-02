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

package runtime

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/volcengine/vei-driver-sdk-go/internal/status"
	"github.com/volcengine/vei-driver-sdk-go/pkg/contracts"
	"github.com/volcengine/vei-driver-sdk-go/pkg/interfaces"
	"github.com/volcengine/vei-driver-sdk-go/pkg/interfaces/mocks"
	"github.com/volcengine/vei-driver-sdk-go/pkg/logger"
)

func MockStatusManager(devices map[string]*contracts.Device) interfaces.StatusManager {
	mockStatusManager := &mocks.StatusManager{}
	mockStatusManager.On("OnAddDevice", mock.Anything).Run(
		func(args mock.Arguments) {
			deviceName := args[0].(string)
			devices[deviceName] = &contracts.Device{Name: deviceName}
		},
	)
	mockStatusManager.On("OnRemoveDevice", mock.Anything).Run(
		func(args mock.Arguments) {
			deviceName := args[0].(string)
			delete(devices, deviceName)
		},
	)
	mockStatusManager.On("UpdateDeviceStatus", mock.Anything, mock.Anything, mock.Anything).Run(
		func(args mock.Arguments) {
			deviceName, status, reason := args[0].(string), args[1].(string), args[2].(string)
			device := devices[deviceName]
			if device != nil {
				device.OperatingState = contracts.OperatingState(status)
				device.Message = reason
			}
		},
	)
	mockStatusManager.On("OnHandleCommandsSuccessfully", mock.Anything, mock.Anything).Return(nil)
	mockStatusManager.On("OnHandleCommandsFailed", mock.Anything, mock.Anything).Return(nil)
	return mockStatusManager
}

func BenchmarkPostProcessRequests100(b *testing.B) {
	_, sm := status.Default(nil)
	a := &Agent{StatusManager: sm}

	length := 100
	reqs := make([]contracts.ReadRequest, 0, length)
	for i := 0; i < length; i++ {
		cr := models.CommandRequest{DeviceResourceName: fmt.Sprintf("res%03d", i), Type: common.ValueTypeInt32}
		req := contracts.NewReadRequest(cr)
		req.SetResult(contracts.NewSimpleResult(int32(i)))
		reqs = append(reqs, req)
	}

	for i := 0; i < b.N; i++ {
		responses := make([]*models.CommandValue, 0)
		err := a.PostProcessRequests("", reqs, false, &responses)
		require.NoError(b, err)
	}
}

func TestAsyncReport(t *testing.T) {
	asyncCh := make(chan *models.AsyncValues, 1)
	ctx, cancel := context.WithCancel(context.Background())
	a := &Agent{
		asyncCh:       asyncCh,
		log:           logger.D,
		ctx:           ctx,
		stop:          cancel,
		wg:            &sync.WaitGroup{},
		async:         make(chan *contracts.AsyncValues, 1),
		StatusManager: MockStatusManager(nil),
	}

	go a.HandleAsyncResults(a.ctx, a.wg)
	go func() {
		_ = <-asyncCh
	}()

	err := a.ReportEvent(&contracts.AsyncValues{})
	require.NoError(t, err)

	time.Sleep(time.Second * 2)
	cancel()
}

func TestDeviceLifeCycleWithoutHandler(t *testing.T) {
	devices := make(map[string]*contracts.Device, 0)
	mockStatusManager := MockStatusManager(devices)
	noHandlerAgent := &Agent{StatusManager: mockStatusManager, log: logger.D}

	deviceName := "device-001"
	err := noHandlerAgent.AddDevice(deviceName, nil, "")
	require.NoError(t, err)
	require.NotNil(t, devices[deviceName])
	require.Equal(t, deviceName, devices[deviceName].Name)

	err = noHandlerAgent.UpdateDevice(deviceName, nil, "")
	require.NoError(t, err)

	err = noHandlerAgent.RemoveDevice(deviceName, nil)
	require.NoError(t, err)
	require.Nil(t, devices[deviceName])
}

func TestDeviceLifeCycleWithHandler(t *testing.T) {
	mockHandler := &mocks.DeviceHandler{}
	mockHandler.On("AddDevice", mock.Anything).Return(nil)
	mockHandler.On("RemoveDevice", mock.Anything).Return(nil)
	mockHandler.On("UpdateDevice", mock.Anything).Run(
		func(args mock.Arguments) {
			device := args[0].(*contracts.Device)
			if device.Name == "fake-device" {
				device.OperatingState = contracts.UNREACHABLE
				device.Message = "update failed"
			}
		},
	).Return(fmt.Errorf("establish connection failed"))

	devices := make(map[string]*contracts.Device, 0)
	mockStatusManager := MockStatusManager(devices)
	hasHandlerAgent := &Agent{handler: mockHandler, StatusManager: mockStatusManager, log: logger.D}

	deviceName := "fake-device"
	err := hasHandlerAgent.AddDevice(deviceName, nil, "")
	require.NoError(t, err)
	require.NotNil(t, devices[deviceName])
	require.Equal(t, deviceName, devices[deviceName].Name)
	require.Empty(t, devices[deviceName].OperatingState)

	err = hasHandlerAgent.UpdateDevice(deviceName, nil, "")
	require.Error(t, err)
	require.NotNil(t, devices[deviceName])
	require.Equal(t, contracts.UNREACHABLE, devices[deviceName].OperatingState)
	require.Equal(t, "update failed", devices[deviceName].Message)

	err = hasHandlerAgent.RemoveDevice(deviceName, nil)
	require.NoError(t, err)
	require.Nil(t, devices[deviceName])
}

func BenchmarkGroupRequestByCategory10(b *testing.B) {
	length := 10
	reqs := make([]models.CommandRequest, 0, length)
	for i := 0; i < length; i++ {
		if i%10 == 0 {
			reqs = append(reqs, models.CommandRequest{Attributes: map[string]interface{}{contracts.CategoryKey: contracts.Service}})
		} else {
			reqs = append(reqs, models.CommandRequest{Attributes: map[string]interface{}{contracts.CategoryKey: contracts.Property}})
		}
	}

	for n := 0; n < b.N; n++ {
		read, call, err := GroupRequestByCategory(reqs)
		require.NoError(b, err)
		require.Equal(b, int(float32(length)*0.9), len(read))
		require.Equal(b, int(float32(length)*0.1), len(call))
	}
}

func BenchmarkGroupRequestByCategory100(b *testing.B) {
	length := 100
	reqs := make([]models.CommandRequest, 0, length)
	for i := 0; i < length; i++ {
		if i%10 == 0 {
			reqs = append(reqs, models.CommandRequest{Attributes: map[string]interface{}{contracts.CategoryKey: contracts.Service}})
		} else {
			reqs = append(reqs, models.CommandRequest{Attributes: map[string]interface{}{contracts.CategoryKey: contracts.Property}})
		}
	}

	for n := 0; n < b.N; n++ {
		read, call, err := GroupRequestByCategory(reqs)
		require.NoError(b, err)
		require.Equal(b, int(float32(length)*0.9), len(read))
		require.Equal(b, int(float32(length)*0.1), len(call))
	}
}
