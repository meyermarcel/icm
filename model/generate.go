package model

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")

func Generate() ContainerNumber {

	ownerCode := randStringRunes(3, letterRunes)

	serialNumber := randStringRunes(6, numberRunes)

	equipmentCategoryIdentifier := randStringRunes(1, []rune(EquipmentCategoryIds))

	return NewUncheckedContainerNumber(ownerCode,
		equipmentCategoryIdentifier,
		serialNumber)

}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randStringRunes(n int, runes []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
