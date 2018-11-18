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
	"testing"
)

func TestNewRandomUniqueGenerator(t *testing.T) {
	type args struct {
		count int
		codes []OwnerCode
	}
	tests := []struct {
		name    string
		args    args
		want    *RandomUniqueGenerator
		wantErr bool
	}{
		{
			"Constructor returns correct RandomUniqueGenerator",
			args{3, []OwnerCode{NewOwnerCode("ABC"), NewOwnerCode("DEF"), NewOwnerCode("GHI")}},
			&RandomUniqueGenerator{
				codes:      []OwnerCode{NewOwnerCode("ABC"), NewOwnerCode("DEF"), NewOwnerCode("GHI")},
				lenCodes:   3,
				randOffset: 5577006791947779410,
			},
			false,
		},
		{
			"Exceed maximum of unique container numbers",
			args{909092, []OwnerCode{NewOwnerCode("ABC")}},
			nil,
			true,
		},
		{
			"Count is lower than minimum count 1",
			args{0, []OwnerCode{NewOwnerCode("ABC")}},
			nil,
			true,
		},
		{
			"Minimum one owner code is needed",
			args{1, []OwnerCode{}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(1)
			got, err := NewRandomUniqueGenerator(tt.args.count, tt.args.codes)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRandomUniqueGenerator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRandomUniqueGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomUniqueGenerator_Generate(t *testing.T) {
	tests := []struct {
		name  string
		count int
		want  int
	}{
		{
			"Generate 1 container number",
			1,
			1,
		},
		{
			"Generate 12 unique container numbers",
			12,
			12,
		},
		{
			"Exceed maximum amount of unique container numbers",
			909092,
			909091,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, _ := NewRandomUniqueGenerator(tt.count, []OwnerCode{NewOwnerCode("ABC"), NewOwnerCode("DEF")})

			contNumbers := map[string]bool{}
			for i := 0; i < tt.count; i++ {
				contNumbers[contNumToString(g.Generate())] = true
			}

			if got := len(contNumbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomUniqueGenerator.Generate() = %v, want %v", got, tt.count)
			}
		})
	}
}

func contNumToString(contNum Number) string {
	return contNum.OwnerCode().Value() +
		fmt.Sprintf("%d", contNum.CheckDigit()) +
		contNum.SerialNumber().String() +
		fmt.Sprintf("%d", contNum.CheckDigit())
}
