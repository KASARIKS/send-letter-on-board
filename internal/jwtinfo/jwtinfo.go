package jwtinfo

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtKey = "opiqwjk;njorgeqoijreqgopkftgeolp;kdvfoldkp;afv"
const jwtExistingTimeMinutes = 15

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
}

func GetCookieWithJwtByNickname(nickname string) (*http.Cookie, error) {
	signedToken, err := getSignedToken(nickname)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     "jwtToken",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Minute * jwtExistingTimeMinutes),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}
	return cookie, nil
}

func getSignedToken(nickname string) (string, error) {
	claims := jwt.MapClaims{
		"nickname": nickname,
		"exp":      time.Now().Add(time.Minute * jwtExistingTimeMinutes).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(jwtKey))
	return signedToken, err
}
