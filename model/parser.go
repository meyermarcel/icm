package model

import (
	"regexp"
	"strings"
	"strconv"
)

type InputValidation struct {
	Input                       string
	InputPossibleValidLength    bool
	InputLengthIs11             bool
	OwnerCode                   string
	OwnerCodeIsLetter           bool
	EquipmentCategoryId         string
	EquipmentCategoryIdIsLetter bool
	EquipmentCategoryIdIsValid  bool
	SerialNumber                string
	SerialNumberIsNumber        bool
	CheckDigit                  string
	CheckDigitIsNumber          bool
	CalculatedCheckDigit        int
	CheckDigitIsValid           bool
}

const MaxInputLength = 15
const MaxAlphaNumericCharacters = 11

func (iv InputValidation) IsValidAlphanumeric() bool {
	return iv.OwnerCodeIsLetter && iv.EquipmentCategoryIdIsValid && iv.SerialNumberIsNumber
}

var IsLetter = regexp.MustCompile("^[A-Z]+$").MatchString
var NoAlphaNumeric = regexp.MustCompile("[^A-Z0-9]+")

func Parse(input string) InputValidation {

	inputValidation := InputValidation{}
	inputValidation.Input = input


	if len(input) > MaxInputLength {
		return inputValidation
	}
	inputValidation.InputPossibleValidLength = true

	inputUpper := strings.ToUpper(input)

	filteredInput := NoAlphaNumeric.ReplaceAllString(inputUpper, "")
	inputValidation.Input = filteredInput

	if len(filteredInput) != MaxAlphaNumericCharacters {
		return inputValidation
	}
	inputValidation.InputLengthIs11 = true

	equipmentCategoryId := string(filteredInput[3])
	ownerCode := filteredInput[0:3]
	serialNumber := filteredInput[4:10]
	checkDigit := string(filteredInput[10])

	inputValidation.OwnerCode = ownerCode
	inputValidation.EquipmentCategoryId = equipmentCategoryId
	inputValidation.SerialNumber = serialNumber
	inputValidation.CheckDigit = checkDigit

	inputValidation.OwnerCodeIsLetter = IsLetter(ownerCode)

	inputValidation.EquipmentCategoryIdIsLetter = IsLetter(equipmentCategoryId)

	_, err := strconv.Atoi(serialNumber)
	inputValidation.SerialNumberIsNumber = err == nil

	checkDigitAsNumber, err := strconv.Atoi(checkDigit)

	inputValidation.CheckDigitIsNumber = err == nil

	cn := NewContainerNumber(ownerCode, equipmentCategoryId, serialNumber, checkDigitAsNumber)

	isValidEquipmentId := cn.hasValidEquipmentCategoryIdentifier()
	inputValidation.EquipmentCategoryIdIsValid = isValidEquipmentId

	if inputValidation.IsValidAlphanumeric() && isValidEquipmentId {
		inputValidation.CheckDigitIsValid = cn.hasValidCheckDigit()
		inputValidation.CalculatedCheckDigit = cn.CalculatedCheckDigit()
	}

	return inputValidation
}
