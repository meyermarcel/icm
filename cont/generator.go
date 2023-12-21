package cont

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

// GeneratorBuilder is the struct for the builder.
// Use NewUniqueGeneratorBuilder to create a new one.
type GeneratorBuilder struct {
	rand                 *rand.Rand
	codes                []string
	sepOE, sepES, sepSC  string
	count                int
	start                int
	end                  int
	exclCheckDigit10     bool
	exclTranspositionErr bool
}

// NewUniqueGeneratorBuilder returns a new random unique container number generator.
// If possible maximum unique container numbers are exceeded, count is less than 1 or
// no owner codes are passed then nil and error is returned.
func NewUniqueGeneratorBuilder(rand *rand.Rand) *GeneratorBuilder {
	return &GeneratorBuilder{
		rand:  rand,
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

// Separators sets the strings between owner code, equipment category ID, serial number and check digit.
func (gb *GeneratorBuilder) Separators(sepOE, sepES, sepSC string) *GeneratorBuilder {
	gb.sepOE, gb.sepES, gb.sepSC = sepOE, sepES, sepSC
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

// Build returns a new UniqueGenerator if all requirements are met.
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

	startIsSet := gb.start > -1
	endIsSet := gb.end > -1

	if startIsSet && endIsSet {
		serialNumIt = newSeqSerialNumIt(gb.start)
		if gb.start <= gb.end {
			count = gb.end + 1 - gb.start
		} else {
			count = gb.end + 1000000 + 1 - gb.start
		}
	} else if startIsSet && !endIsSet {
		serialNumIt = newSeqSerialNumIt(gb.start)
		count = gb.count
	} else if !startIsSet && endIsSet {
		serialNumIt = newSeqSerialNumIt(gb.end + 1 - gb.count)
		count = gb.count
	} else if !startIsSet && !endIsSet {
		if gb.count < 1 {
			return nil, fmt.Errorf("count %d is lower than minimum count 1", gb.count)
		}
		serialNumIt = newRandSerialNumIt(gb.rand.Int())
		count = gb.count
	}

	gb.rand.Shuffle(lenCodes, func(i, j int) {
		gb.codes[i], gb.codes[j] = strings.ToUpper(gb.codes[j]), strings.ToUpper(gb.codes[i])
	})

	return &UniqueGenerator{
		codes:                gb.codes,
		sepOE:                gb.sepOE,
		sepES:                gb.sepES,
		sepSC:                gb.sepSC,
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
	sepOE, sepES, sepSC  string
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
// which will then be available through the ContNumFmt method. It returns false
// when the generation stops by reaching the count of generated container numbers.
func (g *UniqueGenerator) Generate() bool {
	if g.generatedCount == g.count {
		return false
	}

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
	if g.exclTranspositionErr && len(CheckTransposition(code, "U", serialNum, checkDigit)) > 0 {
		return g.Generate()
	}
	g.contNum = Number{code, "U", serialNum, checkDigit % 10}
	g.generatedCount++

	return true
}

// ContNumFmt returns a generated container number with formatting.
func (g *UniqueGenerator) ContNumFmt() NumberFmt {
	return NumberFmt{
		g.contNum, g.sepOE, g.sepES, g.sepSC,
	}
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

func newRandSerialNumIt(randomOffset int) serialNumIt {
	return &randSerialNumIt{
		randOffset: randomOffset,
	}
}

func (r *randSerialNumIt) num() int {
	return permSerialNum((permSerialNum(r.it) + r.randOffset) % 1000000)
}

func (r *randSerialNumIt) increment() {
	r.it++
}

func (r *randSerialNumIt) isLast() bool {
	return (r.it+1)%1000000 == 0
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
