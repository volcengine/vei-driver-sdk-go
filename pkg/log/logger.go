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

package log

import (
	"io"

	"github.com/sirupsen/logrus"

	"github.com/volcengine/vei-driver-sdk-go/pkg/log/format"
	"github.com/volcengine/vei-driver-sdk-go/pkg/log/writer"
)

type Logger interface {
	GeneralLogger
	FormattedLogger
	SetLevel(level logrus.Level)
}

type GeneralLogger interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type FormattedLogger interface {
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type LoggerConfig struct {
	LogLevel  logrus.Level
	Formatter logrus.Formatter
	Writer    io.Writer
}

var D Logger

func NewLogger(cfg LoggerConfig) Logger {
	if cfg.Writer == nil {
		cfg.Writer = writer.Stderr()
	}

	if cfg.Formatter == nil {
		cfg.Formatter = format.NewColorFormatter()
	}

	logger := logrus.New()
	logger.SetLevel(cfg.LogLevel)
	logger.SetFormatter(cfg.Formatter)
	logger.SetOutput(cfg.Writer)
	logger.SetReportCaller(true)
	return logger
}

func init() {
	D = NewLogger(LoggerConfig{LogLevel: logrus.InfoLevel})
}
