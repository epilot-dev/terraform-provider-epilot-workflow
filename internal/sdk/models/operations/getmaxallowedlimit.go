// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/models/shared"
	"net/http"
)

type GetMaxAllowedLimitResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Other errors
	ErrorResp *shared.ErrorResp
	// A combo of current number of workflows, and the max allowed number of workflows.
	MaxAllowedLimit *shared.MaxAllowedLimit
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetMaxAllowedLimitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMaxAllowedLimitResponse) GetErrorResp() *shared.ErrorResp {
	if o == nil {
		return nil
	}
	return o.ErrorResp
}

func (o *GetMaxAllowedLimitResponse) GetMaxAllowedLimit() *shared.MaxAllowedLimit {
	if o == nil {
		return nil
	}
	return o.MaxAllowedLimit
}

func (o *GetMaxAllowedLimitResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMaxAllowedLimitResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
