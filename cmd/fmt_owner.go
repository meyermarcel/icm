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

package cmd

import (
	"fmt"
	"strings"

	"github.com/meyermarcel/icm/internal/cont"
)

func fmtOwnerCode(oce ownerCodeOptEquipCatIDIn, fn func(count int) []cont.OwnerCode) string {

	b := strings.Builder{}
	b.WriteString(indent)

	b.WriteString(fmtOwnerCodeIn(oce.ownerCodeIn))

	b.WriteString(fmtCheckMark(oce.ownerCodeIn.isValidFmt()))

	b.WriteString(fmt.Sprintln())

	b.WriteString(fmtTextsWithArrows(ownerCodeTxt(oce.ownerCodeIn, fn)))

	return b.String()
}

func ownerCodeTxt(ownerCodeIn ownerCodeIn, randomOwnerCodes func(count int) []cont.OwnerCode) posTxt {
	if !ownerCodeIn.isValidFmt() {
		return newPosHint(indentSize+1, fmt.Sprintf("%s must be %s and %s (e.g. %s)",
			underline("ownerDecodeUpdater code"),
			bold("3 letters"),
			bold("registered"),
			underline(randomOwnerCodes(1)[0].Value())))
	}
	return newPosInfo(indentSize+1,
		ownerCodeIn.Owner.Company,
		ownerCodeIn.Owner.City,
		ownerCodeIn.Owner.Country)
}

func fmtOwnerCodeIn(ownerCodeIn ownerCodeIn) string {
	if ownerCodeIn.isValidFmt() {
		return fmt.Sprintf("%s", green(ownerCodeIn.value))
	}
	return fmtIn(ownerCodeIn.input)
}
