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

package cont

import (
	"log"
	"regexp"
	"unicode/utf8"
)

// Owner has a code and associated company with its location in the form of country and city.
type Owner struct {
	Code    OwnerCode
	Company string
	City    string
	Country string
}

// OwnerCode represents the container owner.
type OwnerCode struct {
	value string
}

// Value returns the string value of an OwnerCode.
func (c OwnerCode) Value() string {
	return c.value
}

// NewOwnerCode creates a new OwnerCode from a string value.
func NewOwnerCode(value string) OwnerCode {

	if utf8.RuneCountInString(value) != 3 {
		log.Fatalf("'%s' is not three characters", value)
	}

	if !regexp.MustCompile(`[A-Z]{3}`).MatchString(value) {
		log.Fatalf("'%s' must be 3 letters", value)
	}
	return OwnerCode{value}
}
