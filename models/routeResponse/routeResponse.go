package routeResponse

import (
	. "VID-Card-Backend/controllers/errorHandling"
	. "VID-Card-Backend/routes/v1/_config"
	"encoding/json"
	"net/http"
	"time"
)

type RouteResponse struct {
	ApiVersion     string
	Value          json.RawMessage
	ReqSatisfiedOn time.Time
}


func CreateRouteResp(value json.RawMessage) RouteResponse {
	return RouteResponse{
		ApiVersion:     ApiVersion,
		ReqSatisfiedOn: time.Now(),
		Value:          value,
	}
}

func SendResponse(value interface{}, w http.ResponseWriter) {
	e, err := json.Marshal(value)
	println("SendResponse: Responding with following value -> "+string(e))
	LogErr(err)
	LogErr(json.NewEncoder(w).Encode(CreateRouteResp(e)))
}
