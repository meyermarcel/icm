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

package main

import (
	"log"
	"regexp"
	"unicode/utf8"
)

type ownerCode struct {
	value string
}

func (c ownerCode) Value() string {
	return c.value
}

func newOwnerCode(value string) ownerCode {

	if utf8.RuneCountInString(value) != 3 {
		log.Fatalf("'%s' is not three characters", value)
	}

	if !regexp.MustCompile(`[A-Z]{3}`).MatchString(value) {
		log.Fatalf("'%s' must be 3 letters", value)
	}
	return ownerCode{value}
}

func resolver(pathToDB string) func(code ownerCode) (owner, bool) {
	return func(code ownerCode) (owner, bool) {
		db := openDB(pathToDB)
		defer db.Close()
		return getOwner(db, code)
	}
}

func getRandomOwnerCodes(pathToDB string, count int) []ownerCode {
	db := openDB(pathToDB)
	defer db.Close()
	return getRandomCodes(db, count)
}
