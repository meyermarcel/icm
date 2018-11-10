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
	"github.com/meyermarcel/icm/internal/cont"
)

// OwnerDecodeUpdater is the interface that groups the basic Decode and Update methods.
type OwnerDecodeUpdater interface {
	OwnerDecoder
	OwnerUpdater
}

// OwnerDecoder represents a data source for owner specific data.
type OwnerDecoder interface {
	Decode(code string) (bool, cont.Owner)

	GenerateRandomCodes(count int) []string
}

// OwnerUpdater updates owners of an implemented source.
type OwnerUpdater interface {
	Update(newOwners map[string]cont.Owner) error
}

// EquipCatDecoder represents a data source for equipment category ID specific data.
type EquipCatDecoder interface {
	Decode(ID string) (bool, cont.EquipCat)

	AllCatIDs() []string
}

// LengthDecoder encapsulates data sources for length
type LengthDecoder interface {
	Decode(code string) (bool, cont.Length)
}

// HeightWidthDecoder encapsulates data sources for height and width
type HeightWidthDecoder interface {
	Decode(code string) (bool, cont.HeightWidth)
}

// TypeDecoder represents a data source for type data
type TypeDecoder interface {
	Decode(code string) (bool, cont.TypeAndGroup)
}

// TimestampUpdater updates a timestamp with an implemented time.
type TimestampUpdater interface {
	Update() error
}
