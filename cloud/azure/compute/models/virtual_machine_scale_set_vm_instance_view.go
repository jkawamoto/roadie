package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// VirtualMachineScaleSetVMInstanceView The instance view of a virtual machine scale set VM.
// swagger:model VirtualMachineScaleSetVMInstanceView
type VirtualMachineScaleSetVMInstanceView struct {

	// The boot diagnostics.
	BootDiagnostics *BootDiagnosticsInstanceView `json:"bootDiagnostics,omitempty"`

	// The disks information.
	Disks []*DiskInstanceView `json:"disks"`

	// The extensions information.
	Extensions []*VirtualMachineExtensionInstanceView `json:"extensions"`

	// The placement group in which the VM is running. If the VM is deallocated it will not have a placementGroupId.
	PlacementGroupID string `json:"placementGroupId,omitempty"`

	// The Fault Domain count.
	PlatformFaultDomain int32 `json:"platformFaultDomain,omitempty"`

	// The Update Domain count.
	PlatformUpdateDomain int32 `json:"platformUpdateDomain,omitempty"`

	// The Remote desktop certificate thumbprint.
	RdpThumbPrint string `json:"rdpThumbPrint,omitempty"`

	// The resource status information.
	Statuses []*InstanceViewStatus `json:"statuses"`

	// The VM Agent running on the virtual machine.
	VMAgent *VirtualMachineAgentInstanceView `json:"vmAgent,omitempty"`
}

// Validate validates this virtual machine scale set VM instance view
func (m *VirtualMachineScaleSetVMInstanceView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootDiagnostics(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatuses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVMAgent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineScaleSetVMInstanceView) validateBootDiagnostics(formats strfmt.Registry) error {

	if swag.IsZero(m.BootDiagnostics) { // not required
		return nil
	}

	if m.BootDiagnostics != nil {

		if err := m.BootDiagnostics.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VirtualMachineScaleSetVMInstanceView) validateDisks(formats strfmt.Registry) error {

	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {

		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {

			if err := m.Disks[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VirtualMachineScaleSetVMInstanceView) validateExtensions(formats strfmt.Registry) error {

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

func (m *VirtualMachineScaleSetVMInstanceView) validateStatuses(formats strfmt.Registry) error {

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

func (m *VirtualMachineScaleSetVMInstanceView) validateVMAgent(formats strfmt.Registry) error {

	if swag.IsZero(m.VMAgent) { // not required
		return nil
	}

	if m.VMAgent != nil {

		if err := m.VMAgent.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
