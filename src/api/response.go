package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func EncodeJSONResponse(requestData interface{}) []byte {
	response, errorStatus := json.Marshal(requestData)
	if errorStatus != nil {
		return nil
	}
	return response
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, requestData []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, errorStatus := w.Write(requestData)
	if errorStatus != nil {
		http.Error(w, errorStatus.Error(), http.StatusInternalServerError)
	}
}
