package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// TenantIDDescription Tenant Id information.
// swagger:model TenantIdDescription
type TenantIDDescription struct {

	// The fully qualified ID of the tenant. For example, /tenants/00000000-0000-0000-0000-000000000000.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The tenant ID. For example, 00000000-0000-0000-0000-000000000000.
	// Read Only: true
	TenantID string `json:"tenantId,omitempty"`
}

// Validate validates this tenant Id description
func (m *TenantIDDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
