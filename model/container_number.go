package model

import (
	"fmt"
	"strings"
	"errors"
	"unicode/utf8"
	"regexp"
	"strconv"
)

var EquipmentCategoryIds = []rune("UJZ")

type OwnerCode struct {
	Value string
}

func NewOwnerCode(value string) (OwnerCode, error) {

	var err error

	if utf8.RuneCountInString(value) != 3 {
		err = errors.New(fmt.Sprintf("'%s' is not three characters", value))
		return OwnerCode{}, err
	}
	
	if !regexp.MustCompile(`[A-Z]{3}`).MatchString(value) {
		err = errors.New(fmt.Sprintf("'%s' must be 3 letters", value))
		return OwnerCode{}, err
	}
	return OwnerCode{value}, err
}

type EquipmentCategoryId struct {
	Value string
}

func NewEquipmentCategoryId(value string) (EquipmentCategoryId, error) {

	var err error
	if utf8.RuneCountInString(value) != 1 {
		err = errors.New(fmt.Sprintf("'%s' is not one character", value))
		return EquipmentCategoryId{}, err
	}
	if !regexp.MustCompile(`[UJZ]`).MatchString(value) {
		err = errors.New(fmt.Sprintf("'%s' must be U, J or Z", value))
		return EquipmentCategoryId{}, err
	}
	return EquipmentCategoryId{value}, err
}

type SerialNumber struct {
	Value string
}

func NewSerialNumber(value string) (SerialNumber, error) {

	var err error

	if utf8.RuneCountInString(value) != 6 {
		err = errors.New(fmt.Sprintf("'%s' is not six characters", value))
		return SerialNumber{}, err
	}

	if !regexp.MustCompile(`\d{6}`).MatchString(value) {
		err = errors.New(fmt.Sprintf("'%s' must be 6 numbers", value))
		return SerialNumber{}, err
	}
	return SerialNumber{value}, err
}

type CheckDigit struct {
	value int
}

func NewCheckDigit(value string) (CheckDigit, error) {

	var err error

	if utf8.RuneCountInString(value) != 1 {
		err = errors.New(fmt.Sprintf("'%s' is not one character", value))
		return CheckDigit{}, err
	}

	cd, err := strconv.Atoi(value)

	if err != nil {
		return CheckDigit{}, err
	}

	return CheckDigit{cd}, err
}

/*
This method is a modified version of a Go code sample from
https://en.wikipedia.org/wiki/ISO_6346#Code_Sample_(Go)
*/
func CalculateCheckDigit(ownerCode OwnerCode, equipmentCategoryId EquipmentCategoryId, serialNumber SerialNumber) int {

	concatenated := ownerCode.Value + equipmentCategoryId.Value + serialNumber.Value

	n := 0.0
	d := 0.5
	for _, character := range concatenated {
		d *= 2
		n += d * float64(strings.IndexRune("0123456789A?BCDEFGHIJK?LMNOPQRSTU?VWXYZ", character))
	}
	return (int(n) - int(n/11)*11) % 10
}
