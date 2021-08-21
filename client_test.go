package mobilenig

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("default configuration is used when no option is set", func(t *testing.T) {
		// act
		client := New()

		// assert
		assert.NotEmpty(t, client.environment)
		assert.NotEmpty(t, client.common)

		assert.Empty(t, client.username)
		assert.Empty(t, client.apiKey)

		assert.NotNil(t, client.httpClient)
		assert.NotNil(t, client.Bills)
	})

	t.Run("single configuration value can be set using options", func(t *testing.T) {
		// Arrange
		env := TestEnvironment

		// Act
		client := New(WithEnvironment(env))

		// Assert
		assert.NotNil(t, client.environment)
		assert.Equal(t, env.String(), client.environment.String())
	})

	t.Run("multiple configuration values can be set using options", func(t *testing.T) {
		// Arrange
		env := TestEnvironment
		newHTTPClient := &http.Client{Timeout: 422}

		// Act
		client := New(WithEnvironment(env), WithHTTPClient(newHTTPClient))

		// Assert
		assert.NotEmpty(t, client.environment)
		assert.Equal(t, env.String(), client.environment.String())

		assert.NotNil(t, client.httpClient)
		assert.Equal(t, newHTTPClient.Timeout, client.httpClient.Timeout)
	})

	t.Run("it sets the Bills service correctly", func(t *testing.T) {
		// Arrange
		client := New()

		// Assert
		assert.NotNil(t, client.Bills)
		assert.Equal(t, client.environment.String(), client.Bills.client.environment.String())
	})
}
