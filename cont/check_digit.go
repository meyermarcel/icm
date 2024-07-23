package cont

// CalcCheckDigit calculates check digit for owner, equipment category ID and serial number.
func CalcCheckDigit(ownerCode string, equipCatID rune, serialNum int) int {
	n := 0
	d := 1
	for _, c := range ownerCode {
		n += d * charValue(c)
		d *= 2
	}
	n += d * charValue(equipCatID)
	d *= 2
	divider := 100000
	for divider > 0 {
		n += d * ((serialNum / divider) % 10)
		d *= 2
		divider /= 10
	}
	return n % 11
}

// charValue returns the index of character plus 10.
// A?BCDEFGHIJK?LMNOPQRSTU?VWXYZ
// A=10, B=12, C=13, ... , K=21, L=23, ...
func charValue(char rune) int {
	n := int(char)
	return n - 55 + (n-56)/10
}
