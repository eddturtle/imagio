package main

import (
	"fmt"
	"log"
	"net/http"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}