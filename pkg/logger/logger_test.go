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

package logger

import (
	"io"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/volcengine/vei-driver-sdk-go/pkg/logger/format"
	"github.com/volcengine/vei-driver-sdk-go/pkg/logger/writer"
)

func TestDefaultLogger(t *testing.T) {
	D.SetLevel(logrus.TraceLevel)
	msg := "this is %d test message on level %s"
	D.Tracef(msg, 1, "trace")
	D.Debugf(msg, 1, "debug")
	D.Infof(msg, 1, "info")
	D.Warnf(msg, 1, "warn")
	D.Errorf(msg, 1, "error")
	D.Trace(1, 2, 3, 4)
	D.Debug(1, 2, 3, 4)
	D.Info(1, 2, 3, 4)
	D.Warn(1, 2, 3, 4)
	D.Error(1, 2, 3, 4)
}

func TestFileLogger(t *testing.T) {
	fileWriter, _ := writer.NewFileWriter(filepath.Join(t.TempDir(), "test.log"), 3)
	logger := NewLogger(LoggerConfig{
		LogLevel:  logrus.TraceLevel,
		Formatter: format.NewColorFormatter(),
		Writer:    io.MultiWriter(writer.Stderr(), fileWriter),
	})
	msg := "this is %d test message on level %s"
	logger.Tracef(msg, 1, "trace")
	logger.Debugf(msg, 1, "debug")
	logger.Infof(msg, 1, "info")
	logger.Warnf(msg, 1, "warn")
	logger.Errorf(msg, 1, "error")
	logger.Trace(1, 2, 3, 4)
	logger.Debug(1, 2, 3, 4)
	logger.Info(1, 2, 3, 4)
	logger.Warn(1, 2, 3, 4)
	logger.Error(1, 2, 3, 4)
}
