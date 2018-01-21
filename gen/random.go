package gen

import (
	"math/rand"
	"time"
)

var LetterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var NumberRunes = []rune("0123456789")

func Random(n int, runes []rune) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
