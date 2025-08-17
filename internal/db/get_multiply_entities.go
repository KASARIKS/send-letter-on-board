package db

import (
	"database/sql"

	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbletter"
)

func (db *DB) GetLimitedNumberOfLetters(startPosition, lettersAmount int) ([]*dbletter.DbLetter, error) {
	rows, err := db.db.Query("SELECT * FROM letters LIMIT :startPosition, :lettersAmount",
		sql.Named("startPosition", startPosition),
		sql.Named("lettersAmount", lettersAmount))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gottenLetters []*dbletter.DbLetter

	for rows.Next() {
		var letter dbletter.DbLetter

		if err := rows.Scan(&letter.Id, &letter.Header, &letter.Text, &letter.Owner_id); err != nil {
			return gottenLetters, err
		}
		gottenLetters = append(gottenLetters, &letter)
	}

	return gottenLetters, nil
}
