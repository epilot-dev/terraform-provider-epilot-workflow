// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// Section - A group of Steps that define the progress of the Workflow
type Section struct {
	ID    *string  `json:"id,omitempty"`
	Name  string   `json:"name"`
	Order float64  `json:"order"`
	Steps []Step   `json:"steps"`
	Type  ItemType `json:"type"`
}

func (o *Section) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Section) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Section) GetOrder() float64 {
	if o == nil {
		return 0.0
	}
	return o.Order
}

func (o *Section) GetSteps() []Step {
	if o == nil {
		return []Step{}
	}
	return o.Steps
}

func (o *Section) GetType() ItemType {
	if o == nil {
		return ItemType("")
	}
	return o.Type
}
