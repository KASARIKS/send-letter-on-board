package dbuser

import (
	"crypto/sha256"
	"encoding/hex"
)

type DbUser struct {
	Id                int    `json:"id"`
	Nickname          string `json:"nickname"`
	EncryptedPassword string `json:"password"`
}

func NewDbUser(id int, nickname, unencrypted_password string) *DbUser {
	encryptedPassword := sha256.Sum256([]byte(unencrypted_password))
	newUser := &DbUser{
		Id:                id,
		Nickname:          nickname,
		EncryptedPassword: hex.EncodeToString(encryptedPassword[:]),
	}

	return newUser
}
