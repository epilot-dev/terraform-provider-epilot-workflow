// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ItemType string

const (
	ItemTypeStep    ItemType = "STEP"
	ItemTypeSection ItemType = "SECTION"
)

func (e ItemType) ToPointer() *ItemType {
	return &e
}
func (e *ItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STEP":
		fallthrough
	case "SECTION":
		*e = ItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ItemType: %v", v)
	}
}
