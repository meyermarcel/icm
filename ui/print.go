package ui

import (
	"fmt"
	"github.com/meyermarcel/iso6346/cont"
	"github.com/meyermarcel/iso6346/parser"
)

type Separators struct {
	OwnerEquip  string
	EquipSerial string
	SerialCheck string
}

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

func PrintGen(cn cont.Number, separators Separators) {
	fmt.Printf("%s%s%s%s%06d%s%d",
		cn.OwnerCode().Value(),
		separators.OwnerEquip,
		cn.EquipCatId().Value(),
		separators.EquipSerial,
		cn.SerialNumber().Value(),
		separators.SerialCheck,
		cn.CheckDigit())
	fmt.Println()
}
