// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package errors

import (
	"fmt"
	"net/http"
)

type SDKError struct {
	Message     string
	StatusCode  int
	Body        string
	RawResponse *http.Response
}

var _ error = &SDKError{}

func NewSDKError(message string, statusCode int, body string, httpRes *http.Response) *SDKError {
	return &SDKError{
		Message:     message,
		StatusCode:  statusCode,
		Body:        body,
		RawResponse: httpRes,
	}
}

func (e *SDKError) Error() string {
	body := ""
	if len(e.Body) > 0 {
		body = fmt.Sprintf("\n%s", e.Body)
	}

	return fmt.Sprintf("%s: Status %d%s", e.Message, e.StatusCode, body)
}
