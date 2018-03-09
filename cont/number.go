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
