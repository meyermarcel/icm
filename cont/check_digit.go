package cont

import (
	"strings"
)

// CalcCheckDigit calculates check digit for owner, equipment category ID and serial number.
func CalcCheckDigit(ownerCode string, equipCatID string, serialNum string) int {
	concat := ownerCode + equipCatID + serialNum

	n := 0.0
	d := 0.5
	for _, character := range concat {
		d *= 2
		n += d * float64(strings.IndexRune("0123456789A?BCDEFGHIJK?LMNOPQRSTU?VWXYZ", character))
	}
	return int(n) - int(n/11)*11
}
