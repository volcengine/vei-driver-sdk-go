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

package interfaces

import (
	sdkmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

type Driver interface {
	Initialize(lc logger.LoggingClient, asyncCh chan<- *sdkmodels.AsyncValues, deviceCh chan<- []sdkmodels.DiscoveredDevice) error
	HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest) ([]*sdkmodels.CommandValue, error)
	HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest, params []*sdkmodels.CommandValue) error
	HandlerServiceCall(deviceName string, protocols map[string]models.ProtocolProperties, req sdkmodels.CommandRequest, data []byte) ([]*sdkmodels.CommandValue, error)
	Stop(force bool) error
}

type DeviceHandler interface {
	AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error
	UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error
	RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error
}

type Discovery interface {
	Discover()
}
