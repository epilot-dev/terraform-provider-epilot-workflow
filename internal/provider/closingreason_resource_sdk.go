// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package provider

import (
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *ClosingReasonResourceModel) ToSharedClosingReason() *shared.ClosingReason {
	creationTime := new(string)
	if !r.CreationTime.IsUnknown() && !r.CreationTime.IsNull() {
		*creationTime = r.CreationTime.ValueString()
	} else {
		creationTime = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	lastUpdateTime := new(string)
	if !r.LastUpdateTime.IsUnknown() && !r.LastUpdateTime.IsNull() {
		*lastUpdateTime = r.LastUpdateTime.ValueString()
	} else {
		lastUpdateTime = nil
	}
	status := shared.ClosingReasonsStatus(r.Status.ValueString())
	title := r.Title.ValueString()
	out := shared.ClosingReason{
		CreationTime:   creationTime,
		ID:             id,
		LastUpdateTime: lastUpdateTime,
		Status:         status,
		Title:          title,
	}
	return &out
}

func (r *ClosingReasonResourceModel) RefreshFromSharedClosingReason(resp *shared.ClosingReason) {
	if resp != nil {
		r.CreationTime = types.StringPointerValue(resp.CreationTime)
		r.ID = types.StringPointerValue(resp.ID)
		r.LastUpdateTime = types.StringPointerValue(resp.LastUpdateTime)
		r.Status = types.StringValue(string(resp.Status))
		r.Title = types.StringValue(resp.Title)
	}
}
