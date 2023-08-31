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

package resource

import (
	"strings"
)

const (
	DefaultModule = "default"
	Separator     = ":"
)

func CompatibleName(module string, function string) string {
	if module == DefaultModule {
		return function
	}
	return strings.Join([]string{module, function}, Separator)
}

func SplitName(name string) (module string, function string) {
	ss := strings.Split(name, Separator)
	if len(ss) == 2 {
		return ss[0], ss[1]
	}
	return DefaultModule, ss[0]
}
