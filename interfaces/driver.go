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

	"github.com/volcengine/vei-driver-sdk-go/pkg/log"
	"github.com/volcengine/vei-driver-sdk-go/pkg/models"
)

// Driver is a minimal interface to interact with a specific class of devices.
type Driver interface {
	// Initialize performs protocol-specific initialization for the device service.
	// The given *AsyncValues channel can be used to push asynchronous readings to Core Data.
	Initialize(logger log.Logger, asyncCh chan<- *models.AsyncValues) error
	// ReadProperty passes a slice of ReadRequest each representing a read operation for the specific device resource.
	// The device service has the flexibility to set the final result of the request.
	// SetResult: set the value read by driver instance.
	// Failed: set error encountered when performing a read operation.
	// Skip: just skip the request.
	ReadProperty(device *models.Device, reqs []models.ReadRequest) error
	// WriteProperty passes a slice of WriteRequest each representing a write operation for the specific device resource.
	// The parameter can be obtained through the Param function.
	WriteProperty(device *models.Device, reqs []models.WriteRequest) error
	// CallService passes a CallRequest which representing a service call operation for the specific device resource.
	// The payload can be obtained through the Payload function.
	CallService(device *models.Device, reqs []models.CallRequest) error
	// Stop instructs the protocol-specific DS code to shut down gracefully, or
	// if the force parameter is 'true', immediately. The driver is responsible
	// for closing any in-use channels, including the channel used to send async
	// readings (if supported).
	Stop(force bool) error
}

// DeviceHandler is an optional interface to handle the system event of device
type DeviceHandler interface {
	// AddDevice is a callback function that is invoked when a new device associated with this driver is added
	AddDevice(device *models.Device) error
	// UpdateDevice is a callback function that is invoked when a device associated with this driver is updated
	UpdateDevice(device *models.Device) error
	// RemoveDevice is a callback function that is invoked when a device associated with this driver is removed
	RemoveDevice(device *models.Device) error
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
