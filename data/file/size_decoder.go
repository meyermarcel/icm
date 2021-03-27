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

const sizeFileName = "size.json"

//go:embed size.json
var lengthHeightWidthJSON []byte

type size struct {
	Length      map[string]string      `json:"length"`
	HeightWidth map[string]heightWidth `json:"heightWidth"`
}

type heightWidth struct {
	Width  string `json:"height"`
	Height string `json:"width"`
}

// NewSizeDecoder writes last update lengths, height and width file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewSizeDecoder(path string) (data.LengthDecoder, data.HeightWidthDecoder, error) {
	pathToSizes := filepath.Join(path, sizeFileName)
	if err := initFile(pathToSizes, lengthHeightWidthJSON); err != nil {
		return nil, nil, err
	}
	b, err := os.ReadFile(pathToSizes)
	if err != nil {
		return nil, nil, err
	}

	var size size
	if err := json.Unmarshal(b, &size); err != nil {
		return nil, nil, err
	}
	for lengthCode := range size.Length {
		if err := cont.IsLengthCode(lengthCode); err != nil {
			return nil, nil, err
		}
	}
	for heightWidthCode := range size.HeightWidth {
		if err := cont.IsHeightWidthCode(heightWidthCode); err != nil {
			return nil, nil, err
		}
	}
	return &lengthDecoder{size.Length}, &heightWidthDecoder{size.HeightWidth}, nil
}

type lengthDecoder struct {
	lengths map[string]string
}

// Decode returns length for a given length code.
func (l *lengthDecoder) Decode(code string) (bool, cont.Length) {
	if val, ok := l.lengths[code]; ok {
		return true, cont.Length{Length: val}
	}
	return false, cont.Length{}
}

type heightWidthDecoder struct {
	heightWidths map[string]heightWidth
}

// Decode returns height and width for given height and width code.
func (hw *heightWidthDecoder) Decode(code string) (bool, cont.HeightWidth) {
	if val, ok := hw.heightWidths[code]; ok {
		return true, cont.HeightWidth{Width: val.Width, Height: val.Height}
	}
	return false, cont.HeightWidth{}
}
