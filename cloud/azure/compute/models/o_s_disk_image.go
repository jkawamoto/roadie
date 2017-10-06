package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// OSDiskImage Contains the os disk image information.
// swagger:model OSDiskImage
type OSDiskImage struct {

	// The operating system of the osDiskImage.
	// Required: true
	OperatingSystem *string `json:"operatingSystem"`
}

// Validate validates this o s disk image
func (m *OSDiskImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperatingSystem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oSDiskImageTypeOperatingSystemPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Windows","Linux"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oSDiskImageTypeOperatingSystemPropEnum = append(oSDiskImageTypeOperatingSystemPropEnum, v)
	}
}

const (
	// OSDiskImageOperatingSystemWindows captures enum value "Windows"
	OSDiskImageOperatingSystemWindows string = "Windows"
	// OSDiskImageOperatingSystemLinux captures enum value "Linux"
	OSDiskImageOperatingSystemLinux string = "Linux"
)

// prop value enum
func (m *OSDiskImage) validateOperatingSystemEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, oSDiskImageTypeOperatingSystemPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *OSDiskImage) validateOperatingSystem(formats strfmt.Registry) error {

	if err := validate.Required("operatingSystem", "body", m.OperatingSystem); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperatingSystemEnum("operatingSystem", "body", *m.OperatingSystem); err != nil {
		return err
	}

	return nil
}
