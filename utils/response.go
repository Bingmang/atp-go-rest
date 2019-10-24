package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int		`json:"status_code"`
	Message string		`json:"message"`
	Data    interface{}	`json:"data,omitempty"`
}

func ResponseWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}
