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

package ui

import (
	"reflect"
	"testing"
	"fmt"
)

func TestNewPosHint(t *testing.T) {
	type args struct {
		pos   int
		lines []string
	}
	tests := []struct {
		name string
		args args
		want PosTxt
	}{
		{
			"Constructor creates correct positioned hint",
			args{2, []string{"line1", "line2"}},
			PosTxt{2, 0, []string{"line1", "line2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPosHint(tt.args.pos, tt.args.lines...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPosHint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPosInfo(t *testing.T) {
	type args struct {
		pos   int
		lines []string
	}
	tests := []struct {
		name string
		args args
		want PosTxt
	}{
		{
			"Constructor creates correct positioned info",
			args{2, []string{"line1", "line2"}},
			PosTxt{2, 1, []string{"line1", "line2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPosInfo(tt.args.pos, tt.args.lines...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPosInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fmtTextsWithArrows(t *testing.T) {
	type args struct {
		texts []PosTxt
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Format vertical line",
			args{[]PosTxt{
				NewPosInfo(0, ""),
			}},
			`
│
└─ `,
		},
		{
			"Format arrow",
			args{[]PosTxt{
				NewPosHint(0, ""),
			}},
			`
↑
└─ `,
		},
		{
			"Format indent",
			args{[]PosTxt{
				NewPosHint(2, ""),
			}},
			`
  ↑
  └─ `,
		},
		{
			"Format text",
			args{[]PosTxt{
				NewPosHint(0, "text"),
			}},
			`
↑
└─ text`,
		},
		{
			"Format with multiple lines",
			args{[]PosTxt{
				NewPosHint(0, "line0", "line1"),
			}},
			`
↑
└─ line0
   line1`,
		},
		{
			"Format multiple texts",
			args{[]PosTxt{
				NewPosHint(0, "text"),
				NewPosInfo(2, "text"),
			}},
			`
↑ │
│ └─ text
│
└─ text`,
		},
		{
			"Format multiple lines in second text",
			args{[]PosTxt{
				NewPosHint(0, "text"),
				NewPosInfo(2, "line0", "line1"),
			}},
			`
↑ │
│ └─ line0
│    line1
│
└─ text`,
		},
		{
			"Format multiple lines in multiple texts",
			args{[]PosTxt{
				NewPosInfo(0, "pos0line0", "pos0line1", "pos0line2"),
				NewPosHint(1, "pos1line0", "pos1line1"),
				NewPosInfo(7, "pos7line0", "pos7line1", "pos7line2"),
				NewPosHint(22, "pos22line0"),
			}},
			`
│↑     │              ↑
││     │              └─ pos22line0
││     │
││     └─ pos7line0
││        pos7line1
││        pos7line2
││
│└─ pos1line0
│   pos1line1
│
└─ pos0line0
   pos0line1
   pos0line2`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fmt.Sprintln() + fmtTextsWithArrows(tt.args.texts...); got != tt.want {
				t.Errorf("fmtTextsWithArrows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateSpaces(t *testing.T) {
	type args struct {
		texts []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"No indent",
			args{[]int{0}},
			[]string{""},
		},
		{
			"No space between two positions",
			args{[]int{0, 1}},
			[]string{"", ""},
		},
		{
			"Spaces between two positions",
			args{[]int{0, 3}},
			[]string{"", "  "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSpaces(tt.args.texts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
