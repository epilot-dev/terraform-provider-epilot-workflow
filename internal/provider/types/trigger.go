// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type Trigger struct {
	AutomationTrigger        *AutomationTrigger        `tfsdk:"automation_trigger" tfPlanOnly:"true"`
	JourneySubmissionTrigger *JourneySubmissionTrigger `tfsdk:"journey_submission_trigger" tfPlanOnly:"true"`
	ManualTrigger            *ManualTrigger            `tfsdk:"manual_trigger" tfPlanOnly:"true"`
}
