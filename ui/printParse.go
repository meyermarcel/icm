package ui

import (
	"fmt"
	"iso6346/model"
)

func PrintParse(pi model.ParsedInput) {

	out := ""
	out += FormatParsedInput(pi)
	out += fmt.Sprintln()
	out += fmt.Sprintln()
	out += FormatValidatedInput(pi.ToValidatedInput())
	out += fmt.Sprintln()
	fmt.Printf(out)
}
