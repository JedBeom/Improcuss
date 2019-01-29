package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
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

	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusNotFound)

		pushStatic(w)

		ra := rand.New(rand.NewSource(time.Now().UnixNano()))
		x := ra.Intn(len(saying))

		data := struct {
			Path string
			Saying
		}{
			r.URL.Path,
			saying[x],
		}

		t := getTmpl("404")
		err := t.Execute(w, data)
		if err != nil {
			log.Println(err)
		}
	}
}
