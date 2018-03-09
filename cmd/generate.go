// Copyright © 2017 Marcel Meyer meyer@synyx.de
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

import (
	"github.com/meyermarcel/iso6346/cont"
	"github.com/meyermarcel/iso6346/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate a random container number",
	Long: `
Generate a random container number with ISO 6346 specification.
Only real owners are used. Serial number is pseudo random.
Multiple generated container numbers are unique.

Output can be formatted:

  ABC U 123456 0
     ↑ ↑      ↑
     │ │      └─ separator between serial number and check digit
     │ │
     │ └─ separator between equipment category id and serial number
     │
     └─ separator between owner code and equipment category id`,
	Example: `
  iso6346 generate

  iso6346 generate --count 5000

  iso6346 generate --sep-owner-equip '' --sep-serial-check '-'
`,
	Args: cobra.NoArgs,
	Run:  generate,
}

func generate(cmd *cobra.Command, args []string) {
	c := make(chan cont.Number)
	go cont.Gen(pathToDB, count, c)

	for contNum := range c {
		ui.PrintGen(contNum, ui.Separators{
			viper.GetString(sepOwnerEquip),
			viper.GetString(sepEquipSerial),
			viper.GetString(sepSerialCheck),
		})
	}
	os.Exit(0)
}

var count int

func init() {
	generateCmd.Flags().IntVarP(&count, "count", "c", 1, "count of container numbers")
	RootCmd.AddCommand(generateCmd)
}
