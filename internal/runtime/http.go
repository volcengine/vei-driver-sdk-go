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
	"net/http"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/gorilla/mux"

	"github.com/volcengine/vei-driver-sdk-go/internal/controller/debug"
	"github.com/volcengine/vei-driver-sdk-go/internal/controller/discovery"
)

const (
	CustomizedServerPort = ":9999"
)

const (
	ApiDebugRoute     = common.ApiBase + "/debug"
	ApiDebugLogging   = common.ApiBase + "/logging"
	ApiDiscoveryRoute = common.ApiBase + "/device/discovery"
)

type Route struct {
	route   string
	handler func(http.ResponseWriter, *http.Request)
	method  []string
}

func (a *Agent) RegisterRoutes() error {
	routes := []Route{
		{route: ApiDebugRoute, handler: debug.Debug(a.debugger), method: []string{http.MethodPost}},
		{route: ApiDebugLogging, handler: debug.SetDefaultLogLevel, method: []string{http.MethodGet, http.MethodPost}},
		{route: ApiDiscoveryRoute, handler: discovery.Discover(a.discovery), method: []string{http.MethodGet, http.MethodPost}},
	}

	router := mux.NewRouter()
	for _, route := range routes {
		router.HandleFunc(route.route, route.handler).Methods(route.method...)
	}

	server := &http.Server{
		Addr:    CustomizedServerPort,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	return nil
}
