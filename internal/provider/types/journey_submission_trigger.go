// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type JourneySubmissionTrigger struct {
	AutomationID types.String `tfsdk:"automation_id"`
	ID           types.String `tfsdk:"id"`
	JourneyID    types.String `tfsdk:"journey_id"`
	Type         types.String `tfsdk:"type"`
}
