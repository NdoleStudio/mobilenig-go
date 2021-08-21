package mobilenig

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithHTTPClient(t *testing.T) {
	t.Run("httpClient is not set when the httpClient is nil", func(t *testing.T) {
		// Arrange
		config := defaultClientConfig()

		// Act
		WithHTTPClient(nil).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
	})

	t.Run("httpClient is set when the httpClient is not nil", func(t *testing.T) {
		// Arrange
		config := defaultClientConfig()
		newClient := &http.Client{Timeout: 300}

		// Act
		WithHTTPClient(newClient).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
		assert.Equal(t, newClient.Timeout, config.httpClient.Timeout)
	})
}

func TestWithEnvironment(t *testing.T) {
	t.Run("environment is set successfully", func(t *testing.T) {
		// Arrange
		environment := TestEnvironment
		config := defaultClientConfig()

		// Act
		WithEnvironment(environment).apply(config)

		// Assert
		assert.NotNil(t, config.environment)
		assert.Equal(t, environment.String(), config.environment.String())
	})

	t.Run("environment is not set if it's not equal to LIVE or TEST", func(t *testing.T) {
		// Arrange
		var environment Environment = "DEV"
		config := defaultClientConfig()

		// Act
		WithEnvironment(environment).apply(config)

		// Assert
		assert.NotNil(t, config.environment)
		assert.Equal(t, defaultClientConfig().environment.String(), config.environment.String())
	})
}

func TestWithUsername(t *testing.T) {
	t.Run("username is set successfully", func(t *testing.T) {
		// Arrange
		config := defaultClientConfig()
		username := "username"

		// Act
		WithUsername(username).apply(config)

		// Assert
		assert.Equal(t, username, config.username)
	})
}

func TestWithAPIKey(t *testing.T) {
	t.Run("apiKey is set successfully", func(t *testing.T) {
		// Arrange
		config := defaultClientConfig()
		apiKey := "apiKey"

		// Act
		WithAPIKey(apiKey).apply(config)

		// Assert
		assert.Equal(t, apiKey, config.apiKey)
	})
}

func TestWithBaseURL(t *testing.T) {
	t.Run("it sets the baseURL successfully", func(t *testing.T) {
		// Arrange
		config := defaultClientConfig()
		endpoint, _ := url.Parse("https://example.com")

		// Act
		WithBaseURL(endpoint).apply(config)

		// Assert
		assert.Equal(t, endpoint.String(), config.baseURL)
	})

	t.Run("it does not set the baseURL when it is nil", func(t *testing.T) {
		// Arrange
		config := defaultClientConfig()

		// Act
		WithBaseURL(nil).apply(config)

		// Assert
		assert.Equal(t, baseURL, config.baseURL)
	})
}
