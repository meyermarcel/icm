package ui

import (
	"fmt"
	"github.com/meyermarcel/iso6346/cont"
	"github.com/meyermarcel/iso6346/parser"
	"github.com/meyermarcel/iso6346/sizetype"
)

type Separators struct {
	OwnerEquip  string
	EquipSerial string
	SerialCheck string
}

func PrintContNum(cn parser.ContNum, seps Separators) {

	fmt.Println(fmtRegexIn(cn.RegexIn))
	fmt.Println()
	fmt.Println(fmtParsedContNum(cn, seps))
	fmt.Println()
}

func PrintOwnerCode(oce parser.OwnerCodeOptEquipCat, sepOwnerEquip string) {

	fmt.Println(fmtRegexIn(oce.RegexIn))
	fmt.Println()
	fmt.Println(fmtOwnerCodeOptEquipCat(oce, sepOwnerEquip))
	fmt.Println()
}

func PrintSizeType(st parser.SizeType) {

	fmt.Println(fmtRegexIn(st.RegexIn))
	fmt.Println()
	fmt.Println(fmtParsedSizeType(st))
	fmt.Println()
}

func PrintGen(cn cont.Number, seps Separators) {
	fmt.Printf("%s%s%s%s%06d%s%d",
		cn.OwnerCode().Value(),
		seps.OwnerEquip,
		cn.EquipCatId().Value(),
		seps.EquipSerial,
		cn.SerialNumber().Value(),
		seps.SerialCheck,
		cn.CheckDigit())
	fmt.Println()
}

func PrintSizeTypeDefs(typeSizDef sizetype.Def) {
	fmt.Println(fmtSizeTypeDef(typeSizDef))
}
