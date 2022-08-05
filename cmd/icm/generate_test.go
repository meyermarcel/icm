package cmd

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/meyermarcel/icm/configs"
)

func Test_generateCmd(t *testing.T) {
	type configOverride struct {
		name  string
		value string
	}
	type flag struct {
		name  string
		value string
	}
	tests := []struct {
		name            string
		configOverrides []configOverride
		flags           []flag
		wantErr         bool
		wantWriter      string
	}{
		{
			"Generate 1 random container number",
			nil,
			nil,
			false,
			`RAN U 724553 3
`,
		},
		{
			"Generate 1 random container number with custom owner",
			nil,
			[]flag{{
				name:  "owner",
				value: "ABC",
			}},
			false,
			`ABC U 724553 6
`,
		},
		{
			"Generate 3 random container number",
			nil,
			[]flag{{
				name:  "count",
				value: "3",
			}},
			false,
			`RAN U 724553 3
RAN U 165715 3
RAN U 489155 0
`,
		},
		{
			"Generate 3 container number with sequential serial number excluding check digit 10",
			nil,
			[]flag{
				{
					name:  "start",
					value: "0",
				},
				{
					name:  "count",
					value: "3",
				},
				{
					name:  "exclude-check-digit-10",
					value: "true",
				},
			},
			false,
			`RAN U 000000 9
RAN U 000001 4
RAN U 000003 5
`,
		},
		{
			"Generate 3 container number with sequential serial number with start",
			nil,
			[]flag{
				{
					name:  "count",
					value: "3",
				},
				{
					name:  "start",
					value: "0",
				},
			},
			false,
			`RAN U 000000 9
RAN U 000001 4
RAN U 000002 0
`,
		},
		{
			"Generate 3 container number with sequential serial number with end",
			nil,
			[]flag{
				{
					name:  "count",
					value: "3",
				},
				{
					name:  "end",
					value: "0",
				},
			},
			false,
			`RAN U 999998 0
RAN U 999999 6
RAN U 000000 9
`,
		},
		{
			"Generate 3 container number with sequential serial number with range",
			nil,
			[]flag{
				{
					name:  "start",
					value: "0",
				},
				{
					name:  "end",
					value: "2",
				},
			},
			false,
			`RAN U 000000 9
RAN U 000001 4
RAN U 000002 0
`,
		},
		{
			"Generate 1 random container number with custom separators",
			[]configOverride{
				{
					name:  configs.FlagNames.SepOE,
					value: "***",
				},
				{
					name:  configs.FlagNames.SepES,
					value: "+++",
				},
				{
					name:  configs.FlagNames.SepSC,
					value: "‧‧‧",
				},
			},
			nil,
			false,
			`RAN***U+++724553‧‧‧3
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			writerErr := &bytes.Buffer{}

			config, _ := configs.ReadConfig(configs.DefaultConfig())
			for _, override := range tt.configOverrides {
				config.Map[override.name] = override.value
			}

			cmd := newGenerateCmd(writer, writerErr, config, &dummyOwnerDecodeUpdater{})
			for _, flag := range tt.flags {
				_ = cmd.Flags().Set(flag.name, flag.value)
			}
			rand.Seed(1)
			if got := cmd.RunE(cmd, nil); (got == nil) == tt.wantErr {
				t.Errorf("got = %v, wantErr is %v", got, tt.wantErr)
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
