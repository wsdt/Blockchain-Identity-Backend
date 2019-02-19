package main

import (
	"github.com/VID-Card-Backend/controllers/globalConfig"
	"github.com/VID-Card-Backend/routes/v1/get/showStatus"
	"github.com/VID-Card-Backend/routes/v1/post/oauthToken"
	"github.com/VID-Card-Backend/routes/v1/post/sendVerificationMail"
	"github.com/gorilla/mux"
)

var config = globalConfig.GetConfigApi()

func defineRoutes(router *mux.Router)  {
	router.HandleFunc("/", showStatus.ShowStatus).Methods("GET")
	router.HandleFunc(config.BaseRoute+"/verify_mail", sendVerificationMail.SendVerificationMail).Methods("POST")
	router.HandleFunc(config.BaseRoute+"/oauth_token", oauthToken.OAuthToken).Methods("POST")
	println("Routes defined.")
}
