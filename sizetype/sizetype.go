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

package sizetype

type Size struct {
	Code        string
	HeightWidth MappedHeightAndWidth
	Length      MappedLength
}

type MappedLength struct {
	Code   string
	Length string
}

type MappedHeightAndWidth struct {
	Code   string
	Width  string
	Height string
}

type MappedType struct {
	Code     string
	TypeInfo string
}

type MappedGroup struct {
	Code      string
	GroupInfo string
}

type MappedTypeAndGroup struct {
	MappedType  MappedType
	MappedGroup MappedGroup
}

type Def struct {
	Lengths       []MappedLength
	HeightsWidths []MappedHeightAndWidth
	Types         []MappedType
	Groups        []MappedGroup
}

var mappedLengths = []MappedLength{

	{"1", "2991 mm"},
	{"2", "6068 mm"},
	{"3", "9125 mm"},
	{"4", "12192 mm"},
	{"A", "7150 mm"},
	{"B", "7315 mm"},
	{"C", "7430 mm"},
	{"D", "7450 mm"},
	{"E", "7820 mm"},
	{"F", "8100 mm"},
	{"G", "12500 mm"},
	{"H", "13106 mm"},
	{"K", "13600 mm"},
	{"L", "13716 mm"},
	{"M", "14630 mm"},
	{"N", "14935 mm"},
	{"P", "16154 mm"},
}

func GetLength(code string) (MappedLength, bool) {

	for _, lengthElement := range mappedLengths {
		if lengthElement.Code == code {
			return lengthElement, true
		}
	}
	return MappedLength{}, false
}

var mappedHeightAndWidths = []MappedHeightAndWidth{
	{"0", "2436 mm", "2438 mm"},
	{"2", "2436 mm", "2591 mm"},
	{"4", "2436 mm", "2743 mm"},
	{"5", "2436 mm", "2895 mm"},
	{"6", "2436 mm", "> 2895 mm"},
	{"8", "2436 mm", "1295 mm"},
	{"9", "2436 mm", "< 1219 mm"},
	{"C", "> 2438 mm and ≤ 2500 mm", "2591 mm"},
	{"D", "> 2438 mm and ≤ 2500 mm", "2743 mm"},
	{"E", "> 2438 mm and ≤ 2500 mm", "2895 mm"},
	{"F", "> 2438 mm and ≤ 2500 mm", "> 2895 mm"},
	{"L", "> 2500 mm", "2591 mm"},
	{"M", "> 2500 mm", "2743 mm"},
	{"N", "> 2500 mm", "2895 mm"},
	{"P", "> 2500 mm", "> 2895 mm"},
}

func GetHeightAndWidth(code string) (MappedHeightAndWidth, bool) {

	for _, heightAndWidthElement := range mappedHeightAndWidths {
		if heightAndWidthElement.Code == code {
			return heightAndWidthElement, true

		}
	}
	return MappedHeightAndWidth{}, false
}

