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
	"github.com/meyermarcel/iso6346/owner"
	"github.com/meyermarcel/iso6346/parser"
	"github.com/meyermarcel/iso6346/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse a container number",
	Long:  `Parse a container number.

Output can be formatted:

  ABC U 123456 0
     ↑ ↑      ↑
     │ │      └─ separator between serial number and check digit
     │ │
     │ └─ separator between equipment category id and serial number
     │
     └─ separator between owner code and equipment category id`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		num := parser.ParseContNum(args[0])

		num.OwnerCodeIn.Resolve(owner.Resolver(pathToDB))

		ui.PrintContNum(num, ui.Separators{
			viper.GetString(sepOwnerEquip),
			viper.GetString(sepEquipSerial),
			viper.GetString(sepSerialCheck),
		})

		if num.CheckDigitIn.IsValidCheckDigit {
			os.Exit(0)
		}
		os.Exit(1)
	},
}

func init() {

	RootCmd.AddCommand(parseCmd)
}
