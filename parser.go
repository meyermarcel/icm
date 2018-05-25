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
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const caseInsensitive = `(?i)`

//                        %s   %s Owner codes, for example: LAB|LBI|LXA|MTB|MTV|BHF|...
const ownerCodeRegex = `[^%s]*(%s)`

func ownerCodeRegexResolved() string {
	return fmt.Sprintf(ownerCodeRegex, getRegexPartOwners(), getRegexPartOwners())
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
	return fmt.Sprintf(sizeTypeRegex, getRegexPartHeightAndWidths(),
		getRegexPartHeightAndWidths(), getRegexPartTypes(), getRegexPartTypes())
}

func ownerMatcher() *regexp.Regexp {
	return regexp.MustCompile(caseInsensitive + ownerCodeRegexResolved())
}

func contNumMatcher() *regexp.Regexp {
	return regexp.MustCompile(
		caseInsensitive +
			ownerCodeRegexResolved() +
			fmt.Sprintf(equipCatSerialCheckDigitRegex, getRegexPartEquipCatIDs(), getRegexPartEquipCatIDs()) +
			fmt.Sprintf(optSizeType, getRegexPartLengths(), getRegexPartLengths()) +
			sizeTypeRegexResolved())
}

func sizeTypeMatcher() *regexp.Regexp {
	return regexp.MustCompile(caseInsensitive +
		fmt.Sprintf(onlySizeType, getRegexPartLengths()) +
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

func (pi regexIn) Input() string {
	return pi.input
}

func (pi regexIn) MatchRanges() []int {
	return pi.matchRanges
}

type input struct {
	value    string
	validLen int
}

func (in input) isValidFmt() bool {
	return len(in.value) == in.validLen
}

func (in input) Value() string {
	return in.value
}

func (in input) ValidLen() int {
	return in.validLen
}

func newIn(value string, validLen int) input {
	return input{value: value, validLen: validLen}
}

type ownerCodeIn struct {
	input
	Owner owner
}

func (oi *ownerCodeIn) resolve(fn func(code ownerCode) owner) *ownerCodeIn {

	if oi.isValidFmt() {
		oi.Owner = fn(newOwnerCode(oi.Value()))
	}
	return oi
}

type equipCatIDIn struct {
	input
	EquipCatID equipCatID
}

func (eci *equipCatIDIn) resolve(fn func(code string) equipCatID) *equipCatIDIn {

	if eci.isValidFmt() {
		eci.EquipCatID = fn(eci.Value())
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
	Length length
}

func (li *lengthIn) resolve(fn func(code string) length) *lengthIn {

	if li.isValidFmt() {
		li.Length = fn(li.Value())
	}
	return li
}

type heightWidthIn struct {
	input
	HeightWidth heightAndWidth
}

func (hwi *heightWidthIn) resolve(fn func(code string) heightAndWidth) *heightWidthIn {

	if hwi.isValidFmt() {
		hwi.HeightWidth = fn(hwi.Value())
	}
	return hwi
}

type typeAndGroupIn struct {
	input
	TypeAndGroup mappedTypeAndGroup
}

func (tgi *typeAndGroupIn) resolve(fn func(code string) mappedTypeAndGroup) *typeAndGroupIn {

	if tgi.isValidFmt() {
		tgi.TypeAndGroup = fn(tgi.Value())
	}
	return tgi
}

func (cdi *checkDigitIn) calcCheckDigit(ocIn ownerCodeIn, eciIn equipCatIDIn, snIn input) {

	cdi.CalcCheckDigit = calcCheckDigit(newOwnerCode(ocIn.input.Value()), newEquipCatIDFrom(eciIn.Value()), serialNumFrom(snIn.Value()))
	if cdi.isValidFmt() {
		cd, _ := strconv.Atoi(cdi.Value())
		cdi.IsValidCheckDigit = cd == cdi.CalcCheckDigit
	}
}

type ownerCodeOptEquipCatIDIn struct {
	RegexIn     regexIn
	ownerCodeIn ownerCodeIn
}

type contNumOptSizeTypeIn struct {
	RegexIn        regexIn
	OwnerCodeIn    ownerCodeIn
	EquipCatIDIn   equipCatIDIn
	SerialNumIn    input
	CheckDigitIn   checkDigitIn
	sizeTypeExists bool
	LengthIn       lengthIn
	HeightWidthIn  heightWidthIn
	TypeAndGroupIn typeAndGroupIn
}

type sizeTypeIn struct {
	RegexIn        regexIn
	lengthIn       lengthIn
	heightWidthIn  heightWidthIn
	typeAndGroupIn typeAndGroupIn
}

func (cn contNumOptSizeTypeIn) isCheckDigitCalculable() bool {
	return cn.OwnerCodeIn.isValidFmt() && cn.EquipCatIDIn.isValidFmt() && cn.SerialNumIn.isValidFmt()
}

func (cn contNumOptSizeTypeIn) isValid() bool {
	return cn.CheckDigitIn.IsValidCheckDigit &&
		((!cn.LengthIn.isValidFmt() && !cn.HeightWidthIn.isValidFmt() && !cn.TypeAndGroupIn.isValidFmt()) ||
			(cn.LengthIn.isValidFmt() && cn.HeightWidthIn.isValidFmt() && cn.TypeAndGroupIn.isValidFmt()))
}

func parseOwnerCodeOptEquipCat(in string) ownerCodeOptEquipCatIDIn {
	ownerOptCat := ownerCodeOptEquipCatIDIn{}
	parse := parse(in, ownerMatcher())
	ownerOptCat.RegexIn = parse
	ownerOptCat.ownerCodeIn = ownerCodeIn{input: newIn(parse.getMatch(0), 3)}
	return ownerOptCat
}

func parseContNum(in string) contNumOptSizeTypeIn {
	cni := contNumOptSizeTypeIn{}
	parse := parse(in, contNumMatcher())
	cni.RegexIn = parse
	cni.OwnerCodeIn = ownerCodeIn{input: newIn(parse.getMatch(0), 3)}
	cni.EquipCatIDIn = equipCatIDIn{input: newIn(parse.getMatch(1), 1)}
	cni.SerialNumIn = newIn(parse.getMatches(2, 8), 6)

	cni.CheckDigitIn = checkDigitIn{input: newIn(parse.getMatch(8), 1)}
	if cni.isCheckDigitCalculable() {
		cni.CheckDigitIn.calcCheckDigit(cni.OwnerCodeIn, cni.EquipCatIDIn, cni.SerialNumIn)
	}

	cni.LengthIn = lengthIn{input: newIn(parse.getMatch(9), 1)}
	cni.HeightWidthIn = heightWidthIn{input: newIn(parse.getMatch(10), 1)}
	cni.TypeAndGroupIn = typeAndGroupIn{input: newIn(parse.getMatch(11), 2)}

	return cni
}

func parseSizeType(in string) sizeTypeIn {
	sizeType := sizeTypeIn{}
	parse := parse(in, sizeTypeMatcher())
	sizeType.RegexIn = parse
	sizeType.lengthIn = lengthIn{input: newIn(parse.getMatch(0), 1)}
	sizeType.heightWidthIn = heightWidthIn{input: newIn(parse.getMatch(1), 1)}
	sizeType.typeAndGroupIn = typeAndGroupIn{input: newIn(parse.getMatch(2), 2)}
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
