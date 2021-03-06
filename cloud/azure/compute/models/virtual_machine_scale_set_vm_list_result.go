package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VirtualMachineScaleSetVMListResult The List Virtual Machine Scale Set VMs operation response.
// swagger:model VirtualMachineScaleSetVMListResult
type VirtualMachineScaleSetVMListResult struct {

	// The uri to fetch the next page of Virtual Machine Scale Set VMs. Call ListNext() with this to fetch the next page of VMSS VMs
	NextLink string `json:"nextLink,omitempty"`

	// The list of virtual machine scale sets VMs.
	// Required: true
	Value []*VirtualMachineScaleSetVM `json:"value"`
}

// Validate validates this virtual machine scale set VM list result
func (m *VirtualMachineScaleSetVMListResult) Validate(formats strfmt.Registry) error {
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

func (m *VirtualMachineScaleSetVMListResult) validateValue(formats strfmt.Registry) error {

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
