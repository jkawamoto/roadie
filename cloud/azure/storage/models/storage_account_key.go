package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StorageAccountKey An access key for the storage account.
// swagger:model StorageAccountKey
type StorageAccountKey struct {

	// Name of the key.
	// Read Only: true
	KeyName string `json:"keyName,omitempty"`

	// Permissions for the key -- read-only or full permissions.
	// Read Only: true
	Permissions string `json:"permissions,omitempty"`

	// Base 64-encoded value of the key.
	// Read Only: true
	Value string `json:"value,omitempty"`
}

// Validate validates this storage account key
func (m *StorageAccountKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var storageAccountKeyTypePermissionsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Read","Full"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageAccountKeyTypePermissionsPropEnum = append(storageAccountKeyTypePermissionsPropEnum, v)
	}
}

const (
	// StorageAccountKeyPermissionsRead captures enum value "Read"
	StorageAccountKeyPermissionsRead string = "Read"
	// StorageAccountKeyPermissionsFull captures enum value "Full"
	StorageAccountKeyPermissionsFull string = "Full"
)

// prop value enum
func (m *StorageAccountKey) validatePermissionsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storageAccountKeyTypePermissionsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StorageAccountKey) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	// value enum
	if err := m.validatePermissionsEnum("permissions", "body", m.Permissions); err != nil {
		return err
	}

	return nil
}
