package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// VirtualMachineExtensionProperties Describes the properties of a Virtual Machine Extension.
// swagger:model VirtualMachineExtensionProperties
type VirtualMachineExtensionProperties struct {

	// Whether the extension handler should be automatically upgraded across minor versions.
	AutoUpgradeMinorVersion bool `json:"autoUpgradeMinorVersion,omitempty"`

	// How the extension handler should be forced to update even if the extension configuration has not changed.
	ForceUpdateTag string `json:"forceUpdateTag,omitempty"`

	// The virtual machine extension instance view.
	InstanceView *VirtualMachineExtensionInstanceView `json:"instanceView,omitempty"`

	// Json formatted protected settings for the extension.
	ProtectedSettings interface{} `json:"protectedSettings,omitempty"`

	// The provisioning state, which only appears in the response.
	// Read Only: true
	ProvisioningState string `json:"provisioningState,omitempty"`

	// The name of the extension handler publisher.
	Publisher string `json:"publisher,omitempty"`

	// Json formatted public settings for the extension.
	Settings interface{} `json:"settings,omitempty"`

	// The type of the extension handler.
	Type string `json:"type,omitempty"`

	// The type version of the extension handler.
	TypeHandlerVersion string `json:"typeHandlerVersion,omitempty"`
}

// Validate validates this virtual machine extension properties
func (m *VirtualMachineExtensionProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceView(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineExtensionProperties) validateInstanceView(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceView) { // not required
		return nil
	}

	if m.InstanceView != nil {

		if err := m.InstanceView.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
