package handlers

import (
	"net/http"

	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbuser"
	"github.com/kasariks/send-letter-on-board/internal/jwtinfo"
)

func AuthorizationPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/authorization.html")
}

func Authorization(w http.ResponseWriter, r *http.Request) {
	nickname := r.PostFormValue("nickname")
	password := r.PostFormValue("password")
	inputUser := dbuser.NewDbUser(
		0,
		nickname,
		password,
	)

	if err := handlersDb.CheckUserPassword(inputUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookie, err := jwtinfo.GetCookieWithJwtByNickname(inputUser.Nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
