package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// InboundNatRule Inbound NAT rule of the loadbalancer
// swagger:model InboundNatRule
type InboundNatRule struct {
	SubResource

	// A unique read-only string that changes whenever the resource is updated
	Etag string `json:"etag,omitempty"`

	// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
	Name string `json:"name,omitempty"`

	// properties
	Properties *InboundNatRulePropertiesFormat `json:"properties,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InboundNatRule) UnmarshalJSON(raw []byte) error {
	var data struct {
		Etag string `json:"etag,omitempty"`

		Name string `json:"name,omitempty"`

		Properties *InboundNatRulePropertiesFormat `json:"properties,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.Etag = data.Etag

	m.Name = data.Name

	m.Properties = data.Properties

	var aO0 SubResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SubResource = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InboundNatRule) MarshalJSON() ([]byte, error) {
	var _parts [][]byte
	var data struct {
		Etag string `json:"etag,omitempty"`

		Name string `json:"name,omitempty"`

		Properties *InboundNatRulePropertiesFormat `json:"properties,omitempty"`
	}

	data.Etag = m.Etag

	data.Name = m.Name

	data.Properties = m.Properties

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	aO0, err := swag.WriteJSON(m.SubResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this inbound nat rule
func (m *InboundNatRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.SubResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundNatRule) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {

		if err := m.Properties.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
