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
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"github.com/meyermarcel/icm/configs"
	"github.com/meyermarcel/icm/internal/cont"
	"github.com/meyermarcel/icm/internal/data"
	"github.com/meyermarcel/icm/internal/input"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	yellow    = color.New(color.FgYellow).SprintFunc()
	green     = color.New(color.FgGreen).SprintFunc()
	bold      = color.New(color.Bold).SprintFunc()
	underline = color.New(color.Underline).SprintFunc()
)

type errValidate struct {
	message string
}

func newErrValidate(message string) error {
	return &errValidate{message: message}
}

func (e *errValidate) Error() string {
	return e.message
}

const (
	auto                   = "auto"
	containerNumber        = "container-number"
	owner                  = "owner"
	ownerEquipmentCategory = "owner-equipment-category"
	sizeType               = "size-type"
)

const patternModesInfo string = `                    ` + auto + ` = matches automatically a pattern
        ` + containerNumber + ` = matches a container number
                   ` + owner + ` = matches a three letter owner code
` + ownerEquipmentCategory + ` = matches a three letter owner code with equipment category ID
               ` + sizeType + ` = matches length, width+height and type code`

type patternValue struct {
	value    string
	patterns map[string]func(decoders decoders) [][]func() input.Input
}

func (p *patternValue) String() string {
	return p.value
}

func (p *patternValue) Set(value string) error {
	if pattern := p.patterns[value]; pattern == nil {
		return fmt.Errorf("%s is not \n%s", value, patternModesInfo)
	}
	p.value = value
	return nil
}

func (*patternValue) Type() string {
	return "mode"
}

func newPatternValue() *patternValue {
	return &patternValue{
		value: configs.PatternDefVal,
		patterns: map[string]func(decoders decoders) [][]func() input.Input{
			auto:                   newAutoPattern,
			containerNumber:        newContNumPattern,
			owner:                  newOwnerPattern,
			ownerEquipmentCategory: newOwnerEquipCatPattern,
			sizeType:               newSizeTypePattern,
		},
	}
}

func (p *patternValue) newPatterns(value string) func(decoders decoders) [][]func() input.Input {
	return p.patterns[value]
}

var pValue = newPatternValue()

const (
	outputAuto  = "auto"
	outputFancy = "fancy"
	outputCSV   = "csv"
)

type outputValue struct {
	value    string
	printers map[string]newPrinter
}

func newOutputValue() *outputValue {
	return &outputValue{
		value: configs.OutputDefVal,
		printers: map[string]newPrinter{
			outputAuto:  newAutoPrinter,
			outputFancy: newFancyPrinter,
			outputCSV:   newCSVPrinter,
		},
	}
}

type newPrinter func(writer io.Writer, viperCfg *viper.Viper, isSingleLine bool) input.Printer

func (o *outputValue) String() string {
	return o.value
}

func (o *outputValue) Set(value string) error {
	if printer := o.printers[value]; printer == nil {
		return fmt.Errorf("%s is not \n%s", value, outputModesInfo)
	}
	o.value = value
	return nil
}

func (o *outputValue) Type() string {
	return "string"
}

const outputModesInfo string = ` ` + outputAuto + ` = for a single line '` + outputFancy +
	`' and for multiple lines '` + outputCSV + `' output 
  ` + outputCSV + ` = machine readable CSV output
` + outputFancy + ` = human readable fancy output`

func (o *outputValue) newPrinter(value string) newPrinter {
	return o.printers[value]
}

var oValue = newOutputValue()

