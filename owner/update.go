package owner

import (
	"time"
)

func Update() {

	c := make(chan time.Time)

	go getUpdatedTimeRemote(c)

	db := InitDB()
	defer db.Close()
	updatedTimeDB := getUpdatedTime(db)

	updatedTimeRemote := <-c

	if updatedTimeRemote.After(updatedTimeDB) {

		c := make(chan []Owner)
		go getOwnersRemote(c)

		saveUpdatedTimeNow(db)
		deleteOwners(db)

		owners := <- c
		saveOwners(db, owners)
	}
}
