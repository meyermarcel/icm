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
	"github.com/meyermarcel/iso6346/parser"
	"github.com/meyermarcel/iso6346/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var sizetypeValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a container number",
	Long: `Validate a container number.

` + sepHelp,
	Example: `  ` + appName + ` sizetype validate '20G1'

  ` + appName + ` sizetype validate --` + sepST + ` '' '20G1'`,
	Args: cobra.ExactArgs(1),
	Run:  validateSizeType,
}

func init() {
	sizetypeValidateCmd.Flags().String(sepST, "",
		"20(*)G1  (*) separator between size and type")

	viper.BindPFlag(sepST, sizetypeValidateCmd.Flags().Lookup(sepST))

	sizetypeCmd.AddCommand(sizetypeValidateCmd)
}

func validateSizeType(cmd *cobra.Command, args []string) {
	st := parser.ParseSizeType(args[0])

	ui.PrintSizeType(st, viper.GetString(sepST))

	if st.TypeIn.IsValidFmt() {
		os.Exit(0)
	}
	os.Exit(1)
}
