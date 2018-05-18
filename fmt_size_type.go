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

package main

import (
	"fmt"
	"strings"
)

func fmtParsedSizeType(st sizeTypeIn, sepSizeType string) string {

	var texts []posTxt

	texts = append(texts, lengthTxt(0, st.lengthIn))
	texts = append(texts, heightWidthTxt(0, st.heightWidthIn))
	texts = append(texts, typeAndGroupTxt(0, st.typeAndGroupIn, sepSizeType))

	b := strings.Builder{}
	b.WriteString(indent)

	b.WriteString(fmtIn(st.lengthIn.input))
	b.WriteString(fmtIn(st.heightWidthIn.input))
	b.WriteString(sepSizeType)
	b.WriteString(fmtIn(st.typeAndGroupIn.input))

	b.WriteString(fmtCheckMark(st.typeAndGroupIn.isValidFmt()))

	b.WriteString(fmt.Sprintln())

	b.WriteString(fmtTextsWithArrows(texts...))

	return b.String()
}

func lengthTxt(offset int, lengthIn lengthIn) posTxt {
	if !lengthIn.isValidFmt() {
		return newPosHint(offset+len(indent),
			fmt.Sprintf("%s must be a %s", underline("length code"), bold("valid number")))
	}
	if lengthIn.Found {
		return newPosInfo(offset+len(indent),
			fmt.Sprintf("%s: ", underline("length"))+
				lengthIn.MappedLength.Length)
	}
	return newPosInfo(offset+len(indent),
		fmt.Sprintf("%s not found", underline("length code")))
}

func heightWidthTxt(offset int, heightWidthIn heightWidthIn) posTxt {
	if !heightWidthIn.isValidFmt() {
		return newPosHint(offset+len(indent)+1,
			fmt.Sprintf("%s must be a %s", underline("height and width code"), bold("valid number")))
	}
	if heightWidthIn.Found {
		return newPosInfo(offset+len(indent)+1,
			fmt.Sprintf("%s:  ", underline("width"))+
				heightWidthIn.MappedHeightWidth.Width,
			fmt.Sprintf("%s: ", underline("height"))+
				heightWidthIn.MappedHeightWidth.Height)
	}
	return newPosInfo(offset+len(indent)+1, fmt.Sprintf("%s not found", underline("height and width code")))
}

func typeAndGroupTxt(offset int, typeAndGroupIn typeAndGroupIn, sepSizeType string) posTxt {
	if !typeAndGroupIn.isValidFmt() {
		return newPosHint(offset+len(indent+sepSizeType)+2,
			fmt.Sprintf("%s must be a %s", underline("type code"), bold("valid type")))
	}
	if typeAndGroupIn.Found {
		return newPosInfo(offset+len(indent+sepSizeType)+2,
			fmt.Sprintf("%s ", underline("group"))+
				typeAndGroupIn.MappedTypeAndGroup.MappedGroup.Code+": "+
				typeAndGroupIn.MappedTypeAndGroup.MappedGroup.GroupInfo,
			fmt.Sprintf("%s ", underline("type"))+
				typeAndGroupIn.MappedTypeAndGroup.MappedType.Code+": "+
				typeAndGroupIn.MappedTypeAndGroup.MappedType.TypeInfo)
	}
	return newPosInfo(offset+len(indent+sepSizeType)+2,
		fmt.Sprintf("%s not found", underline("type code")))
}
