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

package cont

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func TestGeneratorBuilder(t *testing.T) {
	type fields struct {
		codes            []string
		count            int
		rangeStart       int
		rangeEnd         int
		exclCheckDigit10 bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    *UniqueGenerator
		wantErr bool
	}{
		{
			"Build unique container number generator with random serial iterator",
			fields{
				[]string{"ABC"},
				2,
				-1,
				-1,
				true,
			},
			&UniqueGenerator{
				codes:    []string{"ABC"},
				lenCodes: 1,
				serialNumIt: &randSerialNumIt{
					randOffset: 5577006791947779410,
				},
				count:            2,
				exclCheckDigit10: true,
			},
			false,
		},
		{
			"Build unique container number generator with sequential serial iterator and only range start",
			fields{
				codes:      []string{"ABC"},
				count:      3,
				rangeStart: 2,
				rangeEnd:   -1,
			},
			&UniqueGenerator{
				codes:       []string{"ABC"},
				lenCodes:    1,
				serialNumIt: newSeqSerialNumIt(2),
				count:       3,
			},
			false,
		},
		{
			"Build unique container number generator with sequential serial iterator and only range end",
			fields{
				codes:      []string{"ABC"},
				count:      4,
				rangeStart: -1,
				rangeEnd:   2,
			},
			&UniqueGenerator{
				codes:       []string{"ABC"},
				lenCodes:    1,
				serialNumIt: newSeqSerialNumIt(-1),
				count:       4,
			},
			false,
		},
		{
			"Build unique container number generator with sequential serial iterator and start and end range",
			fields{
				codes:      []string{"ABC"},
				count:      1,
				rangeStart: 2,
				rangeEnd:   5,
			},
			&UniqueGenerator{
				codes:       []string{"ABC"},
				lenCodes:    1,
				serialNumIt: newSeqSerialNumIt(2),
				count:       4,
			},
			false,
		},
		{
			"Build returns error for no owner codes",
			fields{
				count:      1,
				rangeStart: -1,
				rangeEnd:   -1,
			},
			nil,
			true,
		},
		{
			"Build returns error for count less than zero",
			fields{
				codes:      []string{"ABC"},
				rangeStart: -1,
				rangeEnd:   -1,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(1)
			gb := NewUniqueGeneratorBuilder().
				OwnerCodes(tt.fields.codes).
				Count(tt.fields.count).
				Start(tt.fields.rangeStart).
				End(tt.fields.rangeEnd).
				ExcludeCheckDigit10(tt.fields.exclCheckDigit10)
			got, err := gb.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratorBuilder.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratorBuilder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueGenerator(t *testing.T) {
	type fields struct {
		codes            []string
		count            int
		rangeStart       int
		rangeEnd         int
		exclCheckDigit10 bool
	}
	tests := []struct {
		name         string
		fields       fields
		seqSerialNum bool
		want         int
	}{
		{
			"Generate 3 unique container numbers with random serial numbers",
			fields{
				codes:      []string{"ABC"},
				count:      3,
				rangeStart: -1,
				rangeEnd:   -1,
			},
			false,
			3,
		},
		{
			"Generate 1 container number",
			fields{
				codes:      []string{"ABC"},
				count:      3,
				rangeStart: 1,
				rangeEnd:   1,
			},
			true,
			1,
		},
		{
			"Generate 3 unique container numbers with sequential serial numbers and start 1",
			fields{
				codes:      []string{"ABC"},
				count:      3,
				rangeStart: 1,
				rangeEnd:   -1,
			},
			true,
			3,
		},
		{
			"Generate 4 unique container numbers with sequential serial numbers and end 2",
			fields{
				codes:      []string{"ABC"},
				count:      4,
				rangeStart: -1,
				rangeEnd:   2,
			},
			true,
			4,
		},
		{
			"Generate 5 unique container numbers with sequential serial numbers and start 1 and end 5",
			fields{
				codes:      []string{"ABC"},
				rangeStart: 1,
				rangeEnd:   5,
			},
			true,
			5,
		},
		{
			"Generate 6 unique container numbers with sequential serial numbers and start 999997 and end 2",
			fields{
				codes:      []string{"ABC"},
				rangeStart: 999997,
				rangeEnd:   2,
			},
			true,
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(1)
			gb := NewUniqueGeneratorBuilder().
				OwnerCodes(tt.fields.codes).
				Count(tt.fields.count).
				Start(tt.fields.rangeStart).
				End(tt.fields.rangeEnd).
				ExcludeCheckDigit10(tt.fields.exclCheckDigit10)
			g, _ := gb.Build()

			lastNum := g.serialNumIt.num() - 1
			diff := 0
			contNumbers := map[string]bool{}
			for g.Generate() {
				number, _ := strconv.Atoi(g.ContNum().SerialNumber())
				if number < 0 || number > 999999 {
					t.Errorf("UniqueGenerator.Generate() generated a serial number out of range, %v", number)
					return
				}
				diff += ((lastNum + 1) % 1000000) - (number)
				lastNum = number
				contNumbers[toString(g.ContNum())] = true
			}

			if tt.seqSerialNum && diff != 0 {
				t.Errorf("UniqueGenerator.Generate() generated not sequential serial numbers, diff is %v", diff)
				return
			}

			if got := len(contNumbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueGenerator.Generate() generated %v container numbers, want %v", got, tt.want)
			}
		})
	}
}

func toString(contNum Number) string {
	return contNum.OwnerCode() +
		fmt.Sprintf("%d", contNum.CheckDigit()) +
		contNum.SerialNumber() +
		fmt.Sprintf("%d", contNum.CheckDigit())
}
