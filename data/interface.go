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
	Decode(code string) (bool, cont.Height, cont.Width)
}

// TypeDecoder decodes a code to type and group information.
type TypeDecoder interface {
	Decode(code string) (bool, cont.TypeInfo, cont.GroupInfo)
}

// TimestampUpdater updates a timestamp with an implemented time.
type TimestampUpdater interface {
	Update() error
}
