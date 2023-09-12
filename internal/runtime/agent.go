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
	sdkinterfaces "github.com/edgexfoundry/device-sdk-go/v2/pkg/interfaces"
	sdkmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/service"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"

	"github.com/volcengine/vei-driver-sdk-go/internal/status"
	"github.com/volcengine/vei-driver-sdk-go/pkg/interfaces"
	"github.com/volcengine/vei-driver-sdk-go/pkg/log"
)

type Agent struct {
	name      string
	version   string
	driver    interfaces.Driver
	handler   interfaces.DeviceHandler
	discovery interfaces.Discovery
	debugger  interfaces.Debugger
	eventCb   interfaces.EventCallback
	service   sdkinterfaces.DeviceServiceSDK
	asyncCh   chan<- *sdkmodels.AsyncValues
	deviceCh  chan<- []sdkmodels.DiscoveredDevice
	log       logger.LoggingClient

	OfflineDecision status.OfflineDecision
	StatusManager   interfaces.Manager
}

func (a *Agent) Initialize(lc logger.LoggingClient, asyncCh chan<- *sdkmodels.AsyncValues,
	deviceCh chan<- []sdkmodels.DiscoveredDevice) error {
	a.log = log.C
	a.asyncCh = asyncCh
	a.deviceCh = deviceCh
	a.service = service.RunningService()
	a.eventCb = a

	if a.StatusManager == nil {
		manager, err := status.NewManager(a.OfflineDecision, a.service)
		if err != nil {
			a.OfflineDecision, a.StatusManager = status.Default(a.service)
			a.log.Infof("Use the default status manager with offline decision: %+v", a.OfflineDecision)
		} else {
			a.StatusManager = manager
			a.log.Infof("New status manager with offline decision: %+v", a.OfflineDecision)
		}
	}

	if err := RegisterRoutes(); err != nil {
		return err
	}

	return a.driver.Initialize(a.log, a.asyncCh, a.deviceCh, a.eventCb)
}

func (a *Agent) Stop(force bool) error {
	a.log.Infof("Driver '%s' is stopping...", a.name)
	return a.driver.Stop(force)
}
