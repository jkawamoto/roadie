package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// VirtualMachineScaleSetProperties Describes the properties of a Virtual Machine Scale Set.
// swagger:model VirtualMachineScaleSetProperties
type VirtualMachineScaleSetProperties struct {

	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	Overprovision bool `json:"overprovision,omitempty"`

	// The provisioning state, which only appears in the response.
	// Read Only: true
	ProvisioningState string `json:"provisioningState,omitempty"`

	// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
	SinglePlacementGroup bool `json:"singlePlacementGroup,omitempty"`

	// The upgrade policy.
	UpgradePolicy *UpgradePolicy `json:"upgradePolicy,omitempty"`

	// The virtual machine profile.
	VirtualMachineProfile *VirtualMachineScaleSetVMProfile `json:"virtualMachineProfile,omitempty"`
}

// Validate validates this virtual machine scale set properties
func (m *VirtualMachineScaleSetProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpgradePolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVirtualMachineProfile(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineScaleSetProperties) validateUpgradePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.UpgradePolicy) { // not required
		return nil
	}

	if m.UpgradePolicy != nil {

		if err := m.UpgradePolicy.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VirtualMachineScaleSetProperties) validateVirtualMachineProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualMachineProfile) { // not required
		return nil
	}

	if m.VirtualMachineProfile != nil {

		if err := m.VirtualMachineProfile.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}