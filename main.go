package main

import (
	. "VID-Card-Backend/controllers/errorHandling"
	. "VID-Card-Backend/routes/v1/_config"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	defineRoutes(router)
	LogErr(http.ListenAndServe(ApiPort, router))
}
