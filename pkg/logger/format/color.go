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

package format

import (
	"bytes"
	"fmt"
	"path"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"

	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

type ColorFormatter struct{}

func NewColorFormatter() logrus.Formatter {
	return &ColorFormatter{}
}

func (f *ColorFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.TraceLevel:
		levelColor = int(color.FgWhite)
	case logrus.DebugLevel:
		levelColor = int(color.FgGreen)
	case logrus.InfoLevel:
		levelColor = int(color.FgCyan)
	case logrus.WarnLevel:
		levelColor = int(color.FgYellow)
	default:
		levelColor = int(color.FgRed)
	}
	buffer := utils.Ternary(entry.Buffer != nil, entry.Buffer, &bytes.Buffer{})
	timestamp := entry.Time.Format("2006-01-02 15:04:05.999")
	if entry.HasCaller() {
		location := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		_, _ = fmt.Fprintf(buffer, "\x1b[%dm[%s]\x1b[0m [%s] %s %s\n", levelColor, entry.Level, timestamp, location, entry.Message)
	} else {
		_, _ = fmt.Fprintf(buffer, "\x1b[%dm[%s]\x1b[0m [%s] %s\n", levelColor, entry.Level, timestamp, entry.Message)
	}
	return buffer.Bytes(), nil
}
