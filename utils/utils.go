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

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// JSONMarshal marshals given interface to JSON without escaping string for HTML and handles error.
func JSONMarshal(t interface{}) []byte {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	CheckErr(err)
	var fmt bytes.Buffer
	json.Indent(&fmt, buffer.Bytes(), "", "  ")
	return fmt.Bytes()
}

// JSONUnmarshal wraps known unmarshal method and handles error.
func JSONUnmarshal(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	CheckErr(err)
}

// InitFile writs file tp path if it not exists.
func InitFile(path string, content []byte) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := ioutil.WriteFile(path, content, 0644)
		CheckErr(err)
	}
	return path
}

// CheckErr prints error exits with exit code 1 if err is not nil.
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
