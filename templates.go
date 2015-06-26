package main

import (
	"log"
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func getView(w http.ResponseWriter, name string, object interface{}) {
	err := templates.ExecuteTemplate(w, name, object)
	if err != nil {
		log.Fatal("Cannot Get View ", err)
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}