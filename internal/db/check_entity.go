package db

import (
	"errors"

	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbuser"
)

func (db *DB) CheckUserPassword(inputUser *dbuser.DbUser) error {
	realUser, err := db.GetUserByNickname(inputUser.Nickname)
	if err != nil {
		return err
	}

	if realUser.EncryptedPassword != inputUser.EncryptedPassword {
		return errors.New("wrong password")
	}

	return nil
}
