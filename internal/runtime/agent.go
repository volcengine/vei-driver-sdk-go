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
	"context"
	"sync"

	sdkinterfaces "github.com/edgexfoundry/device-sdk-go/v2/pkg/interfaces"
	sdkmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/service"
	lc "github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"

	"github.com/volcengine/vei-driver-sdk-go/internal/status"
	"github.com/volcengine/vei-driver-sdk-go/pkg/contracts"
	"github.com/volcengine/vei-driver-sdk-go/pkg/interfaces"
	"github.com/volcengine/vei-driver-sdk-go/pkg/logger"
	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

type Agent struct {
	name      string
	version   string
	driver    interfaces.Driver
	handler   interfaces.DeviceHandler
	discovery interfaces.Discovery
	debugger  interfaces.Debugger
	reporter  interfaces.Reporter
	service   sdkinterfaces.DeviceServiceSDK
	asyncCh   chan<- *sdkmodels.AsyncValues // used by agent
	deviceCh  chan<- []sdkmodels.DiscoveredDevice
	log       logger.Logger

	ctx   context.Context
	stop  context.CancelFunc
	wg    *sync.WaitGroup
	async chan *contracts.AsyncValues // used by driver

	// if the driver is in strict mode, any error in request will be returned, ignored otherwise.
	StrictMode bool
	// now only the decision of 'ConsecutiveErrorNum' can be used
	OfflineDecision status.OfflineDecision
	// if no StatusManager is specified, a default one will be initialized
	StatusManager interfaces.StatusManager
}

func (a *Agent) Initialize(_ lc.LoggingClient, asyncCh chan<- *sdkmodels.AsyncValues,
	deviceCh chan<- []sdkmodels.DiscoveredDevice) error {
	logger.D.Infof("Initialize Driver Agent...")

	a.log = logger.D
	a.asyncCh = asyncCh
	a.deviceCh = deviceCh
	a.service = service.RunningService()
	a.reporter = a

	a.ctx, a.stop = context.WithCancel(context.Background())
	a.wg = &sync.WaitGroup{}

	bufferSize := utils.GetIntEnv("DEVICE_ASYNCBUFFERSIZE", 10)
	a.async = make(chan *contracts.AsyncValues, bufferSize)
	go a.HandleAsyncResults(a.ctx, a.wg)
	a.log.Infof("Set async buffer size: %d", bufferSize)

	deviceNames := make([]string, 0)
	for _, device := range a.service.Devices() {
		deviceNames = append(deviceNames, device.Name)
	}

	if a.StatusManager == nil {
		manager, err := status.NewManager(deviceNames, a.OfflineDecision)
		if err != nil {
			a.OfflineDecision, a.StatusManager = status.Default(deviceNames)
			a.log.Infof("Use the default status manager with offline decision: %+v", a.OfflineDecision)
		} else {
			a.StatusManager = manager
			a.log.Infof("New status manager with offline decision: %+v", a.OfflineDecision)
		}
	}

	if err := RegisterRoutes(); err != nil {
		return err
	}

	return a.driver.Initialize(a.log, a.async)
}

func (a *Agent) Stop(force bool) error {
	a.log.Infof("Driver %s is stopping...", a.name)
	a.stop()
	a.log.Infof("Wait for all goroutines stop...")
	a.wg.Wait()
	return a.driver.Stop(force)
}
