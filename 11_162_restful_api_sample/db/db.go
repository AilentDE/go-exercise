package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db?_foreign_keys=on")
	if err != nil {
		panic("Cluld not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	var err error
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err = DB.Exec(createUsersTable)
	if err != nil {
		log.Panicln("Could not create users table.", err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Panicln("Could not create events table.", err)
	}

	createRegisterationsTable := `
	CREATE TABLE IF NOT EXISTS registerations (
		if INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegisterationsTable)
	if err != nil {
		log.Panicln("Could not create registeration table.", err)
	}

}

// [IMPORTANT] If you are using Windows system, make sure you set CGO_ENABLED environment variable and installed gcc
// ref: https://blog.csdn.net/rejoicewindow/article/details/134015972
