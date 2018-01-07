package ui

import (
	"fmt"
	"iso6346/model"
)

/*

  ABC U 123456 0
     ↑ ↑      ↑
     │ │      └─ 3rd separator
     │ │
     │ └─ 2nd separator
     │
     └─ 1st separator
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

func (fcn ContainerNumberFormatted) FirstSep(ABC_U1234560 string) ContainerNumberFormatted {
	fcn.firstSep = ABC_U1234560
	return fcn
}

func (fcn ContainerNumberFormatted) SecondSep(ABCU_1234560 string) ContainerNumberFormatted {
	fcn.secondSep = ABCU_1234560
	return fcn
}

func (fcn ContainerNumberFormatted) ThirdSep(ABCU123456_0 string) ContainerNumberFormatted {
	fcn.thirdSep = ABCU123456_0
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
