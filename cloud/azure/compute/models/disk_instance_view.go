package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// DiskInstanceView The instance view of the disk.
// swagger:model DiskInstanceView
type DiskInstanceView struct {

	// The disk name.
	Name string `json:"name,omitempty"`

	// The resource status information.
	Statuses []*InstanceViewStatus `json:"statuses"`
}

// Validate validates this disk instance view
func (m *DiskInstanceView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatuses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiskInstanceView) validateStatuses(formats strfmt.Registry) error {

	if swag.IsZero(m.Statuses) { // not required
		return nil
	}

	for i := 0; i < len(m.Statuses); i++ {

		if swag.IsZero(m.Statuses[i]) { // not required
			continue
		}

		if m.Statuses[i] != nil {

			if err := m.Statuses[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
