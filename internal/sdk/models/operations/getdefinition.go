// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/models/shared"
	"net/http"
)

type GetDefinitionRequest struct {
	// Short uuid (length 8) to identify the Workflow Definition.
	DefinitionID string `pathParam:"style=simple,explode=false,name=definitionId"`
}

func (o *GetDefinitionRequest) GetDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.DefinitionID
}

type GetDefinitionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Definition Not found
	DefinitionNotFoundResp *shared.DefinitionNotFoundResp
	// Validation Errors
	ErrorResp *shared.ErrorResp
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns the Workflow definition
	WorkflowDefinition *shared.WorkflowDefinition
}

func (o *GetDefinitionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDefinitionResponse) GetDefinitionNotFoundResp() *shared.DefinitionNotFoundResp {
	if o == nil {
		return nil
	}
	return o.DefinitionNotFoundResp
}

func (o *GetDefinitionResponse) GetErrorResp() *shared.ErrorResp {
	if o == nil {
		return nil
	}
	return o.ErrorResp
}

func (o *GetDefinitionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDefinitionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDefinitionResponse) GetWorkflowDefinition() *shared.WorkflowDefinition {
	if o == nil {
		return nil
	}
	return o.WorkflowDefinition
}
