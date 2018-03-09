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
	"fmt"
	"os"

	"github.com/meyermarcel/iso6346/owner"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
)

var defaultCfg = []byte(
	sepOwnerEquip + ": ' '\n" +
		sepEquipSerial + ": ' '\n" +
		sepSerialCheck + ": ' '\n")

var RootCmd = &cobra.Command{
	Use:     "iso6346",
	Version: "0.1.0-beta",
	Short:   "Parse or generate all ISO 6346 related data",
	Long: `Parse or generate all ISO 6346 related data.

Output can be formatted:

  ABC U 123456 0
     ↑ ↑      ↑
     │ │      └─ separator between serial number and check digit
     │ │
     │ └─ separator between equipment category id and serial number
     │
     └─ separator between owner code and equipment category id`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const (
	appDir         = ".iso6346"
	dbName         = "iso6346.db"
	cfgName        = "separators"
	sepOwnerEquip  = "sep-owner-equip"
	sepEquipSerial = "sep-equip-serial"
	sepSerialCheck = "sep-serial-check"
)

var pathToDB string

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().String(sepOwnerEquip, "",
		"ABC(*)U1234560  (*) separator between owner code and equipment category id")
	RootCmd.PersistentFlags().String(sepEquipSerial, "",
		"ABCU(*)1234560  (*) separator between equipment category id and serial number")
	RootCmd.PersistentFlags().String(sepSerialCheck, "",
		"ABCU123456(*)0  (*) separator between serial number and check digit")

	viper.BindPFlags(RootCmd.PersistentFlags())
}

func initConfig() {

	appDirPath := initDir(getPathToAppDir(appDir))

	pathToDB = filepath.Join(appDirPath, dbName)

	owner.InitDB(pathToDB)

	initDefaultCfg(filepath.Join(appDirPath, cfgName+".yml"))

	viper.AddConfigPath(appDirPath)
	viper.SetConfigName(cfgName)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func getPathToAppDir(appDir string) string {
	homeDir, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, appDir)
}

func initDir(path string) string {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir|0700)
	}
	return path
}

func initDefaultCfg(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := ioutil.WriteFile(path, defaultCfg, 0644); err != nil {
			fmt.Println("Can't write default config:", err)
			os.Exit(1)
		}
	}
}
