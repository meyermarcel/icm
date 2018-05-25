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

import (
	"log"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

const equipCatIDsFileName = "equipment-category-ids.json"
const equipCatIDsJSON = `{
  "U": "freight container",
  "J": "detachable freight container-related equipment",
  "Z": "trailer and chassis"
}
`

var loadedCfgEquipCatIDs map[string]string

func initCfgEquipCatIDs(appDirPath string) {
	pathToEquipCatIDs := initFile(filepath.Join(appDirPath, equipCatIDsFileName), []byte(equipCatIDsJSON))
	jsonUnmarshal(readFile(pathToEquipCatIDs), &loadedCfgEquipCatIDs)
}

type equipCatID struct {
	Value string
	Info  string
}

func newEquipCatIDU() equipCatID {
	return equipCatID{"U", ""}
}

func newEquipCatIDFrom(value string) equipCatID {

	if utf8.RuneCountInString(value) != 1 {
		log.Fatalf("'%s' is not one character", value)
	}
	return equipCatID{value, ""}
}

func getEquipCatID(code string) equipCatID {
	info := loadedCfgEquipCatIDs[code]

	return equipCatID{code, info}
}

func getRegexPartEquipCatIDs() string {
	keys := make([]string, 0, len(loadedCfgEquipCatIDs))
	for k := range loadedCfgEquipCatIDs {
		keys = append(keys, k)
	}
	return strings.Join(keys, "")
}

func getEquipCatIDs() []string {

	keys := make([]string, 0, len(loadedCfgEquipCatIDs))
	for k := range loadedCfgEquipCatIDs {
		keys = append(keys, k)
	}
	return keys
}
