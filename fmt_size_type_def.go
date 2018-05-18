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

package main

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func fmtSizeTypeDef(def sizeTypeDef) string {

	b := strings.Builder{}

	tableSize := tablewriter.NewWriter(os.Stdout)
	tableSize.SetHeader([]string{"Code", "Length", "Code", "Height", "Width"})
	tableSize.SetAlignment(tablewriter.ALIGN_RIGHT)

	for i, element := range def.Lengths {
		var row []string

		row = append([]string{element.Code, element.Length})

		if i < len(def.HeightsWidths) {
			row = append(row, def.HeightsWidths[i].Code, def.HeightsWidths[i].Height, def.HeightsWidths[i].Width)
		} else {
			row = append(row, "", "", "")
		}

		tableSize.Append(row)
	}

	tableSize.Render()

	tableType := tablewriter.NewWriter(os.Stdout)
	tableType.SetHeader([]string{"Code", "Type"})
	tableType.SetAlignment(tablewriter.ALIGN_LEFT)
	tableType.SetAutoWrapText(false)

	for _, element := range def.Types {
		tableType.Append([]string{element.Code, element.TypeInfo})
	}

	tableType.Render()

	tableGroup := tablewriter.NewWriter(os.Stdout)
	tableGroup.SetHeader([]string{"Code", "Group"})
	tableGroup.SetAlignment(tablewriter.ALIGN_LEFT)
	tableGroup.SetAutoWrapText(false)

	for _, group := range def.Groups {

		tableGroup.Append([]string{group.Code, group.GroupInfo})
	}

	tableGroup.Render()
	return b.String()
}
