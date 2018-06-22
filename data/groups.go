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

package data

import (
	"path/filepath"

	"io/ioutil"

	"github.com/meyermarcel/iso6346/iso6346"
	"github.com/meyermarcel/iso6346/utils"
)

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

var loadedGroups map[string]string

// InitGroupsData writes group file to path if it not exists and loads its data to memory.
func InitGroupsData(path string) {
	pathToGroups := utils.InitFile(filepath.Join(path, groupsFileName), []byte(groupsJSON))
	b, err := ioutil.ReadFile(pathToGroups)
	utils.CheckErr(err)
	utils.JSONUnmarshal(b, &loadedGroups)
}

func getGroup(code string) (iso6346.TypeGroup, bool) {
	info, exists := loadedGroups[code]

	return iso6346.TypeGroup{Code: code, Info: info}, exists
}
