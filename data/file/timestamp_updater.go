// Copyright © 2018 Marcel Meyer meyermarcel@posteo.de
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

package file

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/meyermarcel/icm/data"

	"io/ioutil"
)

const (
	timeout            = time.Duration(5 * time.Minute)
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
	b, err := ioutil.ReadFile(pathToFile)
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
		err := ioutil.WriteFile(filepath.Join(lu.path, lastUpdateFileName), []byte(now.Format(dateFormat)+"\n"), 0644)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("timeout is set to %v to relieve server load, try in %v again",
		timeout, -time.Duration(now.Sub(loaded)-timeout).Round(time.Second))
}
