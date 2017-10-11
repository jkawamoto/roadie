package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// MetadataItem A name-value pair associated with a Batch service resource.
//
// The Batch service does not assign any meaning to this metadata; it is solely for the use of user code.
// swagger:model MetadataItem
type MetadataItem struct {

	// The name of the metadata item.
	// Required: true
	Name *string `json:"name"`

	// The value of the metadata item.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this metadata item
func (m *MetadataItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetadataItem) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MetadataItem) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}