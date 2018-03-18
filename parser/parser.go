package parser

import (
	"github.com/meyermarcel/iso6346/cont"
	"github.com/meyermarcel/iso6346/equip_cat"
	"github.com/meyermarcel/iso6346/owner"
	"regexp"
	"strconv"
	"strings"
)

const ownerCodeRegex = `(?i)([A-Z])[^A-Z]*([A-Z])?[^A-Z]*([A-Z])`

var ownerCodeOptEquipCatIdMatcher = regexp.MustCompile(ownerCodeRegex)

const contNumRegex = ownerCodeRegex + `[^JUZ]*([JUZ])?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?`

var contNumMatcher = regexp.MustCompile(contNumRegex)

const sizeTypeRegex = `(?i)([1234ABCDEFGHKLMNP])[^0245689CDEFLMNP]*([0245689CDEFLMNP])?[^ABGHKNPRSUV]*(A0|B0|B1|B2|B3|B4|B5|B6|B7|B8|B9|G1|G2|G3|G4|G5|G6|G7|G8|G9|H0|H1|H2|H3|H4|H5|H6|H7|H8|H9|K0|K1|K2|K3|K4|K5|K6|K7|K8|K9|N0|N1|N2|N3|N4|N5|N6|N7|N8|N9|N9|P0|P1|P2|P3|P4|P5|P6|P7|P8|P9|R0|R1|R2|R3|R4|R5|R6|R7|R8|R9|S0|S1|S2|S3|S4|S5|S6|S7|S8|S9|U0|U1|U2|U3|U4|U5|U6|U7|U8|U9|V0|V1|V2|V3|V4|V5|V6|V7|V8|V9)?`

var sizeTypeMatcher = regexp.MustCompile(sizeTypeRegex)

type RegexIn struct {
	matches     []string
	input       string
	matchRanges []int
}

func (pi RegexIn) getMatches(start, end int) (value string) {

	if len(pi.matches) == 0 {
		return
	}
	for _, element := range pi.matches[start:end] {
		value += element
	}
	return strings.ToUpper(value)
}

func (pi RegexIn) getMatch(pos int) string {

	return pi.getMatches(pos, pos+1)
}

func (pi RegexIn) Input() string {
	return pi.input
}

func (pi RegexIn) MatchRanges() []int {
	return pi.matchRanges
}

type In struct {
	value    string
	validLen int
}

func (in In) IsValidFmt() bool {
	return len(in.value) == in.validLen
}

func (in In) Value() string {
	return in.value
}

func (in In) ValidLen() int {
	return in.validLen
}

func NewIn(value string, validLen int) In {
	return In{value: value, validLen: validLen}
}

type OwnerCodeIn struct {
	In
	OwnerFound bool
	FoundOwner owner.Owner
}

func (oi *OwnerCodeIn) Resolve(fn func(code owner.Code) (owner.Owner, bool)) *OwnerCodeIn {

	if oi.IsValidFmt() {
		foundOwner, found := fn(owner.NewCode(oi.Value()))
		oi.OwnerFound = found
		oi.FoundOwner = foundOwner
	}
	return oi
}

type CheckDigitIn struct {
	In
	IsValidCheckDigit bool
	CalcCheckDigit    int
}

func (cdi *CheckDigitIn) calcCheckDigit(ocIn OwnerCodeIn, eciIn In, snIn In) {

	cdi.CalcCheckDigit = cont.CalcCheckDigit(owner.NewCode(ocIn.In.Value()), equip_cat.NewIdFrom(eciIn.Value()), cont.SerialNumFrom(snIn.Value()))
	if cdi.IsValidFmt() {
		cd, _ := strconv.Atoi(cdi.Value())
		cdi.IsValidCheckDigit = cd == cdi.CalcCheckDigit
	}
}

type OwnerCode struct {
	RegexIn     RegexIn
	OwnerCodeIn OwnerCodeIn
}

type ContNum struct {
	RegexIn      RegexIn
	OwnerCodeIn  OwnerCodeIn
	EquipCatIdIn In
	SerialNumIn  In
	CheckDigitIn CheckDigitIn
}

type SizeType struct {
	RegexIn       RegexIn
	LengthIn      In
	HeightWidthIn In
	TypeIn        In
}

func (cn ContNum) IsCheckDigitCalculable() bool {
	return cn.OwnerCodeIn.IsValidFmt() && cn.EquipCatIdIn.IsValidFmt() && cn.SerialNumIn.IsValidFmt()
}

func ParseOwnerCodeOptEquipCat(in string) OwnerCode {
	ownerOptCat := OwnerCode{}
	parse := parse(in, *ownerCodeOptEquipCatIdMatcher)
	ownerOptCat.RegexIn = parse
	ownerOptCat.OwnerCodeIn = OwnerCodeIn{In: NewIn(parse.getMatches(0, 3), 3)}
	return ownerOptCat
}

func ParseContNum(in string) ContNum {
	cni := ContNum{}
	parse := parse(in, *contNumMatcher)
	cni.RegexIn = parse
	cni.OwnerCodeIn = OwnerCodeIn{In: NewIn(parse.getMatches(0, 3), 3)}
	cni.EquipCatIdIn = NewIn(parse.getMatch(3), 1)
	cni.SerialNumIn = NewIn(parse.getMatches(4, 10), 6)
	cni.CheckDigitIn = CheckDigitIn{In: NewIn(parse.getMatch(10), 1)}
	if cni.IsCheckDigitCalculable() {
		cni.CheckDigitIn.calcCheckDigit(cni.OwnerCodeIn, cni.EquipCatIdIn, cni.SerialNumIn)
	}
	return cni
}

func ParseSizeType(in string) SizeType {
	sizeType := SizeType{}
	parse := parse(in, *sizeTypeMatcher)
	sizeType.RegexIn = parse
	sizeType.LengthIn = NewIn(parse.getMatch(0, ), 1)
	sizeType.HeightWidthIn = NewIn(parse.getMatch(1), 1)
	sizeType.TypeIn = NewIn(parse.getMatch(2), 2)
	return sizeType
}

func parse(in string, matcher regexp.Regexp) RegexIn {

	regexIn := RegexIn{input: in}

	subMatch := matcher.FindAllStringSubmatch(in, -1)

	if len(subMatch) == 0 {
		return regexIn
	}

	regexIn.matches = subMatch[0][1:]

	regexIn.matchRanges = matcher.FindAllStringSubmatchIndex(in, -1)[0][2:]

	return regexIn
}
