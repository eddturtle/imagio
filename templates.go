package main

import (
	"net/http"
	"text/template"
)

func getView(w http.ResponseWriter, name string, object interface{}) {
	t := template.Must(template.ParseFiles("templates/"+name+".html"))
	t.Execute(w, object)
}