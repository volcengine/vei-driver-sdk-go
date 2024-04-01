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

package contracts

import (
	"fmt"
	"strings"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
)

// ResourceCategory indicates the category of device resource
type ResourceCategory string

const (
	Property ResourceCategory = "property"
	Service  ResourceCategory = "service"
	Event    ResourceCategory = "event"
)

func (c ResourceCategory) String() string {
	return string(c)
}

const (
	CategoryKey   = "category"
	DefaultModule = "default"
	Separator     = ":"
)

// GetResourceCategory get the category of resource from the request. Property is returned by default for the compatibility.
func GetResourceCategory(req models.CommandRequest) ResourceCategory {
	if req.Attributes == nil || req.Attributes[CategoryKey] == nil {
		return Property
	}
	c := fmt.Sprintf("%v", req.Attributes[CategoryKey])
	if c != string(Property) && c != string(Service) && c != string(Event) {
		return Property
	}
	return ResourceCategory(c)
}

// SetResourceCategory set the resource category into attributes.
func SetResourceCategory(attributes map[string]interface{}, category ResourceCategory) map[string]interface{} {
	if attributes == nil {
		attributes = make(map[string]interface{}, 0)
	}
	attributes[CategoryKey] = category
	return attributes
}

// IsProperty checks whether the device resource is a property.
func IsProperty(req models.CommandRequest) bool {
	return GetResourceCategory(req) == Property
}

// IsService checks whether the device resource is a service.
func IsService(req models.CommandRequest) bool {
	return GetResourceCategory(req) == Service
}

// IsEvent checks whether the device resource is an event.
func IsEvent(req models.CommandRequest) bool {
	return GetResourceCategory(req) == Event
}

// ConcatResourceName will concat the module and resource defined in vei console to edgex resource name.
func ConcatResourceName(module string, resource string) string {
	if module == DefaultModule {
		return resource
	}
	return strings.Join([]string{module, resource}, Separator)
}

// SplitResourceName will split the edgex resource name to the module and resource defined in vei console.
func SplitResourceName(full string) (module string, resource string) {
	ss := strings.Split(full, Separator)
	if len(ss) == 2 {
		return ss[0], ss[1]
	}
	return DefaultModule, ss[0]
}
