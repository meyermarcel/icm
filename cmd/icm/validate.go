package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/logrusorgru/aurora/v4"
	"github.com/mattn/go-isatty"
	"github.com/meyermarcel/icm/configs"
	"github.com/meyermarcel/icm/cont"
	"github.com/meyermarcel/icm/data"
	"github.com/meyermarcel/icm/input"
	"github.com/spf13/cobra"
)

var au *aurora.Aurora

func init() {
	au = aurora.New(aurora.WithColors(os.Getenv("NO_COLOR") == "" && isatty.IsTerminal(os.Stdout.Fd())))
}

type validateError struct {
	message string
}

func newValidateError(message string) error {
	return &validateError{message: message}
}

func (e *validateError) Error() string {
	return e.message
}

const (
	auto                   = "auto"
	containerNumber        = "container-number"
	owner                  = "owner"
	ownerEquipmentCategory = "owner-equipment-category"
	sizeType               = "size-type"
)

const patternsInfo string = `                    ` + auto + ` = matches automatically a pattern
        ` + containerNumber + ` = matches a container number
                   ` + owner + ` = matches a three letter owner code
` + ownerEquipmentCategory + ` = matches a three letter owner code with equipment category ID
               ` + sizeType + ` = matches length, width+height and type code`

type patterns = [][]func() input.Input

type patternValue struct {
	config   *configs.Config
	decoders decoders
	value    string
}

func newPatternValue(config *configs.Config, decoders decoders) *patternValue {
	return &patternValue{
		config:   config,
		decoders: decoders,
	}
}

func (p *patternValue) String() string {
	return p.value
}

func (p *patternValue) Set(value string) error {
	switch value {
	case auto, containerNumber, owner, ownerEquipmentCategory, sizeType:
		p.value = value
		return nil
	default:
		return fmt.Errorf("%s is not \n%s", value, patternsInfo)
	}
}

func (*patternValue) Type() string {
	return "string"
}

func (p *patternValue) getPatterns(value string) patterns {
	switch value {

	case containerNumber:
		return newContNumPattern(p.config, p.decoders)
	case owner:
		return newOwnerPattern(p.decoders)
	case ownerEquipmentCategory:
		return newOwnerEquipCatPattern(p.decoders)
	case sizeType:
		return newSizeTypePattern(p.decoders)
	case auto:
		fallthrough
	default:
		return newAutoPattern(p.config, p.decoders)
	}
}

const (
	outputAuto  = "auto"
	outputFancy = "fancy"
	outputCSV   = "csv"
)

type outputValue struct {
	config *configs.Config
	value  string
}

func newOutputValue(config *configs.Config) *outputValue {
	return &outputValue{
		config: config,
	}
}

func (o *outputValue) String() string {
	return o.value
}

func (o *outputValue) Set(value string) error {
	switch value {
	case outputAuto, outputFancy, outputCSV:
		o.value = value
		return nil
	}
	return fmt.Errorf("%s is not \n%s", value, outputsInfo)
}

func (o *outputValue) Type() string {
	return "string"
}

const outputsInfo string = ` ` + outputAuto + ` = for a single line '` + outputFancy +
	`' and for multiple lines '` + outputCSV + `' output 
  ` + outputCSV + ` = machine readable CSV output
` + outputFancy + ` = human readable fancy output`

func (o *outputValue) getPrinter(value string, writer io.Writer, isSingleLine bool) input.Printer {
	switch value {
	case outputFancy:
		return newFancyPrinter(writer, o.config)
	case outputCSV:
		return newCSVPrinter(writer, o.config)
	case outputAuto:
		fallthrough
	default:
		if isSingleLine {
			return newFancyPrinter(writer, o.config)
		}
		return newCSVPrinter(writer, o.config)
	}
}

