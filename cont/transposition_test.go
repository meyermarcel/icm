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
	"reflect"
	"testing"
)

func TestCheckTransposition(t *testing.T) {
	type args struct {
		ownerCode  string
		equipCatID string
		serialNum  string
	}
	tests := []struct {
		name string
		args args
		want []Number
	}{
		{
			name: "Test CMA U 163912 (1)0",
			args: args{
				ownerCode:  "CMA",
				equipCatID: "U",
				serialNum:  "163912",
			},
			want: []Number{
				newNum("CMA", "U", "169312", 0),
				newNum("CMA", "U", "163192", 0),
			},
		},
		{
			name: "Test RCB U 001130 0",
			args: args{
				ownerCode:  "RCB",
				equipCatID: "U",
				serialNum:  "001130",
			},
			want: []Number{
				newNum("RCB", "U", "010130", 0),
			},
		},
		{ // WSL U 801743 0
			name: "Test WSL U 801743 (1)0",
			args: args{
				ownerCode:  "WSL",
				equipCatID: "U",
				serialNum:  "801743",
			},
			want: []Number{
				newNum("WSL", "U", "810743", 0),
				newNum("WSL", "U", "807143", 0),
				newNum("WSL", "U", "801740", 3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckTransposition(tt.args.ownerCode, tt.args.equipCatID, tt.args.serialNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckTransposition() = %v, want %v", got, tt.want)
			}
		})
	}
}
