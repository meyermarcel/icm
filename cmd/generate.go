// Copyright © 2018 Marcel Meyer meyermarcel@posteo.de
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
	"fmt"
	"io"
	"path/filepath"

	"github.com/meyermarcel/icm/configs"
	"github.com/meyermarcel/icm/internal/cont"
	"github.com/meyermarcel/icm/internal/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newGenerateCmd(writer, writerErr io.Writer, viper *viper.Viper, ownerDecoder data.OwnerDecoder) *cobra.Command {

	var count int

	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a random container number",
		Long: `Owners specified in

  ` + filepath.Join("$HOME", appDir) + `

are used. Container numbers with check digit 10 are excluded.
Equipment category ID 'U' is used for every container number.

` + sepHelp,
		Example: `  ` + appName + ` generate

  ` + appName + ` generate --count 5000

  ` + appName + ` generate --` + configs.SepOE + ` '' --` + configs.SepSC + ` ''`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			generator, err := cont.NewRandomUniqueGenerator(count, ownerDecoder.GenerateRandomCodes(count))
			if err != nil {
				return err
			}
			for i := 0; i < count; i++ {
				contNumber := generator.Generate()
				_, err := io.WriteString(writer,
					fmt.Sprintf("%s%s%s%s%s%s%d\n",
						contNumber.OwnerCode(), viper.GetString(configs.SepOE),
						contNumber.EquipCatID(), viper.GetString(configs.SepES),
						contNumber.SerialNumber(), viper.GetString(configs.SepSC),
						contNumber.CheckDigit()))
				writeErr(writerErr, err)
			}
			return nil
		},
	}

	generateCmd.Flags().String(configs.SepOE, configs.SepOEDefVal,
		"ABC(*)U1234560  (*) separates owner code and equipment category id")
	generateCmd.Flags().String(configs.SepES, configs.SepESDefVal,
		"ABCU(*)1234560  (*) separates equipment category id and serial number")
	generateCmd.Flags().String(configs.SepSC, configs.SepSCDefVal,
		"ABCU123456(*)0  (*) separates serial number and check digit")

	err := viper.BindPFlags(generateCmd.Flags())
	writeErr(writerErr, err)

	generateCmd.Flags().IntVarP(&count, "count", "c", 1, "count of container numbers")

	return generateCmd
}
