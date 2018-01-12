package model

import (
	"regexp"
	"unicode/utf8"
	"strings"
)

type ParsedInput struct {
	Matches        []string
	Input          string
	MatchesIndices map[int]bool
}

func (pi ParsedInput) ToValidatedInput() ValidatedInput {
	if len(pi.Matches) == 0 {
		return NewValidatedInput("", "", "", "")
	}

	return NewValidatedInput(strings.ToUpper(sliceToString(pi.Matches[:3])),
		strings.ToUpper(pi.Matches[3]),
		strings.ToUpper(sliceToString(pi.Matches[4:10])),
		strings.ToUpper(pi.Matches[10]))
}

var Match = regexp.MustCompile(`([A-Za-z])[^A-Za-z\d]*([A-Za-z])?[^A-Za-z\d]*([A-Za-z])?[^JUZjuz\d]*([JUZjuz])?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?`)

func Parse(input string) ParsedInput {
	subMatch := Match.FindAllStringSubmatch(input, -1)

	parsedInput := ParsedInput{Input: input}

	if len(subMatch) == 0 {
		return parsedInput
	}

	parsedInput.Matches = subMatch[0][1:]

	matchRanges := [22]int{}

	copy(matchRanges[:], Match.FindAllStringSubmatchIndex(input, -1)[0][2:])

	parsedInput.MatchesIndices = byteToRuneIndex(input, matchRanges)
	return parsedInput
}

func byteToRuneIndex(input string, matchRanges [22]int) map[int]bool {
	matchesIndices := [11]int{}

	for i := 0; i < len(matchRanges)/2; i++ {
		matchesIndices[i] = matchRanges[i*2]
	}

	byteShiftsForIndices := [11]int{}

	for i := 0; i < len(input); i++ {
		if !utf8.RuneStart(input[i]) {
			for pos, element := range matchesIndices {
				if element > i {
					byteShiftsForIndices[pos]--
				}
			}
		}
	}

	// apply byte shift indices
	for pos, element := range matchesIndices {
		matchesIndices[pos] = element + byteShiftsForIndices[pos]
	}
	var matchesIndicesMap = map[int]bool{}

	for _, element := range matchesIndices {
		if element >= 0 {
			matchesIndicesMap[element] = true
		}
	}
	return matchesIndicesMap
}

func sliceToString(slice []string) string {

	value := ""
	for _, element := range slice {
		value += element
	}
	return value
}
