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

package cmd

const ymlCfgName = "config"
const ymlCfgFileName = ymlCfgName + ".yml"

func cfgSeparators() []byte {
	return []byte(`# Config
#
#  Separators
#
#  ABC U 123456 0   20 G1
#     ↑ ↑      ↑  ↑   ↑
#     │ │      │  │   └─ ` + sepST + `
#     │ │      │  │
#     │ │      │  └─ ` + sepCS + `
#     │ │      │
#     │ │      └─ ` + sepSC + `
#     │ │
#     │ └─ ` + sepES + `
#     │
#     └─ ` + sepOE + `
#
` + sepOE + `: ' '
` + sepES + `: ' '
` + sepSC + `: ' '
` + sepCS + `: '   '
` + sepST + `: ' '
`)
}
