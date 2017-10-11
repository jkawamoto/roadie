package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// InstanceViewStatus Instance view status.
// swagger:model InstanceViewStatus
type InstanceViewStatus struct {

	// The status code.
	Code string `json:"code,omitempty"`

	// The short localizable label for the status.
	DisplayStatus string `json:"displayStatus,omitempty"`

	// The level code.
	Level string `json:"level,omitempty"`

	// The detailed status message, including for alerts and error messages.
	Message string `json:"message,omitempty"`

	// The time of the status.
	Time strfmt.DateTime `json:"time,omitempty"`
}

// Validate validates this instance view status
func (m *InstanceViewStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var instanceViewStatusTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Info","Warning","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceViewStatusTypeLevelPropEnum = append(instanceViewStatusTypeLevelPropEnum, v)
	}
}

const (
	// InstanceViewStatusLevelInfo captures enum value "Info"
	InstanceViewStatusLevelInfo string = "Info"
	// InstanceViewStatusLevelWarning captures enum value "Warning"
	InstanceViewStatusLevelWarning string = "Warning"
	// InstanceViewStatusLevelError captures enum value "Error"
	InstanceViewStatusLevelError string = "Error"
)

// prop value enum
func (m *InstanceViewStatus) validateLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceViewStatusTypeLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceViewStatus) validateLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}