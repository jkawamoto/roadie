package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// TagCount Tag count.
// swagger:model TagCount
type TagCount struct {

	// Type of count.
	Type string `json:"type,omitempty"`

	// Value of count.
	Value int64 `json:"value,omitempty"`
}

// Validate validates this tag count
func (m *TagCount) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
