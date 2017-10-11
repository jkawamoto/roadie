package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ResourceGroupProperties The resource group properties.
// swagger:model ResourceGroupProperties
type ResourceGroupProperties struct {

	// The provisioning state.
	// Read Only: true
	ProvisioningState string `json:"provisioningState,omitempty"`
}

// Validate validates this resource group properties
func (m *ResourceGroupProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}