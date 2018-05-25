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

package main

import (
	"fmt"
	"os"

	"io/ioutil"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	appName = "iso6346"
	appDir  = "." + appName
	dbName  = appName + ".db"
	cfgName = "separators"
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

  ` + filepath.Join("$HOME", appDir, cfgName+".yml")

var iso6346Cmd = &cobra.Command{
	Use:     appName,
	Version: "0.1.0-beta",
	Short:   "Parse or generate ISO 6346 related data",
	Long:    "Parse or generate ISO 6346 related data.",
}

func execute() {
	if err := iso6346Cmd.Execute(); err != nil {
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
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan contNumber)
		go genContNum(pathToDB, count, c)

		for contNum := range c {
			printGen(contNum, separators{
				viper.GetString(sepOE),
				viper.GetString(sepES),
				viper.GetString(sepSC),
				"", "",
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
	Run: func(cmd *cobra.Command, args []string) {
		num := parseContNum(args[0])

		num.OwnerCodeIn.resolve(resolver(pathToDB))
		num.LengthIn.resolve(getLength)
		num.HeightWidthIn.resolve(getHeightAndWidth)
		num.TypeAndGroupIn.resolve(getTypeAndGroup)

		printContNum(num, separators{
			viper.GetString(sepOE),
			viper.GetString(sepES),
			viper.GetString(sepSC),
			viper.GetString(sepCS),
			viper.GetString(sepST),
		})

		if num.CheckDigitIn.IsValidCheckDigit {
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

	viper.BindPFlag(sepOE, generateCmd.Flags().Lookup(sepOE))
	viper.BindPFlag(sepES, generateCmd.Flags().Lookup(sepES))
	viper.BindPFlag(sepSC, generateCmd.Flags().Lookup(sepSC))

	iso6346Cmd.AddCommand(generateCmd)

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

	viper.BindPFlag(sepOE, validateCmd.Flags().Lookup(sepOE))
	viper.BindPFlag(sepES, validateCmd.Flags().Lookup(sepES))
	viper.BindPFlag(sepSC, validateCmd.Flags().Lookup(sepSC))
	viper.BindPFlag(sepCS, validateCmd.Flags().Lookup(sepCS))
	viper.BindPFlag(sepST, validateCmd.Flags().Lookup(sepST))

	iso6346Cmd.AddCommand(validateCmd)
}

var pathToDB string

func initConfig() {

	appDirPath := initDir(getPathToAppDir(appDir))

	pathToDB = filepath.Join(appDirPath, dbName)

	initDB(pathToDB)

	initFile(filepath.Join(appDirPath, cfgName+".yml"), configSeparators())

	viper.AddConfigPath(appDirPath)
	viper.SetConfigName(cfgName)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
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
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, appDir)
}

func initFile(path string, content []byte) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := ioutil.WriteFile(path, content, 0644); err != nil {
			fmt.Println("Cannot write ", path, ":", err)
			os.Exit(1)
		}
	}
}
