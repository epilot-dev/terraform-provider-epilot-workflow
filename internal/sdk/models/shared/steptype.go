// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type StepType string

const (
	StepTypeManual     StepType = "MANUAL"
	StepTypeAutomation StepType = "AUTOMATION"
)

func (e StepType) ToPointer() *StepType {
	return &e
}
func (e *StepType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MANUAL":
		fallthrough
	case "AUTOMATION":
		*e = StepType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StepType: %v", v)
	}
}
