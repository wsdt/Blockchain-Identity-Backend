package main

import (
	. "VID-Card-Backend/routes/v1/_config"
	"VID-Card-Backend/routes/v1/get/showStatus"
	"VID-Card-Backend/routes/v1/post/sendVerificationMail"
	"github.com/gorilla/mux"
)

func defineRoutes(router *mux.Router)  {
	router.HandleFunc("/", showStatus.ShowStatus).Methods("GET")
	router.HandleFunc(ApiBaseRoute+"/verify_mail", sendVerificationMail.SendVerificationMail).Methods("POST")
	println("Routes defined.")
}
