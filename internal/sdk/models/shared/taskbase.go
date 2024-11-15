// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type TaskBase struct {
	AssignedTo []string `json:"assigned_to,omitempty"`
	// Longer information regarding Task
	Description *StepDescription `json:"description,omitempty"`
	DueDate     *string          `json:"due_date,omitempty"`
	// Set due date for the task based on a dynamic condition
	DynamicDueDate *DueDateConfig `json:"dynamic_due_date,omitempty"`
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
	TaskType     TaskType            `json:"task_type"`
	// Taxonomy ids that are associated with this workflow and used for filtering
	Taxonomies []string `json:"taxonomies,omitempty"`
}

func (o *TaskBase) GetAssignedTo() []string {
	if o == nil {
		return nil
	}
	return o.AssignedTo
}

func (o *TaskBase) GetDescription() *StepDescription {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TaskBase) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *TaskBase) GetDynamicDueDate() *DueDateConfig {
	if o == nil {
		return nil
	}
	return o.DynamicDueDate
}

func (o *TaskBase) GetEcp() *ECPDetails {
	if o == nil {
		return nil
	}
	return o.Ecp
}

func (o *TaskBase) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TaskBase) GetInstaller() *ECPDetails {
	if o == nil {
		return nil
	}
	return o.Installer
}

func (o *TaskBase) GetJourney() *StepJourney {
	if o == nil {
		return nil
	}
	return o.Journey
}

func (o *TaskBase) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *TaskBase) GetPhaseID() *string {
	if o == nil {
		return nil
	}
	return o.PhaseID
}

func (o *TaskBase) GetRequirements() []EnableRequirement {
	if o == nil {
		return nil
	}
	return o.Requirements
}

func (o *TaskBase) GetTaskType() TaskType {
	if o == nil {
		return TaskType("")
	}
	return o.TaskType
}

func (o *TaskBase) GetTaxonomies() []string {
	if o == nil {
		return nil
	}
	return o.Taxonomies
}
