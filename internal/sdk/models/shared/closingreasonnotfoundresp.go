// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

// ClosingReasonNotFoundResp - Closing reason could be not found
type ClosingReasonNotFoundResp struct {
	Message *string `json:"message,omitempty"`
}

func (o *ClosingReasonNotFoundResp) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}
