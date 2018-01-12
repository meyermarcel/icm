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

type ContainerNumber struct {
	ownerCode            OwnerCode
	equipmentCategoryId  EquipmentCategoryId
	serialNumber         SerialNumber
	checkDigit           int
	calculatedCheckDigit int
}

type OwnerCode struct {
	value string
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
	value string
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
	value string
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

func NewUncheckedContainerNumber(ownerCode OwnerCode,
	equipmentCategoryId EquipmentCategoryId,
	serialNumber SerialNumber) ContainerNumber {

	calculatedCheckDigit := calculateCheckDigit(ownerCode, equipmentCategoryId, serialNumber)

	return ContainerNumber{ownerCode,
		equipmentCategoryId,
		serialNumber,
		calculatedCheckDigit,
		calculatedCheckDigit}

}

func NewContainerNumber(ownerCode OwnerCode,
	equipmentCategoryId EquipmentCategoryId,
	serialNumber SerialNumber, checkDigit CheckDigit) ContainerNumber {

	calculatedCheckDigit := calculateCheckDigit(ownerCode, equipmentCategoryId, serialNumber)

	return ContainerNumber{ownerCode,
		equipmentCategoryId,
		serialNumber,
		checkDigit.value,
		calculatedCheckDigit}
}

func (cn ContainerNumber) OwnerCode() OwnerCode {
	return cn.ownerCode
}

func (cn ContainerNumber) EquipmentCategoryId() EquipmentCategoryId {
	return cn.equipmentCategoryId
}

func (cn ContainerNumber) SerialNumber() SerialNumber {
	return cn.serialNumber
}

func (cn ContainerNumber) CheckDigit() int {
	return cn.checkDigit
}

func (cn ContainerNumber) ValidCheckDigit() int {
	return cn.calculatedCheckDigit
}

/*
This method is a modified version of an Go code sample from
https://en.wikipedia.org/wiki/ISO_6346#Code_Sample_(Go)
*/
func calculateCheckDigit(ownerCode OwnerCode, equipmentCategoryId EquipmentCategoryId, serialNumber SerialNumber) int {

	concatenated := ownerCode.value + equipmentCategoryId.value + serialNumber.value

	n := 0.0
	d := 0.5
	for _, character := range concatenated {
		d *= 2
		n += d * float64(strings.IndexRune("0123456789A?BCDEFGHIJK?LMNOPQRSTU?VWXYZ", character))
	}
	return (int(n) - int(n/11)*11) % 10
}

func (cn ContainerNumber) hasValidCheckDigit() bool {
	return cn.checkDigit == cn.calculatedCheckDigit
}

func (cn ContainerNumber) String() string {
	return fmt.Sprintf("%s%s%s%d", cn.ownerCode, cn.equipmentCategoryId, cn.serialNumber, cn.checkDigit)
}
