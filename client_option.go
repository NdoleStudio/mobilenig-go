package mobilenig

import (
	"net/http"
	"net/url"
)

// ClientOption are options for constructing a client
type ClientOption interface {
	apply(config *clientConfig)
}

type clientOptionFunc func(config *clientConfig)

func (fn clientOptionFunc) apply(config *clientConfig) {
	fn(config)
}

// WithHTTPClient sets the underlying HTTP client used for requests.
// By default, http.DefaultClient is used.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		if httpClient != nil {
			config.httpClient = httpClient
		}
	})
}

// WithEnvironment sets the MobileNig endpoint for API requests
// By default, LiveEnvironment is used.
func WithEnvironment(environment Environment) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		if environment == LiveEnvironment || environment == TestEnvironment {
			config.environment = environment
		}
	})
}

// WithUsername sets the MobileNig API username
func WithUsername(username string) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		config.username = username
	})
}

// WithAPIKey sets the MobileNig API password
func WithAPIKey(apiKey string) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		config.apiKey = apiKey
	})
}

// WithBaseURL sets the MobileNig API base URL
func WithBaseURL(baseURL *url.URL) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		if baseURL != nil {
			config.baseURL = baseURL.String()
		}
	})
}
