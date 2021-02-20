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

	"github.com/meyermarcel/icm/data"

	"io/ioutil"

	"github.com/meyermarcel/icm/cont"
)

const sizeFileName = "size.json"

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
	if err := initFile(pathToSizes, []byte(lengthHeightWidthJSON)); err != nil {
		return nil, nil, err
	}
	b, err := ioutil.ReadFile(pathToSizes)
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

const lengthHeightWidthJSON = `{
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
  "heightWidth": {
    "0": {
      "height": "2438 mm",
      "width": "2436 mm"
    },
    "2": {
      "height": "2591 mm",
      "width": "2436 mm"
    },
    "4": {
      "height": "2743 mm",
      "width": "2436 mm"
    },
    "5": {
      "height": "2895 mm",
      "width": "2436 mm"
    },
    "6": {
      "height": "> 2895 mm",
      "width": "2436 mm"
    },
    "8": {
      "height": "1295 mm",
      "width": "2436 mm"
    },
    "9": {
      "height": "< 1219 mm",
      "width": "2436 mm"
    },
    "C": {
      "height": "2591 mm",
      "width": "> 2438 mm and ≤ 2500 mm"
    },
    "D": {
      "height": "2743 mm",
      "width": "> 2438 mm and ≤ 2500 mm"
    },
    "E": {
      "height": "2895 mm",
      "width": "> 2438 mm and ≤ 2500 mm"
    },
    "F": {
      "height": "> 2895 mm",
      "width": "> 2438 mm and ≤ 2500 mm"
    },
    "L": {
      "height": "2591 mm",
      "width": "> 2500 mm"
    },
    "M": {
      "height": "2743 mm",
      "width": "> 2500 mm"
    },
    "N": {
      "height": "2895 mm",
      "width": "> 2500 mm"
    },
    "P": {
      "height": "> 2895 mm",
      "width": "> 2500 mm"
    }
  }
}
`
