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
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/volcengine/vei-driver-sdk-go/pkg/contracts"
	"github.com/volcengine/vei-driver-sdk-go/pkg/interfaces"
)

var (
	_ interfaces.StatusManager = (*Manager)(nil)
)

func TestDefaultManager(t *testing.T) {
	_, manager := Default(nil)
	require.NotNil(t, manager)
	require.Equal(t, ExceedConsecutiveErrorNum, manager.decision.policy)
	require.Equal(t, int64(10), manager.decision.threshold)
}

func TestExceedConsecutiveErrorNum(t *testing.T) {
	_, manager := Default(nil)
	require.NotNil(t, manager)

	deviceName := "device1"
	manager.OnAddDevice(deviceName)

	manager.OnHandleCommandsSuccessfully(deviceName, 10)
	device := manager.getManagedDevice(deviceName)
	require.Equal(t, string(contracts.UP), device.Status)

	manager.OnHandleCommandsFailed(deviceName, 20)
	device = manager.getManagedDevice(deviceName)
	require.Equal(t, string(contracts.DOWN), device.Status)

	manager.SetDeviceOnline(deviceName)
	device = manager.getManagedDevice(deviceName)
	require.Equal(t, string(contracts.UP), device.Status)

	manager.SetDeviceOffline(deviceName, "")
	device = manager.getManagedDevice(deviceName)
	require.Equal(t, string(contracts.DOWN), device.Status)

	manager.UpdateDeviceStatus(deviceName, string(contracts.UNREACHABLE), "")
	device = manager.getManagedDevice(deviceName)
	require.Equal(t, string(contracts.UNREACHABLE), device.Status)

	manager.OnRemoveDevice(deviceName)
}
