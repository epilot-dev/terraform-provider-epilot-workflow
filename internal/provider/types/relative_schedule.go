// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type RelativeSchedule struct {
	Direction types.String `tfsdk:"direction"`
	Duration  types.Number `tfsdk:"duration"`
	Mode      types.String `tfsdk:"mode"`
	Reference Reference    `tfsdk:"reference"`
	Unit      types.String `tfsdk:"unit"`
}
