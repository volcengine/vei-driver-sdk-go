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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/volcengine/vei-driver-sdk-go/pkg/resource"
	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

func (a *Agent) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties,
	reqs []sdkmodels.CommandRequest) ([]*sdkmodels.CommandValue, error) {

	responses := make([]*sdkmodels.CommandValue, 0)
	properties, services, _ := GroupByCategory(reqs)

	if len(properties) > 0 {
		result, err := a.driver.HandleReadCommands(deviceName, protocols, properties)
		if err != nil {
			a.StatusManager.OnHandleCommandsFailed(deviceName)
			return nil, err
		}
		responses = append(responses, result...)
	}

	for _, srv := range services {
		data, edgexErr := utils.ParametersFromURLRawQuery(srv)
		if edgexErr != nil {
			return nil, edgexErr
		}
		result, err := a.driver.HandleServiceCall(deviceName, protocols, srv, data)
		if err != nil {
			a.StatusManager.OnHandleCommandsFailed(deviceName)
			return nil, err
		}
		responses = append(responses, result)
	}

	if len(responses) == 0 {
		a.StatusManager.OnHandleCommandsFailed(deviceName)
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "empty responses", nil)
	}

	a.StatusManager.OnHandleCommandsSuccessfully(deviceName)
	return responses, nil
}

func (a *Agent) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties,
	reqs []sdkmodels.CommandRequest, params []*sdkmodels.CommandValue) error {
	err := a.driver.HandleWriteCommands(deviceName, protocols, reqs, params)
	if err != nil {
		a.StatusManager.OnHandleCommandsFailed(deviceName)
	} else {
		a.StatusManager.OnHandleCommandsSuccessfully(deviceName)
	}
	return err
}

func (a *Agent) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	a.log.Infof("[AddDevice]: device '%s' is added", deviceName)
	a.StatusManager.OnAddDevice(deviceName)
	if a.handler == nil {
		return nil
	}
	return a.handler.AddDevice(deviceName, protocols, adminState)
}

func (a *Agent) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	a.log.Infof("[UpdateDevice]: device '%s' is updated", deviceName)
	if a.handler == nil {
		return nil
	}
	return a.handler.UpdateDevice(deviceName, protocols, adminState)
}

func (a *Agent) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	a.log.Infof("[RemoveDevice]: device '%s' is removed", deviceName)
	a.StatusManager.OnRemoveDevice(deviceName)
	if a.handler == nil {
		return nil
	}
	return a.handler.RemoveDevice(deviceName, protocols)
}

func GroupByCategory(reqs []sdkmodels.CommandRequest) (properties, services, events []sdkmodels.CommandRequest) {
	for _, req := range reqs {
		switch resource.GetCategory(req) {
		case resource.Property:
			properties = append(properties, req)
		case resource.Service:
			services = append(services, req)
		case resource.Event:
			events = append(events, req)
		}
	}
	return properties, services, events
}
