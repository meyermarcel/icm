package ui

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/meyermarcel/iso6346/equip_cat"
	"github.com/meyermarcel/iso6346/parser"
	"strings"
)

var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var grey = color.New(color.FgBlack).SprintFunc()

var bold = color.New(color.Bold).SprintFunc()
var underline = color.New(color.Underline).SprintFunc()

const missingCharacter = "_"
const indent = " "

func fmtRegexIn(pi parser.RegexIn) string {

	b := strings.Builder{}
	b.WriteString("'")
	inAsRunes := []rune(pi.Value())
	for pos, char := range inAsRunes {
		if pi.Match(pos) {
			b.WriteString(fmt.Sprintf("%s", string(char)))
		} else {
			b.WriteString(fmt.Sprintf("%s", grey(string(char))))
		}
	}
	b.WriteString("'")
	return b.String()
}

func fmtOwnerCodeOptEquipCat(oce parser.OwnerCodeOptEquipCat, sepOwnerEquip string) string {

	b := strings.Builder{}
	b.WriteString(indent)

	b.WriteString(fmtOwnerCode(oce.OwnerCodeIn))

	if oce.EquipCatIn.IsValidFmt() {
		b.WriteString(sepOwnerEquip)
		b.WriteString(fmtIn(oce.EquipCatIn.In))
	}

	b.WriteString(fmtCheckMark(oce.OwnerCodeIn.IsValidFmt()))

	b.WriteString(fmt.Sprintln())

	b.WriteString(fmtTextsWithArrows(ownerCodeTxt(oce.OwnerCodeIn)))

	return b.String()
}

func fmtParsedContNum(cn parser.ContNum, seps Separators) string {

	b := strings.Builder{}

	b.WriteString(fmtContNum(cn, seps))

	b.WriteString(fmtCheckMark(cn.CheckDigitIn.IsValidCheckDigit))

	b.WriteString(fmt.Sprintln())

	var texts []PosTxt

	texts = append(texts, ownerCodeTxt(cn.OwnerCodeIn))

	if !cn.EquipCatIdIn.IsValidFmt() {
		texts = append(texts, NewPosHint(len(indent)+len(seps.OwnerEquip)+3, fmt.Sprintf("%s must be %s", underline("equipment category id"), equipCatIdsAsList())))
	}
	if !cn.SerialNumIn.IsValidFmt() {
		texts = append(texts, NewPosHint(len(indent)+len(seps.OwnerEquip)+len(seps.EquipSerial)+6, fmt.Sprintf("%s must be %s", underline("serial number"), bold("6 numbers"))))
	}

	cdPos := len(indent) + len(seps.OwnerEquip) + len(seps.EquipSerial) + len(seps.SerialCheck) + 10
	if !cn.CheckDigitIn.IsValidCheckDigit {
		if cn.IsCheckDigitCalculable() {
			if cn.CheckDigitIn.IsValidFmt() {
				texts = append(texts, NewPosHint(cdPos, fmt.Sprintf("%s is incorrect (correct: %s)", underline("check digit"),
					green(cn.CheckDigitIn.CalcCheckDigit))))
			} else {
				texts = append(texts, NewPosHint(cdPos, fmt.Sprintf("%s must be a %s (correct: %s)", underline("check digit"), bold("number"),
					green(cn.CheckDigitIn.CalcCheckDigit))))
			}
		} else {
			texts = append(texts, NewPosHint(cdPos, fmt.Sprintf("%s is not calculable", underline("check digit"))))
		}
	}
	b.WriteString(fmtTextsWithArrows(texts...))

	return b.String()
}

func ownerCodeTxt(ownerCodeIn parser.OwnerCodeIn) PosTxt {
	if !ownerCodeIn.IsValidFmt() {
		return NewPosHint(len(indent)+1, fmt.Sprintf("%s must be %s", underline("owner code"), bold("3 letters")))
	}
	if ownerCodeIn.OwnerFound {
		return NewPosInfo(len(indent)+1, ownerCodeIn.FoundOwner.Company(), ownerCodeIn.FoundOwner.City(), ownerCodeIn.FoundOwner.Country())
	}
	return NewPosInfo(len(indent)+1, fmt.Sprintf("%s not found", underline("owner")))

}

func fmtContNum(cni parser.ContNum, seps Separators) string {

	b := strings.Builder{}

	b.WriteString(indent)
	b.WriteString(fmtOwnerCode(cni.OwnerCodeIn))
	b.WriteString(seps.OwnerEquip)
	b.WriteString(fmtIn(cni.EquipCatIdIn.In))
	b.WriteString(seps.EquipSerial)
	b.WriteString(fmtIn(cni.SerialNumIn.In))
	b.WriteString(seps.SerialCheck)

	if !cni.CheckDigitIn.IsValidCheckDigit && cni.CheckDigitIn.IsValidFmt() {
		b.WriteString(fmt.Sprintf("%s", yellow(string(cni.CheckDigitIn.Value()))))
	} else {
		b.WriteString(fmtIn(cni.CheckDigitIn.In))
	}

	return b.String()
}

func fmtOwnerCode(ownerCodeIn parser.OwnerCodeIn) string {
	if ownerCodeIn.IsValidFmt() {
		if ownerCodeIn.OwnerFound {
			return fmt.Sprintf("%s", green(ownerCodeIn.Value()))
		}
		return fmt.Sprintf("%s", yellow(ownerCodeIn.Value()))
	}
	return fmtIn(ownerCodeIn.In)
}

func fmtIn(in parser.In) string {

	if in.IsValidFmt() {
		return fmt.Sprintf("%s", green(in.Value()))
	}

	b := strings.Builder{}

	startIndexMissingCharacters := 0
	for pos, element := range in.Value() {
		b.WriteString(fmt.Sprintf("%s", yellow(string(element))))
		startIndexMissingCharacters = pos + 1
	}

	for i := startIndexMissingCharacters; i < in.ValidLen(); i++ {
		b.WriteString(fmt.Sprintf("%s", red(missingCharacter)))
	}

	return b.String()
}

func fmtCheckMark(valid bool) string {

	b := strings.Builder{}
	b.WriteString("  ")

	if valid {
		b.WriteString(fmt.Sprintf("%s", green("✔")))
		return b.String()
	}
	b.WriteString(fmt.Sprintf("%s", red("✘")))
	return b.String()
}

func equipCatIdsAsList() string {
	ujz := equip_cat.Ids
	return fmt.Sprintf("%s, %s or %s", green(string(ujz[0])), green(string(ujz[1])), green(string(ujz[2])))
}
