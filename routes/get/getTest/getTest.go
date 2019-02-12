package getTest

import (
	"net/http"
	"encoding/json"
)

func GetTest(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{reza:test}")
}