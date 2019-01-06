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
