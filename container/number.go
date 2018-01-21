package container

import (
	"iso6346/owner"
	"iso6346/equip_cat"
)

type Number struct {
	ownerCode    owner.Code
	equipCatId   equip_cat.Id
	serialNumber SerialNumber
	checkDigit   int
}

func (cn Number) OwnerCode() owner.Code {
	return cn.ownerCode
}

func (cn Number) EquipCatId() equip_cat.Id {
	return cn.equipCatId
}

func (cn Number) SerialNumber() SerialNumber {
	return cn.serialNumber
}

func (cn Number) CheckDigit() int {
	return cn.checkDigit
}

func NewContainerNumber(ownerCode owner.Code,
	equipCatId equip_cat.Id,
	serialNumber SerialNumber,
	checkDigit int) Number {

	return Number{ownerCode: ownerCode,
		equipCatId: equipCatId,
		serialNumber: serialNumber,
		checkDigit: checkDigit}
}
