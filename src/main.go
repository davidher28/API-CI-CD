package main

import (
	"fmt"
	"net/http"
	"src/api"
)

func main() {
	http.HandleFunc("/", api.HelloWorldHandler)

	errorStatus := http.ListenAndServe(":8080", nil)
	if errorStatus != nil {
		fmt.Printf("Error Starting HTTP Server: %s\n", errorStatus)
		return
	}
}
