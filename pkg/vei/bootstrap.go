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

package vei

import (
	"github.com/volcengine/vei-driver-sdk-go/internal/runtime"
	"github.com/volcengine/vei-driver-sdk-go/internal/status"
	"github.com/volcengine/vei-driver-sdk-go/pkg/interfaces"
)

func Bootstrap(name string, version string, driver interface{}, opts ...runtime.Option) {
	runtime.Startup(name, version, driver, opts...)
}

func WithConsecutiveErrorStatusManager(threshold int64) runtime.Option {
	return func(agent *runtime.Agent) {
		agent.OfflineDecision = status.NewOfflineDecision(status.ExceedConsecutiveErrorNum, threshold)
	}
}

func WithContinuousErrorStatusManager(threshold int64) runtime.Option {
	return func(agent *runtime.Agent) {
		agent.OfflineDecision = status.NewOfflineDecision(status.ExceedContinuousErrorDuration, threshold)
	}
}

func WithStatusManager(manager interfaces.StatusManager) runtime.Option {
	return func(agent *runtime.Agent) {
		agent.StatusManager = manager
	}
}

func WithStrictMode(strict bool) runtime.Option {
	return func(agent *runtime.Agent) {
		agent.StrictMode = strict
	}
}
