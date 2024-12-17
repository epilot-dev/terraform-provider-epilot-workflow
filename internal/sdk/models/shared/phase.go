// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Phase struct {
	AssignedTo []string `json:"assigned_to,omitempty"`
	DueDate    *string  `json:"due_date,omitempty"`
	// Set due date for the task based on a dynamic condition
	DueDateConfig *DueDateConfig `json:"due_date_config,omitempty"`
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	// Taxonomy ids that are associated with this workflow and used for filtering
	Taxonomies []string `json:"taxonomies,omitempty"`
}

func (o *Phase) GetAssignedTo() []string {
	if o == nil {
		return nil
	}
	return o.AssignedTo
}

func (o *Phase) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *Phase) GetDueDateConfig() *DueDateConfig {
	if o == nil {
		return nil
	}
	return o.DueDateConfig
}

func (o *Phase) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Phase) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Phase) GetTaxonomies() []string {
	if o == nil {
		return nil
	}
	return o.Taxonomies
}
