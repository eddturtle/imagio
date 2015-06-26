package main

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	p := Page{Title:"Home", Body:"Hello"}
	getView(w, "index", p)
}

