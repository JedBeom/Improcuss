package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	saying []Saying
)

func init() {
	loadSaying()
}

type Saying struct {
	Sentence string `json:"sentence"`
	Author   string `json:"author"`
}

func loadSaying() {
	file, err := ioutil.ReadFile("./static/saying.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &saying)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func NotFound() http.HandlerFunc {

	var (
		once sync.Once
		t    *template.Template
	)

	return func(w http.ResponseWriter, r *http.Request) {

		once.Do(func() {
			t = getTmpl("404")
		})

		w.WriteHeader(http.StatusNotFound)

		// pushStatic(w)

		ra := rand.New(rand.NewSource(time.Now().UnixNano()))
		x := ra.Intn(len(saying))

		data := struct {
			Path string
			Saying
		}{
			r.URL.Path,
			saying[x],
		}

		err := t.Execute(w, data)
		if err != nil {
			log.Println(err)
		}
	}
}
