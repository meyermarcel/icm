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
	"strings"
)

/*
This method is a modified version of a Go code sample from
https://en.wikipedia.org/wiki/ISO_6346#Code_Sample_(Go)
*/
func CalcCheckDigit(ownerCode owner.Code, equipCatId equip_cat.Id, serialNum SerialNum) int {

	concat := ownerCode.Value() + equipCatId.Value()

	n := 0.0
	d := 0.5
	for _, character := range concat {
		d *= 2
		n += d * float64(strings.IndexRune("??????????A?BCDEFGHIJK?LMNOPQRSTU?VWXYZ", character))
	}
	div := 100000.0
	for i := 0; i < 6; i++ {
		d *= 2
		n += d * float64(int(float64(serialNum.Value())/div)%10)
		div /= 10
	}
	return (int(n) - int(n/11)*11) % 10
}
