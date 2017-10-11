package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// VirtualMachineScaleSetStorageProfile Describes a virtual machine scale set storage profile.
// swagger:model VirtualMachineScaleSetStorageProfile
type VirtualMachineScaleSetStorageProfile struct {

	// The data disks.
	DataDisks []*VirtualMachineScaleSetDataDisk `json:"dataDisks"`

	// The image reference.
	ImageReference *ImageReference `json:"imageReference,omitempty"`

	// The OS disk.
	OsDisk *VirtualMachineScaleSetOSDisk `json:"osDisk,omitempty"`
}

// Validate validates this virtual machine scale set storage profile
func (m *VirtualMachineScaleSetStorageProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataDisks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImageReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOsDisk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineScaleSetStorageProfile) validateDataDisks(formats strfmt.Registry) error {

	if swag.IsZero(m.DataDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.DataDisks); i++ {

		if swag.IsZero(m.DataDisks[i]) { // not required
			continue
		}

		if m.DataDisks[i] != nil {

			if err := m.DataDisks[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VirtualMachineScaleSetStorageProfile) validateImageReference(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageReference) { // not required
		return nil
	}

	if m.ImageReference != nil {

		if err := m.ImageReference.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VirtualMachineScaleSetStorageProfile) validateOsDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.OsDisk) { // not required
		return nil
	}

	if m.OsDisk != nil {

		if err := m.OsDisk.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}