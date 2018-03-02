package ui

import (
	"github.com/meyermarcel/iso6346/sizetype"
	"strings"
	"os"
	"github.com/olekukonko/tablewriter"
)

func fmtSizeTypeDef(def sizetype.Def) string {

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
		tableType.Append([]string{element.Code, element.Info})
	}

	tableType.Render()

	tableGroup := tablewriter.NewWriter(os.Stdout)
	tableGroup.SetHeader([]string{"Code", "Group",})
	tableGroup.SetAlignment(tablewriter.ALIGN_LEFT)
	tableGroup.SetAutoWrapText(false)

	for _, group := range def.Groups {

		tableGroup.Append([]string{group.Code, group.Info})
	}

	tableGroup.Render()
	return b.String()
}
