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
	"io"

	"github.com/meyermarcel/icm/internal/data"

	"github.com/spf13/cobra"
)

func newOwnerCmd(
	writer io.Writer,
	owner data.OwnerDecodeUpdater,
	timestampUpdater data.TimestampUpdater,
	ownerURL string) *cobra.Command {
	ownerCmd := &cobra.Command{
		Use:   "owner",
		Short: "Validate an owner or update owners.",
		Long:  "Validate an owner or update owners.",
	}
	ownerCmd.AddCommand(newValidateOwnerCmd(writer, owner))
	ownerCmd.AddCommand(newUpdateOwnerCmd(owner, timestampUpdater, ownerURL))

	return ownerCmd
}

func newValidateOwnerCmd(writer io.Writer, owner data.OwnerDecoder) *cobra.Command {
	command := &cobra.Command{
		Use:   "validate",
		Short: "Validate an owner",
		Long: `Validate an owner.

` + sepHelp,
		Example: `  ` + appName + ` owner validate 'ABCU'`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			in, err := parseOwnerCodeOptEquipCat(args[0], owner)
			printOwnerCode(writer, in, owner)
			return err
		},
	}
	return command
}

func newUpdateOwnerCmd(
	ownerUpdater data.OwnerUpdater,
	timestampUpdater data.TimestampUpdater,
	ownerURL string) *cobra.Command {
	ownerUpdateCmd := &cobra.Command{
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
		RunE: func(cmd *cobra.Command, args []string) error {
			return update(ownerUpdater, timestampUpdater, ownerURL)
		},
	}
	return ownerUpdateCmd
}
