package cmd

import (
	"bytes"
	"testing"

	"github.com/meyermarcel/icm/configs"
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

func Test_validateCmd(t *testing.T) {
	type configOverride struct {
		name  string
		value string
	}
	tests := []struct {
		name            string
		args            []string
		configOverrides []configOverride
		wantErr         bool
		wantWriter      string
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
			[]configOverride{{configs.FlagNames.Pattern, containerNumber}},
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
			[]configOverride{{configs.FlagNames.Pattern, owner}},
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
			[]configOverride{{configs.FlagNames.Pattern, ownerEquipmentCategory}},
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
			[]configOverride{{configs.FlagNames.Pattern, sizeType}},
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
			[]configOverride{
				{configs.FlagNames.SepOE, "***"},
				{configs.FlagNames.SepES, "+++"},
				{configs.FlagNames.SepSC, "‧‧‧"},
				{configs.FlagNames.SepCS, ">>>"},
				{configs.FlagNames.SepST, "---"},
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
			[]configOverride{
				{configs.FlagNames.SepST, "***"},
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
		{
			"Validate container-number with wrong check digit",
			[]string{"abc u 123123 1"},
			[]configOverride{{configs.FlagNames.Pattern, containerNumber}},
			true,
			`
  ABC U 123123 1  ✘
   ↑  ↑        ↑
   │  │        └─ calculated check digit is 7
   │  │
   │  └─ some-equip-cat-ID
   │
   └─ some-company
      some-city
      some-country

`,
		},
		{
			"Validate container-number with letter for check digit",
			[]string{"abc u 123123 a"},
			[]configOverride{{configs.FlagNames.Pattern, containerNumber}},
			true,
			`
  ABC U 123123 _  ✘
   ↑  ↑        ↑
   │  │        └─ check digit must be a number (calculated: 7)
   │  │
   │  └─ some-equip-cat-ID
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
			d := decoders{
				ownerDecodeUpdater: &dummyOwnerDecodeUpdater{},
				equipCatDecoder:    &dummyEquipCatDecoder{},
				sizeTypeDecoders: sizeTypeDecoders{
					&dummyLengthDecoder{},
					&dummyHeightWidthDecoder{},
					&dummyTypeDecoder{},
				},
			}

			config, _ := configs.ReadConfig(configs.DefaultConfig())
			for _, override := range tt.configOverrides {
				config.Map[override.name] = override.value
			}

			cmd := newValidateCmd(nil, writer, config, d)

			if got := cmd.RunE(cmd, tt.args); (got == nil) == tt.wantErr {
				t.Errorf("got = %v, wantErr is %v", got, tt.wantErr)
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
