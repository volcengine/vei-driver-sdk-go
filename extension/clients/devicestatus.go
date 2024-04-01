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

package http

import (
	"context"
	"net/url"
	"path"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	"github.com/volcengine/vei-driver-sdk-go/extension/interfaces"
	"github.com/volcengine/vei-driver-sdk-go/extension/requests"
	"github.com/volcengine/vei-driver-sdk-go/extension/responses"
)

const (
	ApiDeviceStatusRoute       = common.ApiBase + "/devicestatus"
	ApiAllDeviceStatusRoute    = ApiDeviceStatusRoute + "/" + common.All
	ApiDeviceStatusByNameRoute = ApiDeviceStatusRoute + "/" + common.Name + "/{" + common.Name + "}"
)

type DeviceStatusClient struct {
	baseUrl string
}

// NewDeviceStatusClient creates an instance of DeviceStatusClient
func NewDeviceStatusClient(baseUrl string) interfaces.DeviceStatusClient {
	return &DeviceStatusClient{
		baseUrl: baseUrl,
	}
}

func (sc DeviceStatusClient) Update(ctx context.Context, req requests.UpdateDeviceStatusRequest) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	err = utils.PatchRequest(ctx, &res, sc.baseUrl, ApiDeviceStatusRoute, nil, req)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (sc DeviceStatusClient) AllDeviceStatus(ctx context.Context, offset int, limit int) (res responses.MultiDeviceStatusResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, sc.baseUrl, ApiAllDeviceStatusRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (sc DeviceStatusClient) DeviceStatusByName(ctx context.Context, name string) (res responses.DeviceStatusResponse, err errors.EdgeX) {
	path := path.Join(ApiDeviceStatusRoute, common.Name, name)
	err = utils.GetRequest(ctx, &res, sc.baseUrl, path, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}
