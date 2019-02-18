package sendVerificationMail

import (
	"encoding/json"
	"github.com/VID-Card-Backend/_config"
	. "github.com/VID-Card-Backend/controllers/errorHandling"
	. "github.com/VID-Card-Backend/models/routeResponse"
	"io/ioutil"
	"net/http"
	"net/smtp"
)

type VerificationMailRes struct {
	EmailSent bool
}
type VerificationMailReq struct {
	Email string
}

func SendVerificationMail(w http.ResponseWriter, r *http.Request) {
	var t VerificationMailReq
	b,_ := ioutil.ReadAll(r.Body)
	LogErr(json.Unmarshal(b,&t))
	println("Received mail -> "+t.Email)

	//print("SendVerificationMail: Received mail -> "+reqParams.Email)

	respObj := VerificationMailRes{
		EmailSent: sendMail(t.Email),
	}

	SendResponse(respObj, w)
}

func sendMail(targetAddress string) bool {
	var config = _config.GetConfigMail()

	msg := "From: "+config.Address+"\n"+
		"To: "+targetAddress+"\n"+
		"Subject: VID-Card - Mail Verifikation\n\n"+
		"Hi!\nDanke, dass du VID-Card benutzt."

	println("CRED: "+config.Server+config.Port+" ,, "+config.Address+" // "+config.Pwd)
	// TODO: Create random link and send with mail. (when user clicks on it, then save into db)

	auth := smtp.PlainAuth("", config.Address, config.Pwd, config.Server)
	err := smtp.SendMail(config.Server+config.Port,auth, config.Address, []string{targetAddress}, []byte(msg))
	LogErr(err)
	return err == nil
}