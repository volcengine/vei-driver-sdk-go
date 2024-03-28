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

package dtos

type DeviceStatus struct {
	Id               string  `json:"id,omitempty"`
	DeviceName       string  `json:"deviceName"`
	OperatingState   string  `json:"operatingState"`
	Reason           string  `json:"reason,omitempty"`
	UpTime           int64   `json:"upTime,omitempty"`
	DownTime         int64   `json:"downTime,omitempty"`
	LastReportedTime int64   `json:"lastReportedTime,omitempty"`
	Collected        int64   `json:"collected,omitempty"`
	Failures         int64   `json:"failures,omitempty"`
	Frequency        float64 `json:"frequency,omitempty"`
}

type UpdateDeviceStatus struct {
	Id               *string  `json:"id" validate:"required_without=DeviceName,edgex-dto-uuid"`
	DeviceName       *string  `json:"deviceName" validate:"required_without=Id,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	OperatingState   *string  `json:"operatingState,omitempty"`
	Reason           *string  `json:"reason,omitempty"`
	UpTime           *int64   `json:"upTime,omitempty"`
	DownTime         *int64   `json:"downTime,omitempty"`
	LastReportedTime *int64   `json:"lastReportedTime,omitempty"`
	Collected        *int64   `json:"collected,omitempty"`
	Failures         *int64   `json:"failures,omitempty"`
	Frequency        *float64 `json:"frequency,omitempty"`
}
