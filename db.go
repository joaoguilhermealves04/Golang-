package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func initdb(filepath string) *sql.DB {

	db, err := sql.Open("sqlite", filepath)

	if err != nil {
		log.Fatal(err)
	}
	if db == nil {
		log.Fatal("db is nil")
	}
	createTable(db)
	return db
}
func createTable(db *sql.DB) {

	sqlTable := `CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title text NOT NULL
);`

	_, err := db.Exec(sqlTable)
	if err != nil {
		log.Fatal("Failed to create table: %v", err)
	}
}
