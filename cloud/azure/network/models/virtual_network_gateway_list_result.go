package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// VirtualNetworkGatewayListResult Response for ListVirtualNetworkGateways Api service call
// swagger:model VirtualNetworkGatewayListResult
type VirtualNetworkGatewayListResult struct {

	// Gets the URL to get the next set of results.
	NextLink string `json:"nextLink,omitempty"`

	// Gets List of VirtualNetworkGateways that exists in a resource group
	Value []*VirtualNetworkGateway `json:"value"`
}

// Validate validates this virtual network gateway list result
func (m *VirtualNetworkGatewayListResult) Validate(formats strfmt.Registry) error {
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

func (m *VirtualNetworkGatewayListResult) validateValue(formats strfmt.Registry) error {

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
