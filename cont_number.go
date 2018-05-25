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
	"fmt"
	"log"
	"strconv"
)

type contNumber struct {
	ownerCode    ownerCode
	equipCatID   equipCatID
	serialNumber serialNum
	checkDigit   int
}

func (cn contNumber) OwnerCode() ownerCode {
	return cn.ownerCode
}

func (cn contNumber) EquipCatID() equipCatID {
	return cn.equipCatID
}

func (cn contNumber) SerialNumber() serialNum {
	return cn.serialNumber
}

func (cn contNumber) CheckDigit() int {
	return cn.checkDigit
}

func newContNum(ownerCode ownerCode,
	equipCatID equipCatID,
	serialNumber serialNum,
	checkDigit int) contNumber {

	return contNumber{ownerCode: ownerCode,
		equipCatID:   equipCatID,
		serialNumber: serialNumber,
		checkDigit:   checkDigit}
}

type serialNum struct {
	value int
}

func (sn serialNum) Value() int {
	return sn.value
}

func serialNumFrom(value string) serialNum {

	num, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Could not parse '%s' to number", value)
	}
	return newSerialNum(num)
}

func newSerialNum(value int) serialNum {

	if value < 0 || value > 999999 {
		log.Fatalf("'%d' is not '>= 0' and '<= 999999'", value)
	}
	return serialNum{value}
}

func (sn serialNum) String() string {
	return fmt.Sprintf("%06d", sn.Value())
}
