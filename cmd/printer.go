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

package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/meyermarcel/icm/internal/data"

	"github.com/meyermarcel/icm/internal/cont"
)

type separators struct {
	OwnerEquip  string
	EquipSerial string
	SerialCheck string
	CheckSize   string
	SizeType    string
}

func (s *separators) offsetPosForSizeType() int {
	//     owner                   equipment cat id         serial number            check digit
	return 3 + len(s.OwnerEquip) + 1 + len(s.EquipSerial) + 6 + len(s.SerialCheck) + 1 + len(s.CheckSize)
}

func printContNumVal(writer io.Writer, cn contNumOptSizeTypeIn, data decoders, seps separators) {

	b := strings.Builder{}
	b.WriteString(fmtRegexIn(cn.regexIn))
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmtParsedContNum(cn, data, seps))
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmt.Sprintln())
	io.WriteString(writer, b.String())
}

func printOwnerCode(writer io.Writer, oce ownerCodeOptEquipCatIDIn, owner data.OwnerDecoder) {

	b := strings.Builder{}

	b.WriteString(fmtRegexIn(oce.regexIn))
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmtOwnerCode(oce, owner.GenerateRandomCodes))
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmt.Sprintln())
	io.WriteString(writer, b.String())
}

func printSizeType(writer io.Writer, st sizeTypeIn, sepSizeType string) {

	b := strings.Builder{}
	b.WriteString(fmtRegexIn(st.RegexIn))
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmtParsedSizeType(st, sepSizeType))
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmt.Sprintln())
	io.WriteString(writer, b.String())
}

func printContNum(writer io.Writer, cn cont.Number, seps separators) {
	b := strings.Builder{}
	b.WriteString(
		fmt.Sprintf("%s%s%s%s%06d%s%d",
			cn.OwnerCode().Value(),
			seps.OwnerEquip,
			cn.EquipCatID().Value,
			seps.EquipSerial,
			cn.SerialNumber().Value(),
			seps.SerialCheck,
			cn.CheckDigit()))
	b.WriteString(fmt.Sprintln())
	io.WriteString(writer, b.String())
}
