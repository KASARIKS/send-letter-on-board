package server

import (
	"log"
	"net/http"
	"time"

	"github.com/kasariks/send-letter-on-board/internal/db"
	"github.com/kasariks/send-letter-on-board/internal/handlers"
)

type routerData struct {
	Logger *log.Logger
	Server *http.Server
}

func CreateServer(logger *log.Logger) (*routerData, error) {
	database, err := db.NewDb()
	if err != nil {
		return nil, err
	}
	handlers.InitHandlers(database)

	registerHandlers()
	router := newRouterData(logger)
	return router, nil
}

func registerHandlers() {
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/registrationpage", handlers.RegistrationPage)
	http.HandleFunc("/registration", handlers.Registration)
	http.HandleFunc("/authorizationpage", handlers.AuthorizationPage)
	http.HandleFunc("/authorization", handlers.Authorization)
}

func newRouterData(logger *log.Logger) *routerData {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 5,
	}

	router := &routerData{
		Logger: logger,
		Server: server,
	}

	return router
}
