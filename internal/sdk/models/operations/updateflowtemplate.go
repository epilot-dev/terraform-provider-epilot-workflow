// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/models/shared"
	"net/http"
)

type UpdateFlowTemplateRequest struct {
	// Flow Template payload
	FlowTemplate shared.FlowTemplateInput `request:"mediaType=application/json"`
	FlowID       string                   `pathParam:"style=simple,explode=false,name=flowId"`
}

func (o *UpdateFlowTemplateRequest) GetFlowTemplate() shared.FlowTemplateInput {
	if o == nil {
		return shared.FlowTemplateInput{}
	}
	return o.FlowTemplate
}

func (o *UpdateFlowTemplateRequest) GetFlowID() string {
	if o == nil {
		return ""
	}
	return o.FlowID
}

type UpdateFlowTemplateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Validation Errors
	ErrorResp *shared.ErrorResp
	// Flow template has been updated successfully
	FlowTemplate *shared.FlowTemplate
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateFlowTemplateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateFlowTemplateResponse) GetErrorResp() *shared.ErrorResp {
	if o == nil {
		return nil
	}
	return o.ErrorResp
}

func (o *UpdateFlowTemplateResponse) GetFlowTemplate() *shared.FlowTemplate {
	if o == nil {
		return nil
	}
	return o.FlowTemplate
}

func (o *UpdateFlowTemplateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateFlowTemplateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
