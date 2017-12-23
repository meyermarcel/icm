package ui

import (
	"iso6346/model"
	"fmt"
)
/*

 ABC D 123456 7
    ^ ^      ^
    | |      |
    | |       - thirdSep
    | |
    |  - secondSep
    |
     - firstSep
*/

type ContainerNumberFormatted struct {
	cn        model.ContainerNumber
	firstSep  string
	secondSep string
	thirdSep  string
}

func CnFormatter(cn model.ContainerNumber) ContainerNumberFormatted {

	return ContainerNumberFormatted{cn,
		"",
		"",
		""}
}

func (fcn ContainerNumberFormatted) FirstSep(ABC_D1234567 string) ContainerNumberFormatted {
	fcn.firstSep = ABC_D1234567
	return fcn
}

func (fcn ContainerNumberFormatted) SecondSep(ABCD_1234567 string) ContainerNumberFormatted {
	fcn.secondSep = ABCD_1234567
	return fcn
}

func (fcn ContainerNumberFormatted) ThirdSep(ABCD123456_7 string) ContainerNumberFormatted {
	fcn.thirdSep = ABCD123456_7
	return fcn
}

func (fcn ContainerNumberFormatted) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%d",
		fcn.cn.OwnerCode(),
		fcn.firstSep,
		fcn.cn.EquipmentCategoryId(),
		fcn.secondSep,
		fcn.cn.SerialNumber(),
		fcn.thirdSep,
		fcn.cn.CheckDigit())
}
