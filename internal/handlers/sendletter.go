package handlers

import (
	"errors"
	"net/http"

	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbletter"
	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbuser"
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

	gottenUser, err := getAuthorizedUser(r.Cookies())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

func getAuthorizedUser(cookies []*http.Cookie) (*dbuser.DbUser, error) {
	if len(cookies) == 0 {
		return nil, errors.New("not authorized")
	}

	// In future should change this cookie check
	authCookie := cookies[len(cookies)-1]
	authToken, err := jwtinfo.GetTokenFromCookie(authCookie)
	if err != nil {
		return nil, err
	}

	gottenNickname, err := jwtinfo.GetNicknameFromToken(authToken)
	if err != nil {
		return nil, err
	}

	gottenUser, err := handlersDb.GetUserByNickname(gottenNickname)
	if err != nil {
		return nil, err
	}

	return gottenUser, err
}
