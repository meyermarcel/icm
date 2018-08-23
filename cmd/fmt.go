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

	"github.com/fatih/color"

	"strings"
	"unicode/utf8"
)

var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var grey = color.New(color.FgBlack).SprintFunc()

var bold = color.New(color.Bold).SprintFunc()
var underline = color.New(color.Underline).SprintFunc()

const missingCharacter = "_"
const indent = " "
const indentSize = len(indent)

func fmtRegexIn(pi regexIn) string {

	b := strings.Builder{}
	b.WriteString("'")

	matchRangesIdx := 0

	input := pi.input

	for pos, w := 0, 0; pos < len(input); pos += w {

		if matchRangesIdx < len(pi.matchRanges) && pi.matchRanges[matchRangesIdx] == pos {
			matched := input[pos:pi.matchRanges[matchRangesIdx+1]]
			sumWidthSubStr := 0
			for posSubStr, wSubStr := 0, 0; posSubStr < len(matched); posSubStr += wSubStr {
				runeValue, width := utf8.DecodeRuneInString(matched[posSubStr:])
				b.WriteString(fmt.Sprintf("%s", string(runeValue)))
				wSubStr = width
				sumWidthSubStr += width
			}
			w = sumWidthSubStr
			matchRangesIdx += 2
		} else {
			runeValue, width := utf8.DecodeRuneInString(input[pos:])
			b.WriteString(fmt.Sprintf("%s", grey(string(runeValue))))
			w = width
		}
	}
	b.WriteString("'")

	return b.String()
}

func fmtIn(in input) string {

	if in.isValidFmt() {
		return fmt.Sprintf("%s", green(in.value))
	}

	b := strings.Builder{}

	startIndexMissingCharacters := 0
	for pos, element := range in.value {
		b.WriteString(fmt.Sprintf("%s", yellow(string(element))))
		startIndexMissingCharacters = pos + 1
	}

	for i := startIndexMissingCharacters; i < in.validLen; i++ {
		b.WriteString(fmt.Sprintf("%s", red(missingCharacter)))
	}

	return b.String()
}

func fmtCheckMark(valid bool) string {

	b := strings.Builder{}
	b.WriteString("  ")

	if valid {
		b.WriteString(fmt.Sprintf("%s", green("✔")))
		return b.String()
	}
	b.WriteString(fmt.Sprintf("%s", red("✘")))
	return b.String()
}
