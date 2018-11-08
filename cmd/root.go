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
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/meyermarcel/icm/internal/data"

	"github.com/meyermarcel/icm/configs"

	"path/filepath"

	"os/user"

	"github.com/meyermarcel/icm/internal/cont"
	"github.com/meyermarcel/icm/internal/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type decoders struct {
	ownerDecodeUpdater data.OwnerDecodeUpdater
	equipCatDecoder    data.EquipCatDecoder
	sizeTypeDecoders   sizeTypeDecoders
}

type sizeTypeDecoders struct {
	lengthDecoder         data.LengthDecoder
	heightAndWidthDecoder data.HeightAndWidthDecoder
	typeDecoder           data.TypeDecoder
}

const (
	appName  = "icm"
	appDir   = "." + appName
	ownerURL = "https://www.bic-code.org/bic-letter-search/?resultsperpage=17576&searchterm="
)

var sepHelp = `Configuration for separators is generated first time you
execute a command that requires the configuration.

Flags for output formatting can overridden with a config file.
Edit default configuration:

  ` + filepath.Join("$HOME", appDir, configs.NameWithYmlExt)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	current, err := user.Current()
	checkErr(err)

	appDirPath := initDir(filepath.Join(current.HomeDir, appDir))

	pathToCfg := filepath.Join(appDirPath, configs.NameWithYmlExt)
	if _, err := os.Stat(pathToCfg); os.IsNotExist(err) {
		err = ioutil.WriteFile(pathToCfg, configs.Cfg(), 0644)
		checkErr(err)
	}
	viper.AddConfigPath(appDirPath)
	viper.SetConfigName(configs.Name)
	err = viper.ReadInConfig()
	checkErr(err)

	appDirDataPath := initDir(filepath.Join(appDirPath, "data"))

	ownerDecodeUpdater, err := file.NewOwnerDecoderUpdater(appDirDataPath)
	checkErr(err)

	equipCatDecoder, err := file.NewEquipCatDecoder(appDirDataPath)
	checkErr(err)

	lengthDecoder, err := file.NewLengthDecoder(appDirDataPath)
	checkErr(err)

	heightAndWidthDecoder, err := file.NewHeightAndWidthDecoder(appDirDataPath)
	checkErr(err)

	typeDecoder, err := file.NewTypeDecoder(appDirDataPath)
	checkErr(err)

	timestampUpdater, err := file.NewTimestampUpdater(appDirDataPath)
	checkErr(err)

	rootCmd := newRootCmd(
		version,
		os.Stdout,
		decoders{
			ownerDecodeUpdater,
			equipCatDecoder,
			sizeTypeDecoders{
				lengthDecoder,
				heightAndWidthDecoder,
				typeDecoder},
		},
		timestampUpdater,
		ownerURL)

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// RootCmd represents the base command when called without any subcommands
func newRootCmd(version string, writer io.Writer,
	decoders decoders,
	timestampUpdater data.TimestampUpdater,
	ownerURL string) *cobra.Command {
	rootCmd := &cobra.Command{
		Version:      version,
		Use:          appName,
		Short:        "Validate or generate intermodal container markings",
		Long:         "Validate or generate intermodal container markings.",
		SilenceUsage: true,
	}

	rootCmd.AddCommand(newCompletionCmd(writer, rootCmd))
	rootCmd.AddCommand(newGenerateCmd(writer, decoders.ownerDecodeUpdater))
	rootCmd.AddCommand(newValidateCmd(writer, decoders))
	rootCmd.AddCommand(newOwnerCmd(writer, decoders.ownerDecodeUpdater, timestampUpdater, ownerURL))
	rootCmd.AddCommand(newSizeTypeCmd(writer, decoders.sizeTypeDecoders))

	return rootCmd
}

var count int

func newGenerateCmd(writer io.Writer, ownerData data.OwnerDecoder) *cobra.Command {
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
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag(configs.SepOE, cmd.Flags().Lookup(configs.SepOE))
			viper.BindPFlag(configs.SepES, cmd.Flags().Lookup(configs.SepES))
			viper.BindPFlag(configs.SepSC, cmd.Flags().Lookup(configs.SepSC))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c := make(chan cont.Result)
			go cont.GenNum(count, c, ownerData.GenerateRandomCodes)

			for result := range c {
				if result.Err() != nil {
					return result.Err()
				}
				printContNum(writer, result.ContNum(), separators{
					OwnerEquip:  viper.GetString(configs.SepOE),
					EquipSerial: viper.GetString(configs.SepES),
					SerialCheck: viper.GetString(configs.SepSC),
				})
			}
			return nil
		},
	}
	generateCmd.Flags().IntVarP(&count, "count", "c", 1, "count of container numbers")

	generateCmd.Flags().String(configs.SepOE, "",
		"ABC(*)U1234560  (*) separates owner code and equipment category id")
	generateCmd.Flags().String(configs.SepES, "",
		"ABCU(*)1234560  (*) separates equipment category id and serial number")
	generateCmd.Flags().String(configs.SepSC, "",
		"ABCU123456(*)0  (*) separates serial number and check digit")
	return generateCmd
}

func newValidateCmd(writer io.Writer, data decoders) *cobra.Command {
	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate a container number",
		Long: `Validate a container number.

` + sepHelp,
		Example: `  ` + appName + ` validate 'ABCU 1234560'

  ` + appName + ` validate --` + configs.SepOE + ` '' --` + configs.SepSC + ` '' 'ABCU 1234560'`,
		Args: cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag(configs.SepOE, cmd.Flags().Lookup(configs.SepOE))
			viper.BindPFlag(configs.SepES, cmd.Flags().Lookup(configs.SepES))
			viper.BindPFlag(configs.SepSC, cmd.Flags().Lookup(configs.SepSC))
			viper.BindPFlag(configs.SepCS, cmd.Flags().Lookup(configs.SepCS))
			viper.BindPFlag(configs.SepST, cmd.Flags().Lookup(configs.SepST))
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			parsedContNum, err := parseContNum(args[0], data)

			printContNumVal(writer, parsedContNum, data, separators{
				OwnerEquip:  viper.GetString(configs.SepOE),
				EquipSerial: viper.GetString(configs.SepES),
				SerialCheck: viper.GetString(configs.SepSC),
				CheckSize:   viper.GetString(configs.SepCS),
				SizeType:    viper.GetString(configs.SepST),
			})
			return err
		},
	}
	validateCmd.Flags().String(configs.SepOE, "",
		"ABC(*)U1234560   20G1  (*) separates owner code and equipment category id")
	validateCmd.Flags().String(configs.SepES, "",
		"ABCU(*)1234560   20G1  (*) separates equipment category id and serial number")
	validateCmd.Flags().String(configs.SepSC, "",
		"ABCU123456(*)0   20G1  (*) separates serial number and check digit")
	validateCmd.Flags().String(configs.SepCS, "",
		"ABCU1234560 (*)  20G1  (*) separates check digit and size")
	validateCmd.Flags().String(configs.SepST, "",
		"ABCU1234560   20(*)G1  (*) separates size and type")
	return validateCmd
}

func initDir(path string) string {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir|0700)
	}
	return path
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
