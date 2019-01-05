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

package input

import (
	"bytes"
	"encoding/csv"
	"testing"
)

func TestCSVPrinter_Print(t *testing.T) {
	tests := []struct {
		name       string
		noHeader   bool
		inputs     []Input
		wantWriter string
	}{
		{
			name:     "Print CSV with header",
			noHeader: false,
			inputs: []Input{
				{
					data: []Datum{
						{Header: "header-1", Value: "value-1"},
						{Header: "header-2", Value: "value-2"},
						{Header: "header-3", Value: "value-3"},
					},
				},
			},
			wantWriter: `header-1,header-2,header-3
value-1,value-2,value-3
`,
		},
		{
			name:     "Print CSV without header",
			noHeader: true,
			inputs: []Input{
				{
					data: []Datum{
						{Header: "header-1", Value: "value-1"},
						{Header: "header-2", Value: "value-2"},
					},
				},
			},
			wantWriter: `value-1,value-2
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}

			csvWriter := csv.NewWriter(writer)
			csvPrinter := NewCSVPrinter(csvWriter, tt.noHeader)
			_ = csvPrinter.Print(tt.inputs)

			csvWriter.Flush()

			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
