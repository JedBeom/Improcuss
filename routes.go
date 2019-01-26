package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Server) route() {
	s.Router = chi.NewRouter()

	s.Router.Use(middleware.Logger)

	s.Router.Get("/", mainHandler())

	s.Router.Get("/static/*", staticHandler())
	s.Router.NotFound(NotFound())
	// s.Router.Get("/:name", nameHandler())
}

func mainHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := getTmplDefault("content")
		err := executeContent(t, w, nil)
		if err != nil {
			log.Println(err)
		}
	}

}

func NotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, r.URL.String(), "is not supported.")
	}
}
