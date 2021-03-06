package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// OperationStatusResponse Operation status response
// swagger:model OperationStatusResponse
type OperationStatusResponse struct {

	// End time of the operation
	// Read Only: true
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// Api error
	// Read Only: true
	Error *APIError `json:"error,omitempty"`

	// Operation ID
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Start time of the operation
	// Read Only: true
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// Operation status
	// Read Only: true
	Status string `json:"status,omitempty"`
}

// Validate validates this operation status response
func (m *OperationStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperationStatusResponse) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {

		if err := m.Error.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
