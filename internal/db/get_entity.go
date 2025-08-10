package db

import (
	"database/sql"

	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbuser"
)

func (db *DB) GetUserById(id int) (*dbuser.DbUser, error) {
	row := db.db.QueryRow("SELECT * FROM users WHERE id=:id",
		sql.Named("id", id))

	gottenUser := &dbuser.DbUser{}
	err := row.Scan(&gottenUser.Id, &gottenUser.Nickname, &gottenUser.EncryptedPassword)

	return gottenUser, err
}
