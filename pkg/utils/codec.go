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

package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"

	sdkmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

const (
	URLRawQuery   = "urlRawQuery"
	ServiceParams = "service_params"
	Passage       = "passage"
)

func ParametersFromURLRawQuery(req sdkmodels.CommandRequest) ([]byte, errors.EdgeX) {
	values, err := url.ParseQuery(fmt.Sprint(req.Attributes[URLRawQuery]))
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to parse get command parameter for resource '%s'", req.DeviceResourceName), err)
	}
	param, exists := values[ServiceParams]
	if !exists || len(param) == 0 {
		return []byte{}, nil
	}
	data, err := base64.StdEncoding.DecodeString(param[0])
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to decode '%v' parameter for resource '%s', the value should be json object with base64 encoded", ServiceParams, req.DeviceResourceName), err)
	}
	return data, nil
}

func ObjectToQueryParam(obj interface{}) (map[string]string, error) {
	if obj == nil {
		return nil, fmt.Errorf("the object can not be nil when convert to query param")
	}
	result := make(map[string]string, 0)
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the object in json format: %v", err)
	}
	result[ServiceParams] = base64.StdEncoding.EncodeToString(data)
	return result, nil
}

func PassageFromProtocols(protocols map[string]models.ProtocolProperties) (map[string]string, errors.EdgeX) {
	if passage, ok := protocols[Passage]; ok {
		return passage, nil
	}
	return nil, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to extract passage info from device protocols: key '%s' not found", Passage), nil)
}
