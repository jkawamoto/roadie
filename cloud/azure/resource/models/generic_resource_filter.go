package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GenericResourceFilter Resource filter.
// swagger:model GenericResourceFilter
type GenericResourceFilter struct {

	// The resource type.
	ResourceType string `json:"resourceType,omitempty"`

	// The tag name.
	Tagname string `json:"tagname,omitempty"`

	// The tag value.
	Tagvalue string `json:"tagvalue,omitempty"`
}

// Validate validates this generic resource filter
func (m *GenericResourceFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