var mappedTypes = []MappedType{
	{"A0", "(unassigned)"},
	{"B0", "(unassigned)"},
	{"B1", "(unassigned)"},
	{"B2", "(unassigned)"},
	{"B3", "(unassigned)"},
	{"B4", "(unassigned)"},
	{"B5", "(unassigned)"},
	{"B6", "(unassigned)"},
	{"B7", "(unassigned)"},
	{"B8", "(unassigned)"},
	{"B9", "(unassigned)"},
	{"G0", "(unassigned)"},
	{"G1", "(unassigned)"},
	{"G2", "(unassigned)"},
	{"G3", "(unassigned)"},
	{"G4", "(unassigned)"},
	{"G5", "(unassigned)"},
	{"G6", "(unassigned)"},
	{"G7", "(unassigned)"},
	{"G8", "(unassigned)"},
	{"G9", "(unassigned)"},
	{"H0", "(unassigned)"},
	{"H1", "(unassigned)"},
	{"H2", "(unassigned)"},
	{"H3", "(unassigned)"},
	{"H4", "(unassigned)"},
	{"H5", "(unassigned)"},
	{"H6", "(unassigned)"},
	{"H7", "(unassigned)"},
	{"H8", "(unassigned)"},
	{"H9", "(unassigned)"},
	{"K0", "(unassigned)"},
	{"K1", "(unassigned)"},
	{"K2", "(unassigned)"},
	{"K3", "(unassigned)"},
	{"K4", "(unassigned)"},
	{"K5", "(unassigned)"},
	{"K6", "(unassigned)"},
	{"K7", "(unassigned)"},
	{"K8", "(unassigned)"},
	{"K9", "(unassigned)"},
	{"N0", "(unassigned)"},
	{"N1", "(unassigned)"},
	{"N2", "(unassigned)"},
	{"N3", "(unassigned)"},
	{"N4", "(unassigned)"},
	{"N5", "(unassigned)"},
	{"N6", "(unassigned)"},
	{"N7", "(unassigned)"},
	{"N8", "(unassigned)"},
	{"N9", "(unassigned)"},
	{"P0", "(unassigned)"},
	{"P1", "(unassigned)"},
	{"P2", "(unassigned)"},
	{"P3", "(unassigned)"},
	{"P4", "(unassigned)"},
	{"P5", "(unassigned)"},
	{"P6", "(unassigned)"},
	{"P7", "(unassigned)"},
	{"P8", "(unassigned)"},
	{"P9", "(unassigned)"},
	{"R0", "(unassigned)"},
	{"R1", "(unassigned)"},
	{"R2", "(unassigned)"},
	{"R3", "(unassigned)"},
	{"R4", "(unassigned)"},
	{"R5", "(unassigned)"},
	{"R6", "(unassigned)"},
	{"R7", "(unassigned)"},
	{"R8", "(unassigned)"},
	{"R9", "(unassigned)"},
	{"S0", "(unassigned)"},
	{"S1", "(unassigned)"},
	{"S2", "(unassigned)"},
	{"S3", "(unassigned)"},
	{"S4", "(unassigned)"},
	{"S5", "(unassigned)"},
	{"S6", "(unassigned)"},
	{"S7", "(unassigned)"},
	{"S8", "(unassigned)"},
	{"S9", "(unassigned)"},
	{"U0", "(unassigned)"},
	{"U1", "(unassigned)"},
	{"U2", "(unassigned)"},
	{"U3", "(unassigned)"},
	{"U4", "(unassigned)"},
	{"U5", "(unassigned)"},
	{"U6", "(unassigned)"},
	{"U7", "(unassigned)"},
	{"U8", "(unassigned)"},
	{"U9", "(unassigned)"},
	{"V0", "(unassigned)"},
	{"V1", "(unassigned)"},
	{"V2", "(unassigned)"},
	{"V3", "(unassigned)"},
	{"V4", "(unassigned)"},
	{"V5", "(unassigned)"},
	{"V6", "(unassigned)"},
	{"V7", "(unassigned)"},
	{"V8", "(unassigned)"},
	{"V9", "(unassigned)"},
}

func getType(code string) (MappedType, bool) {

	for _, typeElement := range mappedTypes {
		if typeElement.Code == code {
			return typeElement, true
		}
	}
	return MappedType{}, false
}

var mappedGroups = []MappedGroup{
	{"A", "(unassigned)"},
	{"B", "(unassigned)"},
	{"G", "(unassigned)"},
	{"H", "(unassigned)"},
	{"K", "(unassigned)"},
	{"N", "(unassigned)"},
	{"P", "(unassigned)"},
	{"R", "(unassigned)"},
	{"S", "(unassigned)"},
	{"U", "(unassigned)"},
	{"V", "(unassigned)"},
}

func getGroup(code string) (MappedGroup, bool) {

	for _, groupElement := range mappedGroups {
		if groupElement.Code == code {
			return groupElement, true
		}
	}
	return MappedGroup{}, false
}

func GetTypeAndGroup(code string) (MappedTypeAndGroup, bool) {
	typeAndGroup := MappedTypeAndGroup{}
	typeValue, typeFound := getType(code)
	group, groupFound := getGroup(string(code[0]))

	if !typeFound && !groupFound {
		return typeAndGroup, false
	}

	typeAndGroup.MappedType = typeValue
	typeAndGroup.MappedGroup = group
	return typeAndGroup, true
}

func GetDef() Def {
	return Def{
		mappedLengths,
		mappedHeightAndWidths,
		mappedTypes,
		mappedGroups,
	}
}
