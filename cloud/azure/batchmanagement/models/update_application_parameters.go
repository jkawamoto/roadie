package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// UpdateApplicationParameters Parameters for an ApplicationOperations.UpdateApplication request.
// swagger:model UpdateApplicationParameters
type UpdateApplicationParameters struct {

	// A value indicating whether packages within the application may be overwritten using the same version string.
	AllowUpdates bool `json:"allowUpdates,omitempty"`

	// The package to use if a client requests the application but does not specify a version.
	DefaultVersion string `json:"defaultVersion,omitempty"`

	// The display name for the application.
	DisplayName string `json:"displayName,omitempty"`
}

// Validate validates this update application parameters
func (m *UpdateApplicationParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
