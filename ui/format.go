package ui

import (
	"fmt"
	"github.com/fatih/color"
	"iso6346/equip_cat"
	"iso6346/parser"
)

var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var grey = color.New(color.FgBlack).SprintFunc()

var bold = color.New(color.Bold).SprintFunc()
var underline = color.New(color.Underline).SprintFunc()

const missingCharacter = "_"

func fmtRegexIn(pi parser.RegexIn) string {

	out := "'"

	inAsRunes := []rune(pi.Value())
	for pos, char := range inAsRunes {
		if pi.Match(pos) {
			out += fmt.Sprintf("%s", string(char))
		} else {
			out += fmt.Sprintf("%s", grey(string(char)))
		}
	}
	out += "'"
	return out
}

func fmtOwnerCodeOptEquipCat(oce parser.OwnerCodeOptEquipCat) string {

	out := " "

	out += fmtIn(oce.OwnerCodeIn.In)

	if oce.EquipCatIn.IsValidFmt() {
		out += " "
		out += fmtIn(oce.EquipCatIn.In)
	}

	out += formatCheckMark(oce.OwnerCodeIn.IsValidFmt())

	out += fmt.Sprintln()

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

	out += formatMessagesWithArrows(messages)

	return out
}

func fmtParsedContNum(cn parser.ContNum) string {

	out := ""

	out += fmtContNum(cn)

	out += formatCheckMark(cn.CheckDigitIn.IsValidCheckDigit)

	out += fmt.Sprintln()

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
	out += formatMessagesWithArrows(messages)

	return out
}

func fmtContNum(cni parser.ContNum) string {

	out := " "

	if cni.OwnerCodeIn.IsValidFmt() {
		if cni.OwnerCodeIn.OwnerFound {
			out += fmt.Sprintf("%s", green(cni.OwnerCodeIn.Value()))
		} else {
			out += fmt.Sprintf("%s", yellow(cni.OwnerCodeIn.Value()))
		}
	} else {
		out += fmtIn(cni.OwnerCodeIn.In)
	}

	out += " "
	out += fmtIn(cni.EquipCatIdIn.In)
	out += " "
	out += fmtIn(cni.SerialNumIn.In)
	out += " "

	if !cni.CheckDigitIn.IsValidCheckDigit && cni.CheckDigitIn.IsValidFmt() {
		out += fmt.Sprintf("%s", yellow(string(cni.CheckDigitIn.Value())))
	} else {
		out += fmtIn(cni.CheckDigitIn.In)
	}

	return out
}

func fmtIn(in parser.In) string {

	if in.IsValidFmt() {
		return fmt.Sprintf("%s", green(in.Value()))
	}

	out := ""

	startIndexMissingCharacters := 0
	for pos, element := range in.Value() {
		out += fmt.Sprintf("%s", yellow(string(element)))
		startIndexMissingCharacters = pos + 1
	}

	for i := startIndexMissingCharacters; i < in.ValidLen(); i++ {
		out += fmt.Sprintf("%s", red(missingCharacter))
	}

	return out
}

func formatCheckMark(valid bool) string {

	out := "  "

	if valid {
		return out + fmt.Sprintf("%s", green("✔"))
	}
	return out + fmt.Sprintf("%s", red("✘"))

}

func equipCatIdsAsList() string {
	ujz := equip_cat.Ids
	return fmt.Sprintf("%s, %s or %s", green(string(ujz[0])), green(string(ujz[1])), green(string(ujz[2])))
}
