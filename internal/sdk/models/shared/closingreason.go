// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

// ClosingReason - One Closing reason for a workflow
type ClosingReason struct {
	CreationTime   *string              `json:"creationTime,omitempty"`
	ID             *string              `json:"id,omitempty"`
	LastUpdateTime *string              `json:"lastUpdateTime,omitempty"`
	Status         ClosingReasonsStatus `json:"status"`
	Title          string               `json:"title"`
}

func (o *ClosingReason) GetCreationTime() *string {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *ClosingReason) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ClosingReason) GetLastUpdateTime() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdateTime
}

func (o *ClosingReason) GetStatus() ClosingReasonsStatus {
	if o == nil {
		return ClosingReasonsStatus("")
	}
	return o.Status
}

func (o *ClosingReason) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}
