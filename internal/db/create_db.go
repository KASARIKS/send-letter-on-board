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

	return internalDatabase, nil
}
