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
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/meyermarcel/icm/internal/cont"

	"github.com/meyermarcel/icm/internal/input"

	"github.com/meyermarcel/icm/internal/data"

	"github.com/meyermarcel/icm/configs"

	"path/filepath"

	"os/user"

	"github.com/meyermarcel/icm/internal/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type decoders struct {
	ownerDecodeUpdater data.OwnerDecodeUpdater
	equipCatDecoder    data.EquipCatDecoder
	sizeTypeDecoders
}

type sizeTypeDecoders struct {
	lengthDecoder      data.LengthDecoder
	heightWidthDecoder data.HeightWidthDecoder
	typeDecoder        data.TypeDecoder
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

	lengthDecoder, heightWidthDecoder, err := file.NewSizeDecoder(appDirDataPath)
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
				heightWidthDecoder,
				typeDecoder},
		},
		timestampUpdater,
		ownerURL)

	err = rootCmd.Execute()
	// err is printed by cobra.
	if err != nil {
		os.Exit(1)
	}
}

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
	rootCmd.AddCommand(newUpdateOwnerCmd(decoders.ownerDecodeUpdater, timestampUpdater, ownerURL))

	return rootCmd
}

var count int

func newGenerateCmd(writer io.Writer, ownerDecoder data.OwnerDecoder) *cobra.Command {
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
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := viper.BindPFlag(configs.SepOE, cmd.Flags().Lookup(configs.SepOE))
			err = viper.BindPFlag(configs.SepES, cmd.Flags().Lookup(configs.SepES))
			err = viper.BindPFlag(configs.SepSC, cmd.Flags().Lookup(configs.SepSC))
			if err != nil {
				return err
			}
			return nil
		},
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
				if err != nil {
					return err
				}
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

func newValidateCmd(writer io.Writer, decoders decoders) *cobra.Command {
	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate intermodal container markings",
		Long: `Validate intermodal container markings.

` + sepHelp,
		Example: `  ` + appName + ` validate 'ABCU 1234560'

  ` + appName + ` validate 'ABCU'

  ` + appName + ` validate '20G1'

  ` + appName + ` validate --` + configs.SepOE + ` '' --` + configs.SepSC + ` '' 'ABCU 1234560'`,
		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := viper.BindPFlag(configs.SepOE, cmd.Flags().Lookup(configs.SepOE))
			err = viper.BindPFlag(configs.SepES, cmd.Flags().Lookup(configs.SepES))
			err = viper.BindPFlag(configs.SepSC, cmd.Flags().Lookup(configs.SepSC))
			err = viper.BindPFlag(configs.SepCS, cmd.Flags().Lookup(configs.SepCS))
			err = viper.BindPFlag(configs.SepST, cmd.Flags().Lookup(configs.SepST))
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			reader := strings.NewReader(args[0])
			bufReader := bufio.NewReader(reader)
			peek, _ := bufReader.Peek(bufReader.Buffered())

			validator := newIcmValidator(decoders)

			if isSingleLine(string(peek)) {
				buf := new(bytes.Buffer)
				_, _ = buf.ReadFrom(reader)
				inputs := validator.Validate(buf.String())
				fancyPrinter := input.NewFancyPrinter(writer, inputs)
				fancyPrinter.SetIndent("  ")
				fancyPrinter.SetSeparators(
					viper.GetString(configs.SepOE),
					viper.GetString(configs.SepES),
					viper.GetString(configs.SepSC),
					viper.GetString(configs.SepCS),
					"",
					viper.GetString(configs.SepST),
				)
				err := fancyPrinter.Print()
				if err != nil {
					return err
				}
			}
			return nil
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

func isSingleLine(s string) bool {
	scanner := bufio.NewScanner(strings.NewReader(s))
	counter := 0
	for scanner.Scan() {
		counter++
		if counter > 1 {
			return false
		}
	}
	return true
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
		Example: `  ` + appName + ` update`,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return update(ownerUpdater, timestampUpdater, ownerURL)
		},
	}
	return ownerUpdateCmd
}

func initDir(path string) string {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModeDir|0700)
	}
	return path
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
