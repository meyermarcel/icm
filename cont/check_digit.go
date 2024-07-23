package cont

// CalcCheckDigit calculates check digit for owner, equipment category ID and serial number.
// This function was optimized for fun and has a suboptimal reading experience.
func CalcCheckDigit(ownerCode string, equipCatID rune, serialNum int) int {
	var n uint16
	var d uint16 = 1

	for _, c := range ownerCode {
		n += d * charValue(uint16(c))
		d *= 2
	}

	n += d * charValue(uint16(equipCatID))

	// Handle the case for the serial number when it is
	// out of range of uint16.
	n += 512 * uint16(serialNum%10)
	serialNum /= 10
	n += 256 * uint16(serialNum%10)

	// uint16 can be used because we are now <10000
	s := uint16(serialNum / 10)
	d = 128
	for d >= 16 {
		n += d * (s % 10)
		d /= 2
		s /= 10
	}
	return int(n % 11)
}

// charValue returns the index of character plus 10.
// A?BCDEFGHIJK?LMNOPQRSTU?VWXYZ
// A=10, (no 11) B=12, C=13, ... , K=21, (no 22) L=23, ...
func charValue(char uint16) uint16 {
	return char - 55 + (char-56)/10
}
