// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type StepRequirement struct {
	Condition    types.String `tfsdk:"condition"`
	DefinitionID types.String `tfsdk:"definition_id"`
	Type         types.String `tfsdk:"type"`
}
