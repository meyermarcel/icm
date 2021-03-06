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
	"strings"
	"time"
)

// GeneratorBuilder is the struct for the builder.
// Use NewUniqueGeneratorBuilder to create a new one.
type GeneratorBuilder struct {
	codes                []string
	count                int
	start                int
	end                  int
	exclCheckDigit10     bool
	exclTranspositionErr bool
}

// NewUniqueGeneratorBuilder returns a new random unique container number generator.
// If possible maximum unique container numbers are exceeded, count is less than 1 or
// no owner codes are passed then nil and error is returned.
func NewUniqueGeneratorBuilder() *GeneratorBuilder {
	return &GeneratorBuilder{
		count: 1,
		start: -1,
		end:   -1,
	}
}

// OwnerCodes sets the owner codes for generation.
func (gb *GeneratorBuilder) OwnerCodes(codes []string) *GeneratorBuilder {
	gb.codes = codes
	return gb
}

// Count sets the count of container number.
func (gb *GeneratorBuilder) Count(count int) *GeneratorBuilder {
	gb.count = count
	return gb
}

// Start sets the start of serial number range.
func (gb *GeneratorBuilder) Start(start int) *GeneratorBuilder {
	gb.start = start
	return gb
}

// End sets the end of serial number range.
func (gb *GeneratorBuilder) End(end int) *GeneratorBuilder {
	gb.end = end
	return gb
}

// ExcludeCheckDigit10 sets the exclusion of container numbers with check digit 10.
func (gb *GeneratorBuilder) ExcludeCheckDigit10(exclude bool) *GeneratorBuilder {
	gb.exclCheckDigit10 = exclude
	return gb
}

// ExcludeTranspositionErr sets the exclusion of container numbers with possible transposition errors.
func (gb *GeneratorBuilder) ExcludeTranspositionErr(exclude bool) *GeneratorBuilder {
	gb.exclTranspositionErr = exclude
	return gb
}

// Build returns a new UniqueGenerator if all requirements met.
// Valid combinations a
func (gb *GeneratorBuilder) Build() (*UniqueGenerator, error) {

	lenCodes := len(gb.codes)

	if lenCodes < 1 {
		return nil, errors.New("cannot generate container numbers without owner codes")
	}

	serialNums := 1000000

	if gb.exclCheckDigit10 {
		serialNums = 909091
	}

	if gb.count > lenCodes*serialNums {
		return nil, fmt.Errorf("count %d exceeds limit of %d (%d owners * %d serial numbers)",
			gb.count, lenCodes*serialNums, lenCodes, serialNums)
	}

	var serialNumIt serialNumIt
	var count int

	if gb.start > -1 && gb.end > -1 {
		serialNumIt = newSeqSerialNumIt(gb.start)
		if gb.start <= gb.end {
			count = gb.end + 1 - gb.start
		} else {
			count = gb.end + 1000000 + 1 - gb.start
		}
	}

	if gb.start > -1 && gb.end == -1 {
		serialNumIt = newSeqSerialNumIt(gb.start)
		count = gb.count
	}

	if gb.end > -1 && gb.start == -1 {
		serialNumIt = newSeqSerialNumIt(gb.end + 1 - gb.count)
		count = gb.count
	}

	if gb.start == -1 && gb.end == -1 {
		if gb.count < 1 {
			return nil, fmt.Errorf("count %d is lower than minimum count 1", gb.count)
		}
		serialNumIt = newRandSerialNumIt()
		count = gb.count
	}

	rand.Shuffle(lenCodes, func(i, j int) {
		gb.codes[i], gb.codes[j] = strings.ToUpper(gb.codes[j]), strings.ToUpper(gb.codes[i])
	})

	return &UniqueGenerator{
		codes:                gb.codes,
		lenCodes:             lenCodes,
		serialNumIt:          serialNumIt,
		count:                count,
		exclCheckDigit10:     gb.exclCheckDigit10,
		exclTranspositionErr: gb.exclTranspositionErr,
	}, nil
}

// UniqueGenerator holds state for generating random unique container numbers.
// Use NewUniqueGeneratorBuilder for initialization.
type UniqueGenerator struct {
	codes                []string
	lenCodes             int
	ownerOffset          int
	serialNumIt          serialNumIt
	count                int
	contNum              Number
	generatedCount       int
	exclCheckDigit10     bool
	exclTranspositionErr bool
}

// Generate advances the serial number iterator to the next serial number,
// which will then be available through the ContNum method. It returns false
// when the generation stops by reaching the count of generated container numbers.
func (g *UniqueGenerator) Generate() bool {
	code := g.codes[(g.serialNumIt.num()+g.ownerOffset)%g.lenCodes]
	serialNum := fmt.Sprintf("%06d", g.serialNumIt.num())
	checkDigit := CalcCheckDigit(code, "U", serialNum)

	if g.serialNumIt.isLast() {
		g.ownerOffset++
	}
	g.serialNumIt.increment()

	if g.exclCheckDigit10 && checkDigit == 10 {
		return g.Generate()
	}
	if g.exclTranspositionErr && len(CheckTransposition(code, "U", serialNum)) > 0 {
		return g.Generate()
	}
	g.contNum = newNum(code, "U", serialNum, checkDigit%10)
	g.generatedCount++
	return g.generatedCount <= g.count
}

// ContNum returns generated container number.
func (g *UniqueGenerator) ContNum() Number {
	return g.contNum
}

type serialNumIt interface {
	num() int

	increment()

	isLast() bool
}

type randSerialNumIt struct {
	randOffset int
	it         int
}

func newRandSerialNumIt() serialNumIt {
	return &randSerialNumIt{
		randOffset: rand.Int(),
	}
}

func (r *randSerialNumIt) num() int {
	return permSerialNum((permSerialNum(r.it) + r.randOffset) % 1000000)
}

func (r *randSerialNumIt) increment() {
	r.it++
}

func (r *randSerialNumIt) isLast() bool {
	return r.it+1%1000000 == 1000000
}

type seqSerialNumIt struct {
	start int
	it    int
}

func newSeqSerialNumIt(start int) serialNumIt {
	start = (start + 1000000) % 1000000
	return &seqSerialNumIt{
		start: start,
		it:    start,
	}
}

func (i *seqSerialNumIt) num() int {
	return i.it
}

func (i *seqSerialNumIt) increment() {
	i.it = (i.it + 1) % 1000000
}

func (i *seqSerialNumIt) isLast() bool {
	return (i.it+1)%1000000 == i.start
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

func init() {
	rand.Seed(time.Now().UnixNano())
}
