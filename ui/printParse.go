package ui

import (
	"fmt"
	"iso6346/model"
)

func PrintParse(pi model.ParsedInput) {

	out := ""
	out += FormatParsedInput(pi)
	out += "\n"
	out += "\n"
	out += FormatValidatedInput(pi.ToValidatedInput())
	out += "\n"
	fmt.Printf(out)
}
