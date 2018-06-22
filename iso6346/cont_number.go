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

package iso6346

import (
	"fmt"
	"log"
	"strconv"
)

// ContNumber is a container number with needed properties to conform to ISO 6346.
type ContNumber struct {
	ownerCode    OwnerCode
	equipCatID   EquipCatID
	serialNumber SerialNum
	checkDigit   int
}

// OwnerCode returns the owner code of container number.
func (cn ContNumber) OwnerCode() OwnerCode {
	return cn.ownerCode
}

// EquipCatID returns the equipment category ID of container number.
func (cn ContNumber) EquipCatID() EquipCatID {
	return cn.equipCatID
}

// SerialNumber returns the serial number of container number.
func (cn ContNumber) SerialNumber() SerialNum {
	return cn.serialNumber
}

// CheckDigit returns the check digit of container number.
func (cn ContNumber) CheckDigit() int {
	return cn.checkDigit
}

// NewContNum creates new container number with check digit.
// To have a correct check digit calculate it before.
func NewContNum(ownerCode OwnerCode,
	equipCatID EquipCatID,
	serialNumber SerialNum,
	checkDigit int) ContNumber {

	return ContNumber{ownerCode: ownerCode,
		equipCatID:   equipCatID,
		serialNumber: serialNumber,
		checkDigit:   checkDigit}
}

// SerialNum is a 6 digit number found in a container number.
type SerialNum struct {
	value int
}

// Value returns int value of serial number.
func (sn SerialNum) Value() int {
	return sn.value
}

// NewSerialNumFrom creates new serial number from a string value.
func NewSerialNumFrom(value string) SerialNum {

	num, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Could not parse '%s' to number", value)
	}
	return NewSerialNum(num)
}

// NewSerialNum creates new serial number from an int value.
func NewSerialNum(value int) SerialNum {

	if value < 0 || value > 999999 {
		log.Fatalf("'%d' is not '>= 0' and '<= 999999'", value)
	}
	return SerialNum{value}
}

func (sn SerialNum) String() string {
	return fmt.Sprintf("%06d", sn.Value())
}
