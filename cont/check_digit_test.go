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
	"github.com/meyermarcel/iso6346/equip_cat"
	"github.com/meyermarcel/iso6346/owner"
	"testing"
)

func TestCalcCheckDigit(t *testing.T) {
	type args struct {
		ownerCode  owner.Code
		equipCatId equip_cat.Id
		serialNum  SerialNum
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test ABC U 123456 0",
			args{owner.NewCode("ABC"), equip_cat.NewIdU(), NewSerialNum(123456)},
			0},
		{"Test NYK U 008685 2",
			args{owner.NewCode("NYK"), equip_cat.NewIdU(), NewSerialNum(8685)},
			2},
		{"Test NYK U 000000 0",
			args{owner.NewCode("NYK"), equip_cat.NewIdU(), NewSerialNum(0)},
			0},
		{"Test CMA U 163912 0",
			args{owner.NewCode("CMA"), equip_cat.NewIdU(), NewSerialNum(163912)},
			0},
		{"Test CMA U 169312 0",
			args{owner.NewCode("CMA"), equip_cat.NewIdU(), NewSerialNum(169312)},
			0},
		{"Test CSQ U 305438 3",
			args{owner.NewCode("CSQ"), equip_cat.NewIdU(), NewSerialNum(305438)},
			3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcCheckDigit(tt.args.ownerCode, tt.args.equipCatId, tt.args.serialNum); got != tt.want {
				t.Errorf("CalcCheckDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
