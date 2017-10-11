package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Plan Plan for the resource.
// swagger:model Plan
type Plan struct {

	// The plan ID.
	Name string `json:"name,omitempty"`

	// The offer ID.
	Product string `json:"product,omitempty"`

	// The promotion code.
	PromotionCode string `json:"promotionCode,omitempty"`

	// The publisher ID.
	Publisher string `json:"publisher,omitempty"`
}

// Validate validates this plan
func (m *Plan) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}