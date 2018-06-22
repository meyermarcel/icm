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

	"github.com/meyermarcel/iso6346/iso6346"
)

// Separators describes all separators between independent ISO 6346 data.
type Separators struct {
	OwnerEquip  string
	EquipSerial string
	SerialCheck string
	CheckSize   string
	SizeType    string
}

func (s *Separators) offsetPosForSizeType() int {
	//     owner                   equipment cat id         serial number            check digit
	return 3 + len(s.OwnerEquip) + 1 + len(s.EquipSerial) + 6 + len(s.SerialCheck) + 1 + len(s.CheckSize)
}

func printContNum(cn contNumOptSizeTypeIn, seps Separators) {

	fmt.Println(fmtRegexIn(cn.regexIn))
	fmt.Println()
	fmt.Println(fmtParsedContNum(cn, seps))
	fmt.Println()
}

func printOwnerCode(oce ownerCodeOptEquipCatIDIn) {

	fmt.Println(fmtRegexIn(oce.regexIn))
	fmt.Println()
	fmt.Println(fmtOwnerCode(oce))
	fmt.Println()
}

func printSizeType(st sizeTypeIn, sepSizeType string) {

	fmt.Println(fmtRegexIn(st.RegexIn))
	fmt.Println()
	fmt.Println(fmtParsedSizeType(st, sepSizeType))
	fmt.Println()
}

// PrintContNum prints a container number with given separators and a new line.
func PrintContNum(cn iso6346.ContNumber, seps Separators) {
	fmt.Printf("%s%s%s%s%06d%s%d",
		cn.OwnerCode().Value(),
		seps.OwnerEquip,
		cn.EquipCatID().Value,
		seps.EquipSerial,
		cn.SerialNumber().Value(),
		seps.SerialCheck,
		cn.CheckDigit())
	fmt.Println()
}
