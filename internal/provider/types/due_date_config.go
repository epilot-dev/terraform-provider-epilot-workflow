// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DueDateConfig struct {
	Duration types.Number `tfsdk:"duration"`
	PhaseID  types.String `tfsdk:"phase_id"`
	TaskID   types.String `tfsdk:"task_id"`
	Type     types.String `tfsdk:"type"`
	Unit     types.String `tfsdk:"unit"`
}
