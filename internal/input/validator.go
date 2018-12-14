// Copyright © 2018 Marcel Meyer meyermarcel@posteo.de
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package input

import (
	"strings"
	"unicode/utf8"
)

// Validator validates multiple input patterns.
// Use NewValidator for instantiation.
type Validator struct {
	inputOrders [][]Input
}

// NewValidator returns a new Validator.
func NewValidator(inputOrders [][]Input) *Validator {
	return &Validator{inputOrders: inputOrders}
}

// Validate validates multiple input patterns starting with first pattern in slice.
// If pattern matches this pattern is returned. If no pattern matches, first pattern is returned.
func (v *Validator) Validate(in string) ([]Input, error) {

	var defaultErr error
	defaultInputs := v.inputOrders[0]

	previousValues := make([]string, 0)
	for orderIdx, order := range v.inputOrders {
		inTemp := in
		previousValues = nil
		allValidFmt := true
		for inputIdx, input := range order {
			order[inputIdx].previousValues = previousValues

			matchIndex := input.matchIndex(inTemp)
			if matchIndex != nil {
				matchPart := inTemp[matchIndex[0]:matchIndex[1]]
				if order[inputIdx].toUpper {
					matchPart = strings.ToUpper(matchPart)
				}
				order[inputIdx].value = matchPart
				inTemp = inTemp[matchIndex[1]:]
			}

			previousValues = append([]string{order[inputIdx].value}, previousValues...)

			if orderIdx == 0 {
				order[inputIdx].validateValue()
				if defaultErr == nil {
					defaultErr = order[inputIdx].err
				}
			}
			allValidFmt = allValidFmt && order[inputIdx].isValidFmt()
		}
		if allValidFmt {
			var err error
			for inputIdx := range order {
				order[inputIdx].validateValue()
				if err == nil {
					err = order[inputIdx].err
				}
			}
			return order, err
		}
	}
	return defaultInputs, defaultErr
}

// Input is a structured part of an input string.
type Input struct {
	runeCount      int
	matchIndex     func(in string) []int
	validate       func(value string, previousValues []string) (error, []Info)
	toUpper        bool
	value          string
	previousValues []string
	err            error
	infos          []Info
}

// SetToUpper converts the matched value to upper case.
func (i *Input) SetToUpper() {
	i.toUpper = true
}

// NewInput returns a new Input.
func NewInput(runeCount int,
	matchIndex func(in string) []int,
	validate func(value string, previousValues []string) (error, []Info)) Input {
	return Input{runeCount: runeCount, matchIndex: matchIndex, validate: validate}
}

func (i *Input) validateValue() {
	i.err, i.infos = i.validate(i.value, i.previousValues)
}

func (i *Input) isValidFmt() bool {
	if i.runeCount == 0 {
		return false
	}
	return utf8.RuneCountInString(i.value) == i.runeCount
}

// Info is structured information with formatted text.
type Info struct {
	Title string
	Text  string
}
