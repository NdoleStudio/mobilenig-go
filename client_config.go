package mobilenig

import (
	"net/http"
)

type clientConfig struct {
	httpClient  *http.Client
	environment Environment
	baseURL     string
	apiKey      string
	username    string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient:  http.DefaultClient,
		apiKey:      "",
		username:    "",
		baseURL:     apiBaseURL,
		environment: LiveEnvironment,
	}
}
