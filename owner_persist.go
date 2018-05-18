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
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type owner struct {
	code    string
	company string
	city    string
	country string
}

func (o owner) Code() string {
	return o.code
}

func (o owner) Company() string {
	return o.company
}

func (o owner) City() string {
	return o.city
}

func (o owner) Country() string {
	return o.country
}

func newOwner(code, company, city, country string) owner {
	return owner{code, company, city, country}
}

func deleteOwners(db *sql.DB) {

	tx, err := db.Begin()
	checkErr(err)

	stmt, err := tx.Prepare("DELETE FROM owner")
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	checkErr(err)
	tx.Commit()
}

func getUpdatedTime(db *sql.DB) (lastUpdated time.Time) {
	rows, err := db.Query("SELECT last_updated FROM updated")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&lastUpdated)
	}
	err = rows.Err()
	checkErr(err)
	return
}

func saveUpdatedTimeNow(db *sql.DB) {

	tx, err := db.Begin()
	checkErr(err)

	stmtDel, err := tx.Prepare("DELETE FROM updated")
	checkErr(err)
	defer stmtDel.Close()

	_, err = stmtDel.Exec()
	checkErr(err)

	stmtIns, err := tx.Prepare("INSERT INTO updated (last_updated) VALUES (datetime('now'))")
	checkErr(err)
	defer stmtIns.Close()

	_, err = stmtIns.Exec()
	checkErr(err)

	tx.Commit()
}

func saveOwners(db *sql.DB, owners []owner) {

	tx, err := db.Begin()
	checkErr(err)

	stmt, err := tx.Prepare("INSERT INTO owner (code, company, city, country) VALUES (?, ?, ?, ?)")
	checkErr(err)
	defer stmt.Close()

	for _, owner := range owners {
		_, err = stmt.Exec(owner.Code(), owner.Company(), owner.City(), owner.Country())
		checkErr(err)
	}
	tx.Commit()
}

func getOwner(db *sql.DB, code ownerCode) (owner owner, found bool) {
	stmt, err := db.Prepare("SELECT company, city, country FROM owner WHERE code = ?")
	checkErr(err)
	defer stmt.Close()

	var company string
	var city string
	var country string
	err = stmt.QueryRow(code.Value()).Scan(&company, &city, &country)
	if err != nil {
		return
	}
	return newOwner(code.Value(), company, city, country), true
}

func getRandomCodes(db *sql.DB, count int) []ownerCode {
	stmt, err := db.Prepare(`
		SELECT code
		FROM owner
		ORDER BY RANDOM()
		LIMIT MIN(?, (SELECT COUNT(_rowid_) FROM owner)) 
                              `)
	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(count)
	checkErr(err)

	var codes []ownerCode
	for rows.Next() {
		var code string
		err = rows.Scan(&code)
		checkErr(err)
		codes = append(codes, newOwnerCode(code))
	}
	return codes
}

func initDB(pathToDB string) {

	db := openDB(pathToDB)
	defer db.Close()

	sqlStmtOwnerExists := `SELECT name FROM sqlite_master WHERE type='table' AND name='owner';`
	rows, err := db.Query(sqlStmtOwnerExists)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmtOwnerExists)
	}

	isInitialized := false
	for rows.Next() {
		isInitialized = true
	}

	if !isInitialized {
		_, err = db.Exec(sqlDump())
		if err != nil {
			log.Printf("data initialization error\n")
		}
	}

}

func openDB(pathToDB string) *sql.DB {
	db, err := sql.Open("sqlite3", pathToDB)
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
