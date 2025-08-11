package handlers

import (
	"errors"
	"net/http"

	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbuser"
)

func RegistrationPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/registration.html")
}

func Registration(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		http.Error(w, errors.New("request not from form").Error(), http.StatusBadRequest)
	}

	if err := addFromForm(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}

func addFromForm(r *http.Request) error {
	nickname := r.PostFormValue("nickname")
	password := r.PostFormValue("password")

	newUser := dbuser.NewDbUser(0, nickname, password)
	err := handlersDb.AddUser(*newUser)

	return err
}
