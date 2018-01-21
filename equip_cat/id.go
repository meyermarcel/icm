package equip_cat

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

var Ids = []rune("UJZ")

type Id struct {
	value string
}

func (id Id) Value() string {
	return id.value
}

func NewId(value string) (Id, error) {

	if utf8.RuneCountInString(value) != 1 {
		return Id{}, fmt.Errorf("'%s' is not one character", value)
	}
	if !regexp.MustCompile(`[UJZ]`).MatchString(value) {
		return Id{}, fmt.Errorf("'%s' must be U, J or Z", value)
	}
	return Id{value}, nil
}
