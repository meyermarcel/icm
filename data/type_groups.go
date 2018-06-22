// Copyright © 2018 Marcel Meyer meyermarcel@posteo.de
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package data

import "github.com/meyermarcel/iso6346/iso6346"

// TypeAndGroup encapsulates type and group.
type TypeAndGroup struct {
	typeIso iso6346.Type
	group   iso6346.TypeGroup
}

// GetTypeCode returns type code.
func (mtg TypeAndGroup) GetTypeCode() string {
	return mtg.typeIso.Code
}

// GetTypeInfo returns type information.
func (mtg TypeAndGroup) GetTypeInfo() string {
	return mtg.typeIso.Info
}

// GetGroupCode returns group code.
func (mtg TypeAndGroup) GetGroupCode() string {
	return mtg.group.Code
}

// GetGroupInfo returns group information.
func (mtg TypeAndGroup) GetGroupInfo() string {
	return mtg.group.Info
}

// GetTypeAndGroup returns type and group for type code.
func GetTypeAndGroup(code string) TypeAndGroup {
	typeAndGroup := TypeAndGroup{}
	typeValue, typeFound := getType(code)
	group, groupFound := getGroup(string(code[0]))

	if !typeFound && !groupFound {
		return typeAndGroup
	}

	typeAndGroup.typeIso = typeValue
	typeAndGroup.group = group
	return typeAndGroup
}
