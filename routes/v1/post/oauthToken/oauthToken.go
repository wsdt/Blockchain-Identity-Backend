package oauthToken

import (
	"encoding/json"
	"github.com/VID-Card-Backend/controllers/errorHandling"
	"github.com/VID-Card-Backend/models/routeResponse"
	"io/ioutil"
	"net/http"
)

type VerificationOAuthReq struct {
	UserId string
	Type string
	Token string
}

type VerificationOAuthRes struct {
	isTokenValid bool
}

func OAuthToken(w http.ResponseWriter, r *http.Request) {
	var t VerificationOAuthReq
	b,_ := ioutil.ReadAll(r.Body)
	errorHandling.LogErr(json.Unmarshal(b,&t))

	var isValid = verifyToken(t.Token)
	respObj := VerificationOAuthRes{
		isTokenValid:isValid,
	}
	if isValid {
		// TODO: Hash, encrypt save into db, decode json web token
		// currently unencrypted
		// return successful response http
		println("OAuthToken: Token valid. Saving into db.")
	}

	// Send response to frontend
	routeResponse.SendResponse(respObj, w)
}

func verifyToken(token string) bool {
	// TODO
	return true
}
