// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/internal/utils"
)

type TriggerType string

const (
	TriggerTypeManualTrigger            TriggerType = "ManualTrigger"
	TriggerTypeAutomationTrigger        TriggerType = "AutomationTrigger"
	TriggerTypeJourneySubmissionTrigger TriggerType = "JourneySubmissionTrigger"
)

type Trigger struct {
	ManualTrigger            *ManualTrigger            `queryParam:"inline"`
	AutomationTrigger        *AutomationTrigger        `queryParam:"inline"`
	JourneySubmissionTrigger *JourneySubmissionTrigger `queryParam:"inline"`

	Type TriggerType
}

func CreateTriggerManualTrigger(manualTrigger ManualTrigger) Trigger {
	typ := TriggerTypeManualTrigger

	return Trigger{
		ManualTrigger: &manualTrigger,
		Type:          typ,
	}
}

func CreateTriggerAutomationTrigger(automationTrigger AutomationTrigger) Trigger {
	typ := TriggerTypeAutomationTrigger

	return Trigger{
		AutomationTrigger: &automationTrigger,
		Type:              typ,
	}
}

func CreateTriggerJourneySubmissionTrigger(journeySubmissionTrigger JourneySubmissionTrigger) Trigger {
	typ := TriggerTypeJourneySubmissionTrigger

	return Trigger{
		JourneySubmissionTrigger: &journeySubmissionTrigger,
		Type:                     typ,
	}
}

func (u *Trigger) UnmarshalJSON(data []byte) error {

	var manualTrigger ManualTrigger = ManualTrigger{}
	if err := utils.UnmarshalJSON(data, &manualTrigger, "", true, true); err == nil {
		u.ManualTrigger = &manualTrigger
		u.Type = TriggerTypeManualTrigger
		return nil
	}

	var automationTrigger AutomationTrigger = AutomationTrigger{}
	if err := utils.UnmarshalJSON(data, &automationTrigger, "", true, true); err == nil {
		u.AutomationTrigger = &automationTrigger
		u.Type = TriggerTypeAutomationTrigger
		return nil
	}

	var journeySubmissionTrigger JourneySubmissionTrigger = JourneySubmissionTrigger{}
	if err := utils.UnmarshalJSON(data, &journeySubmissionTrigger, "", true, true); err == nil {
		u.JourneySubmissionTrigger = &journeySubmissionTrigger
		u.Type = TriggerTypeJourneySubmissionTrigger
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Trigger", string(data))
}

func (u Trigger) MarshalJSON() ([]byte, error) {
	if u.ManualTrigger != nil {
		return utils.MarshalJSON(u.ManualTrigger, "", true)
	}

	if u.AutomationTrigger != nil {
		return utils.MarshalJSON(u.AutomationTrigger, "", true)
	}

	if u.JourneySubmissionTrigger != nil {
		return utils.MarshalJSON(u.JourneySubmissionTrigger, "", true)
	}

	return nil, errors.New("could not marshal union type Trigger: all fields are null")
}
