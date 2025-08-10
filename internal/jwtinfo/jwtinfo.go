package jwtinfo

import "github.com/golang-jwt/jwt/v5"

const JwtKey = "opiqwjk;njorgeqoijreqgopkftgeolp;kdvfoldkp;afv"

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return []byte(JwtKey), nil
	})
}
