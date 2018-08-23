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
	"testing"
)

func TestGenNum(t *testing.T) {
	type args struct {
		count  int
		random func(count int) []OwnerCode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Generate no container number",
			args{0, random},
			0,
		},
		{
			"Generate 1 container number",
			args{1, random},
			1,
		},
		{
			"Generate 12 unique container numbers",
			args{12, random},
			12,
		},
		{
			"Generate maximum amount of unique container numbers",
			args{909091, random},
			909091,
		},
		{
			"Generate no container numbers because limit is exceeded",
			args{909092, random},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := make(chan Result)

			go GenNum(tt.args.count, c, tt.args.random)

			contNumbers := map[string]bool{}
			for genContNumber := range c {
				if genContNumber.Err() == nil {
					contNumbers[contNumToString(genContNumber.ContNum())] = true
				}
			}

			if got := len(contNumbers); got != tt.want {
				t.Errorf("GenNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func random(count int) []OwnerCode {
	if count == 0 {
		return []OwnerCode{}
	}
	return []OwnerCode{NewOwnerCode("ABC")}

}

func contNumToString(contNum Number) string {
	return contNum.OwnerCode().Value() +
		fmt.Sprintf("%d", contNum.CheckDigit()) +
		contNum.SerialNumber().String() +
		fmt.Sprintf("%d", contNum.CheckDigit())
}
