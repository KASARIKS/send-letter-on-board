package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/kasariks/send-letter-on-board/internal/db/dbEntities/dbletter"
)

// For templates
type viewData struct {
	Letters []*dbletter.DbLetter
}

func GetAllLetters(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	pageNumber := params.Get("page")
	lettersQuantity := params.Get("letters")

	currPageData, err := newPageData(pageNumber, lettersQuantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	letters, err := handlersDb.GetLimitedNumberOfLetters((currPageData.pageNumber-1)*currPageData.lettersQuantity,
		currPageData.lettersQuantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := viewData{
		Letters: letters,
	}

	tmpl, err := template.ParseFiles("./tmp/html/getletters.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

type pageData struct {
	pageNumber      int
	lettersQuantity int
}

func newPageData(pageNumber, lettersQuantity string) (*pageData, error) {
	tmpPageNumber, err := strconv.Atoi(pageNumber)
	if err != nil {
		return nil, err
	}

	tmpLettersQuantity, err := strconv.Atoi(lettersQuantity)
	if err != nil {
		return nil, err
	}

	if tmpLettersQuantity > 10 {
		tmpLettersQuantity = 10
	}

	return &pageData{
		pageNumber:      tmpPageNumber,
		lettersQuantity: tmpLettersQuantity,
	}, nil
}
