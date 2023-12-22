package cont

import "math"

// CheckTransposition checks for possible transposition errors.
// Not equal adjacent digits including check digit are transposed and checked.
func CheckTransposition(ownerCode string, equipCatID rune, serialNum int, checkDigit int) []Number {
	checkDigit = checkDigit % 10

	contNums := make([]Number, 0)

	// Only container numbers with check digit 0, 10 or 3 are affected
	if checkDigit != 3 && checkDigit != 0 {
		return contNums
	}
	for pos := 5; pos > 0; pos-- {
		swapped, transposedSerialNum := swapDigits(serialNum, pos-1, pos)
		if !swapped {
			continue
		}
		calcCheckDigit := CalcCheckDigit(ownerCode, equipCatID, transposedSerialNum) % 10
		if checkDigit == calcCheckDigit {
			contNums = append(contNums, Number{ownerCode, equipCatID, transposedSerialNum, calcCheckDigit})
		}
	}

	serialNumLastDigit := serialNum % 10
	if checkDigit == serialNumLastDigit {
		return contNums
	}

	transposedCheckDigitSerialNum := ((serialNum / 10) * 10) + checkDigit
	calcCheckDigit := CalcCheckDigit(ownerCode, equipCatID, transposedCheckDigitSerialNum) % 10
	if serialNumLastDigit == calcCheckDigit {
		contNums = append(contNums, Number{ownerCode, equipCatID, transposedCheckDigitSerialNum, serialNumLastDigit})
	}
	return contNums
}

// swapDigits returns true if the digits are different and returns the number with swapped numbers.
// false is returned if the digits are same and 0 is returned.
// Position 0 is the digit beginning on the right side of number.
func swapDigits(number int, pos1, pos2 int) (bool, int) {
	p1 := int(math.Pow10(pos1))
	p2 := int(math.Pow10(pos2))

	digit1 := (number / p1) % 10
	digit2 := (number / p2) % 10
	if digit1 == digit2 {
		return false, 0
	}

	number -= digit1 * p1
	number -= digit2 * p2

	number += digit1 * p2
	number += digit2 * p1

	return true, number
}
