package equip_cat

import (
	"regexp"
	"unicode/utf8"
	"log"
)

var Ids = []rune("UJZ")

type Id struct {
	value string
}

func (id Id) Value() string {
	return id.value
}

func NewId(value string) Id {

	if utf8.RuneCountInString(value) != 1 {
		log.Fatalf("'%s' is not one character", value)
	}
	if !regexp.MustCompile(`[UJZ]`).MatchString(value) {
		log.Fatalf("'%s' must be U, J or Z", value)
	}
	return Id{value}
}
