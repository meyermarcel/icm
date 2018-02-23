package parser

import (
	"regexp"
	"unicode/utf8"
	"strings"
	"iso6346/owner"
	"iso6346/equip_cat"
	"iso6346/cont"
	"strconv"
)

const ownerCodeOptEquipCatIdRegex = `([A-Za-z])[^A-Za-z\d]*([A-Za-z])?[^A-Za-z\d]*([A-Za-z])?[^JUZjuz\d]*([JUZjuz])?`

var ContNumMatcher = regexp.MustCompile(contNumRegex)

const contNumRegex = ownerCodeOptEquipCatIdRegex + `[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?[^\d]*(\d)?`

var OwnerCodeOptEquipCatIdMatcher = regexp.MustCompile(ownerCodeOptEquipCatIdRegex)

type RegexIn struct {
	matches        []string
	value          string
	matchesIndices map[int]bool
}

func (pi RegexIn) getMatches(start, end int) string {

	if len(pi.matches) == 0 {
		return ""
	}
	value := ""
	for _, element := range pi.matches[start:end] {
		value += element
	}
	return strings.ToUpper(value)
}

func (pi RegexIn) getMatch(pos int) string {

	return pi.getMatches(pos, pos+1)
}

func (pi RegexIn) Value() string {
	return pi.value
}

func (pi RegexIn) Match(atPosition int) bool {
	return pi.matchesIndices[atPosition]
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

func (oi OwnerCodeIn) resolve(fn func(code owner.Code) (owner.Owner, bool)) OwnerCodeIn {
	if oi.IsValidFmt() {
		foundOwner, found := fn(owner.NewCode(oi.Value()))
		oi.OwnerFound = found
		oi.FoundOwner = foundOwner
	}
	return oi
}

type EquipCatIdIn struct {
	In
}

type SerialNumIn struct {
	In
}

type CheckDigitIn struct {
	In
	IsValidCheckDigit bool
	CalcCheckDigit    int
}

func (cdi *CheckDigitIn) calcCheckDigit(ocIn OwnerCodeIn, eciIn EquipCatIdIn, snIn SerialNumIn) {

	cdi.CalcCheckDigit = cont.CalcCheckDigit(owner.NewCode(ocIn.In.Value()), equip_cat.NewId(eciIn.Value()), cont.NewSerialNum(snIn.Value()))
	if cdi.IsValidFmt() {
		cd, _ := strconv.Atoi(cdi.Value())
		cdi.IsValidCheckDigit = cd == cdi.CalcCheckDigit
	}
}

type OwnerCodeOptEquipCat struct {
	RegexIn     RegexIn
	OwnerCodeIn OwnerCodeIn
	EquipCatIn  EquipCatIdIn
}

type ContNum struct {
	RegexIn      RegexIn
	OwnerCodeIn  OwnerCodeIn
	EquipCatIdIn EquipCatIdIn
	SerialNumIn  SerialNumIn
	CheckDigitIn CheckDigitIn
}

func (cn ContNum) IsCheckDigitCalculable() bool {
	return cn.OwnerCodeIn.IsValidFmt() && cn.EquipCatIdIn.IsValidFmt() && cn.SerialNumIn.IsValidFmt()
}

func ParseOwnerCodeOptEquipCat(in string) OwnerCodeOptEquipCat {
	ownerOptCat := OwnerCodeOptEquipCat{}
	parse := parse(in, *OwnerCodeOptEquipCatIdMatcher)
	ownerOptCat.RegexIn = parse
	ownerOptCat.OwnerCodeIn = OwnerCodeIn{In:NewIn(parse.getMatches(0, 3), 3)}.resolve(owner.Resolver())
	ownerOptCat.EquipCatIn = EquipCatIdIn{NewIn(parse.getMatch(3), 1)}
	return ownerOptCat
}

func ParseContNum(in string) ContNum {
	cni := ContNum{}
	parse := parse(in, *ContNumMatcher)
	cni.RegexIn = parse
	cni.OwnerCodeIn = OwnerCodeIn{In: NewIn(parse.getMatches(0, 3), 3)}.resolve(owner.Resolver())
	cni.EquipCatIdIn = EquipCatIdIn{NewIn(parse.getMatch(3), 1)}
	cni.SerialNumIn = SerialNumIn{NewIn(parse.getMatches(4, 10), 6)}
	cni.CheckDigitIn = CheckDigitIn{In: NewIn(parse.getMatch(10), 1)}
	if cni.IsCheckDigitCalculable() {
		cni.CheckDigitIn.calcCheckDigit(cni.OwnerCodeIn, cni.EquipCatIdIn, cni.SerialNumIn)
	}
	return cni
}

func parse(in string, matcher regexp.Regexp) RegexIn {

	regexIn := RegexIn{value: in}

	subMatch := matcher.FindAllStringSubmatch(in, -1)

	if len(subMatch) == 0 {
		return regexIn
	}

	regexIn.matches = subMatch[0][1:]

	matchRanges := [22]int{}

	copy(matchRanges[:], ContNumMatcher.FindAllStringSubmatchIndex(in, -1)[0][2:])

	regexIn.matchesIndices = byteToRuneIndex(in, matchRanges)

	return regexIn
}

func byteToRuneIndex(in string, matchRanges [22]int) map[int]bool {
	matchesIndices := [11]int{}

	for i := 0; i < len(matchRanges)/2; i++ {
		matchesIndices[i] = matchRanges[i*2]
	}

	byteShiftsForIndices := [11]int{}

	for i := 0; i < len(in); i++ {
		if !utf8.RuneStart(in[i]) {
			for pos, element := range matchesIndices {
				if element > i {
					byteShiftsForIndices[pos]--
				}
			}
		}
	}

	// apply byte shift indices
	for pos, element := range matchesIndices {
		matchesIndices[pos] = element + byteShiftsForIndices[pos]
	}
	var matchesIndicesMap = map[int]bool{}

	for _, element := range matchesIndices {
		if element >= 0 {
			matchesIndicesMap[element] = true
		}
	}
	return matchesIndicesMap
}
