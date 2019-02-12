package main

import (
	"github.com/gorilla/mux"
	"github/wsdt/VID-Card-Backend/routes/get/getTest"
)

func defineRoutes(router *mux.Router)  {
	router.HandleFunc("/", getTest.GetTest).Methods("GET")
	print("Routes defined.")
}
