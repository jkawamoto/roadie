package availability_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new availability sets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for availability sets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AvailabilitySetsCreateOrUpdate Create or update an availability set.
*/
func (a *Client) AvailabilitySetsCreateOrUpdate(params *AvailabilitySetsCreateOrUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*AvailabilitySetsCreateOrUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAvailabilitySetsCreateOrUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AvailabilitySets_CreateOrUpdate",
		Method:             "PUT",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AvailabilitySetsCreateOrUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AvailabilitySetsCreateOrUpdateOK), nil

}

/*
AvailabilitySetsDelete Delete an availability set.
*/
func (a *Client) AvailabilitySetsDelete(params *AvailabilitySetsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*AvailabilitySetsDeleteOK, *AvailabilitySetsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAvailabilitySetsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AvailabilitySets_Delete",
		Method:             "DELETE",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AvailabilitySetsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *AvailabilitySetsDeleteOK:
		return value, nil, nil
	case *AvailabilitySetsDeleteNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
AvailabilitySetsGet Retrieves information about an availability set.
*/
func (a *Client) AvailabilitySetsGet(params *AvailabilitySetsGetParams, authInfo runtime.ClientAuthInfoWriter) (*AvailabilitySetsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAvailabilitySetsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AvailabilitySets_Get",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AvailabilitySetsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AvailabilitySetsGetOK), nil

}

/*
AvailabilitySetsList Lists all availability sets in a resource group.
*/
func (a *Client) AvailabilitySetsList(params *AvailabilitySetsListParams, authInfo runtime.ClientAuthInfoWriter) (*AvailabilitySetsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAvailabilitySetsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AvailabilitySets_List",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AvailabilitySetsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AvailabilitySetsListOK), nil

}

/*
AvailabilitySetsListAvailableSizes Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing availability set.
*/
func (a *Client) AvailabilitySetsListAvailableSizes(params *AvailabilitySetsListAvailableSizesParams, authInfo runtime.ClientAuthInfoWriter) (*AvailabilitySetsListAvailableSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAvailabilitySetsListAvailableSizesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AvailabilitySets_ListAvailableSizes",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}/vmSizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AvailabilitySetsListAvailableSizesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AvailabilitySetsListAvailableSizesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}