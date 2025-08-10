package db

import (
	"database/sql"

	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbletter"
	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbuser"
)

func (db *DB) AddUser(dbUser dbuser.DbUser) error {
	_, err := db.db.Exec("INSERT INTO users (nickname, password) VALUES (:nickname, :password);",
		sql.Named("nickname", dbUser.Nickname),
		sql.Named("password", dbUser.EncryptedPassword))

	return err
}

func (db *DB) AddLetter(dbLetter dbletter.DbLetter) error {
	if _, err := db.GetUserById(dbLetter.Owner_id); err != nil {
		return err
	}

	_, err := db.db.Exec("INSERT INTO letters (header, text, owner_id) VALUES (:header, :text, :owner_id);",
		sql.Named("header", dbLetter.Header),
		sql.Named("text", dbLetter.Text),
		sql.Named("owner_id", dbLetter.Owner_id))

	return err
}
