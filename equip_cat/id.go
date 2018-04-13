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

package equip_cat

import (
	"log"
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

func NewIdU() Id {
	return Id{"U"}
}

func NewIdFrom(value string) Id {

	if utf8.RuneCountInString(value) != 1 {
		log.Fatalf("'%s' is not one character", value)
	}
	if !regexp.MustCompile(`[UJZ]`).MatchString(value) {
		log.Fatalf("'%s' must be U, J or Z", value)
	}
	return Id{value}
}
