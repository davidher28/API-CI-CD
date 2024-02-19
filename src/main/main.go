package main

import (
	"encoding/json"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Define the response struct
	type Response struct {
		Message string `json:"message"`
	}

	// Create the response
	response := Response{
		Message: "Hello, world!",
	}

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	// Define the route
	http.HandleFunc("/hello", helloWorldHandler)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
