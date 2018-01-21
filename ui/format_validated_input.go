package ui

import (
	"fmt"

	"iso6346/model"
	"github.com/fatih/color"
)

var bold = color.New(color.Bold).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var underline = color.New(color.Underline).SprintFunc()

func FormatValidatedInput(vi model.ValidatedInput) string {

	out := ""

	out += FormatContainerNumber(vi)

	out += "  "

	if vi.IsValidCheckDigit {
		out += fmt.Sprintf("%s", green("✔"))
		out += fmt.Sprintln()
		return out
	}

	out += fmt.Sprintf("%s", red("✘"))
	out += fmt.Sprintln()

	var errorMessages []PositionedMessage

	if !vi.OwnerCode.IsComplete {
		errorMessages = append(errorMessages, PositionedMessage{2, fmt.Sprintf("%s must be %s", underline("owner code"), bold("3 letters"))})
	}
	if !vi.EquipmentCategoryId.IsComplete {
		errorMessages = append(errorMessages, PositionedMessage{5, fmt.Sprintf("%s must be %s", underline("equipment category id"), equipmentIdsAsList())})
	}
	if !vi.SerialNumber.IsComplete {
		errorMessages = append(errorMessages, PositionedMessage{9, fmt.Sprintf("%s must be %s", underline("serial number"), bold("6 numbers"))})
	}

	if !vi.IsValidCheckDigit {
		if vi.IsCheckDigitCalculable() {
			if vi.CheckDigit.IsComplete {
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

func equipmentIdsAsList() string {
	ujz := model.EquipmentCategoryIds
	return fmt.Sprintf("%s, %s or %s", green(string(ujz[0])), green(string(ujz[1])), green(string(ujz[2])))
}
