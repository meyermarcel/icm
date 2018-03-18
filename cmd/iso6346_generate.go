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
	"github.com/meyermarcel/iso6346/cont"
	"github.com/meyermarcel/iso6346/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random container number",
	Long: `Generate a random container number.

` + sepHelp,
	Example: `  ` + appName + ` generate

  ` + appName + ` generate --count 5000

  ` + appName + ` generate --` + sepOE + ` '' --` + sepSC + ` ''`,
	Args: cobra.NoArgs,
	Run:  generate,
}

func generate(cmd *cobra.Command, args []string) {
	c := make(chan cont.Number)
	go cont.Gen(pathToDB, count, c)

	for contNum := range c {
		ui.PrintGen(contNum, ui.Separators{
			viper.GetString(sepOE),
			viper.GetString(sepES),
			viper.GetString(sepSC),
			"", "",
		})
	}
	os.Exit(0)
}

var count int

func init() {
	generateCmd.Flags().IntVarP(&count, "count", "c", 1, "count of container numbers")

	generateCmd.Flags().String(sepOE, "",
		"ABC(*)U1234560  (*) separator between owner code and equipment category id")
	generateCmd.Flags().String(sepES, "",
		"ABCU(*)1234560  (*) separator between equipment category id and serial number")
	generateCmd.Flags().String(sepSC, "",
		"ABCU123456(*)0  (*) separator between serial number and check digit")

	viper.BindPFlag(sepOE, generateCmd.Flags().Lookup(sepOE))
	viper.BindPFlag(sepES, generateCmd.Flags().Lookup(sepES))
	viper.BindPFlag(sepSC, generateCmd.Flags().Lookup(sepSC))

	iso6346Cmd.AddCommand(generateCmd)

}
