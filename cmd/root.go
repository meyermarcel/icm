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

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "iso6346",
	Short: "Parse or generate all ISO 6346 related data.",
	Long:  "Parse or generate all ISO 6346 related data.",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var sepOwnerEquip string
var sepEquipSerial string
var sepSerialCheck string

func init() {
	RootCmd.PersistentFlags().StringVarP(&sepOwnerEquip, "sep-owner-equip", "", " ",
		"ABC(*)U1234560  (*) separator between owner code and equipment category id")
	RootCmd.PersistentFlags().StringVarP(&sepEquipSerial, "sep-equip-serial", "", " ",
		"ABCU(*)1234560  (*) separator between equipment category id and serial number")
	RootCmd.PersistentFlags().StringVarP(&sepSerialCheck, "sep-serial-check", "", " ",
		"ABCU123456(*)0  (*) separator between serial number and check digit")
}
