package handlers

import (
	"net/http"

	"github.com/kasariks/send-letter-on-board/internal/db"
)

var handlersDb *db.DB

func InitHandlers(db *db.DB) {
	handlersDb = db
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./tmp/html/main.html")
}