func newValidateCmd(stdin io.Reader, writer io.Writer, viperCfg *viper.Viper, decoders decoders) *cobra.Command {
	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate intermodal container markings",
		Long: `Validate intermodal container markings.

` + sepHelp,
		Example: `  icm validate ABC
  icm validate ABC --pattern container-number
  icm validate ABC U
  icm validate --sep-owner-equip '' --sep-serial-check '-' ABC U 123456 0
  icm validate ABC U 123456 0 20G1
  icm validate 20G1
  icm generate | icm validate
  icm generate --count 10 | icm validate
  icm generate --count 10 | icm validate --output fancy`,
		Args: cobra.MaximumNArgs(6),
		// https://github.com/spf13/viper/issues/233
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return viperCfg.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			var reader io.Reader
			if len(args) != 0 {
				reader = strings.NewReader(strings.Join(args, " "))
			} else {
				reader = stdin
			}

			bufReader := bufio.NewReader(reader)
			peek, _ := bufReader.Peek(bufReader.Size())
			isSingleLine := isSingleLine(string(peek))

			printer := oValue.newPrinter(viperCfg.GetString(configs.Output))(writer, viperCfg, isSingleLine)

			newPatterns := pValue.newPatterns(viperCfg.GetString(configs.Pattern))(decoders)

			newInputs := input.Match(strings.Split(string(peek), "\n")[0], newPatterns)

			scanner := bufio.NewScanner(bufReader)

			var inputErr error
			var inputs []input.Input

			for scanner.Scan() {
				inputs, inputErr = input.Validate(scanner.Text(), newInputs)
				err := printer.Print(inputs)
				if err != nil {
					return err
				}
			}
			return inputErr
		},
	}
	validateCmd.Flags().VarP(pValue, configs.Pattern, "p",
		fmt.Sprintf("sets pattern matching mode to\n%s\n", patternModesInfo))
	validateCmd.Flags().Var(oValue, configs.Output,
		fmt.Sprintf("sets output to\n%s\n", outputModesInfo))
	validateCmd.Flags().String(configs.SepOE, configs.SepOEDefVal,
		"ABC(*)U1234560   20G1  (*) separates owner code and equipment category id")
	validateCmd.Flags().String(configs.SepES, configs.SepESDefVal,
		"ABCU(*)1234560   20G1  (*) separates equipment category id and serial number")
	validateCmd.Flags().String(configs.SepSC, configs.SepSCDefVal,
		"ABCU123456(*)0   20G1  (*) separates serial number and check digit")
	validateCmd.Flags().String(configs.SepCS, configs.SepCSDefVal,
		"ABCU1234560 (*)  20G1  (*) separates check digit and size")
	validateCmd.Flags().String(configs.SepST, configs.SepSTDefVal,
		"ABCU1234560   20(*)G1  (*) separates size and type")
	validateCmd.Flags().Bool(configs.NoHeader, configs.NoHeaderDefVal,
		"omits header of CSV output")
	return validateCmd
}

func isSingleLine(s string) bool {
	scanner := bufio.NewScanner(strings.NewReader(s))
	counter := 0
	for scanner.Scan() {
		counter++
		if counter > 1 {
			return false
		}
	}
	return true
}

func newAutoPrinter(writer io.Writer, viperCfg *viper.Viper, isSingleLine bool) input.Printer {
	if isSingleLine {
		return newFancyPrinter(writer, viperCfg, isSingleLine)
	}
	return newCSVPrinter(writer, viperCfg, isSingleLine)

}

func newFancyPrinter(writer io.Writer, viperCfg *viper.Viper, isSingleLine bool) input.Printer {
	fancyPrinter := input.NewFancyPrinter(writer)
	fancyPrinter.SetIndent("  ")
	fancyPrinter.SetSeparatorsFunc(func(inputs []input.Input) {
		// only size-type has 3 inputs
		if len(inputs) == 3 {
			fancyPrinter.SetSeparators(
				"",
				viperCfg.GetString(configs.SepST),
			)
		} else {
			fancyPrinter.SetSeparators(
				viperCfg.GetString(configs.SepOE),
				viperCfg.GetString(configs.SepES),
				viperCfg.GetString(configs.SepSC),
				viperCfg.GetString(configs.SepCS),
				"",
				viperCfg.GetString(configs.SepST),
			)
		}
	})
	return fancyPrinter
}

func newCSVPrinter(writer io.Writer, viperCfg *viper.Viper, isSingleLine bool) input.Printer {
	csvWriter := csv.NewWriter(writer)
	csvWriter.Comma = ';'
	return input.NewCSVPrinter(csvWriter, viperCfg.GetBool(configs.NoHeader))
}

func newAutoPattern(decoders decoders) [][]func() input.Input {
	owner := newOwnerInput(decoders.ownerDecodeUpdater)
	equipCat := newEquipCatInput(decoders.equipCatDecoder)
	serialNum := newSerialNumInput()
	checkDigit := newCheckDigitInput()
	length := newLengthInput(decoders.lengthDecoder)
	heightWidth := newHeightWidthInput(decoders.heightWidthDecoder)
	typeAndGroup := newTypeAndGroupInput(decoders.typeDecoder)

	return [][]func() input.Input{
		{owner, equipCat, serialNum, checkDigit, length, heightWidth, typeAndGroup},
		{owner, equipCat, serialNum, checkDigit},
		{owner, equipCat},
		{owner},
		{length, heightWidth, typeAndGroup},
	}
}

