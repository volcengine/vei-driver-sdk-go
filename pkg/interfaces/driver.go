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

// Driver is a minimal interface to interact with a specific class of devices.
type Driver interface {
	// Initialize performs protocol-specific initialization for the device service.
	// The given *AsyncValues channel can be used to push asynchronous events and
	// readings to Core Data. The given []DiscoveredDevice channel is used to send
	// discovered devices that will be filtered and added to Core Metadata asynchronously.
	Initialize(lc logger.LoggingClient, asyncCh chan<- *sdkmodels.AsyncValues, deviceCh chan<- []sdkmodels.DiscoveredDevice, eventCallback EventCallback) error
	// HandleReadCommands passes a slice of CommandRequest struct each representing
	// a ResourceOperation for a specific device resource.
	HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest) ([]*sdkmodels.CommandValue, error)
	// HandleWriteCommands passes a slice of CommandRequest struct each representing
	// a ResourceOperation for a specific device resource. Since the commands are actuation commands,
	// params provide parameters for the individual command.
	HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest, params []*sdkmodels.CommandValue) error
	// HandleServiceCall passes a CommandRequest struct representing a service call operation
	// and a slice of byte representing the input parameters.
	HandleServiceCall(deviceName string, protocols map[string]models.ProtocolProperties, req sdkmodels.CommandRequest, data []byte) (*sdkmodels.CommandValue, error)
	// Stop instructs the protocol-specific DS code to shut down gracefully, or
	// if the force parameter is 'true', immediately. The driver is responsible
	// for closing any in-use channels, including the channel used to send async
	// readings (if supported).
	Stop(force bool) error
}

// DeviceHandler is an optional interface to handle the system event of device
type DeviceHandler interface {
	// AddDevice is a callback function that is invoked when a new device associated with this driver is added
	AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error
	// UpdateDevice is a callback function that is invoked when a device associated with this driver is updated
	UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error
	// RemoveDevice is a callback function that is invoked when a device associated with this driver is removed
	RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error
}

// TODO: Discovery is an optional interface implemented by driver that support dynamic device discovery.
type Discovery interface {
	// Discover triggers protocol specific device discovery.
	Discover() []sdkmodels.DiscoveredDevice
}

// TODO: Debugger is an optional interface implemented by driver that support debugging for device profile.
type Debugger interface {
	Debug()
}
