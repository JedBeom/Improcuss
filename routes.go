package main

import (
	"log"
	"net/http"
	"time"

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

		pushStatic(w)

		t := getTmplDefault("content")

		data := struct {
			Threads []Thread
		}{
			[]Thread{},
		}

		threads := []Thread{
			{
				Title:      "이렇게 긴 글을 쓰면 어떻게 되는 걸까?",
				CreatedAt:  time.Now(),
				ModifiedAt: time.Now(),
				NumUser:    14,
				NumRes:     109,
			},
			{
				Title:      "765 프로의 미래는, 이곳에 있어...!",
				ModifiedAt: time.Now(),
				NumUser:    1000,
				NumRes:     102101,
			},
		}

		data.Threads = threads

		err := executeContent(t, w, data)
		if err != nil {
			log.Println(err)
		}
	}

}

func pushStatic(w http.ResponseWriter) {

	if pusher, ok := w.(http.Pusher); ok {

		if err := pusher.Push("/static/custom.css", nil); err != nil {
			log.Println("Failed to push:", err)
		}

		if err := pusher.Push("/static/bootstrap.4.2.1.min.css", nil); err != nil {
			log.Println("Failed to push:", err)
		}

		if err := pusher.Push("/static/bootstrap.min.css.map", nil); err != nil {
			log.Println("Failed to push:", err)
		}
	} else {
		log.Print("Push is unsupported")
	}

}
