package main

import (
	"encoding/json"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Message string `json:"message"`
	}
	response := Response{
		Message: HelloWorld(),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/hello", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}
