package showStatus

import (
	. "VID-Card-Backend/models/routeResponse"
	"net/http"
)

type Status struct {
	IsOnline bool
}

func ShowStatus(w http.ResponseWriter, r *http.Request) {
	status := Status{
		IsOnline:true,
	}

	SendResponse(status, w)
}