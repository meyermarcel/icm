package ui

import (
	"fmt"
	"github.com/meyermarcel/iso6346/parser"
	"github.com/meyermarcel/iso6346/cont"
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

func PrintGen(cn cont.Number, sepOwnerEquip, sepEquipSerial, sepSerialCheck string) {
	fmt.Printf("%s%s%s%s%06d%s%d",
		cn.OwnerCode().Value(),
		sepOwnerEquip,
		cn.EquipCatId().Value(),
		sepEquipSerial,
		cn.SerialNumber().Value(),
		sepSerialCheck,
		cn.CheckDigit())
	fmt.Println()
}
