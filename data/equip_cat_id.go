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

	"github.com/meyermarcel/icm/cont"
	"github.com/meyermarcel/icm/utils"
)

const equipCatIDsFileName = "equipment-category-ids.json"
const equipCatIDsJSON = `{
  "U": "freight container",
  "J": "detachable freight container-related equipment",
  "Z": "trailer and chassis"
}
`

var loadedEquipCatIDs map[string]string

// InitEquipCatIDsData writes equipment category file to path if it not exists and loads its data to memory.
func InitEquipCatIDsData(path string) {
	pathToEquipCatIDs := utils.InitFile(filepath.Join(path, equipCatIDsFileName), []byte(equipCatIDsJSON))
	b, err := ioutil.ReadFile(pathToEquipCatIDs)
	utils.CheckErr(err)
	utils.JSONUnmarshal(b, &loadedEquipCatIDs)
}

// GetEquipCat returns equipment category for given code.
func GetEquipCat(code string) cont.EquipCat {
	info := loadedEquipCatIDs[code]

	return cont.NewEquipCatID(cont.NewEquipCatIDFrom(code), info)
}

// GetEquipCatIDs returns all equipment category IDs.
func GetEquipCatIDs() []string {
	keys := make([]string, 0, len(loadedEquipCatIDs))
	for k := range loadedEquipCatIDs {
		keys = append(keys, k)
	}
	return keys
}
