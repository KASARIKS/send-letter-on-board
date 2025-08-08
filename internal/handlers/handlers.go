package handlers

import (
	"io"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Basic handler")
}
