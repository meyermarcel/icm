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
)

// Number is a container number with needed properties to conform to the specified standard.
type Number struct {
	ownerCode           string
	equipCatID          string
	serialNumber        string
	checkDigit          int
	sepOE, sepES, sepSC string
}

// SetSeparators sets separator for formatting for Stringer interface.
func (cn *Number) SetSeparators(sepOE, sepES, sepSC string) {
	cn.sepOE, cn.sepES, cn.sepSC = sepOE, sepES, sepSC
}

func (cn Number) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%d",
		cn.ownerCode, cn.sepOE,
		cn.equipCatID, cn.sepES,
		cn.serialNumber, cn.sepSC,
		cn.checkDigit)
}

func newNum(ownerCode string,
	equipCatID string,
	serialNumber string,
	checkDigit int) Number {

	return Number{
		ownerCode:    ownerCode,
		equipCatID:   equipCatID,
		serialNumber: serialNumber,
		checkDigit:   checkDigit,
		sepOE:        " ",
		sepES:        " ",
		sepSC:        " ",
	}
}
