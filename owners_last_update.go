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
	"time"
)

const ownersLastUpdateFileName = "owners-last-update"
const ownersLastUpdate = "2018-06-15" + "\n"

var loadedOwnersLastUpdate string
var pathToOwnersLastUpdate string

func initOwnersLastUpdate(appDirPath string) {
	pathToOwnersLastUpdate = filepath.Join(appDirPath, ownersLastUpdateFileName)
	initFile(pathToOwnersLastUpdate, []byte(ownersLastUpdate))
	loadedOwnersLastUpdate = string(readFile(pathToOwnersLastUpdate))
}

func getOwnersLastUpdate() time.Time {
	dateString := strings.TrimSuffix(loadedOwnersLastUpdate, "\n")

	date, err := time.Parse(dateFormat, dateString)
	if err != nil {
		log.Fatal("Cannot parse time ", dateString, ":", err)
	}
	return date
}

func saveNowForOwnersLastUpdate() {
	writeFile(pathToOwnersLastUpdate, []byte(time.Now().Format(dateFormat)+"\n"))
}
