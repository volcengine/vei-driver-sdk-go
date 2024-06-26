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
	"fmt"
	"os"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/startup"

	"github.com/volcengine/vei-driver-sdk-go/pkg/interfaces"
)

var agent *Agent

type Option func(agent *Agent)

func Startup(name string, version string, proto interface{}, opts ...Option) {
	agent = &Agent{name: name, version: version}

	for _, opt := range opts {
		opt(agent)
	}

	if driver, ok := proto.(interfaces.Driver); ok {
		agent.driver = driver
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "Please implement the driver interface\n")
		os.Exit(1)
	}

	if handler, ok := proto.(interfaces.DeviceHandler); ok {
		agent.handler = handler
	}
	if discovery, ok := proto.(interfaces.Discovery); ok {
		agent.discovery = discovery
	}
	if debugger, ok := proto.(interfaces.Debugger); ok {
		agent.debugger = debugger
	}
	if webhook, ok := proto.(interfaces.Webhook); ok {
		agent.webhook = webhook
	}

	startup.Bootstrap(agent.name, agent.version, agent)
}

func StatusManager() interfaces.StatusManager {
	return agent.StatusManager
}

func Reporter() interfaces.Reporter {
	return agent.reporter
}
