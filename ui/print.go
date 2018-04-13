// Copyright © 2018 Marcel Meyer meyermarcel@posteo.de
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func PrintOwnerCode(oce parser.OwnerCode) {

	fmt.Println(fmtRegexIn(oce.RegexIn))
	fmt.Println()
	fmt.Println(fmtOwnerCode(oce))
	fmt.Println()
}

func PrintSizeType(st parser.SizeType, sepSizeType string) {

	fmt.Println(fmtRegexIn(st.RegexIn))
	fmt.Println()
	fmt.Println(fmtParsedSizeType(st, sepSizeType))
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
