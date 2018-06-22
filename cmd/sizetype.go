// Copyright © 2017 Marcel Meyer meyermarcel@posteo.de
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
	"os"

	"github.com/meyermarcel/iso6346/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sizeTypeCmd = &cobra.Command{
	Use:   "sizetype",
	Short: "Validate or print size and type",
	Long:  "Validate or print size and type.",
}

var sizeTypeValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate size and type codes",
	Long: `Validate size and type codes.

` + sepHelp,
	Example: `  ` + appName + ` sizetype validate '20G1'

  ` + appName + ` sizetype validate --` + sepST + ` '' '20G1'`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		valid := ui.ParseAndPrintSizeType(args[0], viper.GetString(sepST))

		if valid {
			os.Exit(0)
		}
		os.Exit(1)
	},
}

func init() {
	sizeTypeValidateCmd.Flags().String(sepST, "",
		"20(*)G1  (*) separates size and type")

	viper.BindPFlag(sepST, sizeTypeValidateCmd.Flags().Lookup(sepST))

	sizeTypeCmd.AddCommand(sizeTypeValidateCmd)
	RootCmd.AddCommand(sizeTypeCmd)
}
