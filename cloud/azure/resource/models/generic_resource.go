package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GenericResource Resource information.
// swagger:model GenericResource
type GenericResource struct {
	Resource

	// The identity of the resource.
	Identity *Identity `json:"identity,omitempty"`

	// The kind of the resource.
	// Pattern: ^[-\w\._,\(\)]+$
	Kind string `json:"kind,omitempty"`

	// ID of the resource that manages this resource.
	ManagedBy string `json:"managedBy,omitempty"`

	// The plan of the resource.
	Plan *Plan `json:"plan,omitempty"`

	// The resource properties.
	Properties interface{} `json:"properties,omitempty"`

	// The SKU of the resource.
	Sku *Sku `json:"sku,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GenericResource) UnmarshalJSON(raw []byte) error {

	var aO0 Resource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Resource = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GenericResource) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Resource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this generic resource
func (m *GenericResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Resource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSku(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericResource) validateIdentity(formats strfmt.Registry) error {

	if swag.IsZero(m.Identity) { // not required
		return nil
	}

	if m.Identity != nil {

		if err := m.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *GenericResource) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if err := validate.Pattern("kind", "body", string(m.Kind), `^[-\w\._,\(\)]+$`); err != nil {
		return err
	}

	return nil
}

func (m *GenericResource) validatePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	if m.Plan != nil {

		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

func (m *GenericResource) validateSku(formats strfmt.Registry) error {

	if swag.IsZero(m.Sku) { // not required
		return nil
	}

	if m.Sku != nil {

		if err := m.Sku.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sku")
			}
			return err
		}
	}

	return nil
}
