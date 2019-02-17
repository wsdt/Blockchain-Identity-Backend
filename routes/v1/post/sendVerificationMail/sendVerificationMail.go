package sendVerificationMail

import (
	. "VID-Card-Backend/controllers/errorHandling"
	. "VID-Card-Backend/models/routeResponse"
	"encoding/json"
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
	print("Received mail -> "+t.Email)



	//print("SendVerificationMail: Received mail -> "+reqParams.Email)

	respObj := VerificationMailRes{
		EmailSent: sendMail(t.Email),
	}

	SendResponse(respObj, w)
}

func sendMail(address string) bool {

	msg := "From: "+SourceMail+"\n"+
		"To: "+address+"\n"+
		"Subject: VID-Card - Mail Verifikation\n\n"+
		"Hi!\nDanke, dass du VID-Card benutzt."

	// TODO: Create random link and send with mail. (when user clicks on it, then save into db)

	return smtp.SendMail(MailServer,
		smtp.PlainAuth("", SourceMail, SourceMailPwd, MailServerHost),
		SourceMail, []string{address}, []byte(msg)) == nil
}