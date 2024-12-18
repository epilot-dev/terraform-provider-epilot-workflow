// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/models/shared"
	"net/http"
)

type ListFlowTemplatesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Other errors
	ErrorResp *shared.ErrorResp
	// Success - flow templates loaded with success. Empty array if customer has no flows defined.
	FlowTemplatesList *shared.FlowTemplatesList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListFlowTemplatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListFlowTemplatesResponse) GetErrorResp() *shared.ErrorResp {
	if o == nil {
		return nil
	}
	return o.ErrorResp
}

func (o *ListFlowTemplatesResponse) GetFlowTemplatesList() *shared.FlowTemplatesList {
	if o == nil {
		return nil
	}
	return o.FlowTemplatesList
}

func (o *ListFlowTemplatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListFlowTemplatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
