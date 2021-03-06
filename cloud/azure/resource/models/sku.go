package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Sku SKU for the resource.
// swagger:model Sku
type Sku struct {

	// The SKU capacity.
	Capacity int32 `json:"capacity,omitempty"`

	// The SKU family.
	Family string `json:"family,omitempty"`

	// The SKU model.
	Model string `json:"model,omitempty"`

	// The SKU name.
	Name string `json:"name,omitempty"`

	// The SKU size.
	Size string `json:"size,omitempty"`

	// The SKU tier.
	Tier string `json:"tier,omitempty"`
}

// Validate validates this sku
func (m *Sku) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
