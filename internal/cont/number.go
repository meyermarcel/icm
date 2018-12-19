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

// Number is a container number with needed properties to conform to the specified standard.
type Number struct {
	ownerCode    string
	equipCatID   string
	serialNumber string
	checkDigit   int
}

// OwnerCode returns the owner code of container number.
func (cn Number) OwnerCode() string {
	return cn.ownerCode
}

// EquipCatID returns the equipment category ID of container number.
func (cn Number) EquipCatID() string {
	return cn.equipCatID
}

// SerialNumber returns the serial number of container number.
func (cn Number) SerialNumber() string {
	return cn.serialNumber
}

// CheckDigit returns the check digit of container number.
func (cn Number) CheckDigit() int {
	return cn.checkDigit
}

func newNum(ownerCode string,
	equipCatID string,
	serialNumber string,
	checkDigit int) Number {

	return Number{ownerCode: ownerCode,
		equipCatID:   equipCatID,
		serialNumber: serialNumber,
		checkDigit:   checkDigit}
}
