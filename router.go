package main

import (
	"github.com/gorilla/mux"
	"VID-Card-Backend/routes/get/getTest"
)

func defineRoutes(router *mux.Router)  {
	router.HandleFunc("/", getTest.GetTest).Methods("GET")
	print("Routes defined.")
}
