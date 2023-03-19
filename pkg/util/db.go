package util

import (
	"database/sql"

	"log"
)

type SqliteIns struct {
	Type           string
	DataBaseSource string
}

func (sqlite3Ins *SqliteIns) Query(query string) (rows *sql.Rows, err error) {
	db, err := sql.Open(sqlite3Ins.Type, sqlite3Ins.DataBaseSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db.Query(query)

}

func (sqlite3Ins *SqliteIns) QueryRow(query string) (rows *sql.Row) {
	db, err := sql.Open(sqlite3Ins.Type, sqlite3Ins.DataBaseSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db.QueryRow(query)
}
