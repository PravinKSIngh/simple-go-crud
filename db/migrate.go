package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// The `InitDB` function initializes a SQLite database connection and creates a users table if it does
// not already exist.
func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSql := `create table if not exists users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name string,
		email string
	);`
	_, err = db.Exec(createTableSql)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
