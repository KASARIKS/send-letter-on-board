package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbuser"
	"github.com/kasariks/send-letter-on-board/internal/jwtinfo"
)

func AuthorizationPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/authorization.html")
}

func Authorization(w http.ResponseWriter, r *http.Request) {
	nickname := r.PostFormValue("nickname")
	password := r.PostFormValue("password")

	if err := checkInputUser(nickname, password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookie, err := jwtToCookie(nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func checkInputUser(nickname, password string) error {
	err := handlersDb.CheckUserPassword(dbuser.NewDbUser(
		0,
		nickname,
		password,
	))

	return err
}

func jwtToCookie(nickname string) (*http.Cookie, error) {
	claims := jwt.MapClaims{
		"nickname": nickname,
		"exp":      time.Now().Add(time.Minute * jwtinfo.JwtExistingTimeMinutes),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(jwtinfo.JwtKey))
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     "jwtToken",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Minute * jwtinfo.JwtExistingTimeMinutes),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}
	return cookie, nil
}
