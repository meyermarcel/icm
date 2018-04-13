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

type Type int

const (
	HINT Type = iota
	INFO
)

type PosTxt struct {
	pos     int
	txtType Type
	values  []string
}

func NewPosHint(pos int, lines ...string) PosTxt {
	return PosTxt{pos, HINT, lines}
}

func NewPosInfo(pos int, lines ...string) PosTxt {
	return PosTxt{pos, INFO, lines}
}

func fmtTextsWithArrows(texts ...PosTxt) string {

	b := strings.Builder{}

	if len(texts) == 0 {
		return b.String()
	}

	spaces := calculateSpaces(texts)

	for pos, message := range texts {
		b.WriteString(spaces[pos])
		switch message.txtType {
		case HINT:
			b.WriteString("↑")
		case INFO:
			b.WriteString("│")
		}
	}

	for len(texts) != 0 {

		b.WriteString(fmt.Sprintln())
		for pos, txt := range texts {
			b.WriteString(spaces[pos])
			if pos == len(texts)-1 {
				for lineIx, line := range txt.values {
					if lineIx == 0 {
						b.WriteString("└─ ")
						b.WriteString(line)
					}
					if lineIx > 0 {
						b.WriteString(fmt.Sprintln())
						b.WriteString(spaces[pos])
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
		}
		for pos := range texts {
			b.WriteString(spaces[pos])
			b.WriteString("│")
		}
	}
	return b.String()
}

func calculateSpaces(texts []PosTxt) []string {

	var spaces []string
	lastPos := 0
	for pos, element := range texts {
		spacesCount := element.pos - lastPos - 1
		spaces = append(spaces, "")
		for i := 0; i <= spacesCount; i++ {
			spaces[pos] += " "
		}
		lastPos = element.pos + 1
	}
	return spaces
}
