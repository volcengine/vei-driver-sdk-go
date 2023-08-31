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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

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
