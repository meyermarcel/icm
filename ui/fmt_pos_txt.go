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
	"strings"
)

type txtType int

const (
	hint txtType = iota
	info
)

type posTxt struct {
	pos     int
	txtType txtType
	lines   []string
}

func newPosHint(pos int, lines ...string) posTxt {
	return posTxt{pos, hint, lines}
}

func newPosInfo(pos int, lines ...string) posTxt {
	return posTxt{pos, info, lines}
}

func fmtTextsWithArrows(texts ...posTxt) string {

	b := strings.Builder{}

	if len(texts) == 0 {
		return b.String()
	}

	var positions []int
	for _, element := range texts {
		positions = append(positions, element.pos)
	}
	spaces := calculateSpaces(positions)

	for idx, txt := range texts {
		b.WriteString(spaces[idx])
		switch txt.txtType {
		case hint:
			b.WriteString("↑")
		case info:
			b.WriteString("│")
		}
	}

	for len(texts) != 0 {

		b.WriteString(fmt.Sprintln())
		for idx, txt := range texts {
			b.WriteString(spaces[idx])
			if idx == len(texts)-1 {
				for lineIdx, line := range txt.lines {
					if lineIdx == 0 {
						b.WriteString("└─ ")
						b.WriteString(line)
					}
					if lineIdx > 0 {
						b.WriteString(fmt.Sprintln())
						for i := 0; i < idx; i++ {
							b.WriteString(spaces[i])
							b.WriteString("│")
						}
						b.WriteString(spaces[idx])
						b.WriteString("   ")
						b.WriteString(line)
					}
				}
			} else {
				b.WriteString("│")
			}
		}

		texts = texts[:len(texts)-1]

		if len(texts) != 0 {
			b.WriteString(fmt.Sprintln())
			for idx := range texts {
				b.WriteString(spaces[idx])
				b.WriteString("│")
			}
		}

	}
	return b.String()
}

func calculateSpaces(positions []int) []string {

	var spaces []string
	lastPos := 0
	for idx, pos := range positions {
		spacesCount := pos - lastPos - 1
		spaces = append(spaces, "")
		for i := 0; i <= spacesCount; i++ {
			spaces[idx] += " "
		}
		lastPos = pos + 1
	}
	return spaces
}
