package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	database, err := sql.Open("sqlite3", "./c2.db")
	if err != nil {
		panic(err)
	}

	createCommandsTable := `CREATE TABLE IF NOT EXISTS commands (
        id INTEGER PRIMARY KEY, 
        command TEXT, 
        status TEXT
    );`
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        username TEXT UNIQUE,
        password TEXT
    );`

	_, err = database.Exec(createCommandsTable)
	if err != nil {
		panic(err)
	}
	_, err = database.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	return database
}
