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
	"log"
	"math/rand"
	"time"
)

// GenNum creates a specified count of container numbers. Random owner code generator
// is needed to define owner code values in generated container numbers.
func GenNum(count int, c chan Number, randomOwnerCodes func(count int) []OwnerCode) {

	codes := randomOwnerCodes(count)
	randOffset := rand.Int()
	lenCodes := len(codes)

	if count > lenCodes*1000000 {
		log.Fatalf("'%d' exceeds generate limit %d (%d owners * 1000000 serial numbers)", count, lenCodes*1000000, lenCodes)
	}

	equipCatID := NewEquipCatIDU()

	serialNumPasses := count / 1000000
	for ownerOffset := 0; ownerOffset <= serialNumPasses; ownerOffset++ {

		for i := 0; i < count && i < 1000000; i++ {
			serialNum := NewSerialNum(permSerialNum((permSerialNum(i) + randOffset) % 1000000))

			code := codes[(i+ownerOffset)%lenCodes]
			checkDigit := CalcCheckDigit(code, equipCatID, serialNum)

			c <- NewNum(code, equipCatID, serialNum, checkDigit)
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