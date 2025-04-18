// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/internal/utils"
)

type ScheduleType string

const (
	ScheduleTypeDelayedSchedule  ScheduleType = "DelayedSchedule"
	ScheduleTypeRelativeSchedule ScheduleType = "RelativeSchedule"
)

type Schedule struct {
	DelayedSchedule  *DelayedSchedule  `queryParam:"inline"`
	RelativeSchedule *RelativeSchedule `queryParam:"inline"`

	Type ScheduleType
}

func CreateScheduleDelayedSchedule(delayedSchedule DelayedSchedule) Schedule {
	typ := ScheduleTypeDelayedSchedule

	return Schedule{
		DelayedSchedule: &delayedSchedule,
		Type:            typ,
	}
}

func CreateScheduleRelativeSchedule(relativeSchedule RelativeSchedule) Schedule {
	typ := ScheduleTypeRelativeSchedule

	return Schedule{
		RelativeSchedule: &relativeSchedule,
		Type:             typ,
	}
}

func (u *Schedule) UnmarshalJSON(data []byte) error {

	var delayedSchedule DelayedSchedule = DelayedSchedule{}
	if err := utils.UnmarshalJSON(data, &delayedSchedule, "", true, true); err == nil {
		u.DelayedSchedule = &delayedSchedule
		u.Type = ScheduleTypeDelayedSchedule
		return nil
	}

	var relativeSchedule RelativeSchedule = RelativeSchedule{}
	if err := utils.UnmarshalJSON(data, &relativeSchedule, "", true, true); err == nil {
		u.RelativeSchedule = &relativeSchedule
		u.Type = ScheduleTypeRelativeSchedule
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Schedule", string(data))
}

func (u Schedule) MarshalJSON() ([]byte, error) {
	if u.DelayedSchedule != nil {
		return utils.MarshalJSON(u.DelayedSchedule, "", true)
	}

	if u.RelativeSchedule != nil {
		return utils.MarshalJSON(u.RelativeSchedule, "", true)
	}

	return nil, errors.New("could not marshal union type Schedule: all fields are null")
}

type DecisionTask struct {
	AssignedTo []string    `json:"assigned_to,omitempty"`
	Conditions []Condition `json:"conditions"`
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
	Schedule     *Schedule           `json:"schedule,omitempty"`
	TaskType     TaskType            `json:"task_type"`
	// Taxonomy ids that are associated with this workflow and used for filtering
	Taxonomies []string `json:"taxonomies,omitempty"`
}

func (o *DecisionTask) GetAssignedTo() []string {
	if o == nil {
		return nil
	}
	return o.AssignedTo
}

func (o *DecisionTask) GetConditions() []Condition {
	if o == nil {
		return []Condition{}
	}
	return o.Conditions
}

func (o *DecisionTask) GetDescription() *StepDescription {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *DecisionTask) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *DecisionTask) GetDueDateConfig() *DueDateConfig {
	if o == nil {
		return nil
	}
	return o.DueDateConfig
}

func (o *DecisionTask) GetEcp() *ECPDetails {
	if o == nil {
		return nil
	}
	return o.Ecp
}

func (o *DecisionTask) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DecisionTask) GetInstaller() *ECPDetails {
	if o == nil {
		return nil
	}
	return o.Installer
}

func (o *DecisionTask) GetJourney() *StepJourney {
	if o == nil {
		return nil
	}
	return o.Journey
}

func (o *DecisionTask) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DecisionTask) GetPhaseID() *string {
	if o == nil {
		return nil
	}
	return o.PhaseID
}

func (o *DecisionTask) GetRequirements() []EnableRequirement {
	if o == nil {
		return nil
	}
	return o.Requirements
}

func (o *DecisionTask) GetSchedule() *Schedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *DecisionTask) GetTaskType() TaskType {
	if o == nil {
		return TaskType("")
	}
	return o.TaskType
}

func (o *DecisionTask) GetTaxonomies() []string {
	if o == nil {
		return nil
	}
	return o.Taxonomies
}
