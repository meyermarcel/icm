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

package ui

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/meyermarcel/icm/cont"
	"github.com/meyermarcel/icm/data"
)

const caseInsensitive = `(?i)`

//                        %s   %s Owner codes, for example: LAB|LBI|LXA|MTB|MTV|BHF|...
const ownerCodeRegex = `[^%s]*(%s)`

func ownerCodeRegexResolved() string {
	return fmt.Sprintf(ownerCodeRegex, strings.Join(data.GetOwnerCodes(), "|"), strings.Join(data.GetOwnerCodes(), "|"))
}

//                                       %s    %s Equipment category IDs, for example: UJZ
const equipCatSerialCheckDigitRegex = `[^%s]*([%s])?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?`

//                      %s Length codes, for example: 1234ABC...
const onlySizeType = `([%s])`

const optSizeType = `[^%s]*([%s])?`

//                       %s    %s Height and width codes, for example: 0245CDE
//                                     %s,   %s Type codes, for example: A1|B1|B2|...
const sizeTypeRegex = `[^%s]*([%s])?[^(%s)]*(%s)?`

func sizeTypeRegexResolved() string {
	return fmt.Sprintf(sizeTypeRegex, strings.Join(data.GetHeightAndWidthCodes(), ""),
		strings.Join(data.GetHeightAndWidthCodes(), ""), strings.Join(data.GetTypeCodes(), "|"), strings.Join(data.GetTypeCodes(), "|"))
}

func ownerMatcher() *regexp.Regexp {
	return regexp.MustCompile(caseInsensitive + ownerCodeRegexResolved())
}

func contNumMatcher() *regexp.Regexp {
	return regexp.MustCompile(
		caseInsensitive +
			ownerCodeRegexResolved() +
			fmt.Sprintf(equipCatSerialCheckDigitRegex, strings.Join(data.GetEquipCatIDs(), ""), strings.Join(data.GetEquipCatIDs(), "")) +
			fmt.Sprintf(optSizeType, strings.Join(data.GetLengthCodes(), ""), strings.Join(data.GetLengthCodes(), "")) +
			sizeTypeRegexResolved())
}

func sizeTypeMatcher() *regexp.Regexp {
	return regexp.MustCompile(caseInsensitive +
		fmt.Sprintf(onlySizeType, strings.Join(data.GetLengthCodes(), "")) +
		sizeTypeRegexResolved())
}

type regexIn struct {
	matches     []string
	input       string
	matchRanges []int
}

func (pi regexIn) getMatches(start, end int) (value string) {

	if len(pi.matches) == 0 {
		return
	}
	for _, element := range pi.matches[start:end] {
		value += element
	}
	return strings.ToUpper(value)
}

func (pi regexIn) getMatch(pos int) string {

	return pi.getMatches(pos, pos+1)
}

type input struct {
	value    string
	validLen int
}

func (in input) isValidFmt() bool {
	return len(in.value) == in.validLen
}

func newIn(value string, validLen int) input {
	return input{value: value, validLen: validLen}
}

type ownerCodeIn struct {
	input
	Owner cont.Owner
}

func (oi *ownerCodeIn) resolve(fn func(code cont.OwnerCode) cont.Owner) *ownerCodeIn {

	if oi.isValidFmt() {
		oi.Owner = fn(cont.NewOwnerCode(oi.value))
	}
	return oi
}

type equipCatIDIn struct {
	input
	EquipCat cont.EquipCat
}

func (eci *equipCatIDIn) resolve(fn func(code string) cont.EquipCat) *equipCatIDIn {

	if eci.isValidFmt() {
		eci.EquipCat = fn(eci.value)
	}
	return eci
}

type checkDigitIn struct {
	input
	IsValidCheckDigit bool
	CalcCheckDigit    int
}

type lengthIn struct {
	input
	Length cont.Length
}

func (li *lengthIn) resolve(fn func(code string) cont.Length) *lengthIn {

	if li.isValidFmt() {
		li.Length = fn(li.value)
	}
	return li
}

type heightWidthIn struct {
	input
	HeightWidth cont.HeightAndWidth
}

func (hwi *heightWidthIn) resolve(fn func(code string) cont.HeightAndWidth) *heightWidthIn {

	if hwi.isValidFmt() {
		hwi.HeightWidth = fn(hwi.value)
	}
	return hwi
}

type typeAndGroupIn struct {
	input
	TypeAndGroup data.TypeAndGroup
}

func (tgi *typeAndGroupIn) resolve(fn func(code string) data.TypeAndGroup) *typeAndGroupIn {

	if tgi.isValidFmt() {
		tgi.TypeAndGroup = fn(tgi.value)
	}
	return tgi
}

func (cdi *checkDigitIn) calcCheckDigit(ocIn ownerCodeIn, eciIn equipCatIDIn, snIn input) {

	cdi.CalcCheckDigit = cont.CalcCheckDigit(cont.NewOwnerCode(ocIn.input.value), cont.NewEquipCatIDFrom(eciIn.value), cont.NewSerialNumFrom(snIn.value))
	if cdi.isValidFmt() {
		cd, _ := strconv.Atoi(cdi.value)
		cdi.IsValidCheckDigit = cd == cdi.CalcCheckDigit%10
	}
}

type ownerCodeOptEquipCatIDIn struct {
	regexIn     regexIn
	ownerCodeIn ownerCodeIn
}

func (oci ownerCodeOptEquipCatIDIn) isValid() bool {
	return oci.ownerCodeIn.isValidFmt()
}

