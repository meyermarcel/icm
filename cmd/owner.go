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

	"github.com/meyermarcel/iso6346/remote"
	"github.com/meyermarcel/iso6346/ui"
	"github.com/spf13/cobra"
)

var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "Validate an owner or update owners.",
	Long:  "Validate an owner or update owners.",
}

var validateOwnerCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate an owner",
	Long: `Validate an owner.

` + sepHelp,
	Example: `  ` + appName + ` owner validate 'ABCU'`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		valid := ui.ParseAndPrintOwnerCodeOptEquipCat(args[0])

		if valid {
			os.Exit(0)
		}
		os.Exit(1)
	},
}

var ownerUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update information of owners",
	Long: `Update information of owners from remote.
Following information is available:

  Owner code
  Company
  City
  Country`,
	Example: `  ` + appName + ` owner update`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		remote.Update()
		os.Exit(0)
	},
}

func init() {
	ownerCmd.AddCommand(ownerUpdateCmd)
	ownerCmd.AddCommand(validateOwnerCmd)
	RootCmd.AddCommand(ownerCmd)
}
