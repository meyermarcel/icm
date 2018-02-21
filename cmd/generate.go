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
	"github.com/spf13/cobra"
	"iso6346/cont"
	"iso6346/ui"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random container number.",
	Long: `
Generate a random container number.
Example:

  iso6346 generate

generates a random container number with valid check digit:

  ABCU1234560

The ISO 6346 standard specifies that all characters are
alphanumeric except the equipment ID which is U, J or Z.

You can also format your output:

  iso6346 generate --2nd-separator ' ' -3 '-'

generates a formatted random container number:

  ABCU 123456-0


  ABC U 123456 0
     ↑ ↑      ↑
     │ │      └─ 3rd separator
     │ │
     │ └─ 2nd separator
     │
     └─ 1st separator`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		ui.PrintGenerate(cont.Generate(),
			cmd.Flag("1st-separator").Value.String(),
			cmd.Flag("2nd-separator").Value.String(),
			cmd.Flag("3rd-separator").Value.String())

	},
}

func init() {
	firstSepUsage := "ABC(*)U1234560  ->  (*) 1st separator"
	secondSepUsage := "ABCU(*)1234560  ->  (*) 2nd separator"
	thirdSepUsage := "ABCU123456(*)0  ->  (*) 3rd separator"
	generateCmd.Flags().StringP("1st-separator", "1", "", firstSepUsage)
	generateCmd.Flags().StringP("2nd-separator", "2", " ", secondSepUsage)
	generateCmd.Flags().StringP("3rd-separator", "3", " ", thirdSepUsage)
	RootCmd.AddCommand(generateCmd)
}
