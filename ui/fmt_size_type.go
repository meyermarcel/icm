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
	"strings"
	"fmt"
	"github.com/meyermarcel/iso6346/parser"
)

func fmtParsedSizeType(st parser.SizeType, sepSizeType string) string {

	var texts []PosTxt

	texts = append(texts, lengthTxt(0, st.LengthIn))
	texts = append(texts, heightWidthTxt(0, st.HeightWidthIn))
	texts = append(texts, typeAndGroupTxt(0, st.TypeAndGroupIn, sepSizeType))

	b := strings.Builder{}
	b.WriteString(indent)

	b.WriteString(fmtIn(st.LengthIn.In))
	b.WriteString(fmtIn(st.HeightWidthIn.In))
	b.WriteString(sepSizeType)
	b.WriteString(fmtIn(st.TypeAndGroupIn.In))

	b.WriteString(fmtCheckMark(st.TypeAndGroupIn.IsValidFmt()))

	b.WriteString(fmt.Sprintln())

	b.WriteString(fmtTextsWithArrows(texts...))

	return b.String()
}

func lengthTxt(offset int, lengthIn parser.LengthIn) PosTxt {
	if !lengthIn.IsValidFmt() {
		return NewPosHint(offset+len(indent),
			fmt.Sprintf("%s must be a %s", underline("length code"), bold("valid number")))
	}
	if lengthIn.Found {
		return NewPosInfo(offset+len(indent),
			fmt.Sprintf("%s: ", underline("length"))+
				lengthIn.MappedLength.Length)
	}
	return NewPosInfo(offset+len(indent),
		fmt.Sprintf("%s not found", underline("length code")))
}

func heightWidthTxt(offset int, heightWidthIn parser.HeightWidthIn) PosTxt {
	if !heightWidthIn.IsValidFmt() {
		return NewPosHint(offset+len(indent)+1,
			fmt.Sprintf("%s must be a %s", underline("height and width code"), bold("valid number")))
	}
	if heightWidthIn.Found {
		return NewPosInfo(offset+len(indent)+1,
			fmt.Sprintf("%s:  ", underline("width"))+
				heightWidthIn.MappedHeightWidth.Width,
			fmt.Sprintf("%s: ", underline("height"))+
				heightWidthIn.MappedHeightWidth.Height)
	}
	return NewPosInfo(offset+len(indent)+1, fmt.Sprintf("%s not found", underline("height and width code")))
}

func typeAndGroupTxt(offset int, typeAndGroupIn parser.TypeAndGroupIn, sepSizeType string) PosTxt {
	if !typeAndGroupIn.IsValidFmt() {
		return NewPosHint(offset+len(indent+sepSizeType)+2,
			fmt.Sprintf("%s must be a %s", underline("type code"), bold("valid type")))
	}
	if typeAndGroupIn.Found {
		return NewPosInfo(offset+len(indent+sepSizeType)+2,
			fmt.Sprintf("%s ", underline("group"))+
				typeAndGroupIn.MappedTypeAndGroup.MappedGroup.Code+ ": "+
				typeAndGroupIn.MappedTypeAndGroup.MappedGroup.GroupInfo,
			fmt.Sprintf("%s ", underline("type"))+
				typeAndGroupIn.MappedTypeAndGroup.MappedType.Code+ ": "+
				typeAndGroupIn.MappedTypeAndGroup.MappedType.TypeInfo)
	}
	return NewPosInfo(offset+len(indent+sepSizeType)+2,
		fmt.Sprintf("%s not found", underline("type code")))
}
