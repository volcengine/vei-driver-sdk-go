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

package debug

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/volcengine/vei-driver-sdk-go/pkg/logger"
)

const (
	LogLevel = "level"
)

func SetDefaultLogLevel(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	lvl := query.Get(LogLevel)
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte(fmt.Sprintf("failed to update default log level: %v\n", err)))
		return
	}
	logger.D.SetLevel(level)
	_ = logger.C.SetLogLevel(level.String())
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("success"))
	logger.D.Infof("update default log level to '%s'", level.String())
	return
}
