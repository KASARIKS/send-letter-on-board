package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbletter"
)

type ViewData struct {
	Letters []*dbletter.DbLetter
}

func GetAllLetters(w http.ResponseWriter, r *http.Request) {
	// if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
	// 	http.Error(w, errors.New("request not from form").Error(), http.StatusBadRequest)
	// 	return
	// }

	// TODO: add page by get parameter
	page, _ := strconv.Atoi(r.PostFormValue("page"))
	letters, err := handlersDb.GetLimitedAmountOfLetters((page-1)*5, 5)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := ViewData{
		Letters: letters,
	}

	tmpl, err := template.ParseFiles("./tmp/html/getletters.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
