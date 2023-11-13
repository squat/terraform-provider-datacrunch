// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ActionType string

const (
	ActionTypeStart     ActionType = "start"
	ActionTypeShutdown  ActionType = "shutdown"
	ActionTypeHibernate ActionType = "hibernate"
	ActionTypeRestore   ActionType = "restore"
	ActionTypeDelete    ActionType = "delete"
)

func (e ActionType) ToPointer() *ActionType {
	return &e
}

func (e *ActionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "start":
		fallthrough
	case "shutdown":
		fallthrough
	case "hibernate":
		fallthrough
	case "restore":
		fallthrough
	case "delete":
		*e = ActionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ActionType: %v", v)
	}
}
