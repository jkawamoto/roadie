package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// BatchAccountUpdateBaseProperties The properties for a Batch account update.
// swagger:model BatchAccountUpdateBaseProperties
type BatchAccountUpdateBaseProperties struct {

	// The properties related to auto storage account.
	AutoStorage *AutoStorageBaseProperties `json:"autoStorage,omitempty"`
}

// Validate validates this batch account update base properties
func (m *BatchAccountUpdateBaseProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoStorage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchAccountUpdateBaseProperties) validateAutoStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoStorage) { // not required
		return nil
	}

	if m.AutoStorage != nil {

		if err := m.AutoStorage.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
