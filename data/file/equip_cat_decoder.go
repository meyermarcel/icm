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

package file

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/meyermarcel/icm/data"

	"github.com/meyermarcel/icm/cont"
)

const equipCatIDsFileName = "equipment-category-id.json"
const equipCatIDsJSON = `{
  "U": "freight container",
  "J": "detachable freight container-related equipment",
  "Z": "trailer and chassis"
}
`

// NewEquipCatDecoder writes equipment category ID file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewEquipCatDecoder(path string) (data.EquipCatDecoder, error) {
	equipCat := &equipCatDecoder{}
	pathToEquipCat := filepath.Join(path, equipCatIDsFileName)
	if err := initFile(pathToEquipCat, []byte(equipCatIDsJSON)); err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(pathToEquipCat)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &equipCat.categories); err != nil {
		return nil, err
	}
	for ID := range equipCat.categories {
		if err := cont.IsEquipCatID(ID); err != nil {
			return nil, err
		}
	}
	return equipCat, err
}

type equipCatDecoder struct {
	categories map[string]string
}

// Decode decodes ID to equipment category ID.
func (ec *equipCatDecoder) Decode(ID string) (bool, cont.EquipCat) {
	if val, ok := ec.categories[ID]; ok {
		return true, cont.NewEquipCatID(ID, val)
	}
	return false, cont.EquipCat{}
}

// AllCatIDs returns all equipment category IDs.
func (ec *equipCatDecoder) AllCatIDs() []string {
	keys := make([]string, 0, len(ec.categories))
	for k := range ec.categories {
		keys = append(keys, k)
	}
	return keys
}
