// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *ClosingReasonDataSourceModel) RefreshFromSharedClosingReason(resp *shared.ClosingReason) {
	if resp != nil {
		r.CreationTime = types.StringPointerValue(resp.CreationTime)
		r.ID = types.StringPointerValue(resp.ID)
		r.LastUpdateTime = types.StringPointerValue(resp.LastUpdateTime)
		r.Status = types.StringValue(string(resp.Status))
		r.Title = types.StringValue(resp.Title)
	}
}
