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

package main

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/meyermarcel/icm/configs"

	"github.com/spf13/viper"
)

func Test_generateCmd(t *testing.T) {
	type cfgOverride struct {
		name  string
		value string
	}
	type flag struct {
		name  string
		value string
	}
	tests := []struct {
		name         string
		cfgOverrides []cfgOverride
		flags        []flag
		wantErr      bool
		wantWriter   string
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
			[]cfgOverride{
				{
					name:  configs.SepOE,
					value: "***",
				},
				{
					name:  configs.SepES,
					value: "+++",
				},
				{
					name:  configs.SepSC,
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
			viperCfg := viper.New()
			for _, override := range tt.cfgOverrides {
				viperCfg.Set(override.name, override.value)
			}
			cmd := newGenerateCmd(writer, writerErr, viperCfg, &dummyOwnerDecodeUpdater{})
			for _, flag := range tt.flags {
				_ = cmd.Flags().Set(flag.name, flag.value)
			}
			_ = cmd.PreRunE(cmd, nil)
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
