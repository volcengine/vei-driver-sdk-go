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
	"reflect"
	"sync"

	sdkmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/volcengine/vei-driver-sdk-go/pkg/contracts"
)

func (a *Agent) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties,
	reqs []sdkmodels.CommandRequest) ([]*sdkmodels.CommandValue, error) {

	responses := make([]*sdkmodels.CommandValue, 0)
	readRequests, callRequests, err := GroupRequestByCategory(reqs)
	if err != nil {
		a.log.Errorf("group requests by category failed: %v", err)
		return nil, err
	}

	device := contracts.WrapDevice(deviceName, protocols)

	if len(readRequests) > 0 {
		if err = a.driver.ReadProperty(device, readRequests); err != nil {
			a.PostProcessDevice(device, err)
			a.StatusManager.OnHandleCommandsFailed(deviceName, 1)
			return nil, err
		}
		if err = a.PostProcessRequests(deviceName, readRequests, false, &responses); err != nil {
			return responses, err
		}
	}
	if len(callRequests) > 0 {
		if err = a.driver.CallService(device, callRequests); err != nil {
			a.PostProcessDevice(device, err)
			a.StatusManager.OnHandleCommandsFailed(deviceName, 1)
			return nil, err
		}
		if err = a.PostProcessRequests(deviceName, callRequests, false, &responses); err != nil {
			return responses, err
		}
	}

	return responses, nil
}

func (a *Agent) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties,
	reqs []sdkmodels.CommandRequest, params []*sdkmodels.CommandValue) (err error) {

	if len(reqs) == 0 {
		return errors.NewCommonEdgeX(errors.KindServerError, "unexpected empty write request", nil)
	}
	if len(reqs) != len(params) {
		return errors.NewCommonEdgeX(errors.KindServerError, "the length of requests and params is not match", nil)
	}

	device := contracts.WrapDevice(deviceName, protocols)
	requests := make([]contracts.WriteRequest, len(reqs))
	for i := 0; i < len(reqs); i++ {
		requests[i] = contracts.NewWriteRequest(reqs[i], params[i])
	}

	if err = a.driver.WriteProperty(device, requests); err != nil {
		a.PostProcessDevice(device, err)
		a.StatusManager.OnHandleCommandsFailed(deviceName, 1)
		return err
	}

	return a.PostProcessRequests(deviceName, requests, true, nil)
}

func (a *Agent) PostProcessRequests(deviceName string, reqs interface{}, write bool, cvs *[]*sdkmodels.CommandValue) error {
	rValue := reflect.ValueOf(reqs)
	if rValue.Kind() != reflect.Slice {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "requests to be processed not a slice", nil)
	}
	for i := 0; i < rValue.Len(); i++ {
		req, ok := rValue.Index(i).Interface().(contracts.BaseRequest)
		if !ok {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, "request can not assert as BaseRequest", nil)
		}
		if req.Skipped() || !write && req.Error() == nil && req.Result() == nil {
			continue
		}
		if err := req.Error(); err != nil {
			a.StatusManager.OnHandleCommandsFailed(deviceName, 1)
			if a.StrictMode || write {
				return err
			}
			continue
		}
		if result := req.Result(); result != nil && cvs != nil {
			cv, err := result.CommandValue(req.Native().DeviceResourceName, req.Native().Type)
			if err != nil {
				return err
			}
			*cvs = append(*cvs, cv)
		}
		a.StatusManager.OnHandleCommandsSuccessfully(deviceName, 1)
	}
	return nil
}

func (a *Agent) ReportEvent(event *contracts.AsyncValues) error {
	a.async <- event
	return nil
}

func (a *Agent) HandleAsyncResults(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer func() {
		wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			a.log.Infof("Stop handle async results...")
			return
		case result := <-a.async:
			a.StatusManager.OnHandleCommandsSuccessfully(result.DeviceName, int64(len(result.CommandValues)))
			a.asyncCh <- result.Transform()
		}
	}
}

func (a *Agent) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, _ models.AdminState) error {
	a.log.Infof("device '%s' is added", deviceName)
	a.StatusManager.OnAddDevice(deviceName)
	if a.handler == nil {
		return nil
	}
	// Call the interface 'AddDevice' if the driver has implemented the handler.
	device := contracts.WrapDevice(deviceName, protocols)
	err := a.handler.AddDevice(device)
	a.PostProcessDevice(device, err)
	return err
}

func (a *Agent) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, _ models.AdminState) error {
	a.log.Infof("device '%s' is updated", deviceName)
	if a.handler == nil {
		return nil
	}
	// Call the interface 'UpdateDevice' if the driver has implemented the handler.
	device := contracts.WrapDevice(deviceName, protocols)
	err := a.handler.UpdateDevice(device)
	a.PostProcessDevice(device, err)
	return err
}

func (a *Agent) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	a.log.Infof("device '%s' is removed", deviceName)
	a.StatusManager.OnRemoveDevice(deviceName)
	if a.handler == nil {
		return nil
	}
	// Call the interface 'RemoveDevice' if the driver has implemented the handler.
	device := contracts.WrapDevice(deviceName, protocols)
	return a.handler.RemoveDevice(device)
}

func (a *Agent) PostProcessDevice(device *contracts.Device, err error) {
	device.UpdateStateByError(err)
	if device.OperatingState == "" {
		return
	}
	a.StatusManager.UpdateDeviceStatus(device.Name, string(device.OperatingState), device.Message)
}

func GroupRequestByCategory(reqs []sdkmodels.CommandRequest) ([]contracts.ReadRequest, []contracts.CallRequest, error) {
	var (
		read []contracts.ReadRequest
		call []contracts.CallRequest
	)

	for _, req := range reqs {
		switch contracts.GetResourceCategory(req) {
		case contracts.Property:
			request := contracts.NewReadRequest(req)
			read = append(read, request)
		case contracts.Service:
			request, err := contracts.NewCallRequest(req)
			if err != nil {
				return nil, nil, err
			}
			call = append(call, request)
		}
	}

	return read, call, nil
}
