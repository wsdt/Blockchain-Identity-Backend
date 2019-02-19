package main

import (
	. "github.com/VID-Card-Backend/controllers/errorHandling"
	"github.com/VID-Card-Backend/controllers/globalConfig"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	defineRoutes(router)
	LogErr(http.ListenAndServe(globalConfig.GetConfigApi().Port, router))
}
