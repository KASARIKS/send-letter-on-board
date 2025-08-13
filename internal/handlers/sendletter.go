package handlers

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbletter"
	"github.com/kasariks/send-letter-on-board/internal/jwtinfo"
)

func SendLetterPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/sendletter.html")
}

func SendLetter(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		http.Error(w, errors.New("request not from form").Error(), http.StatusBadRequest)
		return
	}

	authCookie := r.Cookies()[0]
	authTokenString := authCookie.Value

	authToken, err := jwtinfo.ParseToken(authTokenString)
	if err != nil {
		http.Error(w, err.Error()+"fff", http.StatusBadRequest)
		return
	}

	if !authToken.Valid {
		http.Error(w, errors.New("invalid token").Error(), http.StatusBadRequest)
		return
	}

	claims, ok := authToken.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, errors.New("invalid token format").Error(), http.StatusBadRequest)
		return
	}

	gottenNickname, ok := claims["nickname"].(string)
	if !ok {
		http.Error(w, errors.New("invalid claims format").Error(), http.StatusBadRequest)
		return
	}

	gottenUser, err := handlersDb.GetUserByNickname(gottenNickname)
	if err != nil {
		http.Error(w, err.Error()+"aaa", http.StatusBadRequest)
	}

	header := r.PostFormValue("header")
	text := r.PostFormValue("text")
	newLetter := dbletter.NewDbLetter(0, header, text, gottenUser.Id)
	if err := handlersDb.AddLetter(*newLetter); err != nil {
		http.Error(w, err.Error()+"ggg", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
