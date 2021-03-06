package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VirtualMachineScaleSetVMInstanceRequiredIds Specifies a list of virtual machine instance IDs from the VM scale set.
// swagger:model VirtualMachineScaleSetVMInstanceRequiredIDs
type VirtualMachineScaleSetVMInstanceRequiredIds struct {

	// The virtual machine scale set instance ids.
	// Required: true
	InstanceIds []string `json:"instanceIds"`
}

// Validate validates this virtual machine scale set VM instance required ids
func (m *VirtualMachineScaleSetVMInstanceRequiredIds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineScaleSetVMInstanceRequiredIds) validateInstanceIds(formats strfmt.Registry) error {

	if err := validate.Required("instanceIds", "body", m.InstanceIds); err != nil {
		return err
	}

	return nil
}