func newValidateCmd(stdin io.Reader, writer io.Writer, config *configs.Config, decoders decoders) (*cobra.Command, error) {
	pValue := newPatternValue(config, decoders)

	oValue := newOutputValue(config)

	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate intermodal container markings",
		Long: `Validate intermodal container markings with single or multi line input.

For single line input a human-readable output is used.

For multi line input CSV output is used. For example this is useful to scan
data sets for possible transposition errors. It is also possible to generate
CSV data sets of random container numbers.

` + sepHelp,
		Example: `icm validate ABC
# Validate with pattern 'container-number' instead of pattern 'auto'
icm validate ABC --pattern container-number
icm validate ABC U
# Validate and use custom format for output
icm validate --sep-owner-equip '' --sep-serial-check '-' ABC U 123456 0
# Validate a type
icm validate 20G1
# Validate a container number with a type
icm validate ABC U 123456 0 20G1
# Validate a random container number
icm generate | icm validate
icm generate --count 10 | icm validate
icm generate --count 10 | icm validate --output fancy
# Generate CSV data set
icm generate --count 1000000 | icm validate
# Validate a container number with 6 (!) possible transposition errors
icm validate APL U 689473 0`,
		Args:              cobra.MaximumNArgs(6),
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			config.Overwrite(cmd.Flags())

			var reader io.Reader
			if len(args) != 0 {
				reader = strings.NewReader(strings.Join(args, " "))
			} else {
				reader = stdin
			}

			bufReader := bufio.NewReader(reader)
			peek, _ := bufReader.Peek(bufReader.Size())
			singleLine := isSingleLine(string(peek))

			printer := oValue.getPrinter(config.Output(), writer, singleLine)

			patterns := pValue.getPatterns(config.Pattern())

			newInputs := input.Match(strings.Split(string(peek), "\n")[0], patterns)

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

	validateCmd.Flags().SortFlags = false

	validateCmd.Flags().VarP(pValue, configs.FlagNames.Pattern, "p",
		fmt.Sprintf("sets pattern matching to %s, %s, %s, %s or %s\n%s\n",
			auto, containerNumber, owner, ownerEquipmentCategory, sizeType,
			patternsInfo))
	err := validateCmd.RegisterFlagCompletionFunc(configs.FlagNames.Pattern, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{auto, containerNumber, owner, ownerEquipmentCategory, sizeType}, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, err
	}
	validateCmd.Flags().Var(oValue, configs.FlagNames.Output,
		fmt.Sprintf("sets output to %s, %s or %s\n%s\n",
			outputAuto, outputFancy, outputCSV,
			outputsInfo))
	err = validateCmd.RegisterFlagCompletionFunc(configs.FlagNames.Output, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{outputAuto, outputFancy, outputCSV}, cobra.ShellCompDirectiveNoFileComp
	})
	if err != nil {
		return nil, err
	}
	validateCmd.Flags().Bool(configs.FlagNames.NoHeader, configs.DefaultValues.NoHeader,
		"omits header of CSV output")
	validateCmd.Flags().String(configs.FlagNames.SepOE, configs.DefaultValues.SepOE,
		"ABC(x)U1234560   20G1  (x) separates owner code and equipment category id")
	validateCmd.Flags().String(configs.FlagNames.SepES, configs.DefaultValues.SepES,
		"ABCU(x)1234560   20G1  (x) separates equipment category id and serial number")
	validateCmd.Flags().String(configs.FlagNames.SepSC, configs.DefaultValues.SepSC,
		"ABCU123456(x)0   20G1  (x) separates serial number and check digit")
	validateCmd.Flags().String(configs.FlagNames.SepCS, configs.DefaultValues.SepCS,
		"ABCU1234560 (x)  20G1  (x) separates check digit and size")
	validateCmd.Flags().String(configs.FlagNames.SepST, configs.DefaultValues.SepST,
		"ABCU1234560   20(x)G1  (x) separates size and type")
	return validateCmd, nil
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

func newFancyPrinter(writer io.Writer, config *configs.Config) input.Printer {
	fancyPrinter := input.NewFancyPrinter(writer)
	fancyPrinter.SetIndent("  ")
	fancyPrinter.SetSeparatorsFunc(func(inputs []input.Input) {
		// only size-type has 3 inputs
		if len(inputs) == 3 {
			fancyPrinter.SetSeparators(
				"",
				config.SepST(),
			)
		} else {
			fancyPrinter.SetSeparators(
				config.SepOE(),
				config.SepES(),
				config.SepSC(),
				config.SepCS(),
				"",
				config.SepST(),
			)
		}
	})
	return fancyPrinter
}

func newCSVPrinter(writer io.Writer, config *configs.Config) input.Printer {
	csvWriter := csv.NewWriter(writer)
	csvWriter.Comma = ';'
	return input.NewCSVPrinter(csvWriter, config.NoHeader())
}

func newAutoPattern(config *configs.Config, decoders decoders) patterns {
	owner := newOwnerInput(decoders.ownerDecodeUpdater)
	equipCat := newEquipCatInput(decoders.equipCatDecoder)
	serialNum := newSerialNumInput()
	checkDigit := newCheckDigitInput(config)
	length := newLengthInput(decoders.lengthDecoder)
	heightWidth := newHeightWidthInput(decoders.heightWidthDecoder)
	typeAndGroup := newTypeAndGroupInput(decoders.typeDecoder)

	return patterns{
		{owner, equipCat, serialNum, checkDigit, length, heightWidth, typeAndGroup},
		{owner, equipCat, serialNum, checkDigit},
		{owner, equipCat},
		{owner},
		{length, heightWidth, typeAndGroup},
	}
}

