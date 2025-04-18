// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TriggerMode string

const (
	TriggerModeManual    TriggerMode = "manual"
	TriggerModeAutomatic TriggerMode = "automatic"
)

func (e TriggerMode) ToPointer() *TriggerMode {
	return &e
}
func (e *TriggerMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "manual":
		fallthrough
	case "automatic":
		*e = TriggerMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TriggerMode: %v", v)
	}
}
