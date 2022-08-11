package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/meyermarcel/icm/data"
)

const (
	timeout            = 5 * time.Minute
	dateFormat         = "2006-01-02T15:04:05Z07:00"
	lastUpdateFileName = "owner-last-update"
	lastUpdate         = "2018-10-29T15:00:00Z" + "\n"
)

type timestampUpdater struct {
	path      string
	timestamp string
}

// NewTimestampUpdater writes last update file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewTimestampUpdater(path string) (data.TimestampUpdater, error) {
	timestampUpdater := &timestampUpdater{path: path}
	pathToFile := filepath.Join(timestampUpdater.path, lastUpdateFileName)
	if err := initFile(pathToFile, []byte(lastUpdate)); err != nil {
		return nil, err
	}
	b, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, err
	}
	timestampUpdater.timestamp = string(b)
	return timestampUpdater, nil
}

// Update writes the recent time to last update file if timeout is exceeded.
func (lu *timestampUpdater) Update() error {
	dateString := strings.TrimSuffix(lu.timestamp, "\n")
	loaded, err := time.Parse(dateFormat, dateString)
	if err != nil {
		return err
	}
	now := time.Now()
	afterTimeout := now.After(loaded.Add(timeout))
	if afterTimeout {
		err := os.WriteFile(filepath.Join(lu.path, lastUpdateFileName), []byte(now.Format(dateFormat)+"\n"), 0o644)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("timeout is set to %v to relieve server load, try in %v again",
		timeout, -(now.Sub(loaded) - timeout).Round(time.Second))
}
