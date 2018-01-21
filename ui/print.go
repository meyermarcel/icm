package ui

import (
	"fmt"
	"iso6346/parser"
	"iso6346/owner"
	"iso6346/container"
)

func PrintParse(pi parser.Input) {

	out := ""
	out += formatParsedInput(pi)
	out += fmt.Sprintln()
	out += fmt.Sprintln()
	out += formatValidatedInput(container.Validate(pi))
	out += fmt.Sprintln()
	fmt.Print(out)
}

func PrintOwnerCodeParse(pi parser.Input) {

	out := ""
	out += formatParsedInput(pi)
	out += fmt.Sprintln()
	out += fmt.Sprintln()
	ownerCode := owner.Validate(pi)
	out += formatOwnerCode(ownerCode)
	out += fmt.Sprintln()
	fmt.Print(out)
}

func PrintGenerate(cn container.Number, firstSep, secondSep, thirdSep string) {
	fmt.Printf("%s%s%s%s%s%s%d",
		cn.OwnerCode().Value(),
		firstSep,
		cn.EquipCatId().Value(),
		secondSep,
		cn.SerialNumber().Value(),
		thirdSep,
		cn.CheckDigit())
}
