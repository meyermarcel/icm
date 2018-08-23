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
	"math/rand"
	"time"
)

// Result is the return value that includes a container number or an error.
type Result struct {
	contNum Number
	err     error
}

// ContNum returns the container number.
func (gcn *Result) ContNum() Number {
	return gcn.contNum
}

// Err returns the error.
func (gcn *Result) Err() error {
	return gcn.err
}

// GenNum creates a specified count of container numbers. Random owner code generator
// is needed to define owner code values in generated container numbers.
func GenNum(count int, c chan Result, random func(count int) []OwnerCode) {

	codes := random(count)
	lenCodes := len(codes)

	// An 11th of 1.000.000 serial numbers produce check digit 10.
	//                  909091 = ( 1.000.000 / 11 ) * 10
	if count > lenCodes*909091 {
		c <- Result{err: fmt.Errorf("'%d' exceeds generate limit %d (%d owners * 909091 serial numbers)",
			count, lenCodes*909091, lenCodes)}
		close(c)
		return
	}

	randOffset := rand.Int()

	equipCatID := NewEquipCatIDU()

	serialNumPasses := count / 1000000
	for ownerOffset := 0; ownerOffset <= serialNumPasses; ownerOffset++ {

		for i := 0; i < count && i < 1000000; i++ {
			serialNum := NewSerialNum(permSerialNum((permSerialNum(i) + randOffset) % 1000000))

			code := codes[(i+ownerOffset)%lenCodes]
			checkDigit := CalcCheckDigit(code, equipCatID, serialNum)
			if checkDigit != 10 {
				c <- Result{contNum: NewNum(code, equipCatID, serialNum, checkDigit)}
			} else {
				count++
			}

		}
		count -= 1000000
	}
	close(c)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// See http://preshing.com/20121224/how-to-generate-a-sequence-of-unique-random-integers
func permSerialNum(x int) int {
	// last prime number before 1000000
	// and satisfies p ≡ 3 mod 4
	const prime = 999983

	if x >= prime {
		return x
	}
	residue := (x * x) % prime
	if x <= prime/2 {
		return residue
	}
	return prime - residue
}
