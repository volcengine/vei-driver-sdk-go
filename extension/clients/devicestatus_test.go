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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"

	"github.com/volcengine/vei-driver-sdk-go/extension/requests"
	"github.com/volcengine/vei-driver-sdk-go/extension/responses"
)

func newTestServer(httpMethod string, apiRoute string, expectedResponse interface{}) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != httpMethod {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.EscapedPath() != apiRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(expectedResponse)
		_, _ = w.Write(b)
	}))
}

func TestPatchDeviceStatus(t *testing.T) {
	ts := newTestServer(http.MethodPatch, ApiDeviceStatusRoute, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewDeviceStatusClient(ts.URL)
	res, err := client.Update(context.Background(), requests.UpdateDeviceStatusRequest{})
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestQueryAllDeviceStatus(t *testing.T) {
	ts := newTestServer(http.MethodGet, ApiAllDeviceStatusRoute, responses.MultiDeviceStatusResponse{})
	defer ts.Close()
	client := NewDeviceStatusClient(ts.URL)
	res, err := client.AllDeviceStatus(context.Background(), 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiDeviceStatusResponse{}, res)
}

func TestQueryDeviceStatusByName(t *testing.T) {
	deviceName := "device"
	path := path.Join(ApiDeviceStatusRoute, common.Name, deviceName)
	ts := newTestServer(http.MethodGet, path, responses.DeviceStatusResponse{})
	defer ts.Close()
	client := NewDeviceStatusClient(ts.URL)
	res, err := client.DeviceStatusByName(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, responses.DeviceStatusResponse{}, res)
}
