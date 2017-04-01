package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// ExpressRouteServiceProviderPropertiesFormat Properties of ExpressRouteServiceProvider
// swagger:model ExpressRouteServiceProviderPropertiesFormat
type ExpressRouteServiceProviderPropertiesFormat struct {

	// Gets or bandwidths offered
	BandwidthsOffered []*ExpressRouteServiceProviderBandwidthsOffered `json:"bandwidthsOffered"`

	// Gets or list of peering locations
	PeeringLocations []string `json:"peeringLocations"`

	// Gets or sets Provisioning state of the resource
	ProvisioningState string `json:"provisioningState,omitempty"`
}

// Validate validates this express route service provider properties format
func (m *ExpressRouteServiceProviderPropertiesFormat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBandwidthsOffered(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePeeringLocations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpressRouteServiceProviderPropertiesFormat) validateBandwidthsOffered(formats strfmt.Registry) error {

	if swag.IsZero(m.BandwidthsOffered) { // not required
		return nil
	}

	for i := 0; i < len(m.BandwidthsOffered); i++ {

		if swag.IsZero(m.BandwidthsOffered[i]) { // not required
			continue
		}

		if m.BandwidthsOffered[i] != nil {

			if err := m.BandwidthsOffered[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ExpressRouteServiceProviderPropertiesFormat) validatePeeringLocations(formats strfmt.Registry) error {

	if swag.IsZero(m.PeeringLocations) { // not required
		return nil
	}

	return nil
}
