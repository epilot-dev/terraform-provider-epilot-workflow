// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type StepJourney struct {
	ID        *string `json:"id,omitempty"`
	JourneyID *string `json:"journeyId,omitempty"`
	Name      *string `json:"name,omitempty"`
}

func (o *StepJourney) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *StepJourney) GetJourneyID() *string {
	if o == nil {
		return nil
	}
	return o.JourneyID
}

func (o *StepJourney) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
