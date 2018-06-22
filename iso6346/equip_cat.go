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

package iso6346

import (
	"log"
	"unicode/utf8"
)

// EquipCat has an ID and additional information for the ID.
type EquipCat struct {
	Value EquipCatID
	Info  string
}

// EquipCatID is the equipment category ID in a container number.
type EquipCatID struct {
	Value string
}

// NewEquipCatIDU creates a new equipment category ID with U as value.
func NewEquipCatIDU() EquipCatID {
	return EquipCatID{"U"}
}

// NewEquipCatIDFrom creates a new equipment category ID from a string value.
func NewEquipCatIDFrom(value string) EquipCatID {
	if utf8.RuneCountInString(value) != 1 {
		log.Fatalf("'%s' is not one character", value)
	}
	return EquipCatID{value}
}

// NewEquipCatID creates a new equipment category with an ID and an information.
func NewEquipCatID(id EquipCatID, info string) EquipCat {
	return EquipCat{id, info}
}
