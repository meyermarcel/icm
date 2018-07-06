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
	"testing"
)

func TestCalcCheckDigit(t *testing.T) {
	type args struct {
		ownerCode  OwnerCode
		equipCatID EquipCatID
		serialNum  SerialNum
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test ABC U 123456 0",
			args{NewOwnerCode("ABC"), NewEquipCatIDU(), NewSerialNum(123456)},
			0},
		{"Test NYK U 008685 2",
			args{NewOwnerCode("NYK"), NewEquipCatIDU(), NewSerialNum(8685)},
			2},
		{"Test NYK U 000000 0",
			args{NewOwnerCode("NYK"), NewEquipCatIDU(), NewSerialNum(0)},
			0},
		{"Test CMA U 163912 0",
			args{NewOwnerCode("CMA"), NewEquipCatIDU(), NewSerialNum(163912)},
			0},
		{"Test CMA U 169312 0",
			args{NewOwnerCode("CMA"), NewEquipCatIDU(), NewSerialNum(169312)},
			0},
		{"Test CSQ U 305438 3",
			args{NewOwnerCode("CSQ"), NewEquipCatIDU(), NewSerialNum(305438)},
			3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcCheckDigit(tt.args.ownerCode, tt.args.equipCatID, tt.args.serialNum); got != tt.want {
				t.Errorf("calcCheckDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
