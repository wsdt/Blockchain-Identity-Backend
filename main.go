package main

import (
	"github.com/VID-Card-Backend/_config"
	. "github.com/VID-Card-Backend/controllers/errorHandling"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	defineRoutes(router)
	LogErr(http.ListenAndServe(_config.GetConfigApi().Port, router))
}
