package ui

import (
	"fmt"
	"iso6346/parser"
	"iso6346/cont"
	"strings"
)

func PrintContNum(cn parser.ContNum) {

	b := strings.Builder{}
	b.WriteString(fmtRegexIn(cn.RegexIn))
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmtParsedContNum(cn))
	b.WriteString(fmt.Sprintln())
	fmt.Print(b.String())
}

func PrintOwnerCode(oce parser.OwnerCodeOptEquipCat) {

	b := strings.Builder{}
	b.WriteString(fmtRegexIn(oce.RegexIn))
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmtOwnerCodeOptEquipCat(oce))
	b.WriteString(fmt.Sprintln())
	fmt.Print(b.String())
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
