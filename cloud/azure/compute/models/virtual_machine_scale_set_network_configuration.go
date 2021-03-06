package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VirtualMachineScaleSetNetworkConfiguration Describes a virtual machine scale set network profile's network configurations.
// swagger:model VirtualMachineScaleSetNetworkConfiguration
type VirtualMachineScaleSetNetworkConfiguration struct {
	SubResource

	// The network configuration name.
	// Required: true
	Name *string `json:"name"`

	// properties
	Properties *VirtualMachineScaleSetNetworkConfigurationProperties `json:"properties,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VirtualMachineScaleSetNetworkConfiguration) UnmarshalJSON(raw []byte) error {
	var data struct {
		Name *string `json:"name"`

		Properties *VirtualMachineScaleSetNetworkConfigurationProperties `json:"properties,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

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
func (m VirtualMachineScaleSetNetworkConfiguration) MarshalJSON() ([]byte, error) {
	var _parts [][]byte
	var data struct {
		Name *string `json:"name"`

		Properties *VirtualMachineScaleSetNetworkConfigurationProperties `json:"properties,omitempty"`
	}

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

// Validate validates this virtual machine scale set network configuration
func (m *VirtualMachineScaleSetNetworkConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.SubResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
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

func (m *VirtualMachineScaleSetNetworkConfiguration) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineScaleSetNetworkConfiguration) validateProperties(formats strfmt.Registry) error {

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
