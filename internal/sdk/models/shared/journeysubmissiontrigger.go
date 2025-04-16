// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type JourneySubmissionTriggerType string

const (
	JourneySubmissionTriggerTypeJourneySubmission JourneySubmissionTriggerType = "journey_submission"
)

func (e JourneySubmissionTriggerType) ToPointer() *JourneySubmissionTriggerType {
	return &e
}
func (e *JourneySubmissionTriggerType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "journey_submission":
		*e = JourneySubmissionTriggerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JourneySubmissionTriggerType: %v", v)
	}
}

type JourneySubmissionTrigger struct {
	AutomationID *string `json:"automation_id,omitempty"`
	ID           *string `json:"id,omitempty"`
	// ID of the journey that will trigger this flow
	JourneyID string                       `json:"journey_id"`
	Type      JourneySubmissionTriggerType `json:"type"`
}

func (o *JourneySubmissionTrigger) GetAutomationID() *string {
	if o == nil {
		return nil
	}
	return o.AutomationID
}

func (o *JourneySubmissionTrigger) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *JourneySubmissionTrigger) GetJourneyID() string {
	if o == nil {
		return ""
	}
	return o.JourneyID
}

func (o *JourneySubmissionTrigger) GetType() JourneySubmissionTriggerType {
	if o == nil {
		return JourneySubmissionTriggerType("")
	}
	return o.Type
}
