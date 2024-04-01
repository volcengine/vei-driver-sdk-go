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
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/rcrowley/go-metrics"

	"github.com/volcengine/vei-driver-sdk-go/extension/clients"
	"github.com/volcengine/vei-driver-sdk-go/extension/dtos"
	"github.com/volcengine/vei-driver-sdk-go/extension/interfaces"
	"github.com/volcengine/vei-driver-sdk-go/extension/requests"
	"github.com/volcengine/vei-driver-sdk-go/pkg/contracts"
	"github.com/volcengine/vei-driver-sdk-go/pkg/logger"
	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

var (
	client   interfaces.DeviceStatusClient
	interval int64
)

func init() {
	metadataHost := os.Getenv("CLIENTS_CORE_METADATA_HOST")
	metadataPort := utils.GetIntEnv("CLIENTS_CORE_METADATA_PORT", 59881)
	baseUrl := fmt.Sprintf("http://%s:%d", metadataHost, metadataPort)
	client = http.NewDeviceStatusClient(baseUrl)

	interval = utils.GetIntEnv("DEVICE_STATUS_REPORT_INTERVAL", 30)
	interval = utils.Ternary(interval < 10, 10, interval)
}

type ManagedDevice struct {
	prev dtos.DeviceStatus

	Status              string
	Reason              string
	UpTime              int64
	DownTime            int64
	LastReportedTime    int64
	Frequency           metrics.GaugeFloat64
	DeltaCollected      metrics.Counter
	DeltaFailures       metrics.Counter
	ConsecutiveErrorNum metrics.Counter

	ctx   context.Context
	stop  context.CancelFunc
	mutex sync.Mutex
	flush chan bool
	last  time.Time
}

func NewManagedDevice(deviceName string) *ManagedDevice {
	ctx, cancel := context.WithCancel(context.Background())

	device := &ManagedDevice{
		prev:                dtos.DeviceStatus{DeviceName: deviceName},
		Frequency:           metrics.NewGaugeFloat64(),
		DeltaCollected:      metrics.NewCounter(),
		DeltaFailures:       metrics.NewCounter(),
		ConsecutiveErrorNum: metrics.NewCounter(),
		ctx:                 ctx,
		stop:                cancel,
		mutex:               sync.Mutex{},
		flush:               make(chan bool),
		last:                time.Now(),
	}

	resp, err := client.DeviceStatusByName(ctx, deviceName)
	if err != nil {
		logger.D.Warnf("[StatusManager] get device status for '%s' failed: %v", deviceName, err)
	} else {
		device.prev = resp.Status
		device.Status = resp.Status.OperatingState
		device.Reason = resp.Status.Reason
		logger.D.Debugf("[StatusManager] new managed device for '%s' with status %v", deviceName, device.Status)
	}

	return device
}

func (md *ManagedDevice) ReportPeriodically() {
	logger.D.Debugf("[StatusManager] device %s report status periodically", md.prev.DeviceName)
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	for {
		select {
		case <-md.ctx.Done():
			return
		case <-md.flush:
			md.report()
		case <-ticker.C:
			md.report()
		}
	}
}

func (md *ManagedDevice) ReportImmediately() {
	logger.D.Debugf("[StatusManager] device %s report status immediately", md.prev.DeviceName)
	md.flush <- true
}

func (md *ManagedDevice) Stop() {
	logger.D.Debugf("[StatusManager] device %s stop reporting status", md.prev.DeviceName)
	md.stop()
}

func (md *ManagedDevice) report() {
	md.mutex.Lock()
	defer md.mutex.Unlock()

	changed := false
	request := dtos.UpdateDeviceStatus{DeviceName: &md.prev.DeviceName}

	if md.Status != md.prev.OperatingState {
		md.prev.OperatingState = md.Status
		request.OperatingState = &md.prev.OperatingState
		changed = true
		logger.D.Infof("[StatusManager] update device '%s' status to [%s] at %s", md.prev.DeviceName,
			md.prev.OperatingState, time.Now().Format(time.RFC3339))

		if md.prev.OperatingState == string(contracts.UP) {
			md.UpTime = time.Now().UnixMilli()
		} else {
			md.DownTime = time.Now().UnixMilli()
		}
	}

	if md.Reason != md.prev.Reason {
		md.prev.Reason = md.Reason
		request.Reason = &md.prev.Reason
		changed = true
	}

	if md.UpTime > md.prev.UpTime {
		md.prev.UpTime = md.UpTime
		request.UpTime = &md.prev.UpTime
		changed = true
	}

	if md.DownTime > md.prev.DownTime {
		md.prev.DownTime = md.DownTime
		request.DownTime = &md.prev.DownTime
		changed = true
	}

	if md.LastReportedTime > md.prev.LastReportedTime {
		md.prev.LastReportedTime = md.LastReportedTime
		request.LastReportedTime = &md.prev.LastReportedTime
		changed = true
	}

	if delta := md.DeltaCollected.Count(); delta > 0 {
		md.DeltaCollected.Clear()
		md.prev.Collected += delta
		request.Collected = &md.prev.Collected
		changed = true

		now := time.Now()
		seconds := now.Sub(md.last).Seconds()
		if seconds > 0 {
			md.last = now
			md.Frequency.Update(float64(delta) / seconds)
		}
	}

	if delta := md.DeltaFailures.Count(); delta > 0 {
		md.DeltaFailures.Clear()
		md.prev.Failures += delta
		request.Failures = &md.prev.Failures
		changed = true
	}

	if freq := md.Frequency.Value(); freq != md.prev.Frequency {
		md.Frequency.Update(0.0)
		md.prev.Frequency = freq
		request.Frequency = &md.prev.Frequency
		changed = true
	}

	if !changed {
		return
	}

	if _, err := client.Update(context.Background(), requests.NewUpdateDeviceStatusRequest(request)); err != nil {
		logger.D.Warnf("[StatusManager] update device '%s' status failed: %v", md.prev.DeviceName, err)
	}
}
