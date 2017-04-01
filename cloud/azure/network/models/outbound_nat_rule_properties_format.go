package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// OutboundNatRulePropertiesFormat Outbound NAT pool of the loadbalancer
// swagger:model OutboundNatRulePropertiesFormat
type OutboundNatRulePropertiesFormat struct {

	// Gets or sets the number of outbound ports to be used for SNAT
	AllocatedOutboundPorts int32 `json:"allocatedOutboundPorts,omitempty"`

	// Gets or sets a reference to a pool of DIPs. Outbound traffic is randomly load balanced across IPs in the backend IPs
	// Required: true
	BackendAddressPool *SubResource `json:"backendAddressPool"`

	// Gets or sets Frontend IP addresses of the load balancer
	FrontendIPConfigurations []*SubResource `json:"frontendIPConfigurations"`

	// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
	ProvisioningState string `json:"provisioningState,omitempty"`
}

// Validate validates this outbound nat rule properties format
func (m *OutboundNatRulePropertiesFormat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackendAddressPool(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFrontendIPConfigurations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutboundNatRulePropertiesFormat) validateBackendAddressPool(formats strfmt.Registry) error {

	if err := validate.Required("backendAddressPool", "body", m.BackendAddressPool); err != nil {
		return err
	}

	if m.BackendAddressPool != nil {

		if err := m.BackendAddressPool.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *OutboundNatRulePropertiesFormat) validateFrontendIPConfigurations(formats strfmt.Registry) error {

	if swag.IsZero(m.FrontendIPConfigurations) { // not required
		return nil
	}

	for i := 0; i < len(m.FrontendIPConfigurations); i++ {

		if swag.IsZero(m.FrontendIPConfigurations[i]) { // not required
			continue
		}

		if m.FrontendIPConfigurations[i] != nil {

			if err := m.FrontendIPConfigurations[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
