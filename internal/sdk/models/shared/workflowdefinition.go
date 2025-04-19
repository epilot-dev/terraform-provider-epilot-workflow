// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/internal/utils"
)

type FlowType string

const (
	FlowTypeSection FlowType = "Section"
	FlowTypeStep    FlowType = "Step"
)

type Flow struct {
	Section *Section `queryParam:"inline"`
	Step    *Step    `queryParam:"inline"`

	Type FlowType
}

func CreateFlowSection(section Section) Flow {
	typ := FlowTypeSection

	return Flow{
		Section: &section,
		Type:    typ,
	}
}

func CreateFlowStep(step Step) Flow {
	typ := FlowTypeStep

	return Flow{
		Step: &step,
		Type: typ,
	}
}

func (u *Flow) UnmarshalJSON(data []byte) error {

	var section Section = Section{}
	if err := utils.UnmarshalJSON(data, &section, "", true, true); err == nil {
		u.Section = &section
		u.Type = FlowTypeSection
		return nil
	}

	var step Step = Step{}
	if err := utils.UnmarshalJSON(data, &step, "", true, true); err == nil {
		u.Step = &step
		u.Type = FlowTypeStep
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Flow", string(data))
}

func (u Flow) MarshalJSON() ([]byte, error) {
	if u.Section != nil {
		return utils.MarshalJSON(u.Section, "", true)
	}

	if u.Step != nil {
		return utils.MarshalJSON(u.Step, "", true)
	}

	return nil, errors.New("could not marshal union type Flow: all fields are null")
}

type WorkflowDefinition struct {
	AssignedTo     []string          `json:"assignedTo,omitempty"`
	ClosingReasons []ClosingReasonID `json:"closingReasons,omitempty"`
	// ISO String Date & Time
	CreationTime *string `json:"creationTime,omitempty"`
	Description  *string `json:"description,omitempty"`
	DueDate      *string `json:"dueDate,omitempty"`
	// set a Duedate for a step then a specific
	DynamicDueDate *DynamicDueDate `json:"dynamicDueDate,omitempty"`
	// Indicates whether this workflow is available for End Customer Portal or not. By default it's not.
	EnableECPWorkflow *bool   `json:"enableECPWorkflow,omitempty"`
	Flow              []Flow  `json:"flow"`
	ID                *string `json:"id,omitempty"`
	// ISO String Date & Time
	LastUpdateTime         *string                  `json:"lastUpdateTime,omitempty"`
	Name                   string                   `json:"name"`
	UpdateEntityAttributes []UpdateEntityAttributes `json:"updateEntityAttributes,omitempty"`
	// This field is deprecated. Please use assignedTo
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	UserIds []float64 `json:"userIds,omitempty"`
}

func (o *WorkflowDefinition) GetAssignedTo() []string {
	if o == nil {
		return nil
	}
	return o.AssignedTo
}

func (o *WorkflowDefinition) GetClosingReasons() []ClosingReasonID {
	if o == nil {
		return nil
	}
	return o.ClosingReasons
}

func (o *WorkflowDefinition) GetCreationTime() *string {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *WorkflowDefinition) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *WorkflowDefinition) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *WorkflowDefinition) GetDynamicDueDate() *DynamicDueDate {
	if o == nil {
		return nil
	}
	return o.DynamicDueDate
}

func (o *WorkflowDefinition) GetEnableECPWorkflow() *bool {
	if o == nil {
		return nil
	}
	return o.EnableECPWorkflow
}

func (o *WorkflowDefinition) GetFlow() []Flow {
	if o == nil {
		return []Flow{}
	}
	return o.Flow
}

func (o *WorkflowDefinition) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WorkflowDefinition) GetLastUpdateTime() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdateTime
}

func (o *WorkflowDefinition) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *WorkflowDefinition) GetUpdateEntityAttributes() []UpdateEntityAttributes {
	if o == nil {
		return nil
	}
	return o.UpdateEntityAttributes
}

func (o *WorkflowDefinition) GetUserIds() []float64 {
	if o == nil {
		return nil
	}
	return o.UserIds
}
