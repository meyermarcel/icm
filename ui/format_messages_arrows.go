package ui

import "fmt"

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

func formatMessagesWithArrows(messages []PosMsg) string {

	out := ""

	if len(messages) == 0 {
		return out
	}

	spaces := calculateSpaces(messages)

	for pos, message := range messages {
		out += spaces[pos]
		switch message.msgType {
		case HINT:
			out += "↑"
		case INFO:
			out += "│"
		}
	}
	out += fmt.Sprintln()

	for len(messages) != 0 {
		for pos := range messages {
			out += spaces[pos]
			out += "│"
		}
		out += fmt.Sprintln()

		for pos, message := range messages {
			out += spaces[pos]
			if pos == len(messages)-1 {
				for lineIx, line := range message.values {
					if lineIx == 0 {
						out += "└─ " + line
					}
					if lineIx > 0 {
						out += "\n"
						out += spaces[pos] + "   " + line
					}
				}
			} else {
				out += "│"
			}
		}
		out += fmt.Sprintln()

		messages = messages[:len(messages)-1]
	}
	return out
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
