package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateTables(db *sqlx.DB) error {

	tx := db.MustBegin()

	tx.MustExec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT,
			password BLOB
		)
	`)

	tx.MustExec(`
		CREATE TABLE IF NOT EXISTS session (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			expires_at TEXT,
			token TEXT,
			FOREIGN KEY (id) REFERENCES users(id)
		)
	`)

	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func PrintTables(db *sqlx.DB) {
	var tables []string
	query := `
		SELECT name
		FROM sqlite_master
		WHERE type='table'
		ORDER BY name;
	`
	err := db.Select(&tables, query)
	if err != nil {
		log.Fatalf("Error querying tables: %v", err)
	}

	fmt.Println("Tables in the database:")
	for _, table := range tables {
		fmt.Println(table)
	}
}
