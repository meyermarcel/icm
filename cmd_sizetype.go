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

package main

import (
	"os"

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
		st := parseSizeType(args[0])

		st.lengthIn.resolve(getLength)
		st.heightWidthIn.resolve(getHeightAndWidth)
		st.typeAndGroupIn.resolve(getTypeAndGroup)

		printSizeType(st, viper.GetString(sepST))

		if st.lengthIn.isValidFmt() && st.heightWidthIn.isValidFmt() && st.typeAndGroupIn.isValidFmt() {
			os.Exit(0)
		}
		os.Exit(1)
	},
}

var sizeTypePrintCmd = &cobra.Command{
	Use:     "print",
	Short:   "Print length, height, width and type codes",
	Long:    "Print length, height, width and type codes.",
	Example: `  ` + appName + ` sizetype print`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		printSizeTypeDefs(getSizeTypeDef())
	},
}

func init() {
	sizeTypeValidateCmd.Flags().String(sepST, "",
		"20(*)G1  (*) separates size and type")

	viper.BindPFlag(sepST, sizeTypeValidateCmd.Flags().Lookup(sepST))

	sizeTypeCmd.AddCommand(sizeTypeValidateCmd)
	sizeTypeCmd.AddCommand(sizeTypePrintCmd)
	iso6346Cmd.AddCommand(sizeTypeCmd)
}
