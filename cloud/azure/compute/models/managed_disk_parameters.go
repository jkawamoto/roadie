package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// ManagedDiskParameters The parameters of a managed disk.
// swagger:model ManagedDiskParameters
type ManagedDiskParameters struct {
	SubResource

	// The Storage Account type.
	StorageAccountType StorageAccountType `json:"storageAccountType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ManagedDiskParameters) UnmarshalJSON(raw []byte) error {
	var data struct {
		StorageAccountType StorageAccountType `json:"storageAccountType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.StorageAccountType = data.StorageAccountType

	var aO0 SubResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SubResource = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ManagedDiskParameters) MarshalJSON() ([]byte, error) {
	var _parts [][]byte
	var data struct {
		StorageAccountType StorageAccountType `json:"storageAccountType,omitempty"`
	}

	data.StorageAccountType = m.StorageAccountType

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	aO0, err := swag.WriteJSON(m.SubResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this managed disk parameters
func (m *ManagedDiskParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.SubResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageAccountType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagedDiskParameters) validateStorageAccountType(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageAccountType) { // not required
		return nil
	}

	if err := m.StorageAccountType.Validate(formats); err != nil {
		return err
	}

	return nil
}
