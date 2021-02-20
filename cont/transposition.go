// Copyright © 2018 Marcel Meyer meyermarcel@posteo.de
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cont

import (
	"fmt"
	"strconv"
)

// CheckTransposition checks for possible transposition errors.
// Not equal adjacent digits including check digit are transposed and checked.
func CheckTransposition(ownerCode string, equipCatID string, serialNum string) []Number {
	contNums := make([]Number, 0)

	checkDigit := CalcCheckDigit(ownerCode, equipCatID, serialNum) % 10
	for pos := range serialNum {
		if pos < 5 && serialNum[pos] != serialNum[pos+1] {
			transposedSerialNum := fmt.Sprintf(
				"%s%c%c%s", serialNum[:pos], serialNum[pos+1], serialNum[pos], serialNum[pos+2:])
			calcCheckDigit := CalcCheckDigit(ownerCode, equipCatID, transposedSerialNum) % 10
			if checkDigit == calcCheckDigit {
				contNums = append(contNums, newNum(ownerCode, equipCatID, transposedSerialNum, calcCheckDigit))
			}
		} else if pos == 5 && string(serialNum[pos]) != strconv.Itoa(checkDigit) {
			transposedSerialNum := serialNum[:5] + strconv.Itoa(checkDigit)
			transposedCheckDigit := CalcCheckDigit(ownerCode, equipCatID, transposedSerialNum) % 10
			lastSerialDigit, _ := strconv.Atoi(string(serialNum[5]))
			if lastSerialDigit == transposedCheckDigit {
				contNums = append(contNums, newNum(ownerCode, equipCatID, transposedSerialNum, lastSerialDigit))
			}

		}
	}
	return contNums
}
