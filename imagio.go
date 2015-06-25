package main

import (
	// "fmt"
	// "os"
	"net/http"
)

func endpoint(w ) {

}

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Could not use port")
	}
}