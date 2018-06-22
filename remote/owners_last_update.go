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

package remote

import (
	"path/filepath"
	"strings"
	"time"

	"io/ioutil"

	"github.com/meyermarcel/iso6346/utils"
)

const dateFormat = "2006-01-02"

const ownersLastUpdateFileName = "owners-last-update"
const ownersLastUpdate = "2018-06-15" + "\n"

var loadedOwnersLastUpdate string
var pathToOwnersLastUpdate string

// InitOwnersLastUpdate writes last update file to path if it not exists and loads its data to memory.
func InitOwnersLastUpdate(appDirPath string) {
	pathToOwnersLastUpdate = filepath.Join(appDirPath, ownersLastUpdateFileName)
	utils.InitFile(pathToOwnersLastUpdate, []byte(ownersLastUpdate))
	b, err := ioutil.ReadFile(pathToOwnersLastUpdate)
	utils.CheckErr(err)
	loadedOwnersLastUpdate = string(b)
}

func getOwnersLastUpdate() time.Time {
	dateString := strings.TrimSuffix(loadedOwnersLastUpdate, "\n")

	date, err := time.Parse(dateFormat, dateString)
	utils.CheckErr(err)
	return date
}

func saveNowForOwnersLastUpdate() {
	err := ioutil.WriteFile(pathToOwnersLastUpdate, []byte(time.Now().Format(dateFormat)+"\n"), 0644)
	utils.CheckErr(err)
}
