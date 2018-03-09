package cont

import (
	"github.com/meyermarcel/iso6346/equip_cat"
	"github.com/meyermarcel/iso6346/owner"
	"log"
	"math/rand"
	"time"
)

func Gen(count int, c chan Number) {

	codes := owner.GetRandomCodes(count)
	randOffset := rand.Int()
	lenCodes := len(codes)

	if count > lenCodes*1000000 {
		log.Fatalf("'%d' exceeds generate limit %d (%d owners * 1000000 serial numbers)", count, lenCodes*1000000, lenCodes)
	}

	equipCatId := equip_cat.NewIdU()

	serialNumPasses := count / 1000000
	for ownerOffset := 0; ownerOffset <= serialNumPasses; ownerOffset++ {

		for i := 0; i < count && i < 1000000; i++ {
			serialNum := NewSerialNum(permSerialNum((permSerialNum(i) + randOffset) % 1000000))

			code := codes[(i+ownerOffset)%lenCodes]
			checkDigit := CalcCheckDigit(code, equipCatId, serialNum)

			c <- NewContNum(code, equipCatId, serialNum, checkDigit)
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
