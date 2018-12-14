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
	"errors"
	"testing"
)

func TestInputHasCorrectValue(t *testing.T) {

	match1 := Input{
		runeCount: 1,
		matchIndex: func(in string) []int {
			return []int{0, 1}
		},
		validate: func(value string, previousValues []string) (error, []Info) {
			return nil, []Info{{Text: "match 1"}}
		},
	}
	match2 := Input{
		runeCount: 2,
		toUpper:   true,
		matchIndex: func(in string) []int {
			return []int{0, 2}
		},
		validate: func(value string, previousValues []string) (error, []Info) {
			return nil, []Info{{Text: "match 2"}}
		},
	}
	validFmtButInvalidMatch := Input{
		runeCount: 1,
		matchIndex: func(in string) []int {
			return []int{0, 1}
		},
		validate: func(value string, previousValues []string) (error, []Info) {
			return errors.New(""), nil
		},
	}
	noMatch := Input{
		matchIndex: func(in string) []int {
			return nil
		},
		validate: func(value string, previousValues []string) (error, []Info) {
			return errors.New(""), nil
		},
	}

	type wantedInput struct {
		value          string
		previousValues []string
		err            bool
		infoTexts      []string
	}

	tests := []struct {
		name         string
		inputOrders  [][]Input
		in           string
		wantedInputs []wantedInput
		wantErr      bool
	}{
		{
			"Match single value",
			[][]Input{
				{match1},
			},
			"a",
			[]wantedInput{
				{
					"a",
					nil,
					false,
					[]string{"match 1"},
				},
			},
			false,
		},
		{
			"Match multiple values",
			[][]Input{
				{match1, match2},
			},
			"abcd",
			[]wantedInput{
				{
					"a",
					nil,
					false,
					[]string{"match 1"},
				},
				{
					"BC",
					[]string{"a"},
					false,
					[]string{"match 2"},
				},
			},
			false,
		},
		{
			"Use first best match",
			[][]Input{
				{noMatch, match1},
				{match1},
				{match2},
			},
			"abcd",
			[]wantedInput{
				{
					"a",
					nil,
					false,
					[]string{"match 1"},
				},
			},
			false,
		},
		{
			"First match is default",
			[][]Input{
				{match1, noMatch, match2},
				{noMatch},
			},
			"abcd",
			[]wantedInput{
				{
					"a",
					nil,
					false,
					[]string{"match 1"},
				},
				{
					"",
					[]string{"a"},
					true,
					nil,
				},
				{
					"BC",
					[]string{"", "a"},
					false,
					[]string{"match 2"},
				},
			},
			true,
		},
		{
			"Match but invalid",
			[][]Input{
				{noMatch},
				{validFmtButInvalidMatch},
			},
			"a",
			[]wantedInput{
				{
					"a",
					nil,
					true,
					nil,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				tt.inputOrders,
			}
			inputs, err := v.Validate(tt.in)
			if len(inputs) != len(tt.wantedInputs) {
				t.Errorf("inputs len %v, want %v", len(inputs), len(tt.wantedInputs))
			}
			if (err == nil) == tt.wantErr {
				t.Errorf("got = %v, wantErr is %v", err, tt.wantErr)
			}
			for i, input := range inputs {
				if input.value != tt.wantedInputs[i].value {
					t.Errorf("value is %v, want %v", input.value, tt.wantedInputs[i].value)
				}
				if len(input.previousValues) != len(tt.wantedInputs[i].previousValues) {
					t.Errorf("previousValues len %v, want %v", len(inputs), len(tt.wantedInputs[i].previousValues))
				}
				if input.previousValues != nil {
					for j, previousValue := range input.previousValues {
						if previousValue != tt.wantedInputs[i].previousValues[j] {
							t.Errorf("previous value is %v, want %v", previousValue, tt.wantedInputs[i].previousValues[j])
						}
					}
				}
				if (input.err != nil) != tt.wantedInputs[i].err {
					t.Errorf("err is %v, want %v", input.err != nil, tt.wantedInputs[i].err)
				}
				if len(input.infos) != len(tt.wantedInputs[i].infoTexts) {
					t.Errorf("input infos len %v, want %v", len(input.infos), len(tt.wantedInputs[i].infoTexts))
				}
				if input.infos != nil {
					for j, info := range input.infos {
						if info.Text != tt.wantedInputs[i].infoTexts[j] {
							t.Errorf("info text is %v, want %v", info.Text, tt.wantedInputs[i].infoTexts[j])
						}
					}
				}
			}
		})
	}
}
