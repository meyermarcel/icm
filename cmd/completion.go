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
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate scripts completion for shells",
	Long:  "Generate scripts completion for shells.",
}

var bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generate bash completion scripts",
	Long:  "Generate bash completion scripts.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenBashCompletion(os.Stdout)
	},
}

var zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generate zsh completion scripts",
	Long:  "Generate zsh completion scripts.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	completionCmd.AddCommand(bashCmd)
	completionCmd.AddCommand(zshCmd)
	RootCmd.AddCommand(completionCmd)
}
