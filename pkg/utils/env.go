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

package utils

import (
	"os"
	"strconv"
)

func GetIntEnv(key string, fallback int64) int64 {
	if s, ok := os.LookupEnv(key); ok {
		value, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return fallback
		}
		return value
	}
	return fallback
}

func GetBoolEnv(key string, fallback bool) bool {
	if s, ok := os.LookupEnv(key); ok {
		value, err := strconv.ParseBool(s)
		if err != nil {
			return fallback
		}
		return value
	}
	return fallback
}
