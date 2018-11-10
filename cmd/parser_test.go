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

package cmd

import (
	"testing"

	"github.com/meyermarcel/icm/internal/cont"
)

type dummyOwnerDecodeUpdater struct {
}

func (dummyOwnerDecodeUpdater) Decode(code cont.OwnerCode) cont.Owner {
	panic("implement me")
}

func (dummyOwnerDecodeUpdater) AllCodes() []string {
	return []string{"ABC"}
}

func (dummyOwnerDecodeUpdater) GenerateRandomCodes(count int) []cont.OwnerCode {
	panic("implement me")
}

func (dummyOwnerDecodeUpdater) Update(newOwners map[string]cont.Owner) error {
	panic("implement me")
}

type dummyLengthDecoder struct {
}

func (dummyLengthDecoder) Decode(code string) cont.Length {
	panic("implement me")
}

func (dummyLengthDecoder) AllCodes() []string {
	return []string{"2"}
}

type dummyHeightWidthDecoder struct {
}

func (dummyHeightWidthDecoder) Decode(code string) cont.HeightWidth {
	panic("implement me")
}

func (dummyHeightWidthDecoder) AllCodes() []string {
	return []string{"0"}
}

type dummyTypeDecoder struct {
}

func (dummyTypeDecoder) Decode(code string) cont.TypeAndGroup {
	panic("implement me")
}

func (dummyTypeDecoder) AllCodes() []string {
	return []string{"A1"}
}

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
