package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// DeploymentExportResult The deployment export result.
// swagger:model DeploymentExportResult
type DeploymentExportResult struct {

	// The template content.
	Template interface{} `json:"template,omitempty"`
}

// Validate validates this deployment export result
func (m *DeploymentExportResult) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
