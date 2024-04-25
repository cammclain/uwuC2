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

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS commands (id INTEGER PRIMARY KEY, command TEXT, status TEXT)")
	statement.Exec()
	return database
}
