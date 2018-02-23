package owner

import (
	"log"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"path/filepath"
	"os"
	"github.com/mitchellh/go-homedir"
	"time"
)

const appDir = ".iso6346"
const dbName = "iso6346.db"
const dirPerm = 0700

type Owner struct {
	code    string
	company string
	city    string
	country string
}

func (o Owner) Code() string {
	return o.code
}

func (o Owner) Company() string {
	return o.company
}

func (o Owner) City() string {
	return o.city
}

func (o Owner) Country() string {
	return o.country
}

func NewOwner(code, company, city, country string) Owner {
	return Owner{code, company, city, country}
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

func saveOwners(db *sql.DB, owners []Owner) {

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

func getOwner(db *sql.DB, code Code) (owner Owner, found bool) {
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
	return NewOwner(code.Value(), company, city, country), true
}

func getPathToAppDir() string {
	homeDir, err := homedir.Dir()
	checkErr(err)
	return filepath.Join(homeDir, appDir)
}

func initDir(path string) string {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir|dirPerm)
	}
	return path
}

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", filepath.Join(initDir(getPathToAppDir()), dbName))
	checkErr(err)

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
		sqlStmtInit := sqlDump()
		_, err = db.Exec(sqlStmtInit)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmtInit)
		}
	}
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
