package ui

import (
	"fmt"
	"iso6346/parser"
	"iso6346/cont"
)

func PrintContNum(cn parser.ContNum) {

	out := ""
	out += fmtRegexIn(cn.RegexIn)
	out += fmt.Sprintln()
	out += fmt.Sprintln()
	out += fmtParsedContNum(cn)
	out += fmt.Sprintln()
	fmt.Print(out)
}

func PrintOwnerCode(oce parser.OwnerCodeOptEquipCat) {

	out := ""
	out += fmtRegexIn(oce.RegexIn)
	out += fmt.Sprintln()
	out += fmt.Sprintln()
	out += fmtOwnerCodeOptEquipCat(oce)
	out += fmt.Sprintln()
	fmt.Print(out)
}

func PrintGenerate(cn cont.Number, firstSep, secondSep, thirdSep string) {
	fmt.Printf("%s%s%s%s%s%s%d",
		cn.OwnerCode().Value(),
		firstSep,
		cn.EquipCatId().Value(),
		secondSep,
		cn.SerialNumber().Value(),
		thirdSep,
		cn.CheckDigit())
}
