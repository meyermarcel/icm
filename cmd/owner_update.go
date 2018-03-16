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
	"github.com/spf13/cobra"
	"os"
)

var ownerUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update information of owners",
	Long: `Update information of owners from remote.
Following information is available:

  Owner code
  Company
  City
  Country`,
	Example: `
  iso6346 owner update`,
	Args: cobra.NoArgs,
	Run:  updateOwners,
}

func updateOwners(cmd *cobra.Command, args []string) {
	owner.Update(pathToDB)
	os.Exit(0)
}

func init() {

	ownerCmd.AddCommand(ownerUpdateCmd)
}
