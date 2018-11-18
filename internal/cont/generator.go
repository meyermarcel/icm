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
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// RandomUniqueGenerator holds state for generating random unique container numbers.
// Use NewRandomUniqueGenerator for initialization.
type RandomUniqueGenerator struct {
	codes       []OwnerCode
	lenCodes    int
	randOffset  int
	ownerOffset int
	serialNumIt int
}

// NewRandomUniqueGenerator returns a new random unique container number generator.
// If possible maximum unique container numbers are exceeded, count is less than 1 or
// no owner codes are passed then nil and error is returned.
func NewRandomUniqueGenerator(count int, randomCodes []OwnerCode) (*RandomUniqueGenerator, error) {

	lenCodes := len(randomCodes)

	if count > lenCodes*909091 {
		return nil, fmt.Errorf("%d exceeds limit of %d (%d owners * 909091 serial numbers)",
			count, lenCodes*909091, lenCodes)
	}

	if count < 1 {
		return nil, fmt.Errorf("%d is lower than minimum count 1", count)
	}

	if lenCodes < 1 {
		return nil, errors.New("cannot generate container numbers without owner codes")
	}

	return &RandomUniqueGenerator{
		codes:      randomCodes,
		lenCodes:   lenCodes,
		randOffset: rand.Int(),
	}, nil
}

// Generate generates a random container number. Container number is unique if Generate is not called more
// than passed count in constructor.
func (g *RandomUniqueGenerator) Generate() Number {

	serialNum := NewSerialNum(permSerialNum((permSerialNum(g.serialNumIt) + g.randOffset) % 1000000))
	code := g.codes[(g.serialNumIt+g.ownerOffset)%g.lenCodes]
	equipCatID := NewEquipCatIDU()
	checkDigit := CalcCheckDigit(code, equipCatID, serialNum)

	if g.serialNumIt < 1000000 {
		g.serialNumIt++
	} else {
		g.serialNumIt = 0
		g.ownerOffset++
	}

	if checkDigit == 10 {
		return g.Generate()
	}
	return NewNum(code, equipCatID, serialNum, checkDigit)
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
