package api

import (
	"fmt"
	"net/http"
	"src/services"
)

func HelloWorldHandler(w http.ResponseWriter, _ *http.Request) {
	response := Response{
		Message: services.HelloWorld(),
	}

	jsonData := EncodeJSONResponse(response)

	WriteJSONResponse(w, http.StatusOK, jsonData)

	fmt.Printf("hello world was executed\n")
}