func newContNumPattern(decoders decoders) [][]func() input.Input {
	owner := newOwnerInput(decoders.ownerDecodeUpdater)
	equipCat := newEquipCatInput(decoders.equipCatDecoder)
	serialNum := newSerialNumInput()
	checkDigit := newCheckDigitInput()

	return [][]func() input.Input{{owner, equipCat, serialNum, checkDigit}}
}

func newOwnerPattern(decoders decoders) [][]func() input.Input {
	owner := newOwnerInput(decoders.ownerDecodeUpdater)
	return [][]func() input.Input{{owner}}
}

func newOwnerEquipCatPattern(decoders decoders) [][]func() input.Input {
	owner := newOwnerInput(decoders.ownerDecodeUpdater)
	equipCat := newEquipCatInput(decoders.equipCatDecoder)

	return [][]func() input.Input{{owner, equipCat}}
}

func newSizeTypePattern(decoders decoders) [][]func() input.Input {
	length := newLengthInput(decoders.lengthDecoder)
	heightWidth := newHeightWidthInput(decoders.heightWidthDecoder)
	typeAndGroup := newTypeAndGroupInput(decoders.typeDecoder)

	return [][]func() input.Input{{length, heightWidth, typeAndGroup}}
}

func newOwnerInput(ownerDecodeUpdater data.OwnerDecodeUpdater) func() input.Input {
	owner := input.NewInput(
		3,
		regexp.MustCompile(`[A-Za-z]{3}`).FindStringIndex,
		func(value string, previousValues []string) (error, []input.Info, []input.Datum) {
			ownerCodeDatum := input.NewDatum("owner-code")
			ownerCompanyDatum := input.NewDatum("company")
			ownerCityDatum := input.NewDatum("city")
			ownerCountryDatum := input.NewDatum("country")

			if value == "" {
				return newErrValidate(fmt.Sprintf("%s is not %s long (e.g. %s)",
						underline("owner code"),
						bold("3 letters"),
						underline(ownerDecodeUpdater.GetAllOwnerCodes()[0]))),
					nil,
					[]input.Datum{ownerCodeDatum, ownerCompanyDatum, ownerCityDatum, ownerCountryDatum}
			}
			found, owner := ownerDecodeUpdater.Decode(value)
			if !found {
				return newErrValidate(fmt.Sprintf("%s is not %s (e.g. %s)",
						underline(value),
						bold("registered"),
						underline(ownerDecodeUpdater.GetAllOwnerCodes()[0]))),
					nil,
					[]input.Datum{ownerCodeDatum, ownerCompanyDatum, ownerCityDatum, ownerCountryDatum}

			}
			return nil,
				[]input.Info{{Text: owner.Company},
					{Text: owner.City},
					{Text: owner.Country},
				},
				[]input.Datum{
					ownerCodeDatum.WithValue(owner.Code),
					ownerCompanyDatum.WithValue(owner.Company),
					ownerCityDatum.WithValue(owner.City),
					ownerCountryDatum.WithValue(owner.Country),
				}
		})
	owner.SetToUpper()
	return func() input.Input { return owner }
}

