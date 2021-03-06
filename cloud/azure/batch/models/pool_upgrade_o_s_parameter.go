package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PoolUpgradeOSParameter Options for upgrading the operating system of compute nodes in a pool.
// swagger:model PoolUpgradeOSParameter
type PoolUpgradeOSParameter struct {

	// The Azure Guest OS version to be installed on the virtual machines in the pool.
	// Required: true
	TargetOSVersion *string `json:"targetOSVersion"`
}

// Validate validates this pool upgrade o s parameter
func (m *PoolUpgradeOSParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetOSVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolUpgradeOSParameter) validateTargetOSVersion(formats strfmt.Registry) error {

	if err := validate.Required("targetOSVersion", "body", m.TargetOSVersion); err != nil {
		return err
	}

	return nil
}
