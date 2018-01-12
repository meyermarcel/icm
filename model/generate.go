package model

import (
	"math/rand"
	"time"
	"log"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")

func Generate() ContainerNumber {

	ownerCode, err := NewOwnerCode(randStringRunes(3, letterRunes))

	if err == nil {
		log.Fatal(err)
	}

	serialNumber, err := NewSerialNumber(randStringRunes(3, numberRunes))

	if err == nil {
		log.Fatal(err)
	}
	equipmentCategoryId, err := NewEquipmentCategoryId(randStringRunes(3, letterRunes))

	if err == nil {
		log.Fatal(err)
	}

	return NewUncheckedContainerNumber(ownerCode,
		equipmentCategoryId,
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
