package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
