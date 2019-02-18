package _config

import (
	"github.com/VID-Card-Backend/controllers/errorHandling"
	"github.com/tkanos/gonfig"
)

type ConfigMail struct {
	Address string
	Pwd string
	Port string
	Server string
}
var configMail ConfigMail = ConfigMail{}
func GetConfigMail() ConfigMail {
	if configMail == (ConfigMail{}) { // is configMail empty?
		errorHandling.LogErr(gonfig.GetConf("./_config/mail.json", &configMail))
		println("getConfigMail: Reparsed mail.json.")
	}
	return configMail
}

type ConfigApi struct {
	Version string
	Port string
	BaseRoute string
}
var configApi ConfigApi = ConfigApi{}
func GetConfigApi() ConfigApi {
	if configApi == (ConfigApi{}) { // is configMail empty?
		errorHandling.LogErr(gonfig.GetConf("./_config/api.json", &configApi))
		println("getConfigApi: Reparsed api.json.")
	}
	return configApi
}

type ConfigDb struct {
	User string
	Pwd string
	Name string
	Host string
	Port string
	ConnStr string
}
var configDb ConfigDb = ConfigDb{}
func GetConfigDb() ConfigDb {
	if configDb == (ConfigDb{}) { // is configDb empty?
		errorHandling.LogErr(gonfig.GetConf("./_config/db.json", &configDb))
		println("getConfigDb: Reparsed db.json.")
	}
	return configDb
}