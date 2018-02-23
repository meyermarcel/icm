package ui

import (
	"fmt"
	"github.com/fatih/color"
	"iso6346/equip_cat"
	"iso6346/parser"
	"strings"
)

var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var grey = color.New(color.FgBlack).SprintFunc()

var bold = color.New(color.Bold).SprintFunc()
var underline = color.New(color.Underline).SprintFunc()

const missingCharacter = "_"

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

func fmtOwnerCodeOptEquipCat(oce parser.OwnerCodeOptEquipCat) string {

	b := strings.Builder{}
	b.WriteString(" ")

	b.WriteString(fmtIn(oce.OwnerCodeIn.In))

	if oce.EquipCatIn.IsValidFmt() {
		b.WriteString(" ")
		b.WriteString(fmtIn(oce.EquipCatIn.In))
	}

	b.WriteString(fmtCheckMark(oce.OwnerCodeIn.IsValidFmt()))

	b.WriteString(fmt.Sprintln())

	var messages []PosMsg

	if !oce.OwnerCodeIn.IsValidFmt() {
		messages = append(messages, NewPosHint(2, fmt.Sprintf("%s must be %s", underline("owner code"), bold("3 letters"))))
	} else {
		if oce.OwnerCodeIn.OwnerFound {
			messages = append(messages, NewPosInfo(2, oce.OwnerCodeIn.FoundOwner.Company(), oce.OwnerCodeIn.FoundOwner.City(), oce.OwnerCodeIn.FoundOwner.Country()))
		} else {
			messages = append(messages, NewPosHint(2, fmt.Sprintf("%s not found", underline("owner"))))
		}
	}

	b.WriteString(fmtMessagesWithArrows(messages))

	return b.String()
}

func fmtParsedContNum(cn parser.ContNum) string {

	b := strings.Builder{}

	b.WriteString(fmtContNum(cn))

	b.WriteString(fmtCheckMark(cn.CheckDigitIn.IsValidCheckDigit))

	b.WriteString(fmt.Sprintln())

	var messages []PosMsg

	if !cn.OwnerCodeIn.IsValidFmt() {
		messages = append(messages, NewPosHint(2, fmt.Sprintf("%s must be %s", underline("owner code"), bold("3 letters"))))
	} else {
		if cn.OwnerCodeIn.OwnerFound {
			messages = append(messages, NewPosInfo(2, cn.OwnerCodeIn.FoundOwner.Company(), cn.OwnerCodeIn.FoundOwner.City(), cn.OwnerCodeIn.FoundOwner.Country()))
		} else {
			messages = append(messages, NewPosHint(2, fmt.Sprintf("%s not found", underline("owner"))))
		}
	}
	if !cn.EquipCatIdIn.IsValidFmt() {
		messages = append(messages, NewPosHint(5, fmt.Sprintf("%s must be %s", underline("equipment category id"), equipCatIdsAsList())))
	}
	if !cn.SerialNumIn.IsValidFmt() {
		messages = append(messages, NewPosHint(9, fmt.Sprintf("%s must be %s", underline("serial number"), bold("6 numbers"))))
	}

	if !cn.CheckDigitIn.IsValidCheckDigit {
		if cn.IsCheckDigitCalculable() {
			if cn.CheckDigitIn.IsValidFmt() {
				messages = append(messages, NewPosHint(14, fmt.Sprintf("%s is incorrect (correct: %s)", underline("check digit"),
					green(cn.CheckDigitIn.CalcCheckDigit))))
			} else {
				messages = append(messages, NewPosHint(14, fmt.Sprintf("%s must be a %s (correct: %s)", underline("check digit"), bold("number"),
					green(cn.CheckDigitIn.CalcCheckDigit))))
			}
		} else {
			messages = append(messages, NewPosHint(14, fmt.Sprintf("%s is not calculable", underline("check digit"))))
		}
	}
	b.WriteString(fmtMessagesWithArrows(messages))

	return b.String()
}

func fmtContNum(cni parser.ContNum) string {

	b := strings.Builder{}
	b.WriteString(" ")

	if cni.OwnerCodeIn.IsValidFmt() {
		if cni.OwnerCodeIn.OwnerFound {
			b.WriteString(fmt.Sprintf("%s", green(cni.OwnerCodeIn.Value())))
		} else {
			b.WriteString(fmt.Sprintf("%s", yellow(cni.OwnerCodeIn.Value())))
		}
	} else {
		b.WriteString(fmtIn(cni.OwnerCodeIn.In))
	}

	b.WriteString(" ")
	b.WriteString(fmtIn(cni.EquipCatIdIn.In))
	b.WriteString(" ")
	b.WriteString(fmtIn(cni.SerialNumIn.In))
	b.WriteString(" ")

	if !cni.CheckDigitIn.IsValidCheckDigit && cni.CheckDigitIn.IsValidFmt() {
		b.WriteString(fmt.Sprintf("%s", yellow(string(cni.CheckDigitIn.Value()))))
	} else {
		b.WriteString(fmtIn(cni.CheckDigitIn.In))
	}

	return b.String()
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
