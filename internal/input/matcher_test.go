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
	"testing"
)

func TestMatcher_Match(t *testing.T) {

	match1 := Input{
		runeCount: 1,
		matchIndex: func(in string) []int {
			return []int{0, 1}
		},
	}
	match2 := Input{
		runeCount: 2,
		matchIndex: func(in string) []int {
			return []int{0, 2}
		},
	}
	noMatch := Input{
		matchIndex: func(in string) []int {
			return nil
		},
	}

	tests := []struct {
		name          string
		inputPatterns [][]Input
		in            string
		wantedLen     int
	}{
		{
			"Use first pattern",
			[][]Input{
				{match1},
				{match1, match2},
			},
			"a",
			1,
		},
		{
			"Use first pattern as default",
			[][]Input{
				{noMatch},
				{match2, noMatch},
			},
			"abcd",
			1,
		},
		{
			"Use first best match",
			[][]Input{
				{noMatch},
				{match1, noMatch},
				{match1, match1, match1},
			},
			"abcd",
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if patternIdx := NewMatcher(tt.inputPatterns).Match(tt.in); len(patternIdx) != tt.wantedLen {
				t.Errorf("Match() = %v, want length %v", patternIdx, tt.wantedLen)
			}
		})
	}
}
