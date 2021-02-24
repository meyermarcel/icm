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

package file

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/meyermarcel/icm/data"

	"github.com/meyermarcel/icm/cont"
)

const (
	typeFileName  = "type.json"
	groupFileName = "group.json"
)

type typeAndGroupDecoder struct {
	types  map[string]string
	groups map[string]string
}

// NewTypeDecoder writes type and group file to path if it not exists and
// returns a struct that uses this file as a data source.
func NewTypeDecoder(path string) (data.TypeDecoder, error) {
	typeAndGroup := &typeAndGroupDecoder{}
	pathToType := filepath.Join(path, typeFileName)
	if err := initFile(pathToType, []byte(typeJSON)); err != nil {
		return nil, err
	}
	b, err := os.ReadFile(pathToType)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &typeAndGroup.types); err != nil {
		return nil, err
	}

	pathToGroup := filepath.Join(path, groupFileName)
	if err := initFile(pathToGroup, []byte(groupJSON)); err != nil {
		return nil, err
	}
	b, err = os.ReadFile(pathToGroup)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &typeAndGroup.groups); err != nil {
		return nil, err
	}
	for typeCode := range typeAndGroup.types {
		if err := cont.IsTypeCode(typeCode); err != nil {
			return nil, err
		}
	}
	return typeAndGroup, nil
}

// Decode returns type and group for type code.
func (tg *typeAndGroupDecoder) Decode(code string) (bool, cont.TypeAndGroup) {
	typeAndGroup := cont.TypeAndGroup{}

	typeInfo, exists := tg.types[code]
	typeValue, typeFound := cont.Type{TypeCode: code, TypeInfo: typeInfo}, exists

	info, exists := tg.groups[string(code[0])]
	group, groupFound := cont.Group{GroupCode: code, GroupInfo: info}, exists

	if !typeFound || !groupFound {
		return false, typeAndGroup
	}

	typeAndGroup.Type = typeValue
	typeAndGroup.Group = group
	return true, typeAndGroup
}

const typeJSON = `{
  "G0": "General - Openings at one or both ends",
  "G1": "General - Passive vents at upper part of cargo space",
  "G2": "General - Openings at one or both ends + full openings on one or both sides",
  "G3": "General - Openings at one or both ends + partial openings on one or both sides",
  "V0": "Fantainer - Non-mechanical, vents at lower and upper parts of cargo space",
  "V2": "Fantainer - Mechanical ventilation system located internally",
  "V4": "Fantainer - Mechanical ventilation system located externally",
  "R0": "Integral Reefer - Mechanically refrigerated",
  "R1": "Integral Reefer - Mechanically refrigerated and heated",
  "R2": "Integral Reefer - Self-powered mechanically refrigerated",
  "R3": "Integral Reefer - Self-powered mechanically refrigerated and heated",
  "H0": "Refrigerated or heated with removable equipment located externally; heat transfer coefficient K=0.4W/M2.K",
  "H1": "Refrigerated or heated with removable equipment located internally",
  "H2": "Refrigerated or heated with removable equipment located externally; heat transfer coefficient K=0.7W/M2.K",
  "H5": "Insulated - Heat transfer coefficient K=0.4W/M2.K",
  "H6": "Insulated - Heat transfer coefficient K=0.7W/M2.K",
  "U0": "Open Top - Openings at one or both ends",
  "U1": "Open Top - Idem + removable top members in end frames",
  "U2": "Open Top - Openings at one or both ends + openings at one or both sides",
  "U3": "Open Top - Idem + removable top members in end frames",
  "U4": "Open Top - Openings at one or both ends + partial on one and full at other side",
  "U5": "Open Top - Complete, fixed side and end walls ( no doors )",
  "T0": "Tank - Non dangerous liquids, minimum pressure 0.45 bar",
  "T1": "Tank - Non dangerous liquids, minimum pressure 1.50 bar",
  "T2": "Tank - Non dangerous liquids, minimum pressure 2.65 bar",
  "T3": "Tank - Dangerous liquids, minimum pressure 1.50 bar",
  "T4": "Tank - Dangerous liquids, minimum pressure 2.65 bar",
  "T5": "Tank - Dangerous liquids, minimum pressure 4.00 bar",
  "T6": "Tank - Dangerous liquids, minimum pressure 6.00 bar",
  "T7": "Tank - Gases, minimum pressure 9.10 bar",
  "T8": "Tank - Gases, minimum pressure 22.00 bar",
  "T9": "Tank - Gases, minimum pressure to be decided",
  "B0": "Bulk - Closed",
  "B1": "Bulk - Airtight",
  "B3": "Bulk - Horizontal discharge, test pressure 1.50 bar",
  "B4": "Bulk - Horizontal discharge, test pressure 2.65 bar",
  "B5": "Bulk - Tipping discharge, test pressure 1.50 bar",
  "B6": "Bulk - Tipping discharge, test pressure 2.65 bar",
  "P0": "Flat or Bolster - Plain platform",
  "P1": "Flat or Bolster - Two complete and fixed ends",
  "P2": "Flat or Bolster - Fixed posts, either free-standing or with removable top member",
  "P3": "Flat or Bolster - Folding complete end structure",
  "P4": "Flat or Bolster - Folding posts, either free-standing or with removable top member",
  "P5": "Flat or Bolster - Open top, open ends (skeletal)",
  "S0": "Livestock carrier",
  "S1": "Automobile carrier",
  "S2": "Live fish carrier"
}
`

const groupJSON = `{
  "A": "Air/surface container",
  "B": "Bulk container",
  "G": "General purpose container",
  "H": "Insulated container",
  "N": "Pressurized and non-pressurized tank container (dry)",
  "P": "Flat",
  "R": "Thermal container",
  "S": "Named cargo container",
  "T": "Tank container",
  "U": "Open-top/hardtop container",
  "V": "Ventilated container"
}
`
