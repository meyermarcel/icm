package ui

import (
	"fmt"
	"strings"
)

type MsgType int

const (
	HINT MsgType = iota
	INFO 
)

type PosMsg struct {
	pos     int
	msgType MsgType
	values  []string
}

func NewPosHint(pos int, lines ...string) PosMsg {
	return PosMsg{pos, HINT, lines}
}

func NewPosInfo(pos int, lines ...string) PosMsg {
	return PosMsg{pos, INFO, lines}
}

func fmtMessagesWithArrows(messages []PosMsg) string {

	b := strings.Builder{}

	if len(messages) == 0 {
		return b.String()
	}

	spaces := calculateSpaces(messages)

	for pos, message := range messages {
		b.WriteString(spaces[pos])
		switch message.msgType {
		case HINT:
			b.WriteString("↑")
		case INFO:
			b.WriteString("│")
		}
	}
	b.WriteString(fmt.Sprintln())

	for len(messages) != 0 {
		for pos := range messages {
			b.WriteString(spaces[pos])
			b.WriteString("│")
		}
		b.WriteString(fmt.Sprintln())

		for pos, message := range messages {
			b.WriteString(spaces[pos])
			if pos == len(messages)-1 {
				for lineIx, line := range message.values {
					if lineIx == 0 {
						b.WriteString("└─ ")
						b.WriteString(line)
					}
					if lineIx > 0 {
						b.WriteString(fmt.Sprintln())
						b.WriteString(spaces[pos])
						b.WriteString("   ")
						b.WriteString(line)
					}
				}
			} else {
				b.WriteString("│")
			}
		}
		b.WriteString(fmt.Sprintln())

		messages = messages[:len(messages)-1]
	}
	return b.String()
}

func calculateSpaces(messages []PosMsg) []string {

	var spaces []string
	lastPos := 0
	for pos, element := range messages {
		spacesCount := element.pos - lastPos - 1
		spaces = append(spaces, "")
		for i := 0; i <= spacesCount; i++ {
			spaces[pos] += " "
		}
		lastPos = element.pos + 1
	}
	return spaces
}
