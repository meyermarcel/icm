package container

import (
	"fmt"
	"unicode/utf8"
	"strconv"
	"iso6346/equip_cat"
	"strings"
	"iso6346/owner"
)

type CheckDigit struct {
	value int
}

func (cd CheckDigit) Value() int {
	return cd.value
}

func NewCheckDigit(value string) (CheckDigit, error) {

	if utf8.RuneCountInString(value) != 1 {
		return CheckDigit{}, fmt.Errorf("'%s' is not one character", value)
	}

	cd, err := strconv.Atoi(value)

	if err != nil {
		return CheckDigit{}, err
	}

	return CheckDigit{cd}, nil
}

/*
This method is a modified version of a Go code sample from
https://en.wikipedia.org/wiki/ISO_6346#Code_Sample_(Go)
*/
func CalculateCheckDigit(ownerCode owner.Code, equipCatId equip_cat.Id, serialNumber SerialNumber) int {

	concatenated := ownerCode.Value() + equipCatId.Value() + serialNumber.Value()

	n := 0.0
	d := 0.5
	for _, character := range concatenated {
		d *= 2
		n += d * float64(strings.IndexRune("0123456789A?BCDEFGHIJK?LMNOPQRSTU?VWXYZ", character))
	}
	return (int(n) - int(n/11)*11) % 10
}
