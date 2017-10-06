package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PurchasePlan Used for establishing the purchase context of any 3rd Party artifact through MarketPlace.
// swagger:model PurchasePlan
type PurchasePlan struct {

	// The plan ID.
	// Required: true
	Name *string `json:"name"`

	// The product ID.
	// Required: true
	Product *string `json:"product"`

	// The publisher ID.
	// Required: true
	Publisher *string `json:"publisher"`
}

// Validate validates this purchase plan
func (m *PurchasePlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublisher(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PurchasePlan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PurchasePlan) validateProduct(formats strfmt.Registry) error {

	if err := validate.Required("product", "body", m.Product); err != nil {
		return err
	}

	return nil
}

func (m *PurchasePlan) validatePublisher(formats strfmt.Registry) error {

	if err := validate.Required("publisher", "body", m.Publisher); err != nil {
		return err
	}

	return nil
}
