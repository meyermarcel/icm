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
	"path/filepath"
	"strings"
)

const sizesFileName = "sizes.json"

type length struct {
	Length string
}

type heightAndWidth struct {
	Width  string `json:"height"`
	Height string `json:"width"`
}

var loadedCfgSizes sizes

type sizes struct {
	Length         map[string]string         `json:"length"`
	HeightAndWidth map[string]heightAndWidth `json:"heightAndWidth"`
}

func initCfgSizes(appDirPath string) {
	pathToSizes := initFile(filepath.Join(appDirPath, sizesFileName), []byte(lengthWidthAndHeightJSON))
	jsonUnmarshal(readFile(pathToSizes), &loadedCfgSizes)
}

func getRegexPartLengths() string {
	keys := make([]string, 0, len(loadedCfgSizes.Length))
	for k := range loadedCfgSizes.Length {
		keys = append(keys, k)
	}
	return strings.Join(keys, "")
}

func getRegexPartHeightAndWidths() string {
	keys := make([]string, 0, len(loadedCfgSizes.HeightAndWidth))
	for k := range loadedCfgSizes.HeightAndWidth {
		keys = append(keys, k)
	}
	return strings.Join(keys, "")
}

func getLength(code string) length {

	mappedLength := loadedCfgSizes.Length[code]
	return length{mappedLength}
}

func getHeightAndWidth(code string) heightAndWidth {
	return loadedCfgSizes.HeightAndWidth[code]
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
