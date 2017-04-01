package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// PublicIPAddressDNSSettings Contains FQDN of the DNS record associated with the public IP address
// swagger:model PublicIPAddressDnsSettings
type PublicIPAddressDNSSettings struct {

	// Gets or sets the Domain name label.The concatenation of the domain name label and the regionalized DNS zone make up the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel string `json:"domainNameLabel,omitempty"`

	// Gets the FQDN, Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
	Fqdn string `json:"fqdn,omitempty"`

	// Gets or Sests the Reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn string `json:"reverseFqdn,omitempty"`
}

// Validate validates this public IP address Dns settings
func (m *PublicIPAddressDNSSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
