// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Operator string

const (
	OperatorEquals              Operator = "equals"
	OperatorNotEquals           Operator = "not_equals"
	OperatorAnyOf               Operator = "any_of"
	OperatorNoneOf              Operator = "none_of"
	OperatorContains            Operator = "contains"
	OperatorNotContains         Operator = "not_contains"
	OperatorStartsWith          Operator = "starts_with"
	OperatorEndsWith            Operator = "ends_with"
	OperatorGreaterThan         Operator = "greater_than"
	OperatorLessThan            Operator = "less_than"
	OperatorGreaterThanOrEquals Operator = "greater_than_or_equals"
	OperatorLessThanOrEquals    Operator = "less_than_or_equals"
	OperatorIsEmpty             Operator = "is_empty"
	OperatorIsNotEmpty          Operator = "is_not_empty"
)

func (e Operator) ToPointer() *Operator {
	return &e
}
func (e *Operator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "equals":
		fallthrough
	case "not_equals":
		fallthrough
	case "any_of":
		fallthrough
	case "none_of":
		fallthrough
	case "contains":
		fallthrough
	case "not_contains":
		fallthrough
	case "starts_with":
		fallthrough
	case "ends_with":
		fallthrough
	case "greater_than":
		fallthrough
	case "less_than":
		fallthrough
	case "greater_than_or_equals":
		fallthrough
	case "less_than_or_equals":
		fallthrough
	case "is_empty":
		fallthrough
	case "is_not_empty":
		*e = Operator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Operator: %v", v)
	}
}
