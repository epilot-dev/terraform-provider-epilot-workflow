// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AutomationConfig struct {
	// Id of the configured automation to run
	FlowID string `json:"flowId"`
}

func (o *AutomationConfig) GetFlowID() string {
	if o == nil {
		return ""
	}
	return o.FlowID
}

// Step - Action that needs to be done in a Workflow
type Step struct {
	AssignedTo       []string          `json:"assignedTo,omitempty"`
	AutomationConfig *AutomationConfig `json:"automationConfig,omitempty"`
	// Longer information regarding Task
	Description *StepDescription `json:"description,omitempty"`
	DueDate     *string          `json:"dueDate,omitempty"`
	// set a Duedate for a step then a specific
	DynamicDueDate *DynamicDueDate `json:"dynamicDueDate,omitempty"`
	// Details regarding ECP for the workflow step
	Ecp           *ECPDetails `json:"ecp,omitempty"`
	ExecutionType *StepType   `json:"executionType,omitempty"`
	ID            *string     `json:"id,omitempty"`
	// Details regarding ECP for the workflow step
	Installer *ECPDetails  `json:"installer,omitempty"`
	Journey   *StepJourney `json:"journey,omitempty"`
	Name      string       `json:"name"`
	Order     float64      `json:"order"`
	// requirements that need to be fulfilled in order to enable the step execution
	Requirements []StepRequirement `json:"requirements,omitempty"`
	Type         ItemType          `json:"type"`
	// This field is deprecated. Please use assignedTo
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	UserIds []float64 `json:"userIds,omitempty"`
}

func (o *Step) GetAssignedTo() []string {
	if o == nil {
		return nil
	}
	return o.AssignedTo
}

func (o *Step) GetAutomationConfig() *AutomationConfig {
	if o == nil {
		return nil
	}
	return o.AutomationConfig
}

func (o *Step) GetDescription() *StepDescription {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Step) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *Step) GetDynamicDueDate() *DynamicDueDate {
	if o == nil {
		return nil
	}
	return o.DynamicDueDate
}

func (o *Step) GetEcp() *ECPDetails {
	if o == nil {
		return nil
	}
	return o.Ecp
}

func (o *Step) GetExecutionType() *StepType {
	if o == nil {
		return nil
	}
	return o.ExecutionType
}

func (o *Step) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Step) GetInstaller() *ECPDetails {
	if o == nil {
		return nil
	}
	return o.Installer
}

func (o *Step) GetJourney() *StepJourney {
	if o == nil {
		return nil
	}
	return o.Journey
}

func (o *Step) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Step) GetOrder() float64 {
	if o == nil {
		return 0.0
	}
	return o.Order
}

func (o *Step) GetRequirements() []StepRequirement {
	if o == nil {
		return nil
	}
	return o.Requirements
}

func (o *Step) GetType() ItemType {
	if o == nil {
		return ItemType("")
	}
	return o.Type
}

func (o *Step) GetUserIds() []float64 {
	if o == nil {
		return nil
	}
	return o.UserIds
}
