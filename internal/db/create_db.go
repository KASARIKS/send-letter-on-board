package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	db *sql.DB
}

func NewDb() (*DB, error) {
	db, err := sql.Open("sqlite", "db.db")
	if err != nil {
		return nil, err
	}

	internalDatabase := &DB{
		db: db,
	}
	if err := createTables(internalDatabase); err != nil {
		return nil, err
	}

	return internalDatabase, nil
}

func createTables(internalDatabase *DB) error {
	if err := createUsersTable(internalDatabase); err != nil {
		return err
	}

	if err := createLettersTable(internalDatabase); err != nil {
		return err
	}

	return nil
}

func createUsersTable(internalDatabase *DB) error {
	_, tableCheck := internalDatabase.db.Exec("SELECT id FROM users LIMIT 1;")
	if tableCheck != nil {
		_, err := internalDatabase.db.Exec("CREATE TABLE users (" +
			"id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, " +
			"nickname TEXT NOT NULL UNIQUE, " +
			"password TEXT NOT NULL" +
			");")
		if err != nil {
			return err
		}
	}

	return nil
}

func createLettersTable(internalDatabase *DB) error {
	_, tableCheck := internalDatabase.db.Exec("SELECT id FROM letters LIMIT 1;")
	if tableCheck != nil {
		_, err := internalDatabase.db.Exec("CREATE TABLE letters (" +
			"id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, " +
			"header TEXT, " +
			"text TEXT, " +
			"owner_id INTEGER NOT NULL, " +
			"FOREIGN KEY (owner_id) REFERENCES users (id) " +
			");")
		if err != nil {
			return err
		}
	}

	return nil
}
