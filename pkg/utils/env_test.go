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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetIntEnv(t *testing.T) {
	key, fallback := "ERROR_NUM_THRESHOLD", int64(10)
	got := GetIntEnv(key, fallback)
	require.Equal(t, fallback, got)

	_ = os.Setenv(key, "abc")
	got = GetIntEnv(key, fallback)
	require.Equal(t, fallback, got)

	_ = os.Setenv(key, "20")
	got = GetIntEnv(key, fallback)
	require.Equal(t, int64(20), got)
}

func TestGetBoolEnv(t *testing.T) {
	key := "STRICTLY_READ"
	got := GetBoolEnv(key, false)
	require.Equal(t, false, got)

	_ = os.Setenv(key, "abc")
	got = GetBoolEnv(key, false)
	require.Equal(t, false, got)

	_ = os.Setenv(key, "true")
	got = GetBoolEnv(key, false)
	require.Equal(t, true, got)
}
