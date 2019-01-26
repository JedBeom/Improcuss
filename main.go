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
	err := http.ListenAndServe(":8080", s.Router)
	if err != nil {
		log.Println("Server Error:", err)
	}
}
