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
	"github.com/spf13/cobra"
	"github.com/meyermarcel/iso6346/ui"
	"github.com/meyermarcel/iso6346/sizetype"
)

var sizeTypePrintCmd = &cobra.Command{
	Use:   "sizetype",
	Short: "Print length, height, width and type codes",
	Long: `Print length, height, width and type codes.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintSizeTypeDefs(sizetype.GetDef())
	},
}

func init() {
	RootCmd.AddCommand(sizeTypePrintCmd)
}
