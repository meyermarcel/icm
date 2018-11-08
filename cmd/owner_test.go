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
	"bytes"
	"testing"

	"github.com/meyermarcel/icm/internal/cont"
)

type OwnerDummySource struct {
}

func (of *OwnerDummySource) Decode(code cont.OwnerCode) cont.Owner {
	return cont.Owner{
		Code:    code,
		Company: "some-company",
		City:    "some-city",
		Country: "some-country",
	}
}

func (of *OwnerDummySource) AllCodes() []string {
	return []string{"ABC", "DEF"}
}

func (of *OwnerDummySource) GenerateRandomCodes(count int) []cont.OwnerCode {
	return []cont.OwnerCode{cont.NewOwnerCode("RAN")}
}

func Test_newValidateOwnerCmd(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantErr    bool
		wantWriter string
	}{
		{
			"Parse and validate existing code returns no error",
			[]string{"abc"},
			false,
			`'abc'

 ABC  ✔
  │
  └─ some-company
     some-city
     some-country

`,
		},
		{
			"Parse and validate existing other code returns no error",
			[]string{"def"},
			false,
			`'def'

 DEF  ✔
  │
  └─ some-company
     some-city
     some-country

`,
		},
		{
			"Parse and validate non existing code returns error",
			[]string{"noe"},
			true,
			`'noe'

 ___  ✘
  ↑
  └─ owner code must be 3 letters and registered (e.g. RAN)

`,
		},
		{
			"Parse and validate existing code returns error",
			[]string{"§$©€% abc &&"},
			false,
			`'§$©€% abc &&'

 ABC  ✔
  │
  └─ some-company
     some-city
     some-country

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if got := newValidateOwnerCmd(writer, &OwnerDummySource{}).RunE(nil, tt.args); (got == nil) == tt.wantErr {
				t.Errorf("newValidateOwnerCmd() = %v, wantErr is %v", got, tt.wantErr)
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("newValidateOwnerCmd() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
