package globalConfig

import (
	"github.com/VID-Card-Backend/controllers/errorHandling"
	"github.com/tkanos/gonfig"
)

const configDir = "../../_config"

// TODO: Change port to 465 for tls, currently 587 without (but works)
type ConfigMail struct {
	Address string
	Pwd string
	Port string
	Server string
}
var configMail ConfigMail = ConfigMail{}
func GetConfigMail() ConfigMail {
	if configMail == (ConfigMail{}) { // is configMail empty?
		errorHandling.LogErr(gonfig.GetConf(configDir+"/mail.json", &configMail))
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
		errorHandling.LogErr(gonfig.GetConf(configDir+"/api.json", &configApi))
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
		errorHandling.LogErr(gonfig.GetConf(configDir+"/db.json", &configDb))
		println("getConfigDb: Reparsed db.json.")
	}
	return configDb
}