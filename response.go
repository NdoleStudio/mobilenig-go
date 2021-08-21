package mobilenig

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"
)

// ErrorResponse is the response that is returned when there is an API error
type ErrorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Response captures the http response
type Response struct {
	HTTPResponse *http.Response
	Body         *[]byte
	Error        *ErrorResponse
}

// Err returns an error if the http request is not successfull
func (r *Response) Err() error {
	switch r.HTTPResponse.StatusCode {
	case 200, 201, 202, 204, 205:
		return nil
	default:
		return errors.New(r.errorMessage())
	}
}

func (r *Response) errorMessage() string {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(r.HTTPResponse.StatusCode))
	buf.WriteString(": ")
	buf.WriteString(http.StatusText(r.HTTPResponse.StatusCode))
	buf.WriteString(", Body: ")
	buf.Write(*r.Body)

	return buf.String()
}
