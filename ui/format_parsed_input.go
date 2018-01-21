package ui

import (
	"fmt"

	"iso6346/model"
	"github.com/fatih/color"
)

var grey = color.New(color.FgBlack).SprintFunc()

func FormatParsedInput(pi model.ParsedInput) string {
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
