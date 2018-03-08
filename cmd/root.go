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

var firstSep string
var secondSep string
var thirdSep string

func init() {
	RootCmd.PersistentFlags().StringVarP(&firstSep, "1st-separator", "1", " ", "ABC(*)U1234560  ->  (*) 1st separator")
	RootCmd.PersistentFlags().StringVarP(&secondSep, "2nd-separator", "2", " ", "ABCU(*)1234560  ->  (*) 2nd separator")
	RootCmd.PersistentFlags().StringVarP(&thirdSep, "3rd-separator", "3", " ", "ABCU123456(*)0  ->  (*) 3rd separator")
}
