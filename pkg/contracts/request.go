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

package contracts

import (
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"

	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

type BaseRequest interface {
	// Native returns the native command request.
	Native() *models.CommandRequest
	// Module returns the thingmodel module where the resource belongs to.
	Module() string
	// Resource returns the identifier of the thingmodel resource.
	Resource() string
	// ValueType returns the value type of the thingmodel resource.
	ValueType() ValueType
	// Attributes returns the attributes defined for the thingmodel resource.
	Attributes() map[string]interface{}
	// SetResult set the handle result of the request.
	SetResult(result Result)
	// Result get the handle result of the request.
	Result() Result
	// Failed set the error encountered when handling the request.
	Failed(error error)
	// Error returns the error encountered when handling the request.
	Error() error
	// Skip will skip the request.
	Skip()
	// Skipped indicates whether the request is skipped or not.
	Skipped() bool
}

type ReadRequest interface {
	BaseRequest
}

type WriteRequest interface {
	BaseRequest
	// Param returns the input parameter of the write request.
	Param() *models.CommandValue
}

type CallRequest interface {
	BaseRequest
	// Payload returns the payload of the call request.
	Payload() []byte
}

func NewReadRequest(req models.CommandRequest) ReadRequest {
	r := newRequest(req)
	return r
}

func NewWriteRequest(req models.CommandRequest, param *models.CommandValue) WriteRequest {
	r := newRequest(req)
	r.param = param
	return r
}

func NewCallRequest(req models.CommandRequest) (CallRequest, error) {
	payload, err := utils.ParametersFromURLRawQuery(req)
	if err != nil {
		return nil, err
	}
	r := newRequest(req)
	r.payload = payload
	return r, nil
}

func newRequest(req models.CommandRequest) *request {
	module, resource := SplitResourceName(req.DeviceResourceName)
	return &request{
		native:     &req,
		module:     module,
		resource:   resource,
		valueType:  ValueType(req.Type),
		attributes: req.Attributes,
	}
}

type request struct {
	native     *models.CommandRequest
	module     string
	resource   string
	valueType  ValueType
	attributes map[string]interface{}
	result     Result
	error      error
	skipped    bool
	param      *models.CommandValue
	payload    []byte
}

func (r *request) Native() *models.CommandRequest {
	return r.native
}

func (r *request) Module() string {
	return r.module
}

func (r *request) Resource() string {
	return r.resource
}

func (r *request) ValueType() ValueType {
	return r.valueType
}

func (r *request) Attributes() map[string]interface{} {
	return r.attributes
}

func (r *request) SetResult(result Result) {
	r.result = result
}

func (r *request) Result() Result {
	return r.result
}

func (r *request) Failed(error error) {
	r.error = error
}

func (r *request) Error() error {
	return r.error
}

func (r *request) Skip() {
	r.skipped = true
}

func (r *request) Skipped() bool {
	return r.skipped
}

func (r *request) Param() *models.CommandValue {
	return r.param
}

func (r *request) Payload() []byte {
	return r.payload
}
