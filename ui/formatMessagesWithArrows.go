package ui

type PositionedMessage struct {
	pos   int
	value string
}

func formatMessagesWithArrows(messages []PositionedMessage) string {

	out := ""

	if len(messages) == 0 {
		return out
	}

	spaces := calculateSpaces(messages)

	for pos := range messages {
		out += spaces[pos]
		out += "↑"
	}
	out += "\n"

	for len(messages) != 0 {
		for pos := range messages {
			out += spaces[pos]
			out += "│"
		}
		out += "\n"

		for pos, message := range messages {
			out += spaces[pos]
			if pos == len(messages)-1 {
				out += "└─ " + message.value
			} else {
				out += "│"
			}
		}
		out += "\n"

		messages = messages[:len(messages)-1]
	}
	return out
}

func calculateSpaces(messages []PositionedMessage) []string {

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
