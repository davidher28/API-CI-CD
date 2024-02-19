// +build integration

package hello_world

import (
	"net/http"
	"testing"
)

func TestHelloWorldIntegration(t *testing.T) {
	// Start the HTTP server
	go main()

	// Make a GET request to the /hello endpoint
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the status code is 200
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v want %v", resp.StatusCode, http.StatusOK)
	}

	// Check the response body
	expected := `{"message":"Hello, world!"}`
	actual := make([]byte, len(expected))
	_, err = resp.Body.Read(actual)
	if err != nil {
		t.Fatal(err)
	}
	if string(actual) != expected {
		t.Errorf("unexpected response body: got %v want %v", string(actual), expected)
	}
}
