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
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/meyermarcel/icm/internal/cont"
	"github.com/meyermarcel/icm/internal/data"
	"github.com/meyermarcel/icm/internal/input"
)

var (
	yellow    = color.New(color.FgYellow).SprintFunc()
	green     = color.New(color.FgGreen).SprintFunc()
	bold      = color.New(color.Bold).SprintFunc()
	underline = color.New(color.Underline).SprintFunc()
)

func newIcmValidator(decoders decoders) *input.Validator {
	owner := input.NewInput(
		3,
		regexp.MustCompile(`[A-Za-z]{3}`).FindStringIndex,
		func(value string, previousValues []string) (bool, []input.Info) {
			if value == "" {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not %s long (e.g. %s)",
					underline("owner code"),
					bold("3 letters"),
					underline(decoders.ownerDecodeUpdater.GenerateRandomCodes(1)[0]))}}
			}
			found, owner := decoders.ownerDecodeUpdater.Decode(value)
			if !found {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not %s (e.g. %s)",
					underline(value),
					bold("registered"),
					underline(decoders.ownerDecodeUpdater.GenerateRandomCodes(1)[0]))}}

			}
			return true, []input.Info{{Text: owner.Company},
				{Text: owner.City},
				{Text: owner.Country}}
		})
	owner.SetToUpper()

	equipCat := input.NewInput(
		1,
		regexp.MustCompile(`[A-Za-z]`).FindStringIndex,
		func(value string, previousValues []string) (bool, []input.Info) {
			if value == "" {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not %s",
					underline("equipment category id"),
					equipCatIDsAsList(decoders.equipCatDecoder))}}
			}

			found, cat := decoders.equipCatDecoder.Decode(value)
			if !found {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not %s",
					underline("equipment category id"),
					equipCatIDsAsList(decoders.equipCatDecoder))}}
			}
			return true, []input.Info{{Text: cat.Info}}
		})
	equipCat.SetToUpper()

	serialNum := input.NewInput(
		6,
		regexp.MustCompile(`\d{6}`).FindStringIndex,
		func(value string, previousValues []string) (bool, []input.Info) {
			if value == "" {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not %s long",
					underline("serial number"),
					bold("6 numbers"))}}
			}
			return true, nil
		})

	checkDigit := input.NewInput(
		1,
		regexp.MustCompile(`\d`).FindStringIndex,
		func(value string, previousValues []string) (bool, []input.Info) {
			if len(strings.Join(previousValues[0:3], "")) != 10 {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not calculable",
					underline("check digit"))}}
			}

			checkDigit := cont.CalcCheckDigit(previousValues[2], previousValues[1], previousValues[0])

			number, err := strconv.Atoi(value)
			if err != nil {
				return false, appendCheckDigit10Info(checkDigit,
					[]input.Info{{Text: fmt.Sprintf("%s must be a %s (calculated: %s)",
						underline("check digit"),
						bold("number"),
						green(checkDigit))}})
			}

			if number != checkDigit%10 {
				return false, appendCheckDigit10Info(checkDigit, []input.Info{{Text: fmt.Sprintf(
					"calculated %s is %s",
					underline("check digit"),
					green(checkDigit%10))}})
			}

			return true, appendCheckDigit10Info(checkDigit, nil)
		})

	length := input.NewInput(
		1,
		regexp.MustCompile(`[A-Za-z\d]`).FindStringIndex,
		func(value string, previousValues []string) (bool, []input.Info) {
			if value == "" {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not a %s or a %s",
					underline("length code"),
					bold("valid number"),
					bold("valid character"))}}
			}

			found, length := decoders.lengthDecoder.Decode(value)
			if !found {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not %s",
					underline("length code"),
					bold("valid"))}}
			}
			return true, []input.Info{{Text: fmt.Sprintf("length: %s", length.Length)}}
		})
	length.SetToUpper()

	heightWidth := input.NewInput(
		1,
		regexp.MustCompile(`[A-Za-z\d]`).FindStringIndex,
		func(value string, previousValues []string) (b bool, infos []input.Info) {
			if value == "" {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not a %s or a %s",
					underline("height and width code"),
					bold("valid number"),
					bold("valid character"))}}
			}

			found, heightWidth := decoders.heightWidthDecoder.Decode(value)
			if !found {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not %s",
					underline("height and width code"),
					bold("valid"))}}
			}
			return true, []input.Info{
				{Text: fmt.Sprintf("height: %s", heightWidth.Height)},
				{Text: fmt.Sprintf("width:  %s", heightWidth.Width)}}
		})
	heightWidth.SetToUpper()

	typeAndGroup := input.NewInput(
		2,
		regexp.MustCompile(`[A-Za-z\d]{2}`).FindStringIndex,
		func(value string, previousValues []string) (b bool, infos []input.Info) {
			if value == "" {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not a %s or a %s",
					underline("type code"),
					bold("valid number"),
					bold("valid character"))}}
			}

			found, typeAndGroup := decoders.typeDecoder.Decode(value)
			if !found {
				return false, []input.Info{{Text: fmt.Sprintf("%s is not %s",
					underline("type code"),
					bold("valid"))}}
			}
			return true, []input.Info{
				{Text: fmt.Sprintf("type:  %s", typeAndGroup.TypeInfo)},
				{Text: fmt.Sprintf("group: %s", typeAndGroup.GroupInfo)}}

		})
	typeAndGroup.SetToUpper()

	return input.NewValidator([][]input.Input{
		{owner, equipCat, serialNum, checkDigit, length, heightWidth, typeAndGroup},
		{owner, equipCat, serialNum, checkDigit},
		{owner, equipCat},
		{owner},
		{length, heightWidth, typeAndGroup},
	})
}

func appendCheckDigit10Info(checkDigit int, infos []input.Info) []input.Info {
	if checkDigit == 10 {
		if infos == nil {
			infos = make([]input.Info, 0)
		}
		infos = append(infos, input.Info{
			Text: fmt.Sprintf("It is not recommended to use a %s", underline("serial number"))})
		infos = append(infos, input.Info{
			Text: fmt.Sprintf("that generates %s %s (0).", underline("check digit"), yellow(10))})
	}
	return infos
}

func equipCatIDsAsList(equipCatDecoder data.EquipCatDecoder) string {

	b := strings.Builder{}

	iDs := equipCatDecoder.AllCatIDs()
	sort.Strings(iDs)
	for i, element := range iDs {
		b.WriteString(green(element))
		if i < len(iDs)-2 {
			b.WriteString(", ")
		}
		if i == len(iDs)-2 {
			b.WriteString(" or ")
		}
	}
	return b.String()
}
