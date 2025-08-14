package handlers

import (
	"errors"
	"net/http"

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
	authToken, err := jwtinfo.GetTokenFromCookie(authCookie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	gottenNickname, err := jwtinfo.GetNicknameFromToken(authToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	gottenUser, err := handlersDb.GetUserByNickname(gottenNickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	header := r.PostFormValue("header")
	text := r.PostFormValue("text")

	newLetter := dbletter.NewDBLetterWithoutId(header, text, gottenUser.Id)
	if err := handlersDb.AddLetter(*newLetter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
