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
	"regexp"
	"strconv"
	"strings"
)

const caseInsensitive = `(?i)`

const ownerCodeRegex = `([A-Z])[^A-Z]*([A-Z])?[^A-Z]*([A-Z])`

var ownerCodeOptEquipCatIDMatcher = regexp.MustCompile(caseInsensitive + ownerCodeRegex)

const sizeTypeRegex = `[^0245689CDEFLMNP]*([0245689CDEFLMNP])?[^ABGHKNPRSUV]*(A0|B0|B1|B2|B3|B4|B5|B6|B7|B8|B9|G1|G2|G3|G4|G5|G6|G7|G8|G9|H0|H1|H2|H3|H4|H5|H6|H7|H8|H9|K0|K1|K2|K3|K4|K5|K6|K7|K8|K9|N0|N1|N2|N3|N4|N5|N6|N7|N8|N9|N9|P0|P1|P2|P3|P4|P5|P6|P7|P8|P9|R0|R1|R2|R3|R4|R5|R6|R7|R8|R9|S0|S1|S2|S3|S4|S5|S6|S7|S8|S9|U0|U1|U2|U3|U4|U5|U6|U7|U8|U9|V0|V1|V2|V3|V4|V5|V6|V7|V8|V9)?`
const onlySizeType = `([1234ABCDEFGHKLMNP])`

var sizeTypeMatcher = regexp.MustCompile(caseInsensitive + onlySizeType + sizeTypeRegex)

const optSizeType = `[^1234ABCDEFGHKLMNP]*([1234ABCDEFGHKLMNP])?`

const contNumRegex = caseInsensitive + ownerCodeRegex +
	`[^JUZ]*([JUZ])?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?` +
	optSizeType + sizeTypeRegex

var contNumMatcher = regexp.MustCompile(contNumRegex)

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

type ownerCodeInResolvable struct {
	input
	OwnerFound bool
	FoundOwner owner
}

func (oi *ownerCodeInResolvable) resolve(fn func(code ownerCode) (owner, bool)) *ownerCodeInResolvable {

	if oi.isValidFmt() {
		foundOwner, found := fn(newOwnerCode(oi.Value()))
		oi.OwnerFound = found
		oi.FoundOwner = foundOwner
	}
	return oi
}

type checkDigitIn struct {
	input
	IsValidCheckDigit bool
	CalcCheckDigit    int
}

type lengthIn struct {
	input
	Found        bool
	MappedLength mappedLength
}

func (li *lengthIn) resolve(fn func(code string) (mappedLength, bool)) *lengthIn {

	if li.isValidFmt() {
		length, found := fn(li.Value())
		li.MappedLength = length
		li.Found = found
	}
	return li
}

type heightWidthIn struct {
	input
	Found             bool
	MappedHeightWidth mappedHeightAndWidth
}

func (hwi *heightWidthIn) resolve(fn func(code string) (mappedHeightAndWidth, bool)) *heightWidthIn {

	if hwi.isValidFmt() {
		heightWidth, found := fn(hwi.Value())
		hwi.Found = found
		hwi.MappedHeightWidth = heightWidth
	}
	return hwi
}

type typeAndGroupIn struct {
	input
	Found              bool
	MappedTypeAndGroup mappedTypeAndGroup
}

func (tgi *typeAndGroupIn) resolve(fn func(code string) (mappedTypeAndGroup, bool)) *typeAndGroupIn {

	if tgi.isValidFmt() {
		typeAndGroup, found := fn(tgi.Value())
		tgi.Found = found
		tgi.MappedTypeAndGroup = typeAndGroup
	}
	return tgi
}

func (cdi *checkDigitIn) calcCheckDigit(ocIn ownerCodeInResolvable, eciIn input, snIn input) {

	cdi.CalcCheckDigit = calcCheckDigit(newOwnerCode(ocIn.input.Value()), newEquipCatIDFrom(eciIn.Value()), serialNumFrom(snIn.Value()))
	if cdi.isValidFmt() {
		cd, _ := strconv.Atoi(cdi.Value())
		cdi.IsValidCheckDigit = cd == cdi.CalcCheckDigit
	}
}

type ownerCodeIn struct {
	RegexIn               regexIn
	ownerCodeInResolvable ownerCodeInResolvable
}

type contNumIn struct {
	RegexIn        regexIn
	OwnerCodeIn    ownerCodeInResolvable
	EquipCatIDIn   input
	SerialNumIn    input
	CheckDigitIn   checkDigitIn
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

func (cn contNumIn) isCheckDigitCalculable() bool {
	return cn.OwnerCodeIn.isValidFmt() && cn.EquipCatIDIn.isValidFmt() && cn.SerialNumIn.isValidFmt()
}

func parseOwnerCodeOptEquipCat(in string) ownerCodeIn {
	ownerOptCat := ownerCodeIn{}
	parse := parse(in, ownerCodeOptEquipCatIDMatcher)
	ownerOptCat.RegexIn = parse
	ownerOptCat.ownerCodeInResolvable = ownerCodeInResolvable{input: newIn(parse.getMatches(0, 3), 3)}
	return ownerOptCat
}

func parseContNum(in string) contNumIn {
	cni := contNumIn{}
	parse := parse(in, contNumMatcher)
	cni.RegexIn = parse
	cni.OwnerCodeIn = ownerCodeInResolvable{input: newIn(parse.getMatches(0, 3), 3)}
	cni.EquipCatIDIn = newIn(parse.getMatch(3), 1)
	cni.SerialNumIn = newIn(parse.getMatches(4, 10), 6)

	cni.CheckDigitIn = checkDigitIn{input: newIn(parse.getMatch(10), 1)}
	if cni.isCheckDigitCalculable() {
		cni.CheckDigitIn.calcCheckDigit(cni.OwnerCodeIn, cni.EquipCatIDIn, cni.SerialNumIn)
	}

	cni.LengthIn = lengthIn{input: newIn(parse.getMatch(11), 1)}
	cni.HeightWidthIn = heightWidthIn{input: newIn(parse.getMatch(12), 1)}
	cni.TypeAndGroupIn = typeAndGroupIn{input: newIn(parse.getMatch(13), 2)}

	return cni
}

func parseSizeType(in string) sizeTypeIn {
	sizeType := sizeTypeIn{}
	parse := parse(in, sizeTypeMatcher)
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
