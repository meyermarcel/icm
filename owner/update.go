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

package owner

import (
	"time"
)

func Update(pathToDB string) {

	c := make(chan time.Time)

	go getUpdatedTimeRemote(c)

	db := openDB(pathToDB)
	defer db.Close()
	updatedTimeDB := getUpdatedTime(db)

	updatedTimeRemote := <-c

	if updatedTimeRemote.After(updatedTimeDB) {
		owners := getOwnersRemote()
		deleteOwners(db)
		saveOwners(db, owners)
		saveUpdatedTimeNow(db)
	}
}
