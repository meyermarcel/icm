package cont

import (
	"log"
	"strconv"
	"fmt"
)

type SerialNum struct {
	value int
}

func (sn SerialNum) Value() int {
	return sn.value
}

func SerialNumFrom(value string) SerialNum {

	num, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Could not parse '%s' to number", value)
	}
	return NewSerialNum(num)
}

func NewSerialNum(value int) SerialNum {

	if value < 0 || value > 999999 {
		log.Fatalf("'%d' is not '>= 0' and '<= 999999'", value)
	}
	return SerialNum{value}
}

func (sn SerialNum) String() string {
	return fmt.Sprintf("%06d", sn.value)
}
