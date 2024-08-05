// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Source string

const (
	SourceWorkflowStatus Source = "workflow_status"
	SourceCurrentSection Source = "current_section"
	SourceCurrentStep    Source = "current_step"
)

func (e Source) ToPointer() *Source {
	return &e
}
func (e *Source) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workflow_status":
		fallthrough
	case "current_section":
		fallthrough
	case "current_step":
		*e = Source(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Source: %v", v)
	}
}

type Target struct {
	EntityAttribute string `json:"entityAttribute"`
	EntitySchema    string `json:"entitySchema"`
}

func (o *Target) GetEntityAttribute() string {
	if o == nil {
		return ""
	}
	return o.EntityAttribute
}

func (o *Target) GetEntitySchema() string {
	if o == nil {
		return ""
	}
	return o.EntitySchema
}

type UpdateEntityAttributes struct {
	Source Source `json:"source"`
	Target Target `json:"target"`
}

func (o *UpdateEntityAttributes) GetSource() Source {
	if o == nil {
		return Source("")
	}
	return o.Source
}

func (o *UpdateEntityAttributes) GetTarget() Target {
	if o == nil {
		return Target{}
	}
	return o.Target
}
