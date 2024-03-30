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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/sirupsen/logrus"

	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

var C logger.LoggingClient

type CompatLogger struct {
	level string
	Logger
}

func NewCompatLogger(cfg LoggerConfig) logger.LoggingClient {
	return &CompatLogger{
		level:  cfg.LogLevel.String(),
		Logger: NewLogger(cfg),
	}
}

func (c *CompatLogger) SetLogLevel(logLevel string) errors.EdgeX {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindServerError, "failed to set log level", err)
	}
	c.level = logLevel
	c.Logger.SetLevel(level)
	return nil
}

func (c *CompatLogger) LogLevel() string {
	return c.level
}

func (c *CompatLogger) Trace(msg string, args ...interface{}) {
	c.Logger.Trace(msg, args)
}

func (c *CompatLogger) Debug(msg string, args ...interface{}) {
	c.Logger.Debug(msg, args)
}

func (c *CompatLogger) Info(msg string, args ...interface{}) {
	c.Logger.Info(msg, args)
}

func (c *CompatLogger) Warn(msg string, args ...interface{}) {
	c.Logger.Warn(msg, args)
}

func (c *CompatLogger) Error(msg string, args ...interface{}) {
	c.Logger.Error(msg, args)
}

func init() {
	level := utils.GetIntEnv("LOG_LEVEL", int64(logrus.InfoLevel))
	C = NewCompatLogger(LoggerConfig{LogLevel: logrus.Level(level)})
}
