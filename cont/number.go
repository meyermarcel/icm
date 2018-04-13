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
	"github.com/meyermarcel/iso6346/equip_cat"
	"github.com/meyermarcel/iso6346/owner"
)

type Number struct {
	ownerCode    owner.Code
	equipCatId   equip_cat.Id
	serialNumber SerialNum
	checkDigit   int
}

func (cn Number) OwnerCode() owner.Code {
	return cn.ownerCode
}

func (cn Number) EquipCatId() equip_cat.Id {
	return cn.equipCatId
}

func (cn Number) SerialNumber() SerialNum {
	return cn.serialNumber
}

func (cn Number) CheckDigit() int {
	return cn.checkDigit
}

func NewContNum(ownerCode owner.Code,
	equipCatId equip_cat.Id,
	serialNumber SerialNum,
	checkDigit int) Number {

	return Number{ownerCode: ownerCode,
		equipCatId:   equipCatId,
		serialNumber: serialNumber,
		checkDigit:   checkDigit}
}
