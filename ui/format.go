package ui

import (
	"fmt"
	"github.com/fatih/color"
	"iso6346/container"
	"iso6346/validator"
	"iso6346/owner"
	"iso6346/equip_cat"
	"iso6346/parser"
)

var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var grey = color.New(color.FgBlack).SprintFunc()

var bold = color.New(color.Bold).SprintFunc()
var underline = color.New(color.Underline).SprintFunc()

const missingCharacter = "_"

func formatParsedInput(pi parser.Input) string {

	out := "'"

	inputAsRunes := []rune(pi.Input())
	for pos, char := range inputAsRunes {
		if pi.Match(pos) {
			out += fmt.Sprintf("%s", string(char))
		} else {
			out += fmt.Sprintf("%s", grey(string(char)))
		}
	}
	out += "'"
	return out
}

func formatOwnerCode(oc owner.CodeOptEquipCatId) string {

	out := " "

	out += formatInput(oc.OwnerCode())

	if oc.EquipCatId().IsComplete() {
		out += " "
		out += formatInput(oc.EquipCatId())
	}

	out += formatCheckMark(oc.OwnerCode().IsComplete())

	out += fmt.Sprintln()

	return out
}

func formatValidatedInput(vi container.NumberValidated) string {

	out := ""

	out += formatContainerNumber(vi)

	validCheckDigit := vi.IsValidCheckDigit

	out += formatCheckMark(validCheckDigit)

	out += fmt.Sprintln()

	if validCheckDigit {
		return out
	}

	var errorMessages []PositionedMessage

	if !vi.OwnerCode.IsComplete() {
		errorMessages = append(errorMessages, PositionedMessage{2, fmt.Sprintf("%s must be %s", underline("owner code"), bold("3 letters"))})
	}
	if !vi.EquipCatId.IsComplete() {
		errorMessages = append(errorMessages, PositionedMessage{5, fmt.Sprintf("%s must be %s", underline("equipment category id"), equipCatIdsAsList())})
	}
	if !vi.SerialNumber.IsComplete() {
		errorMessages = append(errorMessages, PositionedMessage{9, fmt.Sprintf("%s must be %s", underline("serial number"), bold("6 numbers"))})
	}

	if !validCheckDigit {
		if vi.IsCheckDigitCalculable() {
			if vi.CheckDigit.IsComplete() {
				errorMessages = append(errorMessages, PositionedMessage{14, fmt.Sprintf("%s is incorrect (correct: %s)", underline("check digit"),
					green(vi.CalculatedCheckDigit))})
			} else {
				errorMessages = append(errorMessages, PositionedMessage{14, fmt.Sprintf("%s must be a %s (correct: %s)", underline("check digit"), bold("number"),
					green(vi.CalculatedCheckDigit))})
			}
		} else {
			errorMessages = append(errorMessages, PositionedMessage{14, fmt.Sprintf("%s is not calculable", underline("check digit"))})
		}
	}
	out += formatMessagesWithArrows(errorMessages)

	return out
}

func formatContainerNumber(vi container.NumberValidated) string {

	out := " "

	out += formatInput(vi.OwnerCode)
	out += " "
	out += formatInput(vi.EquipCatId)
	out += " "
	out += formatInput(vi.SerialNumber)
	out += " "

	if !vi.IsValidCheckDigit && vi.CheckDigit.IsComplete() {
		out += fmt.Sprintf("%s", yellow(string(vi.CheckDigit.Value())))
		return out
	}

	out += formatInput(vi.CheckDigit)
	return out
}

func formatInput(input validator.Input) string {

	if input.IsComplete() {
		return fmt.Sprintf("%s", green(input.Value()))
	}

	out := ""

	startIndexMissingCharacters := 0
	for pos, element := range input.Value() {
		out += fmt.Sprintf("%s", yellow(string(element)))
		startIndexMissingCharacters = pos + 1
	}

	for i := startIndexMissingCharacters; i < input.ValidLength(); i++ {
		out += fmt.Sprintf("%s", red(missingCharacter))
	}

	return out
}

func formatCheckMark(valid bool) string {

	out := "  "

	if valid {
		return out + fmt.Sprintf("%s", green("✔"))
	}
	return out + fmt.Sprintf("%s", red("✘"))

}

func equipCatIdsAsList() string {
	ujz := equip_cat.Ids
	return fmt.Sprintf("%s, %s or %s", green(string(ujz[0])), green(string(ujz[1])), green(string(ujz[2])))
}
