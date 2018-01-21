package owner

import (
	"unicode/utf8"
	"fmt"
	"regexp"
)

type Code struct {
	value string
}

func (c Code) Value() string {
	return c.value
}

func NewCode(value string) (Code, error) {

	if utf8.RuneCountInString(value) != 3 {
		return Code{}, fmt.Errorf("'%s' is not three characters", value)
	}

	if !regexp.MustCompile(`[A-Z]{3}`).MatchString(value) {
		return Code{}, fmt.Errorf("'%s' must be 3 letters", value)
	}
	return Code{value}, nil
}
