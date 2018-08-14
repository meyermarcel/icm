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
	"sort"
	"strings"

	"github.com/meyermarcel/icm/data"
)

func fmtParsedContNum(cn contNumOptSizeTypeIn, seps Separators) string {

	b := strings.Builder{}

	sizeTypeExists := cn.lengthIn.isValidFmt()
	b.WriteString(fmtContNum(cn, seps, sizeTypeExists))

	b.WriteString(fmtCheckMark(cn.isValid()))

	b.WriteString(fmt.Sprintln())

	var texts []posTxt

	texts = append(texts, ownerCodeTxt(cn.ownerCodeIn))
	texts = append(texts, equipCatIDTxt(0, cn.equipCatIDIn, seps.OwnerEquip))

	if !cn.serialNumIn.isValidFmt() {
		texts = append(texts,
			newPosHint(indentSize+len(seps.OwnerEquip+seps.EquipSerial)+6,
				fmt.Sprintf("%s must be %s",
					underline("serial number"),
					bold("6 numbers"))))
	}

	cdPos := indentSize + len(seps.OwnerEquip+seps.EquipSerial+seps.SerialCheck) + 10
	if !cn.checkDigitIn.IsValidCheckDigit {
		if cn.isCheckDigitCalculable() {
			if cn.checkDigitIn.isValidFmt() {
				texts = append(texts, newPosHint(cdPos,
					fmt.Sprintf("calculated %s is %s",
						underline("check digit"),
						green(cn.checkDigitIn.CalcCheckDigit))))
			} else {
				texts = append(texts, newPosHint(cdPos,
					fmt.Sprintf("%s must be a %s (calculated: %s)",
						underline("check digit"),
						bold("number"),
						green(cn.checkDigitIn.CalcCheckDigit))))
			}
		} else {
			texts = append(texts, newPosHint(cdPos,
				fmt.Sprintf("%s is not calculable",
					underline("check digit"))))
		}
	}

	if cn.checkDigitIn.CalcCheckDigit == 10 {
		texts = append(texts, newPosHint(cdPos, fmt.Sprintf("It is not recommended to use a %s", underline("serial number")),
			fmt.Sprintf("that generates %s %s (0).", underline("check digit"), yellow(10))))
	}

	if sizeTypeExists {
		texts = append(texts, lengthTxt(seps.offsetPosForSizeType(), cn.lengthIn))
		texts = append(texts, heightWidthTxt(seps.offsetPosForSizeType(), cn.heightWidthIn))
		texts = append(texts, typeAndGroupTxt(seps.offsetPosForSizeType(), cn.typeAndGroupIn, seps.SizeType))
	}

	b.WriteString(fmtTextsWithArrows(texts...))

	return b.String()
}

func fmtContNum(cn contNumOptSizeTypeIn, seps Separators, additionalSizeType bool) string {

	b := strings.Builder{}

	b.WriteString(indent)
	b.WriteString(fmtOwnerCodeIn(cn.ownerCodeIn))
	b.WriteString(seps.OwnerEquip)
	b.WriteString(fmtIn(cn.equipCatIDIn.input))
	b.WriteString(seps.EquipSerial)
	b.WriteString(fmtIn(cn.serialNumIn))
	b.WriteString(seps.SerialCheck)

	if !cn.checkDigitIn.IsValidCheckDigit && cn.checkDigitIn.isValidFmt() {
		b.WriteString(fmt.Sprintf("%s", red(string(cn.checkDigitIn.value))))
	} else if cn.checkDigitIn.CalcCheckDigit == 10 {
		b.WriteString(fmt.Sprintf("%s", yellow(string(cn.checkDigitIn.value))))
	} else {
		b.WriteString(fmtIn(cn.checkDigitIn.input))
	}

	if additionalSizeType {
		b.WriteString(seps.CheckSize)
		b.WriteString(fmtIn(cn.lengthIn.input))
		b.WriteString(fmtIn(cn.heightWidthIn.input))
		b.WriteString(seps.SizeType)
		b.WriteString(fmtIn(cn.typeAndGroupIn.input))
	}

	return b.String()
}

func equipCatIDTxt(offset int, in equipCatIDIn, sepOwnerEquip string) posTxt {
	if !in.isValidFmt() {
		return newPosHint(offset+indentSize+len(sepOwnerEquip)+3,
			fmt.Sprintf("%s must be %s",
				underline("equipment category id"),
				equipCatIDsAsList()))
	}
	return newPosInfo(offset+indentSize+len(sepOwnerEquip)+3, in.EquipCat.Info)
}

func equipCatIDsAsList() string {

	b := strings.Builder{}

	iDs := data.GetEquipCatIDs()
	sort.Strings(iDs)
	for i, element := range iDs {
		b.WriteString(green(element))
		if i < len(iDs)-2 {
			b.WriteString(", ")
		}
		if i == len(iDs)-2 {
			b.WriteString(" or ")
		}
	}
	return b.String()
}
