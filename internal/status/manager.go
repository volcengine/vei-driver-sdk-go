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
	"sync"
	"time"

	"github.com/volcengine/vei-driver-sdk-go/pkg/contracts"
	"github.com/volcengine/vei-driver-sdk-go/pkg/logger"
	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

type Manager struct {
	devices  map[string]*ManagedDevice
	decision OfflineDecision
	mutex    sync.Mutex
	logger   logger.Logger
}

func Default(deviceNames []string) (OfflineDecision, *Manager) {
	consecutiveErrorNum := utils.GetIntEnv("ERROR_NUM_THRESHOLD", 10)
	if consecutiveErrorNum <= 0 {
		consecutiveErrorNum = 10
	}
	decision := NewOfflineDecision(ExceedConsecutiveErrorNum, consecutiveErrorNum)
	manager, _ := NewManager(deviceNames, decision)
	return decision, manager
}

func NewManager(deviceNames []string, decision OfflineDecision) (*Manager, error) {
	if err := ValidateOfflineDecision(decision); err != nil {
		return nil, err
	}

	m := &Manager{
		devices:  make(map[string]*ManagedDevice, 0),
		decision: decision,
		mutex:    sync.Mutex{},
		logger:   logger.D,
	}

	for _, deviceName := range deviceNames {
		device := NewManagedDevice(deviceName)
		m.devices[deviceName] = device
		go device.ReportPeriodically()
	}

	return m, nil
}

func (m *Manager) OnAddDevice(deviceName string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	device := NewManagedDevice(deviceName)
	m.devices[deviceName] = device
	go device.ReportPeriodically()
}

func (m *Manager) OnRemoveDevice(deviceName string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	device := m.devices[deviceName]
	if device != nil {
		device.Stop()
		delete(m.devices, deviceName)
	}
}

func (m *Manager) OnHandleCommandsFailed(deviceName string, n int64) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	device := m.getManagedDevice(deviceName)
	device.DeltaFailures.Inc(n)
	device.ConsecutiveErrorNum.Inc(n)

	switch m.decision.policy {
	case ExceedConsecutiveErrorNum:
		if device.ConsecutiveErrorNum.Count() > m.decision.threshold {
			device.Status = string(contracts.DOWN)
		}
	default:
		m.logger.Warnf("offline decision policy [%s] has not been implemented", m.decision.policy)
	}
}

func (m *Manager) OnHandleCommandsSuccessfully(deviceName string, n int64) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	device := m.getManagedDevice(deviceName)
	device.DeltaCollected.Inc(n)
	device.ConsecutiveErrorNum.Clear()

	device.Status = string(contracts.UP)
	device.LastReportedTime = time.Now().UnixMilli()
}

func (m *Manager) SetDeviceOffline(deviceName string, reason string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	device := m.getManagedDevice(deviceName)
	device.Status = string(contracts.DOWN)
	device.Reason = reason
}

func (m *Manager) SetDeviceOnline(deviceName string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	device := m.getManagedDevice(deviceName)
	device.Status = string(contracts.UP)
}

func (m *Manager) UpdateDeviceStatus(deviceName string, status string, reason string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	device := m.getManagedDevice(deviceName)
	device.Status = status
	device.Reason = reason
}

func (m *Manager) getManagedDevice(deviceName string) *ManagedDevice {
	device := m.devices[deviceName]
	if device == nil {
		device = NewManagedDevice(deviceName)
		m.devices[deviceName] = device
	}
	return device
}
