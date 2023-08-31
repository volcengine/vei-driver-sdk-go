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
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"

	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

type Manager struct {
	devices  map[string]*ManagedDevice
	decision OfflineDecision
	mutex    sync.Mutex
	logger   logger.LoggingClient
	ds       interfaces.DeviceServiceSDK
}

func Default(ds interfaces.DeviceServiceSDK) *Manager {
	consecutiveErrorNum := utils.GetIntEnv("ERROR_NUM_THRESHOLD", 10)
	manager, _ := NewManager(NewOfflineDecision(ExceedConsecutiveErrorNum, consecutiveErrorNum), ds)
	return manager
}

func NewManager(decision OfflineDecision, ds interfaces.DeviceServiceSDK) (*Manager, error) {
	m := &Manager{
		devices:  make(map[string]*ManagedDevice, 0),
		decision: decision,
		mutex:    sync.Mutex{},
		logger:   ds.GetLoggingClient(),
		ds:       ds,
	}
	return m, ValidateOfflineDecision(m.decision)
}

func (m *Manager) OnAddDevice(deviceName string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.devices[deviceName] = NewManagedDevice(deviceName, m.ds)
}

func (m *Manager) OnRemoveDevice(deviceName string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.devices, deviceName)
}

func (m *Manager) OnHandleCommandsFailed(deviceName string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	device := m.getManagedDevice(deviceName)
	device.ConsecutiveErrorNum.Inc(1)
	m.trySetDeviceOffline(device)
}

func (m *Manager) OnHandleCommandsSuccessfully(deviceName string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	device := m.getManagedDevice(deviceName)
	device.ConsecutiveErrorNum.Clear()
	m.updateDeviceStatus(device, Online)
}

func (m *Manager) getManagedDevice(deviceName string) *ManagedDevice {
	device := m.devices[deviceName]
	if device == nil {
		device = NewManagedDevice(deviceName, m.ds)
		m.devices[deviceName] = device
	}
	return device
}

func (m *Manager) trySetDeviceOffline(device *ManagedDevice) {
	switch m.decision.policy {
	case ExceedConsecutiveErrorNum:
		if device.ConsecutiveErrorNum.Count() > m.decision.threshold {
			m.updateDeviceStatus(device, Offline)
		}
	default:
		m.logger.Warnf("offline decision policy [%s] has not been implemented", m.decision.policy)
	}
}

func (m *Manager) updateDeviceStatus(device *ManagedDevice, newStatus DeviceStatus) {
	cachedDevice, err := m.ds.GetDeviceByName(device.Name)
	if err != nil {
		m.logger.Errorf("failed to get device by name '%s'", device.Name)
	}

	if device.Status == newStatus {
		return
	}

	device.Status = newStatus
	if newStatus == Online {
		now := time.Now()
		device.OnlineTime = now.UnixMilli()
		m.logger.Infof("change device '%s' status to [%s] at %s", device.Name, device.Status, now.Format(time.RFC3339))
	} else if newStatus == Offline {
		now := time.Now()
		device.OfflineTime = now.UnixMilli()
		m.logger.Infof("change device '%s' status to [%s] at %s", device.Name, device.Status, now.Format(time.RFC3339))
	}

	labelIdx := -1
	for idx, label := range cachedDevice.Labels {
		if strings.HasPrefix(label, LabelPrefix) {
			labelIdx = idx
		}
	}

	content, _ := json.Marshal(device)
	newLabel := LabelPrefix + string(content)

	if labelIdx != -1 {
		cachedDevice.Labels[labelIdx] = newLabel
	} else {
		cachedDevice.Labels = append(cachedDevice.Labels, newLabel)
	}

	if err = m.ds.UpdateDevice(cachedDevice); err != nil {
		m.logger.Errorf("failed to update device %s", device.Name)
	}
}
