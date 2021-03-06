package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VirtualMachineScaleSetListSkusResult The Virtual Machine Scale Set List Skus operation response.
// swagger:model VirtualMachineScaleSetListSkusResult
type VirtualMachineScaleSetListSkusResult struct {

	// The uri to fetch the next page of Virtual Machine Scale Set Skus. Call ListNext() with this to fetch the next page of VMSS Skus.
	NextLink string `json:"nextLink,omitempty"`

	// The list of skus available for the virtual machine scale set.
	// Required: true
	Value []*VirtualMachineScaleSetSku `json:"value"`
}

// Validate validates this virtual machine scale set list skus result
func (m *VirtualMachineScaleSetListSkusResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineScaleSetListSkusResult) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	for i := 0; i < len(m.Value); i++ {

		if swag.IsZero(m.Value[i]) { // not required
			continue
		}

		if m.Value[i] != nil {

			if err := m.Value[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
