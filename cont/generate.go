package cont

import (
	"iso6346/owner"
	"iso6346/equip_cat"
	"iso6346/gen"
)

func Generate() Number {

	ownerCode := owner.NewCode(gen.Random(6, gen.LetterRunes))
	equipCatId := equip_cat.NewId(gen.Random(1, equip_cat.Ids))
	serialNumber := NewSerialNum(gen.Random(6, gen.NumberRunes))

	checkDigit := CalcCheckDigit(ownerCode,
		equipCatId,
		serialNumber)

	return NewContNum(ownerCode,
		equipCatId,
		serialNumber, checkDigit)
}
