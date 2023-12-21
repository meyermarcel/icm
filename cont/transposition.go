package cont

import (
	"fmt"
	"strconv"
)

// CheckTransposition checks for possible transposition errors.
// Not equal adjacent digits including check digit are transposed and checked.
func CheckTransposition(ownerCode string, equipCatID string, serialNum string, checkDigit int) []Number {
	checkDigit = checkDigit % 10

	contNums := make([]Number, 0)

	// Only container numbers with check digit 0, 10 or 3 are affected
	if checkDigit != 3 && checkDigit != 0 {
		return contNums
	}
	for pos := range serialNum {
		if pos < 5 && serialNum[pos] != serialNum[pos+1] {
			transposedSerialNum := fmt.Sprintf(
				"%s%c%c%s", serialNum[:pos], serialNum[pos+1], serialNum[pos], serialNum[pos+2:])
			calcCheckDigit := CalcCheckDigit(ownerCode, equipCatID, transposedSerialNum) % 10
			if checkDigit == calcCheckDigit {
				contNums = append(contNums, Number{ownerCode, equipCatID, transposedSerialNum, calcCheckDigit})
			}
		} else if pos == 5 && int(serialNum[pos]-'0') != checkDigit {
			transposedSerialNum := serialNum[:5] + strconv.Itoa(checkDigit)
			transposedCheckDigit := CalcCheckDigit(ownerCode, equipCatID, transposedSerialNum) % 10
			lastSerialDigit := int(serialNum[5] - '0')
			if lastSerialDigit == transposedCheckDigit {
				contNums = append(contNums, Number{ownerCode, equipCatID, transposedSerialNum, lastSerialDigit})
			}
		}
	}
	return contNums
}
