package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// VirtualMachineScaleSet Describes a Virtual Machine Scale Set.
// swagger:model VirtualMachineScaleSet
type VirtualMachineScaleSet struct {
	Resource

	// The identity of the virtual machine scale set, if configured.
	Identity *VirtualMachineScaleSetIdentity `json:"identity,omitempty"`

	// The purchase plan when deploying a virtual machine scale set from VM Marketplace images.
	Plan *Plan `json:"plan,omitempty"`

	// properties
	Properties *VirtualMachineScaleSetProperties `json:"properties,omitempty"`

	// The virtual machine scale set sku.
	Sku *Sku `json:"sku,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VirtualMachineScaleSet) UnmarshalJSON(raw []byte) error {
	var data struct {
		Identity *VirtualMachineScaleSetIdentity `json:"identity,omitempty"`

		Plan *Plan `json:"plan,omitempty"`

		Properties *VirtualMachineScaleSetProperties `json:"properties,omitempty"`

		Sku *Sku `json:"sku,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.Identity = data.Identity

	m.Plan = data.Plan

	m.Properties = data.Properties

	m.Sku = data.Sku

	var aO0 Resource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Resource = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VirtualMachineScaleSet) MarshalJSON() ([]byte, error) {
	var _parts [][]byte
	var data struct {
		Identity *VirtualMachineScaleSetIdentity `json:"identity,omitempty"`

		Plan *Plan `json:"plan,omitempty"`

		Properties *VirtualMachineScaleSetProperties `json:"properties,omitempty"`

		Sku *Sku `json:"sku,omitempty"`
	}

	data.Identity = m.Identity

	data.Plan = m.Plan

	data.Properties = m.Properties

	data.Sku = m.Sku

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	aO0, err := swag.WriteJSON(m.Resource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this virtual machine scale set
func (m *VirtualMachineScaleSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Resource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
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

func (m *VirtualMachineScaleSet) validateIdentity(formats strfmt.Registry) error {

	if swag.IsZero(m.Identity) { // not required
		return nil
	}

	if m.Identity != nil {

		if err := m.Identity.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VirtualMachineScaleSet) validatePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	if m.Plan != nil {

		if err := m.Plan.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VirtualMachineScaleSet) validateProperties(formats strfmt.Registry) error {

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

func (m *VirtualMachineScaleSet) validateSku(formats strfmt.Registry) error {

	if swag.IsZero(m.Sku) { // not required
		return nil
	}

	if m.Sku != nil {

		if err := m.Sku.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
