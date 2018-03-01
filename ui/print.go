package ui

import (
	"fmt"
	"iso6346/parser"
	"iso6346/cont"
)

func PrintContNum(cn parser.ContNum) {

	fmt.Println(fmtRegexIn(cn.RegexIn))
	fmt.Println()
	fmt.Println(fmtParsedContNum(cn))
	fmt.Println()
}

func PrintOwnerCode(oce parser.OwnerCodeOptEquipCat) {

	fmt.Println(fmtRegexIn(oce.RegexIn))
	fmt.Println()
	fmt.Println(fmtOwnerCodeOptEquipCat(oce))
	fmt.Println()
}

func PrintGen(cn cont.Number, firstSep, secondSep, thirdSep string) {
	fmt.Printf("%s%s%s%s%06d%s%d",
		cn.OwnerCode().Value(),
		firstSep,
		cn.EquipCatId().Value(),
		secondSep,
		cn.SerialNumber().Value(),
		thirdSep,
		cn.CheckDigit())
	fmt.Println()
}