func newContNumPattern(config *configs.Config, decoders decoders) patterns {
	owner := newOwnerInput(decoders.ownerDecodeUpdater)
	equipCat := newEquipCatInput(decoders.equipCatDecoder)
	serialNum := newSerialNumInput()
	checkDigit := newCheckDigitInput(config)

	return patterns{{owner, equipCat, serialNum, checkDigit}}
}

func newOwnerPattern(decoders decoders) patterns {
	owner := newOwnerInput(decoders.ownerDecodeUpdater)
	return patterns{{owner}}
}

func newOwnerEquipCatPattern(decoders decoders) patterns {
	owner := newOwnerInput(decoders.ownerDecodeUpdater)
	equipCat := newEquipCatInput(decoders.equipCatDecoder)

	return patterns{{owner, equipCat}}
}

func newSizeTypePattern(decoders decoders) patterns {
	length := newLengthInput(decoders.lengthDecoder)
	heightWidth := newHeightWidthInput(decoders.heightWidthDecoder)
	typeAndGroup := newTypeAndGroupInput(decoders.typeDecoder)

	return patterns{{length, heightWidth, typeAndGroup}}
}

func newOwnerInput(ownerDecoder data.OwnerDecoder) func() input.Input {
	owner := input.NewInput(
		3,
		regexp.MustCompile(`[A-Za-z]{3}`).FindStringIndex,
		func(value string, previousValues []string) (error, []input.Info, []input.Datum) {
			ownerCodeDatum := input.NewDatum("owner-code")
			ownerCompanyDatum := input.NewDatum("company")
			ownerCityDatum := input.NewDatum("city")
			ownerCountryDatum := input.NewDatum("country")

			if value == "" {
				return newValidateError(fmt.Sprintf("%s is not %s long (e.g. %s)",
						au.Underline("owner code"),
						au.Bold("3 letters"),
						au.Underline(ownerDecoder.GetAllOwnerCodes()[0]))),
					nil,
					[]input.Datum{ownerCodeDatum, ownerCompanyDatum, ownerCityDatum, ownerCountryDatum}
			}
			found, owner := ownerDecoder.Decode(value)
			if !found {
				return newValidateError(fmt.Sprintf("%s is not %s (e.g. %s)",
						au.Underline(value),
						au.Bold("registered"),
						au.Underline(ownerDecoder.GetAllOwnerCodes()[0]))),
					nil,
					[]input.Datum{ownerCodeDatum, ownerCompanyDatum, ownerCityDatum, ownerCountryDatum}
			}
			return nil,
				[]input.Info{
					{Text: owner.Company},
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
				return newValidateError(fmt.Sprintf("%s is not %s",
						au.Underline("equipment category id"),
						equipCatIDsAsList(equipCatDecoder))),
					nil,
					[]input.Datum{equipCatIDDatum, equipCatDatum}
			}

			found, cat := equipCatDecoder.Decode(value)
			if !found {
				return newValidateError(fmt.Sprintf("%s is not %s",
						au.Underline("equipment category id"),
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
		b.WriteString(fmt.Sprint(au.Green(element)))
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
				serialNumData := input.NewDatum("serial-number")
				if value == "" {
					return newValidateError(fmt.Sprintf("%s is not %s long",
							au.Underline("serial number"),
							au.Bold("6 numbers"))),
						nil,
						[]input.Datum{serialNumData}
				}
				return nil, nil, []input.Datum{serialNumData.WithValue(value)}
			})
	}
}

func newCheckDigitInput(config *configs.Config) func() input.Input {
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
					return newValidateError(fmt.Sprintf("%s is not calculable",
							au.Underline("check digit"))),
						nil,
						[]input.Datum{
							checkDigitDatum,
							calcCheckDigitDatum,
							validCheckDigit.WithValue(fmt.Sprintf("%t", false)),
							possibleTranspositionError,
						}
				}

				serialNum, _ := strconv.Atoi(previousValues[0])
				equipCatID, _ := utf8.DecodeRuneInString(previousValues[1])
				checkDigit := cont.CalcCheckDigit(previousValues[2], equipCatID, serialNum)

				infos := appendCheckDigit10Info(checkDigit, nil)

				number, err := strconv.Atoi(value)
				if err != nil {
					return newValidateError(fmt.Sprintf("%s must be a %s (calculated: %s)",
							au.Underline("check digit"),
							au.Bold("number"),
							au.Green(strconv.Itoa(checkDigit)))),
						infos,
						[]input.Datum{
							checkDigitDatum,
							calcCheckDigitDatum.WithValue(strconv.Itoa(checkDigit)),
							validCheckDigit.WithValue(fmt.Sprintf("%t", false)),
							possibleTranspositionError,
						}
				}

				if number != checkDigit%10 {
					return newValidateError(fmt.Sprintf(
							"calculated %s is %s",
							au.Underline("check digit"),
							au.Green(strconv.Itoa(checkDigit%10)))),
						infos,
						[]input.Datum{
							checkDigitDatum,
							calcCheckDigitDatum.WithValue(strconv.Itoa(checkDigit)),
							validCheckDigit.WithValue(fmt.Sprintf("%t", number == checkDigit%10)),
							possibleTranspositionError,
						}
				}

				transposedContNums := cont.CheckTransposition(previousValues[2], equipCatID, serialNum, checkDigit)

				if len(transposedContNums) != 0 {
					infos = append(infos, input.Info{Text: "Possible transposition errors:"})
					builder := strings.Builder{}

					for idx, tcn := range transposedContNums {

						serialNumber := fmt.Sprintf("%06d", tcn.SerialNumber)
						serialNumberFmt := ""
						for i := 0; i < len(serialNumber); i++ {
							if i == tcn.Pos {
								serialNumberFmt += fmt.Sprintf("%c", au.Red(serialNumber[i]))
								// last serial number digit
								if i == len(serialNumber)-1 {
									continue
								}
								serialNumberFmt += fmt.Sprintf("%c", au.Red(serialNumber[i+1]))
								i++
							} else {
								serialNumberFmt += fmt.Sprintf("%c", serialNumber[i])
							}
						}

						var digitFmt string
						if tcn.Pos == 5 {
							digitFmt = fmt.Sprintf("%d", au.Red(tcn.CheckDigit))
						} else {
							digitFmt = fmt.Sprintf("%d", tcn.CheckDigit)
						}

						contNumFmt := fmt.Sprintf("%s%s%s%s%s%s%s",
							tcn.OwnerCode, config.SepOE(),
							string(tcn.EquipCatID), config.SepES(),
							serialNumberFmt, config.SepSC(),
							digitFmt)
						infos = append(infos, input.Info{Text: fmt.Sprintf("  %s", contNumFmt)})
						builder.WriteString(contNumFmt)
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
			Text: fmt.Sprintf("It is not recommended to use a %s", au.Underline("serial number")),
		})
		infos = append(infos, input.Info{
			Text: fmt.Sprintf("that generates %s %s (0).", au.Underline("check digit"), au.Yellow("10")),
		})
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
				return newValidateError(fmt.Sprintf("%s is not a %s or a %s",
						au.Underline("length code"),
						au.Bold("valid number"),
						au.Bold("valid character"))),
					nil,
					[]input.Datum{lengthDatum, lengthDescDatum}
			}

			found, length := lengthDecoder.Decode(value)
			if !found {
				return newValidateError(fmt.Sprintf("%s is not %s",
						au.Underline("length code"),
						au.Bold("valid"))),
					nil,
					[]input.Datum{lengthDatum, lengthDescDatum}
			}
			return nil,
				[]input.Info{{Text: fmt.Sprintf("length: %s", length)}},
				[]input.Datum{lengthDatum, lengthDescDatum.WithValue(string(length))}
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
				return newValidateError(fmt.Sprintf("%s is not a %s or a %s",
						au.Underline("height and width code"),
						au.Bold("valid number"),
						au.Bold("valid character"))),
					nil,
					[]input.Datum{heightWidthDatum, heightDescDatum, widthDescDatum}
			}

			found, height, width := heightWidthDecoder.Decode(value)
			if !found {
				return newValidateError(fmt.Sprintf("%s is not %s",
						au.Underline("height and width code"),
						au.Bold("valid"))),
					nil,
					[]input.Datum{heightWidthDatum, heightDescDatum, widthDescDatum}
			}
			return nil,
				[]input.Info{
					{Text: fmt.Sprintf("height: %s", height)},
					{Text: fmt.Sprintf("width:  %s", width)},
				},
				[]input.Datum{
					heightWidthDatum,
					heightDescDatum.WithValue(string(height)),
					widthDescDatum.WithValue(string(width)),
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
				return newValidateError(fmt.Sprintf("%s is not a %s or a %s",
						au.Underline("type code"),
						au.Bold("valid number"),
						au.Bold("valid character"))),
					nil,
					[]input.Datum{typeDatum, typeDescDatum, groupDescDatum}
			}

			found, typeInfo, groupInfo := typeDecoder.Decode(value)
			if !found {
				return newValidateError(fmt.Sprintf("%s is not %s",
						au.Underline("type code"),
						au.Bold("valid"))),
					nil,
					[]input.Datum{typeDatum, typeDescDatum, groupDescDatum}
			}
			return nil,
				[]input.Info{
					{Text: fmt.Sprintf("type:  %s", typeInfo)},
					{Text: fmt.Sprintf("group: %s", groupInfo)},
				},
				[]input.Datum{
					typeDatum,
					typeDescDatum.WithValue(string(typeInfo)),
					groupDescDatum.WithValue(string(groupInfo)),
				}
		})
	typeAndGroup.SetToUpper()
	return func() input.Input { return typeAndGroup }
}
