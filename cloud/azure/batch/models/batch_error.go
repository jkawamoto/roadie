package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// BatchError An error response received from the Azure Batch service.
// swagger:model BatchError
type BatchError struct {

	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code string `json:"code,omitempty"`

	// A message describing the error, intended to be suitable for display in a user interface.
	Message *ErrorMessage `json:"message,omitempty"`

	// A collection of key-value pairs containing additional details about the error.
	Values []*BatchErrorDetail `json:"values"`
}

// Validate validates this batch error
func (m *BatchError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchError) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if m.Message != nil {

		if err := m.Message.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *BatchError) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {

		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {

			if err := m.Values[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}