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
	Decode(code cont.OwnerCode) cont.Owner

	AllCodes() []string

	GenerateRandomCodes(count int) []cont.OwnerCode
}

// OwnerUpdater updates owners of an implemented source.
type OwnerUpdater interface {
	Update(newOwners map[string]cont.Owner) error
}

// TimestampUpdater updates a timestamp with an implemented time.
type TimestampUpdater interface {
	Update() error
}

// EquipCatDecoder represents a data source for equipment category ID specific data.
type EquipCatDecoder interface {
	Decode(ID string) cont.EquipCat

	AllIDs() []string
}

// LengthDecoder encapsulates data sources for length
type LengthDecoder interface {
	Decode(code string) cont.Length

	AllCodes() []string
}

// HeightAndWidthDecoder encapsulates data sources for height and width
type HeightAndWidthDecoder interface {
	Decode(code string) cont.HeightAndWidth

	AllCodes() []string
}

// TypeDecoder represents a data source for type data
type TypeDecoder interface {
	Decode(code string) cont.TypeAndGroup

	AllCodes() []string
}
