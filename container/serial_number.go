package container

import (
	"unicode/utf8"
	"fmt"
	"regexp"
)

type SerialNumber struct {
	value string
}

func (sn SerialNumber) Value() string {
	return sn.value
}

func NewSerialNumber(value string) (SerialNumber, error) {

	if utf8.RuneCountInString(value) != 6 {
		return SerialNumber{}, fmt.Errorf("'%s' is not six characters", value)
	}

	if !regexp.MustCompile(`\d{6}`).MatchString(value) {
		return SerialNumber{}, fmt.Errorf("'%s' must be 6 numbers", value)
	}
	return SerialNumber{value}, nil
}
