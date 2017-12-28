package ui

import (
	"iso6346/model"
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var grey = color.New(color.FgBlack).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()
var underline = color.New(color.Underline).SprintFunc()

var missingCharacter = red("_")

func PrintParsedInput(pi model.ParsedInput) {

	inputVisibleSpaces := []rune(strings.Replace(pi.Input, " ", "·", -1))

	fmt.Printf(indentation)

	for pos, char := range inputVisibleSpaces {
		if pi.MatchesIndices[pos] {
			fmt.Printf("%s", string(char))
		} else {
			fmt.Printf("%s", grey(string(char)))
		}
	}

	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Printf(indentation)

	printContainerNumberValue(pi.ValidatedInput.OwnerCode)
	fmt.Printf(" ")
	printContainerNumberValue(pi.ValidatedInput.EquipmentCategoryId)
	fmt.Printf(" ")
	printContainerNumberValue(pi.ValidatedInput.SerialNumber)
	fmt.Printf(" ")

	if !pi.ValidatedInput.IsValidCheckDigit && pi.ValidatedInput.CheckDigit.IsComplete() {
		fmt.Printf("%s", red(pi.ValidatedInput.CheckDigit.GetValues()[0]))
	} else {
		printContainerNumberValue(pi.ValidatedInput.CheckDigit)
	}

	errorMessages := [4]string{}

	if !pi.ValidatedInput.OwnerCode.IsComplete() {
		errorMessages[0] = fmt.Sprintf("%s must be %s", underline("owner code"), bold("3 letters"))
	}
	if !pi.ValidatedInput.EquipmentCategoryId.IsComplete() {
		errorMessages[1] = fmt.Sprintf("%s must be %s", underline("equipment category id"), equipmentIdsAsList())
	}
	if !pi.ValidatedInput.SerialNumber.IsComplete() {
		errorMessages[2] = fmt.Sprintf("%s must be %s", underline("serial number"), bold("6 numbers"))
	}

	if !pi.ValidatedInput.IsValidCheckDigit {

		if pi.ValidatedInput.IsCheckDigitCalculable() {
			if pi.ValidatedInput.CheckDigit.IsComplete() {
				errorMessages[3] = fmt.Sprintf("%s is incorrect (correct: %s)", underline("check digit"),
					green(pi.ValidatedInput.CalculatedCheckDigit))
			} else {
				errorMessages[3] = fmt.Sprintf("%s must be a %s (correct: %s)", underline("check digit"), bold("number"),
					green(pi.ValidatedInput.CalculatedCheckDigit))
			}
		} else {
			errorMessages[3] = fmt.Sprintf("%s is not calculable", underline("check digit"))
		}
	}

	ems := cutEmptyErrorsFromEnd(errorMessages[:])

	if len(ems) == 0 {
		fmt.Printf("  %s", green("✔"))
		fmt.Printf("\n")
		fmt.Printf("\n")
	} else {
		fmt.Printf("  %s", red("✘"))
		fmt.Printf("\n")
		printErrorMessages(ems)
	}

}

func printContainerNumberValue(cnv model.ContainerNumberValue) {
	for _, element := range cnv.GetValues() {
		if element == "" {
			fmt.Printf("%s", missingCharacter)
		} else {
			fmt.Printf("%s", green(strings.ToUpper(element)))
		}
	}
}

func cutEmptyErrorsFromEnd(ems []string) []string {

	length := len(ems)

	for i := length - 1; i >= 0; i-- {
		if ems[i] != "" {
			break
		}
		length = i
	}
	return ems[:length]
}

func printErrorMessages(ems []string) {

	printArrows(ems)

	for len(ems) != 0 {
		printEdgeAndErrorMessage(ems)
		ems = ems[:len(ems)-1]
		ems = cutEmptyErrorsFromEnd(ems)
		printVerticalLine(ems)
	}
}

func printArrows(ems []string) {
	fmt.Printf(indentation)

	for pos, errorMessage := range ems {
		printSpaceBetweenArrows(pos)
		if errorMessage != "" {
			fmt.Printf("↑")
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}

func printEdgeAndErrorMessage(ems []string) {
	fmt.Printf(indentation)

	for pos, errorMessage := range ems {
		printSpaceBetweenArrows(pos)
		if pos == len(ems)-1 {
			fmt.Printf("└─ ")
			fmt.Printf(errorMessage)
		} else if errorMessage != "" {
			fmt.Printf("│")
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}

func printVerticalLine(ems []string) {
	fmt.Printf(indentation)

	for pos, errorMessage := range ems {
		printSpaceBetweenArrows(pos)
		if errorMessage != "" {
			fmt.Printf("│")
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}

/*

  ↑  ↑   ↑    ↑
  │  │   │    └─ cannot calculate check digit

^^ ^  ^   ^
││ │  │   4 spaces
││ │  │
││ │  3 spaces
││ │
││ 2 spaces
││
│1 space
│
indentation
 */

const indentation = " "
const fourSpaces = "    "

func printSpaceBetweenArrows(i int) {
	fmt.Printf(fourSpaces[:i+1])
}

func equipmentIdsAsList() string {
	ujz := model.EquipmentCategoryIds
	return fmt.Sprintf("%s, %s or %s", green(string(ujz[0])), green(string(ujz[1])), green(string(ujz[2])))
}
