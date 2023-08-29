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

package main

import (
	"github.com/edgexfoundry/device-sdk-go/v2"
	sdkmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/volcengine/vei-driver-sdk-go/pkg/startup"
)

const (
	serviceName string = "device-simple"
)

func main() {
	sd := SlimDriver{}
	startup.Bootstrap(serviceName, device.Version, &sd)
}

type SlimDriver struct {
}

func (s *SlimDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *sdkmodels.AsyncValues, deviceCh chan<- []sdkmodels.DiscoveredDevice) error {

	return nil
}

func (s *SlimDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest) (res []*sdkmodels.CommandValue, err error) {

	return
}

func (s *SlimDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest,
	params []*sdkmodels.CommandValue) error {

	return nil
}

func (s *SlimDriver) Stop(force bool) error {
	return nil
}
