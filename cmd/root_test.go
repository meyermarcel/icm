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

import "github.com/meyermarcel/icm/internal/cont"

type dummyOwnerDecodeUpdater struct {
	dummyOwnerDecoder
	dummyOwnerUpdater
}

type dummyOwnerDecoder struct {
}

func (dummyOwnerDecoder) Decode(code string) (bool, cont.Owner) {
	if code != "ABC" {
		return false, cont.Owner{}
	}
	return true, cont.Owner{
		Code:    "ABC",
		Company: "some-company",
		City:    "some-city",
		Country: "some-country",
	}
}

type dummyOwnerUpdater struct {
}

func (dummyOwnerUpdater) GetAllOwnerCodes() []string {
	return []string{"RAN"}
}

func (dummyOwnerUpdater) Update(newOwners map[string]cont.Owner) error {
	panic("implement me")
}

type dummyEquipCatDecoder struct {
}

func (dummyEquipCatDecoder) Decode(ID string) (bool, cont.EquipCat) {
	return true, cont.EquipCat{
		Value: ID,
		Info:  "some-equip-cat-ID",
	}
}

func (dummyEquipCatDecoder) AllCatIDs() []string {
	return []string{"U"}
}

type dummyLengthDecoder struct {
}

func (dummyLengthDecoder) Decode(code string) (bool, cont.Length) {
	return true, cont.Length{
		Length: "some-length",
	}
}

type dummyHeightWidthDecoder struct {
}

func (dummyHeightWidthDecoder) Decode(code string) (bool, cont.HeightWidth) {
	return true, cont.HeightWidth{
		Width:  "some-width",
		Height: "some-height",
	}
}

type dummyTypeDecoder struct {
}

func (dummyTypeDecoder) Decode(code string) (bool, cont.TypeAndGroup) {
	return true, cont.TypeAndGroup{
		Type: cont.Type{
			TypeCode: code,
			TypeInfo: "some-type",
		},
		Group: cont.Group{
			GroupCode: code,
			GroupInfo: "some-group",
		},
	}
}
