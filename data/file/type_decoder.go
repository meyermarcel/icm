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
	"os"
	"path/filepath"

	// needed for package embed
	_ "embed"

	"github.com/meyermarcel/icm/cont"
	"github.com/meyermarcel/icm/data"
)

const typeFileName = "type.json"

//go:embed type.json
var typeJSON []byte

const groupFileName = "group.json"

//go:embed group.json
var groupJSON []byte

type typeAndGroupDecoder struct {
	types  map[string]string
	groups map[string]string
}

// NewTypeDecoder writes type and group file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewTypeDecoder(path string) (data.TypeDecoder, error) {
	typeAndGroup := &typeAndGroupDecoder{}
	pathToType := filepath.Join(path, typeFileName)
	if err := initFile(pathToType, typeJSON); err != nil {
		return nil, err
	}
	b, err := os.ReadFile(pathToType)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &typeAndGroup.types); err != nil {
		return nil, err
	}

	pathToGroup := filepath.Join(path, groupFileName)
	if err := initFile(pathToGroup, groupJSON); err != nil {
		return nil, err
	}
	b, err = os.ReadFile(pathToGroup)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &typeAndGroup.groups); err != nil {
		return nil, err
	}
	for typeCode := range typeAndGroup.types {
		if err := cont.IsTypeCode(typeCode); err != nil {
			return nil, err
		}
	}
	return typeAndGroup, nil
}

// Decode returns type and group for type code.
func (tg *typeAndGroupDecoder) Decode(code string) (bool, cont.TypeAndGroup) {
	typeAndGroup := cont.TypeAndGroup{}

	typeInfo, exists := tg.types[code]
	typeValue, typeFound := cont.Type{TypeCode: code, TypeInfo: typeInfo}, exists

	info, exists := tg.groups[string(code[0])]
	group, groupFound := cont.Group{GroupCode: code, GroupInfo: info}, exists

	if !typeFound || !groupFound {
		return false, typeAndGroup
	}

	typeAndGroup.Type = typeValue
	typeAndGroup.Group = group
	return true, typeAndGroup
}
