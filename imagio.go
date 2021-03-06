package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
