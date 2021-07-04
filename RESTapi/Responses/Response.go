package Responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Simple response function.
func BasicRes(res http.ResponseWriter, statusCode int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	err := json.NewEncoder(res).Encode(data)
	if err != nil {
		fmt.Fprintf(res, "%s", err.Error())
	}
}

// Simple error response function.
func ErrRes(res http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		BasicRes(res, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	BasicRes(res, http.StatusBadRequest, nil)
}
