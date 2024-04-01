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

package requests

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/volcengine/vei-driver-sdk-go/extension/dtos"
)

const (
	ExampleUUID    = "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc"
	TestDeviceName = "TestDevice"
)

var testNowTime = time.Now().Unix()

var testUpdateDeviceStatus = UpdateDeviceStatusRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	Status: mockUpdateDeviceStatus(),
}

func mockUpdateDeviceStatus() dtos.UpdateDeviceStatus {
	testId := ExampleUUID
	testDeviceName := TestDeviceName
	testOperatingState := models.Up
	d := dtos.UpdateDeviceStatus{}
	d.Id = &testId
	d.DeviceName = &testDeviceName
	d.OperatingState = &testOperatingState
	d.UpTime = &testNowTime
	d.DownTime = &testNowTime
	d.LastReportedTime = &testNowTime
	return d
}

func TestUpdateDeviceStatusRequest_UnmarshalJSON(t *testing.T) {
	valid := testUpdateDeviceStatus
	resultTestBytes, _ := json.Marshal(testUpdateDeviceStatus)
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		req     UpdateDeviceStatusRequest
		args    args
		wantErr bool
	}{
		{"unmarshal UpdateDeviceStatusRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid UpdateDeviceRequest, empty data", UpdateDeviceStatusRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid UpdateDeviceRequest, string data", UpdateDeviceStatusRequest{}, args{[]byte("Invalid UpdateDeviceStatusRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.req
			err := tt.req.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.req, "Unmarshal did not result in expected UpdateDeviceStatusRequest.", err)
			}
		})
	}
}

func TestUpdateDeviceStatusRequest_Validate(t *testing.T) {
	emptyString := " "
	invalidUUID := "invalidUUID"

	valid := testUpdateDeviceStatus
	noReqId := valid
	noReqId.RequestId = ""
	invalidReqId := valid
	invalidReqId.RequestId = invalidUUID

	// id
	validOnlyId := valid
	validOnlyId.Status.DeviceName = nil
	invalidId := valid
	invalidId.Status.Id = &invalidUUID
	// name
	validOnlyName := valid
	validOnlyName.Status.Id = nil
	//nameAndEmptyId := valid
	//nameAndEmptyId.Status.Id = &emptyString
	invalidEmptyName := valid
	invalidEmptyName.Status.DeviceName = &emptyString
	// no id and name
	noIdAndName := valid
	noIdAndName.Status.Id = nil
	noIdAndName.Status.DeviceName = nil

	tests := []struct {
		name        string
		req         UpdateDeviceStatusRequest
		expectError bool
	}{
		{"valid", valid, false},
		{"valid, no Request Id", noReqId, false},
		{"invalid, Request Id is not an uuid", invalidReqId, true},

		{"valid, only id", validOnlyId, false},
		{"invalid, invalid Id", invalidId, true},
		{"valid, only name", validOnlyName, false},
		//{"valid, name and empty Id", nameAndEmptyId, false},
		{"invalid, empty name", invalidEmptyName, true},

		{"invalid, no Id and name", noIdAndName, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw, _ := json.Marshal(tt.req)
			t.Log(string(raw))
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected updateDeviceStatusRequest validation result.", err)
		})
	}
}

func TestNewUpdateDeviceStatusRequest(t *testing.T) {
	expectedApiVersion := common.ApiVersion

	actual := NewUpdateDeviceStatusRequest(dtos.UpdateDeviceStatus{})

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
}
