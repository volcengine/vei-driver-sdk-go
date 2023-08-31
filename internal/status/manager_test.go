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
	"errors"
	"testing"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/interfaces/mocks"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDefaultManager(t *testing.T) {
	mockDeviceService := &mocks.DeviceServiceSDK{}
	mockDeviceService.On("GetLoggingClient").Return(logger.NewMockClient())

	manager := Default(mockDeviceService)
	require.NotNil(t, manager)
	require.Equal(t, ExceedConsecutiveErrorNum, manager.decision.policy)
	require.Equal(t, int64(10), manager.decision.threshold)
}

func TestExceedConsecutiveErrorNum(t *testing.T) {
	mockDevice := models.Device{Name: "test_device"}
	mockDeviceService := &mocks.DeviceServiceSDK{}
	mockDeviceService.On("GetLoggingClient").Return(logger.NewMockClient())
	mockDeviceService.On("GetDeviceByName", mockDevice.Name).Return(mockDevice, nil)
	mockDeviceService.On("GetDeviceByName", mock.Anything).Return(models.Device{}, errors.New("device not found"))
	mockDeviceService.On("UpdateDevice", mock.Anything).Return(
		func(device models.Device) error {
			mockDevice.Labels = device.Labels
			return nil
		},
	)

	manager, err := NewManager(OfflineDecision{ExceedConsecutiveErrorNum, 0}, mockDeviceService)
	require.Error(t, err)

	threshold := 5
	manager, err = NewManager(OfflineDecision{ExceedConsecutiveErrorNum, int64(threshold)}, mockDeviceService)
	require.NoError(t, err)

	manager.OnAddDevice(mockDevice.Name)
	manager.OnHandleCommandsSuccessfully(mockDevice.Name)
	require.Greater(t, len(mockDevice.Labels), 0)
	require.Contains(t, mockDevice.Labels[0], Online, mockDevice.Labels)

	for i := 0; i <= threshold; i++ {
		manager.OnHandleCommandsFailed(mockDevice.Name)
	}
	require.Greater(t, len(mockDevice.Labels), 0)
	require.Contains(t, mockDevice.Labels[0], Offline, mockDevice.Labels)
	manager.OnRemoveDevice(mockDevice.Name)
}
