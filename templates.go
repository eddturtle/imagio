package main

import (
	"os"
	"log"
	"net/http"
	"text/template"
)

func getView(w http.ResponseWriter, name string, object interface{}) {
	path := "templates/"+name+".html"
	if _, err := os.Stat(path); err == nil {
		t := template.Must(template.ParseFiles(path))
		t.Execute(w, object)
	} else {
		log.Fatal("Missing Template " , err)
	}
}