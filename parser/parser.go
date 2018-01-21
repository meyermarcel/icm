package parser

import (
	"regexp"
	"unicode/utf8"
	"strings"
)

var MatchContainerNumber = regexp.MustCompile(`([A-Za-z])[^A-Za-z\d]*([A-Za-z])?[^A-Za-z\d]*([A-Za-z])?[^JUZjuz\d]*([JUZjuz])?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?`)
var MatchOwnerCodeOptEquipCatId = regexp.MustCompile(`([A-Za-z])[^A-Za-z\d]*([A-Za-z])?[^A-Za-z\d]*([A-Za-z])?[^JUZjuz\d]*([JUZjuz])?`)

type Input struct {
	matches        []string
	input          string
	matchesIndices map[int]bool
}

func (pi Input) GetMatch(start, endExclusive int) string {

	if len(pi.matches) == 0 {
		return ""
	}

	value := ""
	for _, element := range pi.matches[start:endExclusive] {
		value += element
	}
	return strings.ToUpper(value)
}

func (pi Input) GetMatchSingle(pos int) string {

	return pi.GetMatch(pos, pos+1)
}

func (pi Input) NoMatch() bool {
	return len(pi.matches) == 0
}

func (pi Input) Input() string {
	return pi.input
}

func (pi Input) Match(atPosition int) bool {
	return pi.matchesIndices[atPosition]
}

func Parse(input string, match regexp.Regexp) Input {

	parsedInput := Input{input: input}

	subMatch := match.FindAllStringSubmatch(input, -1)

	if len(subMatch) == 0 {
		return parsedInput
	}

	parsedInput.matches = subMatch[0][1:]

	matchRanges := [22]int{}

	copy(matchRanges[:], MatchContainerNumber.FindAllStringSubmatchIndex(input, -1)[0][2:])

	parsedInput.matchesIndices = byteToRuneIndex(input, matchRanges)

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
