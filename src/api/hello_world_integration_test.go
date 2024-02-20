// +build integration

package api

import (
	"net/http"
	"net/http/httptest"
	"src/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldIntegration(t *testing.T) {
	// Arrange
	expected := `{"message":"` + services.HelloWorld() + `"}`

	// Act
	request, _ := http.NewRequest("GET", "/hello", nil)
	requestRecord := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloWorldHandler)
	handler.ServeHTTP(requestRecord, request)

	// Assert
	assert.Equal(t, http.StatusOK, requestRecord.Code)
	assert.Equal(t, expected, requestRecord.Body.String())
}
