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

package input

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

// FancyPrinter prints inputs in a fancy manner. Use NewFancyPrinterFactory to instantiate one.
type FancyPrinter struct {
	writer         io.Writer
	indent         string
	separators     []string
	separatorsFunc func(inputs []Input)
}

// NewFancyPrinter creates a FancyPrinter.
func NewFancyPrinter(writer io.Writer) *FancyPrinter {
	return &FancyPrinter{
		writer: writer,
	}
}

// SetIndent sets the indentation for printing.
func (fp *FancyPrinter) SetIndent(indent string) *FancyPrinter {
	fp.indent = indent
	return fp
}

// SetSeparators sets the separators between inputs. Default separator is ' '.
func (fp *FancyPrinter) SetSeparators(separators ...string) {
	fp.separators = separators
}

// SetSeparatorsFunc sets a function that can set the separators depending on inputs.
func (fp *FancyPrinter) SetSeparatorsFunc(separatorsFunc func(inputs []Input)) {
	fp.separatorsFunc = separatorsFunc
}

// Print writes formatted inputs to writer.
func (fp *FancyPrinter) Print(inputs []Input) error {

	if fp.separatorsFunc != nil {
		fp.separatorsFunc(inputs)
	}

	b := strings.Builder{}
	b.WriteString(fmt.Sprintln())

	b.WriteString(fp.indent)
	pos := len(fp.indent)

	texts := make([]posTxt, 0)

	for idx, input := range inputs {

		b.WriteString(fmtValue(input))

		sep := " "
		switch {
		case idx == len(inputs)-1:
			sep = ""
		case idx > len(inputs)-1:
			sep = " "
		case idx < len(fp.separators):
			sep = fp.separators[idx]
		}
		b.WriteString(sep)

		if input.err != nil || input.infos != nil {
			posTxt := posTxt{
				pos: pos + input.runeCount/2,
			}
			if input.err != nil {
				posTxt.addLines(input.err.Error())
			}
			for _, info := range input.infos {
				posTxt.addLines(info.Text)
			}
			texts = append(texts, posTxt)
		}
		pos += input.runeCount + utf8.RuneCountInString(sep)
	}
	b.WriteString(fmtCheckMark(inputs[len(inputs)-1].err == nil))
	b.WriteString(fmt.Sprintln())
	b.WriteString(fmtTextsWithArrows(texts...))
	b.WriteString(fmt.Sprintln())

	_, _ = io.WriteString(fp.writer, b.String())
	return nil
}

func fmtValue(input Input) string {
	if input.err == nil {
		return green(input.value)
	}
	if input.isValidFmt() {
		return red(input.value)
	}
	return strings.Repeat(red("_"), input.runeCount)
}

type posTxt struct {
	pos   int
	lines []string
}

func (pt *posTxt) addLines(lines ...string) {
	pt.lines = append(pt.lines, lines...)
}

func fmtTextsWithArrows(texts ...posTxt) string {
	sort.Slice(texts, func(i, j int) bool {
		return texts[i].pos < texts[j].pos
	})
	for idx, element := range texts {
		if idx > 0 && texts[idx-1].pos == element.pos {
			texts[idx-1].addLines(element.lines...)
			texts = append(texts[:idx], texts[idx+1:]...)
		}
	}
	return fmtTexts(texts)
}

func fmtTexts(texts []posTxt) string {

	b := strings.Builder{}

	if len(texts) == 0 {
		return b.String()
	}

	var positions []int
	for _, element := range texts {
		positions = append(positions, element.pos)
	}
	spaces := calculateSpaces(positions)

	for idx := range texts {
		b.WriteString(spaces[idx])
		b.WriteString("↑")
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
	b.WriteString(fmt.Sprintln())
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

func fmtCheckMark(valid bool) string {

	b := strings.Builder{}
	b.WriteString("  ")

	if !valid {
		b.WriteString(fmt.Sprintf("%s", red("✘")))
		return b.String()
	}
	b.WriteString(fmt.Sprintf("%s", green("✔")))
	return b.String()
}
