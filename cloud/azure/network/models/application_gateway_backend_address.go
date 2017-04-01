package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ApplicationGatewayBackendAddress Backend Address of application gateway
// swagger:model ApplicationGatewayBackendAddress
type ApplicationGatewayBackendAddress struct {

	// Gets or sets the dns name
	Fqdn string `json:"fqdn,omitempty"`

	// Gets or sets the ip address
	IPAddress string `json:"ipAddress,omitempty"`
}

// Validate validates this application gateway backend address
func (m *ApplicationGatewayBackendAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
