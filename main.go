package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	defineRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