func newEquipCatInput(equipCatDecoder data.EquipCatDecoder) func() input.Input {
	equipCat := input.NewInput(
		1,
		regexp.MustCompile(`[A-Za-z]`).FindStringIndex,
		func(value string, previousValues []string) (error, []input.Info, []input.Datum) {
			equipCatIDDatum := input.NewDatum("equipment-category-id").WithValue(value)
			equipCatDatum := input.NewDatum("equipment-category")
			if value == "" {
				return newErrValidate(fmt.Sprintf("%s is not %s",
						underline("equipment category id"),
						equipCatIDsAsList(equipCatDecoder))),
					nil,
					[]input.Datum{equipCatIDDatum, equipCatDatum}
			}

			found, cat := equipCatDecoder.Decode(value)
			if !found {
				return newErrValidate(fmt.Sprintf("%s is not %s",
						underline("equipment category id"),
						equipCatIDsAsList(equipCatDecoder))),
					nil,
					[]input.Datum{equipCatIDDatum, equipCatDatum}
			}
			return nil,
				[]input.Info{{Text: cat.Info}},
				[]input.Datum{equipCatIDDatum, equipCatDatum.WithValue(cat.Info)}
		})
	equipCat.SetToUpper()
	return func() input.Input { return equipCat }
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

func newSerialNumInput() func() input.Input {
	return func() input.Input {
		return input.NewInput(
			6,
			regexp.MustCompile(`\d{6}`).FindStringIndex,
			func(value string, previousValues []string) (error, []input.Info, []input.Datum) {
				if value == "" {
					return newErrValidate(fmt.Sprintf("%s is not %s long",
							underline("serial number"),
							bold("6 numbers"))),
						nil,
						nil
				}
				return nil, nil, nil
			})
	}
}

func newCheckDigitInput() func() input.Input {
	return func() input.Input {
		return input.NewInput(
			1,
			regexp.MustCompile(`\d`).FindStringIndex,
			func(value string, previousValues []string) (error, []input.Info, []input.Datum) {
				checkDigitDatum := input.NewDatum("check-digit").WithValue(value)
				calcCheckDigitDatum := input.NewDatum("calculated-check-digit")
				validCheckDigit := input.NewDatum("valid-check-digit")
				possibleTranspositionError := input.NewDatum("possible-transposition-error")
				if len(strings.Join(previousValues[0:3], "")) != 10 {
					return newErrValidate(fmt.Sprintf("%s is not calculable",
							underline("check digit"))),
						nil,
						[]input.Datum{
							checkDigitDatum,
							calcCheckDigitDatum,
							validCheckDigit.WithValue(fmt.Sprintf("%t", false)),
							possibleTranspositionError,
						}
				}

				checkDigit := cont.CalcCheckDigit(previousValues[2], previousValues[1], previousValues[0])

				infos := appendCheckDigit10Info(checkDigit, nil)

				number, err := strconv.Atoi(value)
				if err != nil {
					return newErrValidate(fmt.Sprintf("%s must be a %s (calculated: %s)",
							underline("check digit"),
							bold("number"),
							green(checkDigit))),
						infos,
						[]input.Datum{
							checkDigitDatum,
							calcCheckDigitDatum.WithValue(strconv.Itoa(checkDigit)),
							validCheckDigit.WithValue(fmt.Sprintf("%t", false)),
							possibleTranspositionError,
						}
				}

				if number != checkDigit%10 {
					return newErrValidate(fmt.Sprintf(
							"calculated %s is %s",
							underline("check digit"),
							green(checkDigit%10))),
						infos,
						[]input.Datum{
							checkDigitDatum,
							calcCheckDigitDatum.WithValue(strconv.Itoa(checkDigit)),
							validCheckDigit.WithValue(fmt.Sprintf("%t", number == checkDigit%10)),
							possibleTranspositionError,
						}
				}

				transposedContNums := cont.CheckTransposition(previousValues[2], previousValues[1], previousValues[0])

				if len(transposedContNums) != 0 {
					infos = append(infos, input.Info{Text: "Possible transposition errors:"})
					builder := strings.Builder{}
					for idx, contNum := range transposedContNums {
						infos = append(infos, input.Info{Text: fmt.Sprintf("  %s", contNum)})
						builder.WriteString(contNum.String())
						if idx < len(transposedContNums)-1 {
							builder.WriteString(", ")
						}
					}
					return nil,
						infos,
						[]input.Datum{
							checkDigitDatum,
							calcCheckDigitDatum.WithValue(strconv.Itoa(checkDigit)),
							validCheckDigit.WithValue(fmt.Sprintf("%t", number == checkDigit%10)),
							possibleTranspositionError.WithValue(builder.String()),
						}
				}

				return nil,
					infos,
					[]input.Datum{
						checkDigitDatum,
						calcCheckDigitDatum.WithValue(strconv.Itoa(checkDigit)),
						validCheckDigit.WithValue(fmt.Sprintf("%t", number == checkDigit%10)),
						possibleTranspositionError,
					}
			})
	}
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

func newLengthInput(lengthDecoder data.LengthDecoder) func() input.Input {

	length := input.NewInput(
		1,
		regexp.MustCompile(`[A-Za-z\d]`).FindStringIndex,
		func(value string, previousValues []string) (error, []input.Info, []input.Datum) {
			lengthDatum := input.NewDatum("length-code").WithValue(value)
			lengthDescDatum := input.NewDatum("length-description")
			if value == "" {
				return newErrValidate(fmt.Sprintf("%s is not a %s or a %s",
						underline("length code"),
						bold("valid number"),
						bold("valid character"))),
					nil,
					[]input.Datum{lengthDatum, lengthDescDatum}
			}

			found, length := lengthDecoder.Decode(value)
			if !found {
				return newErrValidate(fmt.Sprintf("%s is not %s",
						underline("length code"),
						bold("valid"))),
					nil,
					[]input.Datum{lengthDatum, lengthDescDatum}
			}
			return nil,
				[]input.Info{{Text: fmt.Sprintf("length: %s", length.Length)}},
				[]input.Datum{lengthDatum, lengthDescDatum.WithValue(length.Length)}
		})
	length.SetToUpper()
	return func() input.Input { return length }
}

func newHeightWidthInput(heightWidthDecoder data.HeightWidthDecoder) func() input.Input {

	heightWidth := input.NewInput(
		1,
		regexp.MustCompile(`[A-Za-z\d]`).FindStringIndex,
		func(value string, previousValues []string) (error, []input.Info, []input.Datum) {
			heightWidthDatum := input.NewDatum("height-width-code").WithValue(value)
			heightDescDatum := input.NewDatum("height-description")
			widthDescDatum := input.NewDatum("width-description")
			if value == "" {
				return newErrValidate(fmt.Sprintf("%s is not a %s or a %s",
						underline("height and width code"),
						bold("valid number"),
						bold("valid character"))),
					nil,
					[]input.Datum{heightWidthDatum, heightDescDatum, widthDescDatum}
			}

			found, heightWidth := heightWidthDecoder.Decode(value)
			if !found {
				return newErrValidate(fmt.Sprintf("%s is not %s",
						underline("height and width code"),
						bold("valid"))),
					nil,
					[]input.Datum{heightWidthDatum, heightDescDatum, widthDescDatum}
			}
			return nil,
				[]input.Info{
					{Text: fmt.Sprintf("height: %s", heightWidth.Height)},
					{Text: fmt.Sprintf("width:  %s", heightWidth.Width)},
				},
				[]input.Datum{
					heightWidthDatum,
					heightDescDatum.WithValue(heightWidth.Height),
					widthDescDatum.WithValue(heightWidth.Width),
				}
		})
	heightWidth.SetToUpper()
	return func() input.Input { return heightWidth }
}

func newTypeAndGroupInput(typeDecoder data.TypeDecoder) func() input.Input {

	typeAndGroup := input.NewInput(
		2,
		regexp.MustCompile(`[A-Za-z\d]{2}`).FindStringIndex,
		func(value string, previousValues []string) (error, []input.Info, []input.Datum) {
			typeDatum := input.NewDatum("type-code").WithValue(value)
			typeDescDatum := input.NewDatum("type-description")
			groupDescDatum := input.NewDatum("group-description")
			if value == "" {
				return newErrValidate(fmt.Sprintf("%s is not a %s or a %s",
						underline("type code"),
						bold("valid number"),
						bold("valid character"))),
					nil,
					[]input.Datum{typeDatum, typeDescDatum, groupDescDatum}
			}

			found, typeAndGroup := typeDecoder.Decode(value)
			if !found {
				return newErrValidate(fmt.Sprintf("%s is not %s",
						underline("type code"),
						bold("valid"))),
					nil,
					[]input.Datum{typeDatum, typeDescDatum, groupDescDatum}
			}
			return nil,
				[]input.Info{
					{Text: fmt.Sprintf("type:  %s", typeAndGroup.TypeInfo)},
					{Text: fmt.Sprintf("group: %s", typeAndGroup.GroupInfo)},
				},
				[]input.Datum{
					typeDatum,
					typeDescDatum.WithValue(typeAndGroup.TypeInfo),
					groupDescDatum.WithValue(typeAndGroup.GroupInfo),
				}
		})
	typeAndGroup.SetToUpper()
	return func() input.Input { return typeAndGroup }
}
