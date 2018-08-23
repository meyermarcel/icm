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
	"path/filepath"

	"github.com/meyermarcel/icm/internal/data"

	"io/ioutil"

	"github.com/meyermarcel/icm/internal/cont"
)

const sizeFileName = "size.json"

type size struct {
	Length         map[string]string         `json:"length"`
	HeightAndWidth map[string]heightAndWidth `json:"heightAndWidth"`
}

type heightAndWidth struct {
	Width  string `json:"height"`
	Height string `json:"width"`
}

// NewLengthDecoder writes last update lengths, height and width file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewLengthDecoder(path string) (data.LengthDecoder, error) {
	pathToSizes := filepath.Join(path, sizeFileName)
	if err := initFile(pathToSizes, []byte(lengthWidthAndHeightJSON)); err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(pathToSizes)
	if err != nil {
		return nil, err
	}

	var size size
	if err := json.Unmarshal(b, &size); err != nil {
		return nil, err
	}
	return &lengthDecoder{size.Length}, nil
}

type lengthDecoder struct {
	lengths map[string]string
}

// Decode returns length for a given length code.
func (l *lengthDecoder) Decode(code string) cont.Length {
	return cont.Length{Length: l.lengths[code]}
}

// AllCodes returns all length codes.
func (l *lengthDecoder) AllCodes() []string {
	keys := make([]string, 0, len(l.lengths))
	for k := range l.lengths {
		keys = append(keys, k)
	}
	return keys
}

// NewHeightAndWidthDecoder initializes the file of lengths, height and width and
// returns a new height and width file data source.
func NewHeightAndWidthDecoder(path string) (data.HeightAndWidthDecoder, error) {
	var size size
	pathToSizes := filepath.Join(path, sizeFileName)
	if err := initFile(pathToSizes, []byte(lengthWidthAndHeightJSON)); err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(pathToSizes)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &size); err != nil {
		return nil, err
	}
	return &heightAndWidthDecoder{size.HeightAndWidth}, err
}

type heightAndWidthDecoder struct {
	heightAndWidths map[string]heightAndWidth
}

// Decode returns height and width for given height and width code.
func (hw *heightAndWidthDecoder) Decode(code string) cont.HeightAndWidth {
	heightAndWidth := hw.heightAndWidths[code]
	return cont.HeightAndWidth{Width: heightAndWidth.Width, Height: heightAndWidth.Height}
}

// AllCodes returns all height and width codes.
func (hw *heightAndWidthDecoder) AllCodes() []string {
	keys := make([]string, 0, len(hw.heightAndWidths))
	for k := range hw.heightAndWidths {
		keys = append(keys, k)
	}
	return keys
}

const lengthWidthAndHeightJSON = `{
  "length": {
    "1": "2991 mm",
    "2": "6068 mm",
    "3": "9125 mm",
    "4": "12192 mm",
    "A": "7150 mm",
    "B": "7315 mm",
    "C": "7430 mm",
    "D": "7450 mm",
    "E": "7820 mm",
    "F": "8100 mm",
    "G": "12500 mm",
    "H": "13106 mm",
    "K": "13600 mm",
    "L": "13716 mm",
    "M": "14630 mm",
    "N": "14935 mm",
    "P": "16154 mm"
  },
  "heightAndWidth": {
    "0": {
      "width": "2436 mm",
      "height": "2438 mm"
    },
    "2": {
      "width": "2436 mm",
      "height": "2591 mm"
    },
    "4": {
      "width": "2436 mm",
      "height": "2743 mm"
    },
    "5": {
      "width": "2436 mm",
      "height": "2895 mm"
    },
    "6": {
      "width": "2436 mm",
      "height": "> 2895 mm"
    },
    "8": {
      "width": "2436 mm",
      "height": "1295 mm"
    },
    "9": {
      "width": "2436 mm",
      "height": "< 1219 mm"
    },
    "C": {
      "width": "> 2438 mm and ≤ 2500 mm",
      "height": "2591 mm"
    },
    "D": {
      "width": "> 2438 mm and ≤ 2500 mm",
      "height": "2743 mm"
    },
    "E": {
      "width": "> 2438 mm and ≤ 2500 mm",
      "height": "2895 mm"
    },
    "F": {
      "width": "> 2438 mm and ≤ 2500 mm",
      "height": "> 2895 mm"
    },
    "L": {
      "width": "> 2500 mm",
      "height": "2591 mm"
    },
    "M": {
      "width": "> 2500 mm",
      "height": "2743 mm"
    },
    "N": {
      "width": "> 2500 mm",
      "height": "2895 mm"
    },
    "P": {
      "width": "> 2500 mm",
      "height": "> 2895 mm"
    }
  }
}
`
