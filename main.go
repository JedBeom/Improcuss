package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	_ "github.com/lib/pq"
)

type Server struct {
	DB     *sql.DB
	Router chi.Router
}

var s Server

func main() {

	s.route()
	s.initDb()

	log.Println("Starting server.")

	var err error
	if config.Server.SSL {
		err = http.ListenAndServeTLS(config.Server.Port, config.Server.Cert, config.Server.Key, s.Router)
	} else {
		err = http.ListenAndServe(config.Server.Port, s.Router)
	}

	if err != nil {
		log.Println("Server Error:", err)
	}
}
