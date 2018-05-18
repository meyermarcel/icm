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
	"strings"
)

func fmtOwnerCode(oce ownerCodeIn) string {

	b := strings.Builder{}
	b.WriteString(indent)

	b.WriteString(fmtOwnerCodeIn(oce.ownerCodeInResolvable))

	b.WriteString(fmtCheckMark(oce.ownerCodeInResolvable.isValidFmt()))

	b.WriteString(fmt.Sprintln())

	b.WriteString(fmtTextsWithArrows(ownerCodeTxt(oce.ownerCodeInResolvable)))

	return b.String()
}

func ownerCodeTxt(ownerCodeIn ownerCodeInResolvable) posTxt {
	if !ownerCodeIn.isValidFmt() {
		return newPosHint(len(indent)+1, fmt.Sprintf("%s must be %s", underline("owner code"), bold("3 letters")))
	}
	if ownerCodeIn.OwnerFound {
		return newPosInfo(len(indent)+1, ownerCodeIn.FoundOwner.Company(), ownerCodeIn.FoundOwner.City(), ownerCodeIn.FoundOwner.Country())
	}
	return newPosInfo(len(indent)+1, fmt.Sprintf("%s not found", underline("owner code")))

}

func fmtOwnerCodeIn(ownerCodeIn ownerCodeInResolvable) string {
	if ownerCodeIn.isValidFmt() {
		if ownerCodeIn.OwnerFound {
			return fmt.Sprintf("%s", green(ownerCodeIn.Value()))
		}
		return fmt.Sprintf("%s", yellow(ownerCodeIn.Value()))
	}
	return fmtIn(ownerCodeIn.input)
}
