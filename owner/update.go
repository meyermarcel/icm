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
