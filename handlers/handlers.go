package handlers

import (
	"net/http"
	"text/template"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	}
}