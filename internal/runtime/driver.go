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
	sdkmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/service"
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/startup"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

var (
	d *driver
)

type driver struct {
	Log       logger.LoggingClient
	name      string
	version   string
	driver    sdkmodels.ProtocolDriver
	discovery sdkmodels.ProtocolDiscovery
	asyncCh   chan *sdkmodels.AsyncValues
	deviceCh  chan []sdkmodels.DiscoveredDevice
}

func (d *driver) Initialize(lc logger.LoggingClient, asyncCh chan<- *sdkmodels.AsyncValues, deviceCh chan<- []sdkmodels.DiscoveredDevice) error {
	service.RunningService()
	return nil
}

func (d *driver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest) ([]*sdkmodels.CommandValue, error) {
	// TODO implement me
	panic("implement me")
}

func (d *driver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest, params []*sdkmodels.CommandValue) error {
	// TODO implement me
	panic("implement me")
}

func (d *driver) Stop(force bool) error {
	return nil
}

func (d *driver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	// TODO implement me
	panic("implement me")
}

func (d *driver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	// TODO implement me
	panic("implement me")
}

func (d *driver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	// TODO implement me
	panic("implement me")
}

var _ sdkmodels.ProtocolDriver = (*driver)(nil)

func InitDriver(name, version string, v interface{}) {

	d = &driver{
		Log:     nil,
		name:    name,
		version: version,
	}

	startup.Bootstrap(name, version, d)
}
