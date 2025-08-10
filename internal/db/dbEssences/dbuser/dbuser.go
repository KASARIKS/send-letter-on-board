package dbuser

import (
	"crypto/sha256"
	"encoding/hex"
)

type dbUser struct {
	Id                int    `json:"id"`
	Nickname          string `json:"nickname"`
	EncryptedPassword string `json:"password"`
}

func NewDbUser(id int, nickname, unencrypted_password string) *dbUser {
	encryptedPassword := sha256.Sum256([]byte(unencrypted_password))
	newUser := &dbUser{
		Id:                id,
		Nickname:          nickname,
		EncryptedPassword: hex.EncodeToString(encryptedPassword[:]),
	}

	return newUser
}
