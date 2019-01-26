package main

import (
	"net/http"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	return w.ResponseWriter.Write(b)
}

func staticHandler() http.HandlerFunc {

	fileServer := http.FileServer(http.Dir("static"))
	return func(w http.ResponseWriter, r *http.Request) {

		sW := sWPool.Get().(*statusWriter)
		sW.ResponseWriter = w
		http.StripPrefix("/static/", fileServer).ServeHTTP(sW, r)

		if sW.status != 200 {
			NotFound().ServeHTTP(w, r)
		}

		sW.status = 0
		sWPool.Put(sW)
	}

}
