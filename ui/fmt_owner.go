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
	"github.com/meyermarcel/iso6346/parser"
	"strings"
	"fmt"
)

func fmtOwnerCode(oce parser.OwnerCode) string {

	b := strings.Builder{}
	b.WriteString(indent)

	b.WriteString(fmtOwnerCodeIn(oce.OwnerCodeIn))

	b.WriteString(fmtCheckMark(oce.OwnerCodeIn.IsValidFmt()))

	b.WriteString(fmt.Sprintln())

	b.WriteString(fmtTextsWithArrows(ownerCodeTxt(oce.OwnerCodeIn)))

	return b.String()
}

func ownerCodeTxt(ownerCodeIn parser.OwnerCodeIn) PosTxt {
	if !ownerCodeIn.IsValidFmt() {
		return NewPosHint(len(indent)+1, fmt.Sprintf("%s must be %s", underline("owner code"), bold("3 letters")))
	}
	if ownerCodeIn.OwnerFound {
		return NewPosInfo(len(indent)+1, ownerCodeIn.FoundOwner.Company(), ownerCodeIn.FoundOwner.City(), ownerCodeIn.FoundOwner.Country())
	}
	return NewPosInfo(len(indent)+1, fmt.Sprintf("%s not found", underline("owner code")))

}

func fmtOwnerCodeIn(ownerCodeIn parser.OwnerCodeIn) string {
	if ownerCodeIn.IsValidFmt() {
		if ownerCodeIn.OwnerFound {
			return fmt.Sprintf("%s", green(ownerCodeIn.Value()))
		}
		return fmt.Sprintf("%s", yellow(ownerCodeIn.Value()))
	}
	return fmtIn(ownerCodeIn.In)
}
