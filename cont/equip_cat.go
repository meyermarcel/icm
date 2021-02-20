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

import "fmt"

// EquipCat has an ID and additional information for the ID.
type EquipCat struct {
	Value string
	Info  string
}

// EquipCatID is the equipment category ID in a container number.
type EquipCatID struct {
	Value string
}

// NewEquipCatID creates a new equipment category with an ID and an information.
func NewEquipCatID(id string, info string) EquipCat {
	return EquipCat{id, info}
}

// IsEquipCatID checks if string is one upper case letter.
func IsEquipCatID(ID string) error {
	if len(ID) != 1 {
		return NewErrContValidate(fmt.Sprintf("%s is not 1 letter long", ID))
	}
	if !isUpperLetter(ID) {
		return NewErrContValidate(fmt.Sprintf("%s is not 1 upper case letter", ID))
	}
	return nil
}
