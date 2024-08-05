// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Step struct {
	AssignedTo       []types.String    `tfsdk:"assigned_to"`
	AutomationConfig *AutomationConfig `tfsdk:"automation_config"`
	Description      *StepDescription  `tfsdk:"description"`
	DueDate          types.String      `tfsdk:"due_date"`
	DynamicDueDate   *DynamicDueDate   `tfsdk:"dynamic_due_date"`
	Ecp              *ECPDetails       `tfsdk:"ecp"`
	ExecutionType    types.String      `tfsdk:"execution_type"`
	ID               types.String      `tfsdk:"id"`
	Installer        *ECPDetails       `tfsdk:"installer"`
	Journey          *StepJourney      `tfsdk:"journey"`
	Name             types.String      `tfsdk:"name"`
	Order            types.Number      `tfsdk:"order"`
	Requirements     []StepRequirement `tfsdk:"requirements"`
	Type             types.String      `tfsdk:"type"`
	UserIds          []types.Number    `tfsdk:"user_ids"`
}
