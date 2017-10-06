package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ApplicationPackageReference A reference to an application package to be deployed to compute nodes.
// swagger:model ApplicationPackageReference
type ApplicationPackageReference struct {

	// The ID of the application to deploy.
	// Required: true
	ApplicationID *string `json:"applicationId"`

	// The version of the application to deploy. If omitted, the default version is deployed.
	//
	// If this is omitted, and no default version is specified for this application, the request fails with the error code InvalidApplicationPackageReferences. If you are calling the REST API directly, the HTTP status code is 409.
	Version string `json:"version,omitempty"`
}

// Validate validates this application package reference
func (m *ApplicationPackageReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationPackageReference) validateApplicationID(formats strfmt.Registry) error {

	if err := validate.Required("applicationId", "body", m.ApplicationID); err != nil {
		return err
	}

	return nil
}
