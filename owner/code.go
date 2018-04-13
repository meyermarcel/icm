// Copyright © 2018 Marcel Meyer meyermarcel@posteo.de
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
