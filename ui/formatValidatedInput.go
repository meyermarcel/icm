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

	if iv.EquipmentCategoryIdIsLetter && iv.EquipmentCategoryIdIsValid {
		fmt.Printf(" %s", green(iv.EquipmentCategoryId))
	} else {
		fmt.Printf(" %s", red(iv.EquipmentCategoryId))
	}

	if iv.SerialNumberIsNumber {
		fmt.Printf(" %s", green(iv.SerialNumber))
	} else {
		fmt.Printf(" %s", red(iv.SerialNumber))
	}

	if !iv.CheckDigitIsNumber {
		fmt.Printf(" %s", red(iv.CheckDigit))
	} else {
		if iv.CheckDigitIsValid {
			fmt.Printf(" %s", green(iv.CheckDigit))
		} else {
			fmt.Printf(" %s", red(iv.CheckDigit))
		}
	}

	var errors []func()

	if !iv.OwnerCodeIsLetter {
		errors = append(errors, func() { fmt.Printf("owner code must be %s", bold("letters")) })
	} else {
		errors = append(errors, nil)
	}

	if !iv.EquipmentCategoryIdIsLetter {
		errors = append(errors, func() {
			fmt.Printf("must be a %s (", bold("letter"))
			printEquipmentIds(model.EquipmentCategoryIds)
			fmt.Printf(")")

		})
	} else {
		if !iv.EquipmentCategoryIdIsValid {
			errors = append(errors, func() {
				fmt.Printf("is a letter but incorrect %s (", bold("equipment id"))
				printEquipmentIds(model.EquipmentCategoryIds)
				fmt.Printf(")")
			})
		} else {
			errors = append(errors, nil)
		}
	}

	if !iv.SerialNumberIsNumber {
		errors = append(errors, func() { fmt.Printf("serial number must be a %s", bold("number")) })
	} else {
		errors = append(errors, nil)
	}

	if !iv.CheckDigitIsNumber {
		errors = append(errors, func() { fmt.Printf("check digit is not a %s", bold("number")) })
	} else {
		if !iv.CheckDigitIsValid {
			{
				if iv.IsValidAlphanumeric() {
					errors = append(errors, func() { fmt.Printf("incorrect %s (correct: %s)", bold("check digit"), green(iv.ValidCheckDigit)) })
				} else {
					errors = append(errors, func() { fmt.Printf("incorrect %s", bold("check digit"), ) })
				}
			}
		} else {
			errors = append(errors, nil)
		}
	}

	fmt.Printf("\n")
	printErrors(errors)
	fmt.Printf("\n")

}

func printErrors(errorMessages []func()) {

	errorMessages = filterNilFromEnd(errorMessages)
	if len(errorMessages) == 0 {
		return
	}

	fmt.Printf(indentation + " ")
	for pos, errorMessage := range errorMessages {
		if errorMessage != nil {
			fmt.Printf("^")
			fmt.Printf(fiveSpaces[:pos+2])
		} else {
			fmt.Printf(" ")
			fmt.Printf(fiveSpaces[:pos+2])
		}
	}

	for len(errorMessages) != 0 {

		errorMessages = filterNilFromEnd(errorMessages)
		if len(errorMessages) != 0 {
			printArrows(errorMessages)
			errorMessages = errorMessages[:len(errorMessages)-1]
		}
	}
	fmt.Printf("\n")
}

func filterNilFromEnd(errorMessages []func()) []func() {

	index := len(errorMessages)

	for i := index - 1; i >= 0; i-- {
		if errorMessages[i] != nil {
			break
		}
		index = i
	}

	return errorMessages[:index]
}

func printArrows(errorMessages []func()) {

	fmt.Printf("\n")

	fmt.Printf(indentation + " ")
	for pos, errorMessage := range errorMessages {
		if errorMessage != nil {
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
			errorMessages[i]()
		} else if errorMessages[i] != nil {
			fmt.Printf("│")
			fmt.Printf(fiveSpaces[:i+2])
		} else {
			fmt.Printf(" ")
			fmt.Printf(fiveSpaces[:i+2])
		}
	}

}

func printEquipmentIds(s string) {
	for pos, char := range s {
		fmt.Printf("%s", green(string(char)))
		if pos < len(s)-2 {
			fmt.Printf(", ")
		}

		if pos == len(s)-2 {
			fmt.Printf(" or ")
		}
	}
}
