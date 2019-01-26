package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func getTmpl(names ...string) (t *template.Template) {

	var files []string

	for _, file := range names {
		files = append(files, fmt.Sprintf("tmpl/%s.tmpl", file))
	}
	t = template.Must(template.ParseFiles(files...))
	return
}

func getTmplDefault(names ...string) (t *template.Template) {
	names = append(names, "base")
	return getTmpl(names...)
}

func executeContent(t *template.Template, w http.ResponseWriter, data interface{}) (err error) {
	return t.ExecuteTemplate(w, "base", data)
}
