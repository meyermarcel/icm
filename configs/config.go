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

package configs

import "fmt"

// Name of the config files and keys for configuration and flags.
const (
	Name           = "config"
	NameWithYmlExt = Name + ".yml"
	Pattern        = "pattern"
	PatternDefVal  = "auto"
	NoHeader       = "no-header"
	NoHeaderDefVal = false
	Output         = "output"
	OutputDefVal   = "auto"
	SepOE          = "sep-owner-equip"
	SepOEDefVal    = " "
	SepES          = "sep-equip-serial"
	SepESDefVal    = " "
	SepSC          = "sep-serial-check"
	SepSCDefVal    = " "
	SepCS          = "sep-check-size"
	SepCSDefVal    = "   "
	SepST          = "sep-size-type"
	SepSTDefVal    = " "
)

// Cfg returns default config.
func Cfg() []byte {
	return []byte(`# Pattern matching mode
#                     auto = matches automatically a pattern
#         container-number = matches a container number
#                    owner = matches a three letter owner code
# owner-equipment-category = matches a three letter owner code with equipment category ID
#                size-type = matches length, width+height and type code
` + Pattern + `: ` + PatternDefVal + `

# Output mode
#  auto = for a single line 'fancy' and for multiple lines 'csv' output 
#   csv = machine readable CSV output
# fancy = human readable fancy output
` + Output + `: ` + OutputDefVal + `

# No header for CSV output
` + NoHeader + `: ` + fmt.Sprintf("%t", NoHeaderDefVal) + `

#  Separators
#
#  ABC U 123456 0   20 G1
#     ↑ ↑      ↑  ↑   ↑
#     │ │      │  │   └─ ` + SepST + `
#     │ │      │  │
#     │ │      │  └─ ` + SepCS + `
#     │ │      │
#     │ │      └─ ` + SepSC + `
#     │ │
#     │ └─ ` + SepES + `
#     │
#     └─ ` + SepOE + `
#
` + SepOE + `:  '` + SepOEDefVal + `'
` + SepES + `: '` + SepESDefVal + `'
` + SepSC + `: '` + SepSCDefVal + `'
` + SepCS + `:   '` + SepCSDefVal + ` '
` + SepST + `:    '` + SepSTDefVal + `'
`)
}
