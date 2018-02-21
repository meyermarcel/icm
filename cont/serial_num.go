package cont

import (
	"unicode/utf8"
	"regexp"
	"log"
)

type SerialNum struct {
	value string
}

func (sn SerialNum) Value() string {
	return sn.value
}

func NewSerialNum(value string) SerialNum {

	if utf8.RuneCountInString(value) != 6 {
		log.Fatalf("'%s' is not six characters", value)
	}

	if !regexp.MustCompile(`\d{6}`).MatchString(value) {
		log.Fatalf("'%s' must be 6 numbers", value)
	}
	return SerialNum{value}
}
