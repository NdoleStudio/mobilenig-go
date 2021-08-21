package mobilenig

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	apiBaseURL = "https://mobilenig.com/API"
)

type service struct {
	client *Client
}

// Client is the MobileNig API client.
// Do not instantiate this client with Client{}. Use the New method instead.
type Client struct {
	httpClient  *http.Client
	common      service
	environment Environment
	username    string
	apiKey      string
	baseURL     string
	Bills       *BillsService
}

// New creates and returns a new mobilenig.Client from a slice of mobilenig.ClientOption.
func New(options ...ClientOption) *Client {
	config := defaultClientConfig()

	for _, option := range options {
		option.apply(config)
	}

	client := &Client{
		httpClient:  config.httpClient,
		environment: config.environment,
		username:    config.username,
		baseURL:     config.baseURL,
		apiKey:      config.apiKey,
	}

	client.common.client = client
	client.Bills = (*BillsService)(&client.common)
	return client
}

// newRequest creates an API request. A relative URL can be provided in uri,
// in which case it is resolved relative to the apiBaseURL of the Client.
// URI's should always be specified without a preceding slash.
func (client *Client) newRequest(ctx context.Context, uri string, params map[string]string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, client.baseURL+uri, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	q.Add("username", client.username)
	q.Add("api_key", client.apiKey)

	for key, value := range params {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}

// do carries out an HTTP request and returns a Response
func (client *Client) do(req *http.Request) (*Response, error) {
	httpResponse, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = httpResponse.Body.Close() }()

	resp, err := client.newResponse(httpResponse)
	if err != nil {
		return resp, err
	}

	_, err = io.Copy(ioutil.Discard, httpResponse.Body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// newResponse converts an *http.Response to *Response
func (client *Client) newResponse(httpResponse *http.Response) (*Response, error) {
	resp := new(Response)
	resp.HTTPResponse = httpResponse

	buf, err := ioutil.ReadAll(resp.HTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = &buf

	errResponse := new(ErrorResponse)
	err = json.Unmarshal(*resp.Body, errResponse)
	if err == nil {
		resp.Error = errResponse
	}

	return resp, resp.Err()
}
