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

package cont

// HeightAndWidth describes width and height of first code in specified standard size code.
type HeightAndWidth struct {
	Width  string
	Height string
}

// Length describes length of second code in the specified standard size code.
type Length struct {
	Length string
}

// Type has code and information about the specified standard type.
type Type struct {
	Code string
	Info string
}

// Group has code and information about an specified type group.
type Group struct {
	Code string
	Info string
}

// TypeAndGroup encapsulates type and group.
type TypeAndGroup struct {
	TypeCont Type
	Group    Group
}

// TypeCode returns type code.
func (mtg TypeAndGroup) TypeCode() string {
	return mtg.TypeCont.Code
}

// TypeInfo returns type information.
func (mtg TypeAndGroup) TypeInfo() string {
	return mtg.TypeCont.Info
}

// GroupCode returns group code.
func (mtg TypeAndGroup) GroupCode() string {
	return mtg.Group.Code
}

// GroupInfo returns group information.
func (mtg TypeAndGroup) GroupInfo() string {
	return mtg.Group.Info
}
