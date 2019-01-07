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
	"os"
	"path/filepath"

	"github.com/spf13/cobra/doc"

	"github.com/spf13/cobra"
)

func newMiscCmd(writer io.Writer, rootCmd *cobra.Command) *cobra.Command {
	miscCmd := &cobra.Command{
		Use:   "misc",
		Short: "Miscellaneous generation commands for man page and bash, zsh completions",
		Long:  "Miscellaneous generation commands for man page and bash, zsh completions.",
	}

	bashCmd := &cobra.Command{
		Use:   "bash-completion",
		Short: "Generate bash completion scripts",
		Long:  "Generate bash completion scripts.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.GenBashCompletion(writer)
		},
	}

	zshCmd := &cobra.Command{
		Use:   "zsh-completion",
		Short: "Generate zsh completion scripts",
		Long:  "Generate zsh completion scripts.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.GenZshCompletion(writer)
		},
	}

	// https://unix.stackexchange.com/questions/3586/what-do-the-numbers-in-a-man-page-mean
	// https://docs.brew.sh/Formula-Cookbook -> #{prefix}/share/man
	manCmd := &cobra.Command{
		Use:     "man",
		Short:   "Generate man pages",
		Long:    "Generate man pages.",
		Example: "  icm misc man . && cat *.1",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]

			if _, err := os.Stat(path); os.IsNotExist(err) {
				err := os.MkdirAll(path, os.ModePerm)
				if err != nil {
					return err
				}
			}

			// Root command
			err := writeMan(path, fmt.Sprintf("%s.1", rootCmd.Name()), rootCmd)
			if err != nil {
				return err
			}

			// Sub commands
			for _, subCmd := range rootCmd.Commands() {
				if subCmd.Name() != "help" {
					err := writeMan(path, fmt.Sprintf("%s-%s.1", appName, subCmd.Name()), subCmd)
					if err != nil {
						return err
					}
				}
			}
			return nil
		},
	}

	miscCmd.AddCommand(bashCmd)
	miscCmd.AddCommand(zshCmd)
	miscCmd.AddCommand(manCmd)

	return miscCmd
}

func writeMan(path, name string, cmd *cobra.Command) error {
	file, err := os.Create(filepath.Join(path, name))
	if err != nil {
		return err
	}
	err = doc.GenMan(cmd, nil, file)
	if err != nil {
		return err
	}
	return nil
}
