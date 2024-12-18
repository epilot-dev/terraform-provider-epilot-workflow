// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/models/shared"
	"net/http"
)

type DeleteDefinitionRequest struct {
	// Id of the definition to de deleted.
	DefinitionID string `pathParam:"style=simple,explode=false,name=definitionId"`
}

func (o *DeleteDefinitionRequest) GetDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.DefinitionID
}

type DeleteDefinitionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Failed to authenticate
	ErrorResp *shared.ErrorResp
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteDefinitionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteDefinitionResponse) GetErrorResp() *shared.ErrorResp {
	if o == nil {
		return nil
	}
	return o.ErrorResp
}

func (o *DeleteDefinitionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteDefinitionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
