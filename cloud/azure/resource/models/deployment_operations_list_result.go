package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeploymentOperationsListResult List of deployment operations.
// swagger:model DeploymentOperationsListResult
type DeploymentOperationsListResult struct {

	// The URL to use for getting the next set of results.
	// Read Only: true
	NextLink string `json:"nextLink,omitempty"`

	// An array of deployment operations.
	Value []*DeploymentOperation `json:"value"`
}

// Validate validates this deployment operations list result
func (m *DeploymentOperationsListResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentOperationsListResult) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	for i := 0; i < len(m.Value); i++ {

		if swag.IsZero(m.Value[i]) { // not required
			continue
		}

		if m.Value[i] != nil {

			if err := m.Value[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}