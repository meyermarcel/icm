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

	"github.com/meyermarcel/icm/configs"
	"github.com/spf13/viper"

	"github.com/meyermarcel/icm/internal/cont"
)

func Test_singleLine(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			"Single line",
			"line",
			true,
		},
		{
			"Multiple lines",
			"line1\nline2",
			false,
		},
		{
			"Single line with newline",
			"line\n",
			true,
		},
		{
			"Multiple lines with empty lines and with multiple newlines",
			"line\n\n\n",
			false,
		},
		{
			"Multiple lines with multiple newlines between lines",
			"line1\n\n\nline2\n",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSingleLine(tt.arg); got != tt.want {
				t.Errorf("isSingleLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

type dummyOwnerDecodeUpdater struct {
	dummyOwnerDecoder
	dummyOwnerUpdater
}

type dummyOwnerDecoder struct {
}

func (dummyOwnerDecoder) Decode(code string) (bool, cont.Owner) {
	if code != "ABC" {
		return false, cont.Owner{}
	}
	return true, cont.Owner{
		Code:    "ABC",
		Company: "some-company",
		City:    "some-city",
		Country: "some-country",
	}
}

type dummyOwnerUpdater struct {
}

func (dummyOwnerUpdater) GenerateRandomCodes(count int) []string {
	return []string{"RAN"}
}

func (dummyOwnerUpdater) Update(newOwners map[string]cont.Owner) error {
	panic("implement me")
}

type dummyEquipCatDecoder struct {
}

func (dummyEquipCatDecoder) Decode(ID string) (bool, cont.EquipCat) {
	return true, cont.EquipCat{
		Value: ID,
		Info:  "some-equip-cat-ID",
	}
}

func (dummyEquipCatDecoder) AllCatIDs() []string {
	return []string{"U"}
}

type dummyLengthDecoder struct {
}

func (dummyLengthDecoder) Decode(code string) (bool, cont.Length) {
	return true, cont.Length{
		Length: "some-length",
	}
}

type dummyHeightWidthDecoder struct {
}

func (dummyHeightWidthDecoder) Decode(code string) (bool, cont.HeightWidth) {
	return true, cont.HeightWidth{
		Width:  "some-width",
		Height: "some-height",
	}
}

type dummyTypeDecoder struct {
}

func (dummyTypeDecoder) Decode(code string) (bool, cont.TypeAndGroup) {
	return true, cont.TypeAndGroup{
		Type: cont.Type{
			TypeCode: code,
			TypeInfo: "some-type",
		},
		Group: cont.Group{
			GroupCode: code,
			GroupInfo: "some-group",
		},
	}
}

func Test_validateCmd(t *testing.T) {
	type cfgOverride struct {
		name  string
		value string
	}
	tests := []struct {
		name         string
		args         []string
		cfgOverrides []cfgOverride
		wantErr      bool
		wantWriter   string
	}{
		{
			"Validate owner, equipment category ID, serial number, check digit, size and type",
			[]string{" abc u 123456 0 20 g1 "},
			nil,
			false,
			`
  ABC U 123456 0   20 G1  ✔
   ↑  ↑            ↑↑  ↑
   │  │            ││  └─ type:  some-type
   │  │            ││     group: some-group
   │  │            ││
   │  │            │└─ height: some-height
   │  │            │   width:  some-width
   │  │            │
   │  │            └─ length: some-length
   │  │
   │  └─ some-equip-cat-ID
   │
   └─ some-company
      some-city
      some-country

`,
		},
		{
			"Validate owner, equipment category ID, serial number and check digit",
			[]string{" abc u 123456 0 "},
			nil,
			false,
			`
  ABC U 123456 0  ✔
   ↑  ↑
   │  └─ some-equip-cat-ID
   │
   └─ some-company
      some-city
      some-country

`,
		},
		{
			"Validate owner and equipment category ID",
			[]string{" abc u "},
			nil,
			false,
			`
  ABC U  ✔
   ↑  ↑
   │  └─ some-equip-cat-ID
   │
   └─ some-company
      some-city
      some-country

`,
		},
		{
			"Validate owner",
			[]string{" abc "},
			nil,
			false,
			`
  ABC  ✔
   ↑
   └─ some-company
      some-city
      some-country

`,
		},
		{
			"Validate size and type",
			[]string{" 20 g1 "},
			nil,
			false,
			`
  20 G1  ✔
  ↑↑  ↑
  ││  └─ type:  some-type
  ││     group: some-group
  ││
  │└─ height: some-height
  │   width:  some-width
  │
  └─ length: some-length

`,
		},
		{
			"Validate input with pattern container-number",
			[]string{" abc "},
			[]cfgOverride{{configs.Pattern, containerNumber}},
			true,
			`
  ABC _ ______ _  ✘
   ↑  ↑    ↑   ↑
   │  │    │   └─ check digit is not calculable
   │  │    │
   │  │    └─ serial number is not 6 numbers long
   │  │
   │  └─ equipment category id is not U
   │
   └─ some-company
      some-city
      some-country

`,
		},
		{
			"Validate input with pattern owner",
			[]string{" abc u "},
			[]cfgOverride{{configs.Pattern, owner}},
			false,
			`
  ABC  ✔
   ↑
   └─ some-company
      some-city
      some-country

`,
		},
		{
			"Validate input with pattern owner-equipment-category",
			[]string{" abc "},
			[]cfgOverride{{configs.Pattern, ownerEquipmentCategory}},
			true,
			`
  ABC _  ✘
   ↑  ↑
   │  └─ equipment category id is not U
   │
   └─ some-company
      some-city
      some-country

`,
		},
		{
			"Validate input with pattern size-type",
			[]string{" abc "},
			[]cfgOverride{{configs.Pattern, sizeType}},
			true,
			`
  AB __  ✘
  ↑↑  ↑
  ││  └─ type code is not a valid number or a valid character
  ││
  │└─ height: some-height
  │   width:  some-width
  │
  └─ length: some-length

`,
		},
		{
			"Validate input with custom separators",
			[]string{" abc u 123456 0 20 g1  "},
			[]cfgOverride{
				{configs.SepOE, "***"},
				{configs.SepES, "+++"},
				{configs.SepSC, "‧‧‧"},
				{configs.SepCS, ">>>"},
				{configs.SepST, "---"},
			},
			false,
			`
  ABC***U+++123456‧‧‧0>>>20---G1  ✔
   ↑    ↑                ↑↑    ↑
   │    │                ││    └─ type:  some-type
   │    │                ││       group: some-group
   │    │                ││
   │    │                │└─ height: some-height
   │    │                │   width:  some-width
   │    │                │
   │    │                └─ length: some-length
   │    │
   │    └─ some-equip-cat-ID
   │
   └─ some-company
      some-city
      some-country

`,
		},
		{
			"Validate sizetype input with custom separators",
			[]string{" 20 g1 "},
			[]cfgOverride{
				{configs.SepST, "***"},
			},
			false,
			`
  20***G1  ✔
  ↑↑    ↑
  ││    └─ type:  some-type
  ││       group: some-group
  ││
  │└─ height: some-height
  │   width:  some-width
  │
  └─ length: some-length

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			writerErr := &bytes.Buffer{}
			d := decoders{
				ownerDecodeUpdater: &dummyOwnerDecodeUpdater{},
				equipCatDecoder:    &dummyEquipCatDecoder{},
				sizeTypeDecoders: sizeTypeDecoders{
					&dummyLengthDecoder{},
					&dummyHeightWidthDecoder{},
					&dummyTypeDecoder{},
				},
			}
			viperCfg := viper.New()
			for _, option := range tt.cfgOverrides {
				viperCfg.Set(option.name, option.value)
			}
			cmd := newValidateCmd(writer, writerErr, viperCfg, d)
			if got := cmd.RunE(nil, tt.args); (got == nil) == tt.wantErr {
				t.Errorf("got = %v, wantErr is %v", got, tt.wantErr)
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
