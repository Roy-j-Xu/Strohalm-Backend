package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initPanic(err error, message string) {
	if err != nil {
		panic(message + ": " + err.Error())
	}
}

func InitDatabase() {
	var err error

	db, err = sql.Open("sqlite3", "./data.db")
	initPanic(err, "Failed to open database")

	err = db.Ping()
	initPanic(err, "Failed to ping database")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS customers (
            Id INTEGER PRIMARY KEY AUTOINCREMENT,
            Name TEXT NOT NULL
        )
	`)
	initPanic(err, "Failed to create table `customers`")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS visits (
            Id INTEGER PRIMARY KEY AUTOINCREMENT,
			CustomerId INTEGER,
			VisitDate TEXT NOT NULL,
			FOREIGN KEY (CustomerId) REFERENCES customers(Id)
        )
	`)
	initPanic(err, "Failed to create table `visits`")
}
