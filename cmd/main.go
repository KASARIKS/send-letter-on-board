package main

import (
	"log"
	"net/http"

	"github.com/kasariks/send-letter-on-board/internal/server"
)

func main() {
	logger := &log.Logger{}
	router, err := server.CreateServer(logger)
	if err != nil {
		log.Fatalf("Error with creating the server: %s\n", err.Error())
	}
	if err := http.ListenAndServe(router.Server.Addr, router.Server.Handler); err != nil {
		log.Fatalf("Error with starting the server: %s\n", err.Error())
	}
}
