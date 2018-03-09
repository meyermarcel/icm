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
	"os"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random container number",
	Long: `Example:

  iso6346 generate

generates a random container number with valid check digit:

  ABC U 123456 0

The ISO 6346 standard specifies that all characters are
alphanumeric except the equipment ID which is U, J or Z.

You can also format your output:

  iso6346 generate --sep-owner-equip '' --sep-serial-check '-'

generates a formatted random container number:

  ABCU 123456-0


  ABC U 123456 0
     ↑ ↑      ↑
     │ │      └─ separator between serial number and check digit
     │ │
     │ └─ separator between equipment category id and serial number
     │
     └─ separator between owner code and equipment category id`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan cont.Number)
		go cont.Gen(count, c)

		for contNum := range c {
			ui.PrintGen(contNum, ui.Separators{sepOwnerEquip, sepEquipSerial, sepSerialCheck})
		}
		os.Exit(0)
	},
}

var count int

func init() {
	generateCmd.Flags().IntVarP(&count, "count", "c", 1, "generate container numbers")
	RootCmd.AddCommand(generateCmd)
}
