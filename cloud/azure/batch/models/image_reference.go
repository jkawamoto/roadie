package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ImageReference A reference to an Azure Virtual Machines Marketplace image. To get the list of all imageReferences verified by Azure Batch, see the 'List supported node agent SKUs' operation.
// swagger:model ImageReference
type ImageReference struct {

	// The offer type of the Azure Virtual Machines Marketplace image.
	//
	// For example, UbuntuServer or WindowsServer.
	// Required: true
	Offer *string `json:"offer"`

	// The publisher of the Azure Virtual Machines Marketplace image.
	//
	// For example, Canonical or MicrosoftWindowsServer.
	// Required: true
	Publisher *string `json:"publisher"`

	// The SKU of the Azure Virtual Machines Marketplace image.
	//
	// For example, 14.04.0-LTS or 2012-R2-Datacenter.
	// Required: true
	Sku *string `json:"sku"`

	// The version of the Azure Virtual Machines Marketplace image.
	//
	// A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
	Version string `json:"version,omitempty"`
}

// Validate validates this image reference
func (m *ImageReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOffer(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublisher(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSku(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageReference) validateOffer(formats strfmt.Registry) error {

	if err := validate.Required("offer", "body", m.Offer); err != nil {
		return err
	}

	return nil
}

func (m *ImageReference) validatePublisher(formats strfmt.Registry) error {

	if err := validate.Required("publisher", "body", m.Publisher); err != nil {
		return err
	}

	return nil
}

func (m *ImageReference) validateSku(formats strfmt.Registry) error {

	if err := validate.Required("sku", "body", m.Sku); err != nil {
		return err
	}

	return nil
}
