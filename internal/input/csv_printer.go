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
	"encoding/csv"
)

// Datum represents a datum that is be used by CSVPrinter.
type Datum struct {
	Header string
	Value  string
}

// NewDatum returns a new Datum.
func NewDatum(header string) Datum {
	return Datum{Header: header}
}

// WithValue sets value and returns Datum.
func (d *Datum) WithValue(value string) Datum {
	d.Value = value
	return *d
}

// CSVPrinter prints the set record. Use SetRecord to set a record.
type CSVPrinter struct {
	csvWriter     *csv.Writer
	headers       []string
	record        []string
	headerPrinted bool
	printHeader   bool
}

// NewCSVPrinter creates a new CSVPrinter.
func NewCSVPrinter(csvWriter *csv.Writer, printHeader bool) *CSVPrinter {
	return &CSVPrinter{
		csvWriter:   csvWriter,
		printHeader: printHeader,
	}
}

// SetRecord sets the record that will be printed by Print.
func (cp *CSVPrinter) SetRecord(inputs []Input) {
	cp.headers = nil
	cp.record = nil
	for _, input := range inputs {
		for _, datum := range input.data {
			cp.headers = append(cp.headers, datum.Header)
			cp.record = append(cp.record, datum.Value)
		}
	}
}

// Print writes set record to passed writer.
// No header is printed if printHeader is set to false.
// Print returns an error if writing to writer fails.
func (cp *CSVPrinter) Print() error {
	if cp.printHeader && !cp.headerPrinted {
		err := cp.csvWriter.Write(cp.headers)
		if err != nil {
			return err
		}
		cp.headerPrinted = true
	}
	return cp.csvWriter.Write(cp.record)
}
