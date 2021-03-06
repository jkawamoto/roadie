package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// NodeFile Information about a file or directory on a compute node.
// swagger:model NodeFile
type NodeFile struct {

	// Whether the object represents a directory.
	IsDirectory bool `json:"isDirectory,omitempty"`

	// The file path.
	Name string `json:"name,omitempty"`

	// The file properties.
	Properties *FileProperties `json:"properties,omitempty"`

	// The URL of the file.
	URL string `json:"url,omitempty"`
}

// Validate validates this node file
func (m *NodeFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeFile) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {

		if err := m.Properties.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
