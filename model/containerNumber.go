package model

import (
	"fmt"
	"strings"
)

const EquipmentCategoryIds = "UJZ"

type ContainerNumber struct {
	ownerCode            string
	equipmentCategoryId  string
	serialNumber         string
	checkDigit           int
	calculatedCheckDigit int
}

func NewUncheckedContainerNumber(ownerCode string,
	equipmentCategoryIdentifier string,
	serialNumber string) ContainerNumber {

	calculatedCheckDigit := calculateCheckDigit(ownerCode + equipmentCategoryIdentifier + serialNumber)
	return ContainerNumber{ownerCode,
		equipmentCategoryIdentifier,
		serialNumber,
		calculatedCheckDigit,
		calculatedCheckDigit}

}

func NewContainerNumber(ownerCode string,
	equipmentCategoryIdentifier string,
	serialNumber string, checkDigit int) ContainerNumber {

	calculatedCheckDigit := calculateCheckDigit(ownerCode + equipmentCategoryIdentifier + serialNumber)

	return ContainerNumber{ownerCode,
		equipmentCategoryIdentifier,
		serialNumber,
		checkDigit,
		calculatedCheckDigit}
}

func (cn ContainerNumber) OwnerCode() string {
	return cn.ownerCode
}

func (cn ContainerNumber) EquipmentCategoryId() string {
	return cn.equipmentCategoryId
}

func (cn ContainerNumber) SerialNumber() string {
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
func calculateCheckDigit(cn string) int {

	n := 0.0
	d := 0.5
	for i := 0; i < 10; i++ {
		d *= 2
		n += d * float64(strings.Index("0123456789A?BCDEFGHIJK?LMNOPQRSTU?VWXYZ", string(cn[i])))
	}
	return (int(n) - int(n/11)*11) % 10
}

func (cn ContainerNumber) hasValidCheckDigit() bool {
	return cn.checkDigit == cn.calculatedCheckDigit
}

func (cn ContainerNumber) String() string {
	return fmt.Sprintf("%s%s%s%d", cn.ownerCode, cn.equipmentCategoryId, cn.serialNumber, cn.checkDigit)
}
