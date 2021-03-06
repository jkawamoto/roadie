package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// VirtualMachineScaleSetInstanceView The instance view of a virtual machine scale set.
// swagger:model VirtualMachineScaleSetInstanceView
type VirtualMachineScaleSetInstanceView struct {

	// The extensions information.
	// Read Only: true
	Extensions []*VirtualMachineScaleSetVMExtensionsSummary `json:"extensions"`

	// The resource status information.
	Statuses []*InstanceViewStatus `json:"statuses"`

	// The instance view status summary for the virtual machine scale set.
	// Read Only: true
	VirtualMachine *VirtualMachineScaleSetInstanceViewStatusesSummary `json:"virtualMachine,omitempty"`
}

// Validate validates this virtual machine scale set instance view
func (m *VirtualMachineScaleSetInstanceView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatuses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVirtualMachine(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineScaleSetInstanceView) validateExtensions(formats strfmt.Registry) error {

	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	for i := 0; i < len(m.Extensions); i++ {

		if swag.IsZero(m.Extensions[i]) { // not required
			continue
		}

		if m.Extensions[i] != nil {

			if err := m.Extensions[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VirtualMachineScaleSetInstanceView) validateStatuses(formats strfmt.Registry) error {

	if swag.IsZero(m.Statuses) { // not required
		return nil
	}

	for i := 0; i < len(m.Statuses); i++ {

		if swag.IsZero(m.Statuses[i]) { // not required
			continue
		}

		if m.Statuses[i] != nil {

			if err := m.Statuses[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VirtualMachineScaleSetInstanceView) validateVirtualMachine(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualMachine) { // not required
		return nil
	}

	if m.VirtualMachine != nil {

		if err := m.VirtualMachine.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
