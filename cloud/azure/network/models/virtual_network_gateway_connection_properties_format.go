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

// VirtualNetworkGatewayConnectionPropertiesFormat VirtualNeworkGatewayConnection properties
// swagger:model VirtualNetworkGatewayConnectionPropertiesFormat
type VirtualNetworkGatewayConnectionPropertiesFormat struct {

	// The authorizationKey.
	AuthorizationKey string `json:"authorizationKey,omitempty"`

	// Virtual network Gateway connection status
	ConnectionStatus string `json:"connectionStatus,omitempty"`

	// Gateway connection type -Ipsec/Dedicated/VpnClient/Vnet2Vnet
	ConnectionType string `json:"connectionType,omitempty"`

	// The Egress Bytes Transferred in this connection
	EgressBytesTransferred int64 `json:"egressBytesTransferred,omitempty"`

	// EnableBgp Flag
	EnableBgp bool `json:"enableBgp,omitempty"`

	// The Ingress Bytes Transferred in this connection
	IngressBytesTransferred int64 `json:"ingressBytesTransferred,omitempty"`

	// local network gateway2
	LocalNetworkGateway2 *LocalNetworkGateway `json:"localNetworkGateway2,omitempty"`

	// The reference to peerings resource.
	Peer *SubResource `json:"peer,omitempty"`

	// Gets or sets Provisioning state of the VirtualNetworkGatewayConnection resource Updating/Deleting/Failed
	ProvisioningState string `json:"provisioningState,omitempty"`

	// Gets or sets resource guid property of the VirtualNetworkGatewayConnection resource
	ResourceGUID string `json:"resourceGuid,omitempty"`

	// The Routing weight.
	RoutingWeight int32 `json:"routingWeight,omitempty"`

	// The Ipsec share key.
	SharedKey string `json:"sharedKey,omitempty"`

	// virtual network gateway1
	VirtualNetworkGateway1 *VirtualNetworkGateway `json:"virtualNetworkGateway1,omitempty"`

	// virtual network gateway2
	VirtualNetworkGateway2 *VirtualNetworkGateway `json:"virtualNetworkGateway2,omitempty"`
}

// Validate validates this virtual network gateway connection properties format
func (m *VirtualNetworkGatewayConnectionPropertiesFormat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocalNetworkGateway2(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePeer(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVirtualNetworkGateway1(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVirtualNetworkGateway2(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var virtualNetworkGatewayConnectionPropertiesFormatTypeConnectionStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Connecting","Connected","NotConnected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualNetworkGatewayConnectionPropertiesFormatTypeConnectionStatusPropEnum = append(virtualNetworkGatewayConnectionPropertiesFormatTypeConnectionStatusPropEnum, v)
	}
}

const (
	// VirtualNetworkGatewayConnectionPropertiesFormatConnectionStatusUnknown captures enum value "Unknown"
	VirtualNetworkGatewayConnectionPropertiesFormatConnectionStatusUnknown string = "Unknown"
	// VirtualNetworkGatewayConnectionPropertiesFormatConnectionStatusConnecting captures enum value "Connecting"
	VirtualNetworkGatewayConnectionPropertiesFormatConnectionStatusConnecting string = "Connecting"
	// VirtualNetworkGatewayConnectionPropertiesFormatConnectionStatusConnected captures enum value "Connected"
	VirtualNetworkGatewayConnectionPropertiesFormatConnectionStatusConnected string = "Connected"
	// VirtualNetworkGatewayConnectionPropertiesFormatConnectionStatusNotConnected captures enum value "NotConnected"
	VirtualNetworkGatewayConnectionPropertiesFormatConnectionStatusNotConnected string = "NotConnected"
)

// prop value enum
func (m *VirtualNetworkGatewayConnectionPropertiesFormat) validateConnectionStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, virtualNetworkGatewayConnectionPropertiesFormatTypeConnectionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VirtualNetworkGatewayConnectionPropertiesFormat) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStatusEnum("connectionStatus", "body", m.ConnectionStatus); err != nil {
		return err
	}

	return nil
}

var virtualNetworkGatewayConnectionPropertiesFormatTypeConnectionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPsec","Vnet2Vnet","ExpressRoute","VPNClient"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualNetworkGatewayConnectionPropertiesFormatTypeConnectionTypePropEnum = append(virtualNetworkGatewayConnectionPropertiesFormatTypeConnectionTypePropEnum, v)
	}
}

const (
	// VirtualNetworkGatewayConnectionPropertiesFormatConnectionTypeIpsec captures enum value "IPsec"
	VirtualNetworkGatewayConnectionPropertiesFormatConnectionTypeIpsec string = "IPsec"
	// VirtualNetworkGatewayConnectionPropertiesFormatConnectionTypeVnet2Vnet captures enum value "Vnet2Vnet"
	VirtualNetworkGatewayConnectionPropertiesFormatConnectionTypeVnet2Vnet string = "Vnet2Vnet"
	// VirtualNetworkGatewayConnectionPropertiesFormatConnectionTypeExpressRoute captures enum value "ExpressRoute"
	VirtualNetworkGatewayConnectionPropertiesFormatConnectionTypeExpressRoute string = "ExpressRoute"
	// VirtualNetworkGatewayConnectionPropertiesFormatConnectionTypeVPNClient captures enum value "VPNClient"
	VirtualNetworkGatewayConnectionPropertiesFormatConnectionTypeVPNClient string = "VPNClient"
)

// prop value enum
func (m *VirtualNetworkGatewayConnectionPropertiesFormat) validateConnectionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, virtualNetworkGatewayConnectionPropertiesFormatTypeConnectionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VirtualNetworkGatewayConnectionPropertiesFormat) validateConnectionType(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionTypeEnum("connectionType", "body", m.ConnectionType); err != nil {
		return err
	}

	return nil
}

func (m *VirtualNetworkGatewayConnectionPropertiesFormat) validateLocalNetworkGateway2(formats strfmt.Registry) error {

	if swag.IsZero(m.LocalNetworkGateway2) { // not required
		return nil
	}

	if m.LocalNetworkGateway2 != nil {

		if err := m.LocalNetworkGateway2.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VirtualNetworkGatewayConnectionPropertiesFormat) validatePeer(formats strfmt.Registry) error {

	if swag.IsZero(m.Peer) { // not required
		return nil
	}

	if m.Peer != nil {

		if err := m.Peer.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VirtualNetworkGatewayConnectionPropertiesFormat) validateVirtualNetworkGateway1(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualNetworkGateway1) { // not required
		return nil
	}

	if m.VirtualNetworkGateway1 != nil {

		if err := m.VirtualNetworkGateway1.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VirtualNetworkGatewayConnectionPropertiesFormat) validateVirtualNetworkGateway2(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualNetworkGateway2) { // not required
		return nil
	}

	if m.VirtualNetworkGateway2 != nil {

		if err := m.VirtualNetworkGateway2.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
