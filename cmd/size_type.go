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

	"github.com/meyermarcel/icm/configs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newSizeTypeCmd(writer io.Writer, sizeTypeDecoders sizeTypeDecoders) *cobra.Command {
	sizeTypeCmd := &cobra.Command{
		Use:   "sizetype",
		Short: "Validate or print size and type",
		Long:  "Validate or print size and type.",
	}
	sizeTypeCmd.AddCommand(newSizeTypeValidateCmd(writer, sizeTypeDecoders))
	return sizeTypeCmd
}

func newSizeTypeValidateCmd(writer io.Writer, sizeTypeDecoders sizeTypeDecoders) *cobra.Command {
	sizeTypeValidateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate size and type codes",
		Long: `Validate size and type codes.

` + sepHelp,
		Example: `  ` + appName + ` sizetype validate '20G1'

  ` + appName + ` sizetype validate --` + configs.SepST + ` '' '20G1'`,
		Args: cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag(configs.SepST, cmd.Flags().Lookup(configs.SepST))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			in, err := parseSizeType(args[0], sizeTypeDecoders)
			printSizeType(writer, in, viper.GetString(configs.SepST))
			return err
		},
	}
	sizeTypeValidateCmd.Flags().String(configs.SepST, "", "20(*)G1  (*) separates size and type")

	return sizeTypeValidateCmd
}
