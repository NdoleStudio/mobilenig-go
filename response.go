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
	if r.Error != nil && len(r.Error.Description) > 0 {
		return errors.New(r.errorMessage())
	}
	return nil
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
