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
	checkDigit int,
) Number {
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
