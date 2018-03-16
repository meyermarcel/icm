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

var validateOwnerCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate an owner",
	Long: `Validate an owner.

` + sepHelp,
	Example: `
  iso6346 owner validate 'ABC'

  iso6346 owner validate --` + sepOE + ` '' 'ABCU'`,
	Args: cobra.ExactArgs(1),
	Run:  validateOwner,
}

func init() {

	validateOwnerCmd.Flags().String(sepOE, "",
		"ABC(*)U1234560  (*) separator between owner code and equipment category id")

	viper.BindPFlag(sepOE, validateOwnerCmd.Flags().Lookup(sepOE))

	ownerCmd.AddCommand(validateOwnerCmd)
}

func validateOwner(cmd *cobra.Command, args []string) {
	oce := parser.ParseOwnerCodeOptEquipCat(args[0])

	oce.OwnerCodeIn.Resolve(owner.Resolver(pathToDB))

	ui.PrintOwnerCode(oce, viper.GetString(sepOE))

	if oce.OwnerCodeIn.IsValidFmt() {
		os.Exit(0)
	}
	os.Exit(1)
}
