package cmd

import (
	"os"
	"path"

	"github.com/meyermarcel/icm/http"

	"github.com/meyermarcel/icm/data"
	"github.com/spf13/cobra"
)

type filePathValue struct {
	value string
}

func (v *filePathValue) String() string {
	return v.value
}

func (v *filePathValue) Set(value string) error {
	dir, _ := path.Split(value)
	if dir != "" {
		_, err := os.Stat(dir)
		if err != nil {
			return err
		}
	}
	v.value = value
	return nil
}

func (*filePathValue) Type() string {
	return "string"
}

func newDownloadOwnersCmd(
	writeOwnersCSVFunc data.WriteOwnersCSVFunc,
	timestampUpdater data.TimestampUpdater,
	ownersDownloader http.OwnersDownloader,
	ownerCSVPath string,
) *cobra.Command {
	filePath := filePathValue{value: ownerCSVPath}

	downloadOwnersCmd := &cobra.Command{
		Aliases: []string{"update"},
		Use:     "download-owners",
		Short:   "Download information of owners and write CSV to file",
		Long: `Download information of owners and write CSV to file.
Following information is available:

  Owner code
  Company
  City
  Country`,
		Example: `# Overwrite owner.csv file with newest owners
icm download-owners
# Create custom-owner.csv to have additional custom mapping of owner codes
# Use semicolon as a separator. For using double quotes please see existing
# owner.csv file.
echo 'AAA;my company;my city;my country' >> $HOME/.icm/data/custom-owner.csv`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return overwriteOwnersFile(writeOwnersCSVFunc, timestampUpdater, ownersDownloader, filePath.value)
		},
	}
	downloadOwnersCmd.Flags().VarP(&filePath, "output", "o", "output file")

	return downloadOwnersCmd
}

func overwriteOwnersFile(writeOwnersCSV data.WriteOwnersCSVFunc, timestampUpdater data.TimestampUpdater, ownersDownloader http.OwnersDownloader, filePath string) error {
	if err := timestampUpdater.Update(); err != nil {
		return err
	}

	owners, err := ownersDownloader.Download()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	return writeOwnersCSV(owners, file)
}
