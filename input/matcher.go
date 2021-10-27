package input

// Match returns pattern if all values are valid formatted. If no pattern
// meets the requirement the first pattern is returned.
func Match(in string, newPatterns [][]func() Input) []func() Input {
	for _, newInputs := range newPatterns {
		inTemp := in
		allValidFmt := true
		for _, newInput := range newInputs {
			input := newInput()
			matchIndex := input.matchIndex(inTemp)
			if matchIndex != nil {
				input.value = inTemp[matchIndex[0]:matchIndex[1]]
				inTemp = inTemp[matchIndex[1]:]
			}
			allValidFmt = allValidFmt && input.isValidFmt()
		}
		if allValidFmt {
			return newInputs
		}
	}
	return newPatterns[0]
}
