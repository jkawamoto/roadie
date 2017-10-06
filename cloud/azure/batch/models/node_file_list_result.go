package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// NodeFileListResult The result of listing the files on a compute node, or the files associated with a task on a node.
// swagger:model NodeFileListResult
type NodeFileListResult struct {

	// The URL to get the next set of results.
	OdataNextLink string `json:"odata.nextLink,omitempty"`

	// The list of files.
	Value []*NodeFile `json:"value"`
}

// Validate validates this node file list result
func (m *NodeFileListResult) Validate(formats strfmt.Registry) error {
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

func (m *NodeFileListResult) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	for i := 0; i < len(m.Value); i++ {

		if swag.IsZero(m.Value[i]) { // not required
			continue
		}

		if m.Value[i] != nil {

			if err := m.Value[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
