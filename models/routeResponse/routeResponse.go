package routeResponse

import (
	"encoding/json"
	. "github.com/VID-Card-Backend/controllers/errorHandling"
	"github.com/VID-Card-Backend/controllers/globalConfig"
	"net/http"
	"time"
)

var config = globalConfig.GetConfigApi()

type RouteResponse struct {
	ApiVersion     string
	Value          json.RawMessage
	ReqSatisfiedOn time.Time
}


func CreateRouteResp(value json.RawMessage) RouteResponse {
	return RouteResponse{
		ApiVersion:     config.Version,
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
