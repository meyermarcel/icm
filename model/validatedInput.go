package model

import (
	"strconv"
)

type ValidatedInput struct {
	OwnerCode            OwnerCode
	EquipmentCategoryId  EquipmentCategoryId
	SerialNumber         SerialNumber
	CheckDigit           CheckDigit
	IsValidCheckDigit    bool
	CalculatedCheckDigit int
}

func NewValidatedInput(ownerCode OwnerCode, equipmentCategoryId EquipmentCategoryId, serialNumber SerialNumber,
	checkDigit CheckDigit) ValidatedInput {
	vi := ValidatedInput{OwnerCode: ownerCode, EquipmentCategoryId: equipmentCategoryId, SerialNumber: serialNumber, CheckDigit: checkDigit}
	if !vi.IsCheckDigitCalculable() {
		return vi
	}

	var cn ContainerNumber

	if !vi.CheckDigit.IsComplete() {
		cn = NewUncheckedContainerNumber(vi.OwnerCode.getConcatenatedValue(),
			vi.EquipmentCategoryId.getConcatenatedValue(),
			vi.SerialNumber.getConcatenatedValue())
		vi.CalculatedCheckDigit = cn.calculatedCheckDigit
		return vi
	}

	cd, _ := strconv.Atoi(vi.CheckDigit.getConcatenatedValue())
	cn = NewContainerNumber(vi.OwnerCode.getConcatenatedValue(),
		vi.EquipmentCategoryId.getConcatenatedValue(),
		vi.SerialNumber.getConcatenatedValue(), cd)
	vi.CalculatedCheckDigit = cn.calculatedCheckDigit
	vi.IsValidCheckDigit = cn.hasValidCheckDigit()
	return vi

}

func (vi ValidatedInput) IsCheckDigitCalculable() bool {
	return vi.OwnerCode.IsComplete() && vi.EquipmentCategoryId.IsComplete() && vi.SerialNumber.IsComplete()
}

type ContainerNumberValue interface {
	IsComplete() bool
	GetValues() []string
	getConcatenatedValue() string
}

type OwnerCode struct {
	values [3]string
}

func (oc OwnerCode) getConcatenatedValue() string {
	return concat(oc.values[:])
}

func (oc OwnerCode) GetValues() []string {
	return oc.values[:]
}

func (oc OwnerCode) IsComplete() bool {
	return isComplete(oc.values[:])
}

type EquipmentCategoryId struct {
	value string
}

func (eci EquipmentCategoryId) getConcatenatedValue() string {
	return eci.value
}

func (eci EquipmentCategoryId) GetValues() []string {
	return []string{eci.value}
}

func (eci EquipmentCategoryId) IsComplete() bool {
	return eci.value != ""
}

type SerialNumber struct {
	values [6]string
}

func (sn SerialNumber) getConcatenatedValue() string {
	return concat(sn.values[:])
}

func (sn SerialNumber) GetValues() []string {
	return sn.values[:]
}

func (sn SerialNumber) IsComplete() bool {
	return isComplete(sn.values[:])
}

type CheckDigit struct {
	value string
}

func (cd CheckDigit) getConcatenatedValue() string {
	return cd.value
}

func (cd CheckDigit) GetValues() []string {
	return []string{cd.value}
}

func (cd CheckDigit) IsComplete() bool {
	return cd.value != ""
}

func isComplete(values []string) bool {
	for _, element := range values {
		if element == "" {
			return false
		}
	}
	return true
}

func concat(values []string) string {
	value := ""
	for _, element := range values {
		value += element
	}
	return value
}
