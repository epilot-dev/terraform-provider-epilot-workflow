// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/models/shared"
	"net/http"
)

type SetWorkflowClosingReasonsRequest struct {
	// set all closing reasons for a specific definition
	ClosingReasonsIds shared.ClosingReasonsIds `request:"mediaType=application/json"`
	// ID of a workflow definition
	DefinitionID string `pathParam:"style=simple,explode=false,name=definitionId"`
}

func (o *SetWorkflowClosingReasonsRequest) GetClosingReasonsIds() shared.ClosingReasonsIds {
	if o == nil {
		return shared.ClosingReasonsIds{}
	}
	return o.ClosingReasonsIds
}

func (o *SetWorkflowClosingReasonsRequest) GetDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.DefinitionID
}

type SetWorkflowClosingReasonsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SetWorkflowClosingReasonsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetWorkflowClosingReasonsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetWorkflowClosingReasonsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
