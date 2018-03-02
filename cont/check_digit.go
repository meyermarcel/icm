package cont

import (
	"iso6346/equip_cat"
	"strings"
	"iso6346/owner"
)

/*
This method is a modified version of a Go code sample from
https://en.wikipedia.org/wiki/ISO_6346#Code_Sample_(Go)
*/
func CalcCheckDigit(ownerCode owner.Code, equipCatId equip_cat.Id, serialNum SerialNum) int {

	concat := ownerCode.Value() + equipCatId.Value()

	n := 0.0
	d := 0.5
	for _, character := range concat {
		d *= 2
		n += d * float64(strings.IndexRune("??????????A?BCDEFGHIJK?LMNOPQRSTU?VWXYZ", character))
	}
	div := 100000.0
	for i := 0; i < 6; i++ {
		d *= 2
		n += d * float64(int(float64(serialNum.Value())/div)%10)
		div /= 10
	}
	return (int(n) - int(n/11)*11) % 10
}
