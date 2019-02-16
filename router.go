package main

import (
	"VID-Card-Backend/routes/get/getTest"
	"github.com/gorilla/mux"
)

func defineRoutes(router *mux.Router)  {
	router.HandleFunc("/", getTest.GetTest).Methods("GET")
	print("Routes defined.")
}
