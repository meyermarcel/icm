package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra/doc"

	"github.com/spf13/cobra"
)

func newDocCmd(rootCmd *cobra.Command) *cobra.Command {
	docCmd := &cobra.Command{
		Use:   "doc",
		Short: "Documentation commands for man pages and markdown generation",
		Long:  "Documentation commands for man pages and markdown generation.",
	}

	// https://unix.stackexchange.com/questions/3586/what-do-the-numbers-in-a-man-page-mean
	// https://docs.brew.sh/Formula-Cookbook -> #{prefix}/share/man
	manCmd := &cobra.Command{
		Use:               "man",
		Short:             "Generate man pages",
		Long:              "Generate man pages.",
		Example:           "icm doc man . && cat *.1",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cobra.NoFileCompletions,
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

	mdCmd := &cobra.Command{
		Use:     "markdown",
		Short:   "Generate markdown",
		Long:    "Generate markdown.",
		Example: "icm doc markdown docs/",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return doc.GenMarkdownTree(rootCmd, args[0])
		},
	}

	docCmd.AddCommand(manCmd)
	docCmd.AddCommand(mdCmd)

	return docCmd
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
