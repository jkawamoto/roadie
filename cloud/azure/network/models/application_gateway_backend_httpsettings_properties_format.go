package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ApplicationGatewayBackendHTTPSettingsPropertiesFormat Properties of Backend address pool settings of application gateway
// swagger:model ApplicationGatewayBackendHttpSettingsPropertiesFormat
type ApplicationGatewayBackendHTTPSettingsPropertiesFormat struct {

	// Gets or sets the cookie affinity
	CookieBasedAffinity string `json:"cookieBasedAffinity,omitempty"`

	// Gets or sets the port
	Port int32 `json:"port,omitempty"`

	// Gets or sets probe resource of application gateway
	Probe *SubResource `json:"probe,omitempty"`

	// Gets or sets the protocol
	Protocol string `json:"protocol,omitempty"`

	// Gets or sets Provisioning state of the backend http settings resource Updating/Deleting/Failed
	ProvisioningState string `json:"provisioningState,omitempty"`

	// Gets or sets request timeout
	RequestTimeout int32 `json:"requestTimeout,omitempty"`
}

// Validate validates this application gateway backend Http settings properties format
func (m *ApplicationGatewayBackendHTTPSettingsPropertiesFormat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCookieBasedAffinity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProbe(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var applicationGatewayBackendHttpSettingsPropertiesFormatTypeCookieBasedAffinityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Enabled","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationGatewayBackendHttpSettingsPropertiesFormatTypeCookieBasedAffinityPropEnum = append(applicationGatewayBackendHttpSettingsPropertiesFormatTypeCookieBasedAffinityPropEnum, v)
	}
}

const (
	// ApplicationGatewayBackendHTTPSettingsPropertiesFormatCookieBasedAffinityEnabled captures enum value "Enabled"
	ApplicationGatewayBackendHTTPSettingsPropertiesFormatCookieBasedAffinityEnabled string = "Enabled"
	// ApplicationGatewayBackendHTTPSettingsPropertiesFormatCookieBasedAffinityDisabled captures enum value "Disabled"
	ApplicationGatewayBackendHTTPSettingsPropertiesFormatCookieBasedAffinityDisabled string = "Disabled"
)

// prop value enum
func (m *ApplicationGatewayBackendHTTPSettingsPropertiesFormat) validateCookieBasedAffinityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applicationGatewayBackendHttpSettingsPropertiesFormatTypeCookieBasedAffinityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationGatewayBackendHTTPSettingsPropertiesFormat) validateCookieBasedAffinity(formats strfmt.Registry) error {

	if swag.IsZero(m.CookieBasedAffinity) { // not required
		return nil
	}

	// value enum
	if err := m.validateCookieBasedAffinityEnum("cookieBasedAffinity", "body", m.CookieBasedAffinity); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationGatewayBackendHTTPSettingsPropertiesFormat) validateProbe(formats strfmt.Registry) error {

	if swag.IsZero(m.Probe) { // not required
		return nil
	}

	if m.Probe != nil {

		if err := m.Probe.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var applicationGatewayBackendHttpSettingsPropertiesFormatTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Http","Https"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationGatewayBackendHttpSettingsPropertiesFormatTypeProtocolPropEnum = append(applicationGatewayBackendHttpSettingsPropertiesFormatTypeProtocolPropEnum, v)
	}
}

const (
	// ApplicationGatewayBackendHTTPSettingsPropertiesFormatProtocolHTTP captures enum value "Http"
	ApplicationGatewayBackendHTTPSettingsPropertiesFormatProtocolHTTP string = "Http"
	// ApplicationGatewayBackendHTTPSettingsPropertiesFormatProtocolHTTPS captures enum value "Https"
	ApplicationGatewayBackendHTTPSettingsPropertiesFormatProtocolHTTPS string = "Https"
)

// prop value enum
func (m *ApplicationGatewayBackendHTTPSettingsPropertiesFormat) validateProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applicationGatewayBackendHttpSettingsPropertiesFormatTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationGatewayBackendHTTPSettingsPropertiesFormat) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}
