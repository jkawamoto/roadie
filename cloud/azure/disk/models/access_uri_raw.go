package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// AccessURIRaw This object gets 'bubbled up' through flattening.
// swagger:model AccessUriRaw
type AccessURIRaw struct {

	// A SAS uri for accessing a disk.
	// Read Only: true
	AccessSAS string `json:"accessSAS,omitempty"`
}

// Validate validates this access Uri raw
func (m *AccessURIRaw) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
