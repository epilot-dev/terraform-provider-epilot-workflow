// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AutomationTask struct {
	AssignedTo []string `json:"assigned_to,omitempty"`
	// Configuration for automation execution to run
	AutomationConfig AutomationConfig `json:"automation_config"`
	// Longer information regarding Task
	Description *StepDescription `json:"description,omitempty"`
	DueDate     *string          `json:"due_date,omitempty"`
	// Set due date for the task based on a dynamic condition
	DueDateConfig *DueDateConfig `json:"due_date_config,omitempty"`
	// Details regarding ECP for the workflow step
	Ecp *ECPDetails `json:"ecp,omitempty"`
	ID  string      `json:"id"`
	// Details regarding ECP for the workflow step
	Installer *ECPDetails  `json:"installer,omitempty"`
	Journey   *StepJourney `json:"journey,omitempty"`
	Name      string       `json:"name"`
	PhaseID   *string      `json:"phase_id,omitempty"`
	// requirements that need to be fulfilled in order to enable the task while flow instances are running
	Requirements []EnableRequirement `json:"requirements,omitempty"`
	Schedule     *ActionSchedule     `json:"schedule,omitempty"`
	TaskType     TaskType            `json:"task_type"`
	// Taxonomy ids that are associated with this workflow and used for filtering
	Taxonomies  []string     `json:"taxonomies,omitempty"`
	TriggerMode *TriggerMode `json:"trigger_mode,omitempty"`
}

func (o *AutomationTask) GetAssignedTo() []string {
	if o == nil {
		return nil
	}
	return o.AssignedTo
}

func (o *AutomationTask) GetAutomationConfig() AutomationConfig {
	if o == nil {
		return AutomationConfig{}
	}
	return o.AutomationConfig
}

func (o *AutomationTask) GetDescription() *StepDescription {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AutomationTask) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *AutomationTask) GetDueDateConfig() *DueDateConfig {
	if o == nil {
		return nil
	}
	return o.DueDateConfig
}

func (o *AutomationTask) GetEcp() *ECPDetails {
	if o == nil {
		return nil
	}
	return o.Ecp
}

func (o *AutomationTask) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AutomationTask) GetInstaller() *ECPDetails {
	if o == nil {
		return nil
	}
	return o.Installer
}

func (o *AutomationTask) GetJourney() *StepJourney {
	if o == nil {
		return nil
	}
	return o.Journey
}

func (o *AutomationTask) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AutomationTask) GetPhaseID() *string {
	if o == nil {
		return nil
	}
	return o.PhaseID
}

func (o *AutomationTask) GetRequirements() []EnableRequirement {
	if o == nil {
		return nil
	}
	return o.Requirements
}

func (o *AutomationTask) GetSchedule() *ActionSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *AutomationTask) GetTaskType() TaskType {
	if o == nil {
		return TaskType("")
	}
	return o.TaskType
}

func (o *AutomationTask) GetTaxonomies() []string {
	if o == nil {
		return nil
	}
	return o.Taxonomies
}

func (o *AutomationTask) GetTriggerMode() *TriggerMode {
	if o == nil {
		return nil
	}
	return o.TriggerMode
}
