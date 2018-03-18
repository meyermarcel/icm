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
	"github.com/meyermarcel/iso6346/owner"
	"github.com/meyermarcel/iso6346/parser"
	"github.com/meyermarcel/iso6346/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"github.com/meyermarcel/iso6346/sizetype"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a container number",
	Long: `Validate a container number.

` + sepHelp,
	Example: `  ` + appName + ` validate 'ABCU 1234560'

  ` + appName + ` validate --` + sepOE + ` '' --` + sepSC + ` '' 'ABCU 1234560'`,
	Args: cobra.ExactArgs(1),
	Run:  validateContNum,
}

func init() {
	validateCmd.Flags().String(sepOE, "",
		"ABC(*)U1234560   20G1  (*) separator between owner code and equipment category id")
	validateCmd.Flags().String(sepES, "",
		"ABCU(*)1234560   20G1  (*) separator between equipment category id and serial number")
	validateCmd.Flags().String(sepSC, "",
		"ABCU123456(*)0   20G1  (*) separator between serial number and check digit")
	validateCmd.Flags().String(sepCS, "",
		"ABCU1234560 (*)  20G1  (*) separator between check digit and size")
	validateCmd.Flags().String(sepST, "",
		"ABCU1234560   20(*)G1  (*) separator between size and type")

	viper.BindPFlag(sepOE, validateCmd.Flags().Lookup(sepOE))
	viper.BindPFlag(sepES, validateCmd.Flags().Lookup(sepES))
	viper.BindPFlag(sepSC, validateCmd.Flags().Lookup(sepSC))
	viper.BindPFlag(sepCS, validateCmd.Flags().Lookup(sepCS))
	viper.BindPFlag(sepST, validateCmd.Flags().Lookup(sepST))

	iso6346Cmd.AddCommand(validateCmd)
}

func validateContNum(cmd *cobra.Command, args []string) {
	num := parser.ParseContNum(args[0])

	num.OwnerCodeIn.Resolve(owner.Resolver(pathToDB))
	num.LengthIn.Resolve(sizetype.GetLength)
	num.HeightWidthIn.Resolve(sizetype.GetHeightAndWidth)
	num.TypeAndGroupIn.Resolve(sizetype.GetTypeAndGroup)

	ui.PrintContNum(num, ui.Separators{
		viper.GetString(sepOE),
		viper.GetString(sepES),
		viper.GetString(sepSC),
		viper.GetString(sepCS),
		viper.GetString(sepST),
	})

	if num.CheckDigitIn.IsValidCheckDigit {
		os.Exit(0)
	}
	os.Exit(1)
}
