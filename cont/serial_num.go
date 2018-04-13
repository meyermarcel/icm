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
	"fmt"
	"log"
	"strconv"
)

type SerialNum struct {
	value int
}

func (sn SerialNum) Value() int {
	return sn.value
}

func SerialNumFrom(value string) SerialNum {

	num, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Could not parse '%s' to number", value)
	}
	return NewSerialNum(num)
}

func NewSerialNum(value int) SerialNum {

	if value < 0 || value > 999999 {
		log.Fatalf("'%d' is not '>= 0' and '<= 999999'", value)
	}
	return SerialNum{value}
}

func (sn SerialNum) String() string {
	return fmt.Sprintf("%06d", sn.value)
}
