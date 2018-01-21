package container

import (
	"iso6346/parser"
	"iso6346/validator"
	"iso6346/owner"
	"iso6346/equip_cat"
	"log"
	"os"
)

type NumberValidated struct {
	OwnerCode            validator.Input
	EquipCatId           validator.Input
	SerialNumber         validator.Input
	CheckDigit           validator.Input
	IsValidCheckDigit    bool
	CalculatedCheckDigit int
}

func Validate(input parser.Input) NumberValidated {

	vi := NumberValidated{OwnerCode: validator.NewInput(input.GetMatch(0, 3), 3),
		EquipCatId: validator.NewInput(input.GetMatchSingle(3), 1),
		SerialNumber: validator.NewInput(input.GetMatch(4, 10), 6),
		CheckDigit: validator.NewInput(input.GetMatchSingle(10), 1)}

	if !vi.IsCheckDigitCalculable() {
		return vi
	}

	ownerCode, err := owner.NewCode(vi.OwnerCode.Value())

	equipCatId, err := equip_cat.NewId(vi.EquipCatId.Value())

	serialNumber, err := NewSerialNumber(vi.SerialNumber.Value())

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	vi.CalculatedCheckDigit = CalculateCheckDigit(ownerCode,
		equipCatId,
		serialNumber)

	if !vi.CheckDigit.IsComplete() {
		return vi
	}

	checkDigit, err := NewCheckDigit(vi.CheckDigit.Value())

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	vi.IsValidCheckDigit = vi.CalculatedCheckDigit == checkDigit.value

	return vi

}

func (vi NumberValidated) IsCheckDigitCalculable() bool {
	return vi.OwnerCode.IsComplete() && vi.EquipCatId.IsComplete() && vi.SerialNumber.IsComplete()
}