type contNumOptSizeTypeIn struct {
	regexIn        regexIn
	ownerCodeIn    ownerCodeIn
	equipCatIDIn   equipCatIDIn
	serialNumIn    input
	checkDigitIn   checkDigitIn
	sizeTypeExists bool
	lengthIn       lengthIn
	heightWidthIn  heightWidthIn
	typeAndGroupIn typeAndGroupIn
}

func (cn contNumOptSizeTypeIn) isCheckDigitCalculable() bool {
	return cn.ownerCodeIn.isValidFmt() && cn.equipCatIDIn.isValidFmt() && cn.serialNumIn.isValidFmt()
}

func (cn contNumOptSizeTypeIn) isValid() bool {
	return cn.checkDigitIn.IsValidCheckDigit &&
		((!cn.lengthIn.isValidFmt() && !cn.heightWidthIn.isValidFmt() && !cn.typeAndGroupIn.isValidFmt()) ||
			(cn.lengthIn.isValidFmt() && cn.heightWidthIn.isValidFmt() && cn.typeAndGroupIn.isValidFmt()))
}

type sizeTypeIn struct {
	RegexIn        regexIn
	lengthIn       lengthIn
	heightWidthIn  heightWidthIn
	typeAndGroupIn typeAndGroupIn
}

func (sti sizeTypeIn) isValid() bool {
	return sti.lengthIn.isValidFmt() && sti.heightWidthIn.isValidFmt() && sti.typeAndGroupIn.isValidFmt()
}

// ParseAndPrintOwnerCodeOptEquipCat parses and prints owner code with an optional equipment category ID.
func ParseAndPrintOwnerCodeOptEquipCat(in string) bool {
	parseOwnerCodeOptEquipCat := parseOwnerCodeOptEquipCat(in)
	printOwnerCode(parseOwnerCodeOptEquipCat)
	return parseOwnerCodeOptEquipCat.isValid()
}

func parseOwnerCodeOptEquipCat(in string) ownerCodeOptEquipCatIDIn {
	ownerOptCat := ownerCodeOptEquipCatIDIn{}
	parse := parse(in, ownerMatcher())
	ownerOptCat.regexIn = parse
	ownerOptCat.ownerCodeIn = ownerCodeIn{input: newIn(parse.getMatch(0), 3)}
	ownerOptCat.ownerCodeIn.resolve(data.GetOwner)
	return ownerOptCat
}

// ParseAndPrintContNum parses and prints container number and uses separators for output.
// Given separators are not used for to parse input.
func ParseAndPrintContNum(in string, seps Separators) bool {
	parsedContNum := parseContNum(in)
	printContNum(parsedContNum, seps)
	return parsedContNum.isValid()
}

func parseContNum(in string) contNumOptSizeTypeIn {
	cni := contNumOptSizeTypeIn{}
	parse := parse(in, contNumMatcher())
	cni.regexIn = parse
	cni.ownerCodeIn = ownerCodeIn{input: newIn(parse.getMatch(0), 3)}
	cni.ownerCodeIn.resolve(data.GetOwner)
	cni.equipCatIDIn = equipCatIDIn{input: newIn(parse.getMatch(1), 1)}
	cni.equipCatIDIn.resolve(data.GetEquipCat)
	cni.serialNumIn = newIn(parse.getMatches(2, 8), 6)

	cni.checkDigitIn = checkDigitIn{input: newIn(parse.getMatch(8), 1)}
	if cni.isCheckDigitCalculable() {
		cni.checkDigitIn.calcCheckDigit(cni.ownerCodeIn, cni.equipCatIDIn, cni.serialNumIn)
	}

	cni.lengthIn = lengthIn{input: newIn(parse.getMatch(9), 1)}
	cni.lengthIn.resolve(data.GetLength)
	cni.heightWidthIn = heightWidthIn{input: newIn(parse.getMatch(10), 1)}
	cni.heightWidthIn.resolve(data.GetHeightAndWidth)
	cni.typeAndGroupIn = typeAndGroupIn{input: newIn(parse.getMatch(11), 2)}
	cni.typeAndGroupIn.resolve(data.GetTypeAndGroup)

	return cni
}

// ParseAndPrintSizeType parses and prints size and type and uses separator between size and type for output.
// Separator between size and type is not used for to parse input.
func ParseAndPrintSizeType(in string, sepSizeType string) bool {
	parsedSizeType := parseSizeType(in)
	printSizeType(parsedSizeType, sepSizeType)
	return parsedSizeType.isValid()
}

func parseSizeType(in string) sizeTypeIn {
	sizeType := sizeTypeIn{}
	parse := parse(in, sizeTypeMatcher())
	sizeType.RegexIn = parse
	sizeType.lengthIn = lengthIn{input: newIn(parse.getMatch(0), 1)}
	sizeType.lengthIn.resolve(data.GetLength)
	sizeType.heightWidthIn = heightWidthIn{input: newIn(parse.getMatch(1), 1)}
	sizeType.heightWidthIn.resolve(data.GetHeightAndWidth)
	sizeType.typeAndGroupIn = typeAndGroupIn{input: newIn(parse.getMatch(2), 2)}
	sizeType.typeAndGroupIn.resolve(data.GetTypeAndGroup)

	return sizeType
}

func parse(in string, matcher *regexp.Regexp) regexIn {

	regexIn := regexIn{input: in}

	subMatch := matcher.FindAllStringSubmatch(in, -1)

	if len(subMatch) == 0 {
		return regexIn
	}

	regexIn.matches = subMatch[0][1:]

	regexIn.matchRanges = matcher.FindAllStringSubmatchIndex(in, -1)[0][2:]

	return regexIn
}
