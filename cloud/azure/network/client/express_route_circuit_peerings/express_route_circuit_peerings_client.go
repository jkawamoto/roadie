package express_route_circuit_peerings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new express route circuit peerings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for express route circuit peerings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ExpressRouteCircuitPeeringsCreateOrUpdate The Put Pering operation creates/updates an peering in the specified ExpressRouteCircuits
*/
func (a *Client) ExpressRouteCircuitPeeringsCreateOrUpdate(params *ExpressRouteCircuitPeeringsCreateOrUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*ExpressRouteCircuitPeeringsCreateOrUpdateOK, *ExpressRouteCircuitPeeringsCreateOrUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExpressRouteCircuitPeeringsCreateOrUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ExpressRouteCircuitPeerings_CreateOrUpdate",
		Method:             "PUT",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExpressRouteCircuitPeeringsCreateOrUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ExpressRouteCircuitPeeringsCreateOrUpdateOK:
		return value, nil, nil
	case *ExpressRouteCircuitPeeringsCreateOrUpdateCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
ExpressRouteCircuitPeeringsDelete The delete peering operation deletes the specified peering from the ExpressRouteCircuit.
*/
func (a *Client) ExpressRouteCircuitPeeringsDelete(params *ExpressRouteCircuitPeeringsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*ExpressRouteCircuitPeeringsDeleteOK, *ExpressRouteCircuitPeeringsDeleteAccepted, *ExpressRouteCircuitPeeringsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExpressRouteCircuitPeeringsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ExpressRouteCircuitPeerings_Delete",
		Method:             "DELETE",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExpressRouteCircuitPeeringsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ExpressRouteCircuitPeeringsDeleteOK:
		return value, nil, nil, nil
	case *ExpressRouteCircuitPeeringsDeleteAccepted:
		return nil, value, nil, nil
	case *ExpressRouteCircuitPeeringsDeleteNoContent:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
ExpressRouteCircuitPeeringsGet The GET peering operation retrieves the specified authorization from the ExpressRouteCircuit.
*/
func (a *Client) ExpressRouteCircuitPeeringsGet(params *ExpressRouteCircuitPeeringsGetParams, authInfo runtime.ClientAuthInfoWriter) (*ExpressRouteCircuitPeeringsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExpressRouteCircuitPeeringsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ExpressRouteCircuitPeerings_Get",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExpressRouteCircuitPeeringsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExpressRouteCircuitPeeringsGetOK), nil

}

/*
ExpressRouteCircuitPeeringsList The List peering operation retrieves all the peerings in an ExpressRouteCircuit.
*/
func (a *Client) ExpressRouteCircuitPeeringsList(params *ExpressRouteCircuitPeeringsListParams, authInfo runtime.ClientAuthInfoWriter) (*ExpressRouteCircuitPeeringsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExpressRouteCircuitPeeringsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ExpressRouteCircuitPeerings_List",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExpressRouteCircuitPeeringsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExpressRouteCircuitPeeringsListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
