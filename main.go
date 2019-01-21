package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Server struct {
	db     *sql.DB
	router *router.Router
}

var s Server

func main() {

	initUsersPool()

	s.router = router.New()
	s.route()
	s.initDb()

	log.Println("Starting server.")
	err := fasthttp.ListenAndServe(":8080", s.router.Handler)
	if err != nil {
		log.Println("Server Error:", err)
	}
}
