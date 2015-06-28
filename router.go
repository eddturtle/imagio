package main

import (
	"net/http"

    "github.com/gorilla/mux"
)

const (
	StaticDIR = "/static/"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Server CSS, JS & Images Statically.
    router.
    	PathPrefix(StaticDIR).
    	Handler(http.StripPrefix(StaticDIR, http.FileServer(http.Dir("."+StaticDIR))))

	return router
}