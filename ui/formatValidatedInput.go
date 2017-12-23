package ui

import (
	"iso6346/model"
	"fmt"
	"os"
	"github.com/fatih/color"
)

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()

var indentation = " "
var fiveSpaces = "     "

func Print(iv model.ValidatedInput) {

	fmt.Printf("\n")

	if !iv.InputPossibleValidLength {

		fmt.Printf(" `%s`\n", red(iv.Input))
		fmt.Printf("\n")
		fmt.Printf("%s %s is too long (max. %s).\n", red(len(iv.Input)), bold("characters"), green(model.MaxInputLength))
		os.Exit(1)
	}

	if !iv.InputLengthIs11 {

		fmt.Printf(" `%s`\n", red(iv.Input))
		fmt.Printf("\n")
		fmt.Printf("%s %s are not valid (%s).\n", red(len(iv.Input)), bold("alphanumeric [A-Za-z0-9] characters"), green(model.MaxAlphaNumericCharacters))
		os.Exit(1)
	}

	if iv.OwnerCodeIsLetter {
		fmt.Printf(" %s", green(iv.OwnerCode))
	} else {
		fmt.Printf(" %s", red(iv.OwnerCode))
	}

	if iv.EquipmentCategoryIdIsValid {
		fmt.Printf(" %s", green(iv.EquipmentCategoryId))
	} else {
		fmt.Printf(" %s", red(iv.EquipmentCategoryId))
	}

	if iv.SerialNumberIsNumber {
		fmt.Printf(" %s", green(iv.SerialNumber))
	} else {
		fmt.Printf(" %s", red(iv.SerialNumber))
	}

	if iv.CheckDigitIsValid {
		fmt.Printf(" %s", green(iv.CheckDigit))
	} else {
		fmt.Printf(" %s", red(iv.CheckDigit))
	}

	var errors []string

	if iv.OwnerCodeIsLetter {
		errors = append(errors, "")
	} else {
		errors = append(errors, fmt.Sprintf("owner code must be %s", bold("letters")))
	}

	if iv.EquipmentCategoryIdIsValid {
		errors = append(errors, "")
	} else {
		errors = append(errors,
			fmt.Sprintf("must be %s", printEquipmentIds(model.EquipmentCategoryIds)))
	}

	if iv.SerialNumberIsNumber {
		errors = append(errors, "")
	} else {
		errors = append(errors, fmt.Sprintf("serial number must be a %s", bold("number")))
	}

	if iv.CheckDigitIsValid {
		errors = append(errors, "")
	} else {
		{
			if iv.IsValidAlphanumeric() {
				errors = append(errors, fmt.Sprintf("incorrect %s (correct: %s)", bold("check digit"), green(iv.ValidCheckDigit)))
			} else {
				errors = append(errors, fmt.Sprintf("cannot calculate %s", bold("check digit")))
			}
		}
	}

	fmt.Printf("\n")
	printErrors(errors)
	fmt.Printf("\n")

}

func printErrors(errorMessages []string) {

	errorMessages = filterEmptyErrorsFromEnd(errorMessages)
	if len(errorMessages) == 0 {
		return
	}

	fmt.Printf(indentation + " ")
	for pos, errorMessage := range errorMessages {
		if errorMessage != "" {
			fmt.Printf("^")
			fmt.Printf(fiveSpaces[:pos+2])
		} else {
			fmt.Printf(" ")
			fmt.Printf(fiveSpaces[:pos+2])
		}
	}

	for len(errorMessages) != 0 {

		errorMessages = filterEmptyErrorsFromEnd(errorMessages)
		if len(errorMessages) != 0 {
			printArrows(errorMessages)
			errorMessages = errorMessages[:len(errorMessages)-1]
		}
	}
	fmt.Printf("\n")
}

func filterEmptyErrorsFromEnd(errorMessages []string) []string {

	index := len(errorMessages)

	for i := index - 1; i >= 0; i-- {
		if errorMessages[i] != "" {
			break
		}
		index = i
	}

	return errorMessages[:index]
}

func printArrows(errorMessages []string) {

	fmt.Printf("\n")

	fmt.Printf(indentation + " ")
	for pos, errorMessage := range errorMessages {
		if errorMessage != "" {
			fmt.Printf("│")
			fmt.Printf(fiveSpaces[:pos+2])
		} else {
			fmt.Printf(" ")
			fmt.Printf(fiveSpaces[:pos+2])
		}
	}
	fmt.Printf("\n")

	fmt.Printf(indentation + " ")
	for i := 0; i < len(errorMessages); i++ {

		if i == len(errorMessages)-1 {
			fmt.Printf("└─ ")
			fmt.Printf(errorMessages[i])
		} else if errorMessages[i] != "" {
			fmt.Printf("│")
			fmt.Printf(fiveSpaces[:i+2])
		} else {
			fmt.Printf(" ")
			fmt.Printf(fiveSpaces[:i+2])
		}
	}

}

func printEquipmentIds(s string) string {
	formattedList := ""
	for pos, char := range s {
		formattedList += fmt.Sprintf("%s", green(string(char)))
		if pos < len(s)-2 {
			formattedList += fmt.Sprintf(", ")
		}

		if pos == len(s)-2 {
			formattedList += fmt.Sprintf(" or ")
		}
	}
	return formattedList
}
