// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/internal/utils"
)

type FlowTemplateInput struct {
	AssignedTo []string `json:"assigned_to,omitempty"`
	// Indicates whether this workflow is available for End Customer Portal or not. By default it's not.
	AvailableInEcp *bool                `json:"available_in_ecp,omitempty"`
	ClosingReasons []ClosingReasonInput `json:"closing_reasons,omitempty"`
	// ISO String Date & Time
	CreatedAt   *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	DueDate     *string `json:"due_date,omitempty"`
	// Set due date for the task based on a dynamic condition
	DueDateConfig *DueDateConfig `json:"due_date_config,omitempty"`
	Edges         []Edge         `json:"edges"`
	// Whether the workflow is enabled or not
	Enabled *bool   `default:"true" json:"enabled"`
	ID      *string `json:"id,omitempty"`
	// Whether the workflow is migrated from workflows to flows or not
	IsFlowMigrated *bool   `default:"false" json:"is_flow_migrated"`
	Name           string  `json:"name"`
	OrgID          *string `json:"org_id,omitempty"`
	Phases         []Phase `json:"phases,omitempty"`
	Tasks          []Task  `json:"tasks"`
	// Taxonomy ids that are associated with this workflow and used for filtering
	Taxonomies             []string                 `json:"taxonomies,omitempty"`
	Trigger                *Trigger                 `json:"trigger,omitempty"`
	UpdateEntityAttributes []UpdateEntityAttributes `json:"update_entity_attributes,omitempty"`
	// ISO String Date & Time
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (f FlowTemplateInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FlowTemplateInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FlowTemplateInput) GetAssignedTo() []string {
	if o == nil {
		return nil
	}
	return o.AssignedTo
}

func (o *FlowTemplateInput) GetAvailableInEcp() *bool {
	if o == nil {
		return nil
	}
	return o.AvailableInEcp
}

func (o *FlowTemplateInput) GetClosingReasons() []ClosingReasonInput {
	if o == nil {
		return nil
	}
	return o.ClosingReasons
}

func (o *FlowTemplateInput) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *FlowTemplateInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *FlowTemplateInput) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *FlowTemplateInput) GetDueDateConfig() *DueDateConfig {
	if o == nil {
		return nil
	}
	return o.DueDateConfig
}

func (o *FlowTemplateInput) GetEdges() []Edge {
	if o == nil {
		return []Edge{}
	}
	return o.Edges
}

func (o *FlowTemplateInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *FlowTemplateInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FlowTemplateInput) GetIsFlowMigrated() *bool {
	if o == nil {
		return nil
	}
	return o.IsFlowMigrated
}

func (o *FlowTemplateInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *FlowTemplateInput) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

func (o *FlowTemplateInput) GetPhases() []Phase {
	if o == nil {
		return nil
	}
	return o.Phases
}

func (o *FlowTemplateInput) GetTasks() []Task {
	if o == nil {
		return []Task{}
	}
	return o.Tasks
}

func (o *FlowTemplateInput) GetTaxonomies() []string {
	if o == nil {
		return nil
	}
	return o.Taxonomies
}

func (o *FlowTemplateInput) GetTrigger() *Trigger {
	if o == nil {
		return nil
	}
	return o.Trigger
}

func (o *FlowTemplateInput) GetUpdateEntityAttributes() []UpdateEntityAttributes {
	if o == nil {
		return nil
	}
	return o.UpdateEntityAttributes
}

func (o *FlowTemplateInput) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
