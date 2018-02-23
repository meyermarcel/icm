package cont

import (
	"iso6346/owner"
	"iso6346/equip_cat"
	"time"
	"math/rand"
)

var numRunes = []rune("0123456789")

func Gen() Number {

	ownerCode := owner.GetRandomCode()
	equipCatId := equip_cat.NewId(random(1, equip_cat.Ids))
	serialNumber := NewSerialNum(random(6, numRunes))

	checkDigit := CalcCheckDigit(ownerCode,
		equipCatId,
		serialNumber)

	return NewContNum(ownerCode,
		equipCatId,
		serialNumber, checkDigit)
}

func random(n int, runes []rune) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}