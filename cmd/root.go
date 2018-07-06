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
	"os"

	"path/filepath"

	"github.com/meyermarcel/icm/cont"
	"github.com/meyermarcel/icm/data"
	"github.com/meyermarcel/icm/remote"
	"github.com/meyermarcel/icm/ui"
	"github.com/meyermarcel/icm/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version = "dev"

const (
	appName = "icm"
	appDir  = "." + appName
	sepOE   = "sep-owner-equip"
	sepES   = "sep-equip-serial"
	sepSC   = "sep-serial-check"
	sepCS   = "sep-check-size"
	sepST   = "sep-size-type"
)

var sepHelp = `Configuration for separators is generated first time you
execute a command that requires the configuration.

Flags for output formatting can overridden with a config file.
Edit default configuration:

  ` + filepath.Join("$HOME", appDir, ymlCfgFileName)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     appName,
	Version: version,
	Short:   "Validate or generate intermodal container markings",
	Long:    "Validate or generate intermodal container markings.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var count int

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random container number",
	Long: `Generate a random container number.

` + sepHelp,
	Example: `  ` + appName + ` generate

  ` + appName + ` generate --count 5000

  ` + appName + ` generate --` + sepOE + ` '' --` + sepSC + ` ''`,
	Args: cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag(sepOE, cmd.Flags().Lookup(sepOE))
		viper.BindPFlag(sepES, cmd.Flags().Lookup(sepES))
		viper.BindPFlag(sepSC, cmd.Flags().Lookup(sepSC))
	},
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan cont.Number)
		go cont.GenNum(count, c, data.GetRandomOwnerCodes)

		for contNum := range c {
			ui.PrintContNum(contNum, ui.Separators{
				OwnerEquip:  viper.GetString(sepOE),
				EquipSerial: viper.GetString(sepES),
				SerialCheck: viper.GetString(sepSC),
			})
		}
		os.Exit(0)
	},
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a container number",
	Long: `Validate a container number.

` + sepHelp,
	Example: `  ` + appName + ` validate 'ABCU 1234560'

  ` + appName + ` validate --` + sepOE + ` '' --` + sepSC + ` '' 'ABCU 1234560'`,
	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag(sepOE, cmd.Flags().Lookup(sepOE))
		viper.BindPFlag(sepES, cmd.Flags().Lookup(sepES))
		viper.BindPFlag(sepSC, cmd.Flags().Lookup(sepSC))
		viper.BindPFlag(sepCS, cmd.Flags().Lookup(sepCS))
		viper.BindPFlag(sepST, cmd.Flags().Lookup(sepST))
	},
	Run: func(cmd *cobra.Command, args []string) {

		valid := ui.ParseAndPrintContNum(args[0], ui.Separators{
			OwnerEquip:  viper.GetString(sepOE),
			EquipSerial: viper.GetString(sepES),
			SerialCheck: viper.GetString(sepSC),
			CheckSize:   viper.GetString(sepCS),
			SizeType:    viper.GetString(sepST),
		})

		if valid {
			os.Exit(0)
		}
		os.Exit(1)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	generateCmd.Flags().IntVarP(&count, "count", "c", 1, "count of container numbers")

	generateCmd.Flags().String(sepOE, "",
		"ABC(*)U1234560  (*) separates owner code and equipment category id")
	generateCmd.Flags().String(sepES, "",
		"ABCU(*)1234560  (*) separates equipment category id and serial number")
	generateCmd.Flags().String(sepSC, "",
		"ABCU123456(*)0  (*) separates serial number and check digit")

	RootCmd.AddCommand(generateCmd)

	validateCmd.Flags().String(sepOE, "",
		"ABC(*)U1234560   20G1  (*) separates owner code and equipment category id")
	validateCmd.Flags().String(sepES, "",
		"ABCU(*)1234560   20G1  (*) separates equipment category id and serial number")
	validateCmd.Flags().String(sepSC, "",
		"ABCU123456(*)0   20G1  (*) separates serial number and check digit")
	validateCmd.Flags().String(sepCS, "",
		"ABCU1234560 (*)  20G1  (*) separates check digit and size")
	validateCmd.Flags().String(sepST, "",
		"ABCU1234560   20(*)G1  (*) separates size and type")

	RootCmd.AddCommand(validateCmd)
}

func initConfig() {

	appDirPath := initDir(getPathToAppDir(appDir))

	utils.InitFile(filepath.Join(appDirPath, ymlCfgFileName), cfgSeparators())

	viper.AddConfigPath(appDirPath)
	viper.SetConfigName(ymlCfgName)

	appDirPathData := initDir(filepath.Join(appDirPath, "data"))

	data.InitOwnersData(appDirPathData)
	remote.InitOwnersLastUpdate(appDirPathData)

	data.InitEquipCatIDsData(appDirPathData)
	data.InitSizesData(appDirPathData)
	data.InitTypesData(appDirPathData)
	data.InitGroupsData(appDirPathData)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Cannot read config:", err)
		os.Exit(1)
	}
}

func initDir(path string) string {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir|0700)
	}
	return path
}

func getPathToAppDir(appDir string) string {
	homeDir, err := homedir.Dir()
	utils.CheckErr(err)
	return filepath.Join(homeDir, appDir)
}
