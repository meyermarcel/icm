package ui

import (
	"fmt"
	"iso6346/iso6346"
	"unicode/utf8"
	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()

const missingCharacter = "_"

func FormatContainerNumber(vi iso6346.ValidatedInput) string {

	out := " "

	out += formatContainerNumberValue(vi.OwnerCode.Value, 3)
	out += " "
	out += formatContainerNumberValue(vi.EquipmentCategoryId.Value, 1)
	out += " "
	out += formatContainerNumberValue(vi.SerialNumber.Value, 6)
	out += " "

	if !vi.IsValidCheckDigit && vi.CheckDigit.IsComplete {
		out += fmt.Sprintf("%s", yellow(string(vi.CheckDigit.Value)))
		return out
	}

	out += formatContainerNumberValue(vi.CheckDigit.Value, 1)
	return out
}

func formatContainerNumberValue(value string, validLength int) string {

	if utf8.RuneCountInString(value) == validLength {
		return fmt.Sprintf("%s", green(value))
	}

	out := ""

	startIndexMissingCharacters := 0
	for pos, element := range value {
		out += fmt.Sprintf("%s", yellow(string(element)))
		startIndexMissingCharacters = pos + 1
	}

	for i := startIndexMissingCharacters; i < validLength; i++ {
		out += fmt.Sprintf("%s", red(missingCharacter))
	}

	return out
}
