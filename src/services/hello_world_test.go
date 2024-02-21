//go:build !integration

package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	// Arrange
	expected := "Hello world Claudia and David :)"

	// Act
	result := HelloWorld()

	// Assert
	assert.Equal(t, expected, result, "The two greetings would be the same.")
}
