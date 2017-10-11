package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// VirtualMachineExtensionHandlerInstanceView The instance view of a virtual machine extension handler.
// swagger:model VirtualMachineExtensionHandlerInstanceView
type VirtualMachineExtensionHandlerInstanceView struct {

	// The extension handler status.
	Status *InstanceViewStatus `json:"status,omitempty"`

	// Full type of the extension handler which includes both publisher and type.
	Type string `json:"type,omitempty"`

	// The type version of the extension handler.
	TypeHandlerVersion string `json:"typeHandlerVersion,omitempty"`
}

// Validate validates this virtual machine extension handler instance view
func (m *VirtualMachineExtensionHandlerInstanceView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineExtensionHandlerInstanceView) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {

		if err := m.Status.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}