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
	"context"

	"github.com/volcengine/vei-driver-sdk-go/extension/requests"
	"github.com/volcengine/vei-driver-sdk-go/pkg/contracts"
	"github.com/volcengine/vei-driver-sdk-go/pkg/logger"
)

// Driver is a minimal interface to interact with a specific class of devices.
type Driver interface {
	// Initialize performs protocol-specific initialization for the device service.
	// The given *AsyncValues channel can be used to push asynchronous readings to Core Data.
	Initialize(logger logger.Logger, asyncCh chan<- *contracts.AsyncValues) error
	// ReadProperty passes a slice of ReadRequest each representing a read operation for the specific device resource.
	// The device service has the flexibility to set the final result of the request.
	// SetResult: set the value read by driver instance.
	// Failed: set error encountered when performing a read operation.
	// Skip: just skip the request.
	ReadProperty(device *contracts.Device, reqs []contracts.ReadRequest) error
	// WriteProperty passes a slice of WriteRequest each representing a write operation for the specific device resource.
	// The parameter can be obtained through the Param function.
	WriteProperty(device *contracts.Device, reqs []contracts.WriteRequest) error
	// CallService passes a CallRequest which representing a service call operation for the specific device resource.
	// The payload can be obtained through the Payload function.
	CallService(device *contracts.Device, reqs []contracts.CallRequest) error
	// Stop instructs the protocol-specific DS code to shut down gracefully, or
	// if the force parameter is 'true', immediately. The driver is responsible
	// for closing any in-use channels, including the channel used to send async
	// readings (if supported).
	Stop(force bool) error
}

// DeviceHandler is an optional interface to handle the system event of device
type DeviceHandler interface {
	// AddDevice is a callback function that is invoked when a new device associated with this driver is added
	AddDevice(device *contracts.Device) error
	// UpdateDevice is a callback function that is invoked when a device associated with this driver is updated
	UpdateDevice(device *contracts.Device) error
	// RemoveDevice is a callback function that is invoked when a device associated with this driver is removed
	RemoveDevice(device *contracts.Device) error
}

// Discovery is an optional interface implemented by driver that support dynamic device discovery.
type Discovery interface {
	// Discover triggers protocol specific device discovery.
	Discover(ctx context.Context, param *requests.DiscoveryParameter, deviceCh chan<- *contracts.Device)
}

// TODO: Debugger is an optional interface implemented by driver that support debugging for device profile.
type Debugger interface {
	Debug()
}
