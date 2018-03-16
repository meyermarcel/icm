package owner

import (
	"log"
	"regexp"
	"unicode/utf8"
)

type Code struct {
	value string
}

func (c Code) Value() string {
	return c.value
}

func NewCode(value string) Code {

	if utf8.RuneCountInString(value) != 3 {
		log.Fatalf("'%s' is not three characters", value)
	}

	if !regexp.MustCompile(`[A-Z]{3}`).MatchString(value) {
		log.Fatalf("'%s' must be 3 letters", value)
	}
	return Code{value}
}

func Resolver(pathToDB string) func(code Code) (Owner, bool) {
	return func(code Code) (Owner, bool) {
		db := openDB(pathToDB)
		defer db.Close()
		return getOwner(db, code)
	}
}

func GetRandomCodes(pathToDB string, count int) []Code {
	db := openDB(pathToDB)
	defer db.Close()
	return getRandomCodes(db, count)
}
