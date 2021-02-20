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

import (
	"github.com/meyermarcel/icm/cont"
)

// OwnerDecodeUpdater is the interface that groups the Decode and Update methods.
type OwnerDecodeUpdater interface {
	OwnerDecoder
	OwnerUpdater
}

// OwnerDecoder decodes a code to an owner or generates a random owner.
type OwnerDecoder interface {
	Decode(code string) (bool, cont.Owner)

	GetAllOwnerCodes() []string
}

// OwnerUpdater updates owners of an implemented source.
type OwnerUpdater interface {
	Update(newOwners map[string]cont.Owner) error
}

// EquipCatDecoder decodes an ID to an equipment category.
type EquipCatDecoder interface {
	Decode(ID string) (bool, cont.EquipCat)

	AllCatIDs() []string
}

// LengthDecoder decodes a code to a length.
type LengthDecoder interface {
	Decode(code string) (bool, cont.Length)
}

// HeightWidthDecoder decodes a code to height and width.
type HeightWidthDecoder interface {
	Decode(code string) (bool, cont.HeightWidth)
}

// TypeDecoder decodes a code to type and group.
type TypeDecoder interface {
	Decode(code string) (bool, cont.TypeAndGroup)
}

// TimestampUpdater updates a timestamp with an implemented time.
type TimestampUpdater interface {
	Update() error
}
