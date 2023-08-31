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

package api

import (
	"net/http"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/service"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
)

// AddRoute allows leveraging the existing internal web server to add routes specific to Device Service.
func AddRoute(route string, handler func(http.ResponseWriter, *http.Request), methods ...string) error {
	return service.RunningService().AddRoute(route, handler, methods...)
}

// GetLoggingClient returns the logger.LoggingClient. The name was chosen to avoid conflicts
// with service.DeviceService.LoggingClient struct field.
func GetLoggingClient() logger.LoggingClient {
	return service.RunningService().GetLoggingClient()
}
