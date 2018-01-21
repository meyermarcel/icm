package iso6346

import (
	"math/rand"
	"time"
	"log"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")

type ContainerNumber struct {
	OwnerCode           OwnerCode
	EquipmentCategoryId EquipmentCategoryId
	SerialNumber        SerialNumber
	CheckDigit          int
}

func Generate() ContainerNumber {

	ownerCode, err := NewOwnerCode(randStringRunes(3, letterRunes))

	if err != nil {
		log.Fatal(err)
	}

	equipmentCategoryId, err := NewEquipmentCategoryId(randStringRunes(1, EquipmentCategoryIds))

	if err != nil {
		log.Fatal(err)
	}

	serialNumber, err := NewSerialNumber(randStringRunes(6, numberRunes))

	if err != nil {
		log.Fatal(err)
	}

	checkDigit := CalculateCheckDigit(ownerCode,
		equipmentCategoryId,
		serialNumber)

	return ContainerNumber{ownerCode,
		equipmentCategoryId,
		serialNumber, checkDigit}
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
