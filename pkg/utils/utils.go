package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	err := json.NewDecoder(r.Body).Decode(x)
	if err != nil {
		log.Print("There's a error", err)
	}
}
