package cont

import (
	"fmt"
)

// Number is a container number with needed properties to conform to the specified standard.
type Number struct {
	ownerCode    string
	equipCatID   rune
	serialNumber int
	checkDigit   int
}

type NumberFmt struct {
	Number
	SepOE, SepES, SepSC string
}

func (cn *NumberFmt) String() string {
	return fmt.Sprintf("%s%s%s%s%06d%s%d",
		cn.ownerCode, cn.SepOE,
		string(cn.equipCatID), cn.SepES,
		cn.serialNumber, cn.SepSC,
		cn.checkDigit)
}
