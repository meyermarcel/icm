package ui

import (
	"fmt"

	"iso6346/iso6346"
	"github.com/fatih/color"
)

var grey = color.New(color.FgBlack).SprintFunc()

func FormatParsedInput(pi iso6346.ParsedInput) string {
	out := "'"

	inputAsRunes := []rune(pi.Input)
	for pos, char := range inputAsRunes {
		if pi.MatchesIndices[pos] {
			out += fmt.Sprintf("%s", string(char))
		} else {
			out += fmt.Sprintf("%s", grey(string(char)))
		}
	}
	out += "'"
	return out
}
