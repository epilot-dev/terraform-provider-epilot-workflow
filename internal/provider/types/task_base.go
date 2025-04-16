// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type TaskBase struct {
	AssignedTo    []types.String      `tfsdk:"assigned_to"`
	Description   *StepDescription    `tfsdk:"description"`
	DueDate       types.String        `tfsdk:"due_date"`
	DueDateConfig *DueDateConfig      `tfsdk:"due_date_config"`
	Ecp           *ECPDetails         `tfsdk:"ecp"`
	ID            types.String        `tfsdk:"id"`
	Installer     *ECPDetails         `tfsdk:"installer"`
	Journey       *StepJourney        `tfsdk:"journey"`
	Name          types.String        `tfsdk:"name"`
	PhaseID       types.String        `tfsdk:"phase_id"`
	Requirements  []EnableRequirement `tfsdk:"requirements"`
	TaskType      types.String        `tfsdk:"task_type"`
	Taxonomies    []types.String      `tfsdk:"taxonomies"`
}
