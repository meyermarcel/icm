package container

import (
	"log"
	"iso6346/owner"
	"iso6346/equip_cat"
	"os"
	"iso6346/gen"
)

func Generate() Number {

	ownerCode, err := owner.NewCode(gen.Random(6, gen.LetterRunes))
	equipCatId, err := equip_cat.NewId(gen.Random(1, equip_cat.Ids))
	serialNumber, err := NewSerialNumber(gen.Random(6, gen.NumberRunes))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	checkDigit := CalculateCheckDigit(ownerCode,
		equipCatId,
		serialNumber)

	return NewContainerNumber(ownerCode,
		equipCatId,
		serialNumber, checkDigit)
}
