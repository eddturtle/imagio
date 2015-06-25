package main

import (
	"log"
	"net/http"

    "github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	p := Page{Title:"Home", Body:"Hello"}
	getView(w, "index", p)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}