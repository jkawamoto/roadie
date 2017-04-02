package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// LinuxConfiguration Describes Windows configuration of the OS Profile.
// swagger:model LinuxConfiguration
type LinuxConfiguration struct {

	// Specifies whether password authentication should be disabled.
	DisablePasswordAuthentication bool `json:"disablePasswordAuthentication,omitempty"`

	// The SSH configuration for linux VMs.
	SSH *SSHConfiguration `json:"ssh,omitempty"`
}

// Validate validates this linux configuration
func (m *LinuxConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSSH(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinuxConfiguration) validateSSH(formats strfmt.Registry) error {

	if swag.IsZero(m.SSH) { // not required
		return nil
	}

	if m.SSH != nil {

		if err := m.SSH.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
