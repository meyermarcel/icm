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

package main

import "path/filepath"

const groupsFileName = "groups.json"

const groupsJSON = `{
  "A": "(unassigned)",
  "B": "(unassigned)",
  "G": "(unassigned)",
  "H": "(unassigned)",
  "K": "(unassigned)",
  "N": "(unassigned)",
  "P": "(unassigned)",
  "R": "(unassigned)",
  "S": "(unassigned)",
  "U": "(unassigned)",
  "V": "(unassigned)"
}
`

type group struct {
	Code string
	Info string
}

var loadedCfgGroups map[string]string

func initCfgGroups(appDirPath string) {
	pathToGroups := initFile(filepath.Join(appDirPath, groupsFileName), []byte(groupsJSON))
	jsonUnmarshal(readFile(pathToGroups), &loadedCfgGroups)
}

func getGroup(code string) (group, bool) {
	info, exists := loadedCfgGroups[code]

	return group{code, info}, exists
}
