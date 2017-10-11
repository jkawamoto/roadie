package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// KeyVaultKeyReference Describes a reference to Key Vault Key
// swagger:model KeyVaultKeyReference
type KeyVaultKeyReference struct {

	// The URL referencing a key in a Key Vault.
	// Required: true
	KeyURL *string `json:"keyUrl"`

	// The relative URL of the Key Vault containing the key.
	// Required: true
	SourceVault *SubResource `json:"sourceVault"`
}

// Validate validates this key vault key reference
func (m *KeyVaultKeyReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSourceVault(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyVaultKeyReference) validateKeyURL(formats strfmt.Registry) error {

	if err := validate.Required("keyUrl", "body", m.KeyURL); err != nil {
		return err
	}

	return nil
}

func (m *KeyVaultKeyReference) validateSourceVault(formats strfmt.Registry) error {

	if err := validate.Required("sourceVault", "body", m.SourceVault); err != nil {
		return err
	}

	if m.SourceVault != nil {

		if err := m.SourceVault.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}