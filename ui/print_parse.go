package ui

import (
	"fmt"
	"iso6346/iso6346"
)

func PrintParse(pi iso6346.ParsedInput) {

	out := ""
	out += FormatParsedInput(pi)
	out += fmt.Sprintln()
	out += fmt.Sprintln()
	out += FormatValidatedInput(pi.ToValidatedInput())
	out += fmt.Sprintln()
	fmt.Printf(out)
}
