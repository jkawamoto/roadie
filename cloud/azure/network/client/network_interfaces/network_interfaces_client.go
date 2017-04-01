package network_interfaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new network interfaces API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network interfaces API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
NetworkInterfacesCreateOrUpdate The Put NetworkInterface operation creates/updates a networkInterface
*/
func (a *Client) NetworkInterfacesCreateOrUpdate(params *NetworkInterfacesCreateOrUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*NetworkInterfacesCreateOrUpdateOK, *NetworkInterfacesCreateOrUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkInterfacesCreateOrUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkInterfaces_CreateOrUpdate",
		Method:             "PUT",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NetworkInterfacesCreateOrUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *NetworkInterfacesCreateOrUpdateOK:
		return value, nil, nil
	case *NetworkInterfacesCreateOrUpdateCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
NetworkInterfacesDelete The delete netwokInterface operation deletes the specified netwokInterface.
*/
func (a *Client) NetworkInterfacesDelete(params *NetworkInterfacesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*NetworkInterfacesDeleteOK, *NetworkInterfacesDeleteAccepted, *NetworkInterfacesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkInterfacesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkInterfaces_Delete",
		Method:             "DELETE",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NetworkInterfacesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *NetworkInterfacesDeleteOK:
		return value, nil, nil, nil
	case *NetworkInterfacesDeleteAccepted:
		return nil, value, nil, nil
	case *NetworkInterfacesDeleteNoContent:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
NetworkInterfacesGet The Get ntework interface operation retreives information about the specified network interface.
*/
func (a *Client) NetworkInterfacesGet(params *NetworkInterfacesGetParams, authInfo runtime.ClientAuthInfoWriter) (*NetworkInterfacesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkInterfacesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkInterfaces_Get",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NetworkInterfacesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkInterfacesGetOK), nil

}

/*
NetworkInterfacesGetVirtualMachineScaleSetNetworkInterface The Get ntework interface operation retreives information about the specified network interface in a virtual machine scale set.
*/
func (a *Client) NetworkInterfacesGetVirtualMachineScaleSetNetworkInterface(params *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams, authInfo runtime.ClientAuthInfoWriter) (*NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkInterfaces_GetVirtualMachineScaleSetNetworkInterface",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOK), nil

}

/*
NetworkInterfacesList The List networkInterfaces opertion retrieves all the networkInterfaces in a resource group.
*/
func (a *Client) NetworkInterfacesList(params *NetworkInterfacesListParams, authInfo runtime.ClientAuthInfoWriter) (*NetworkInterfacesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkInterfacesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkInterfaces_List",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NetworkInterfacesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkInterfacesListOK), nil

}

/*
NetworkInterfacesListAll The List networkInterfaces opertion retrieves all the networkInterfaces in a subscription.
*/
func (a *Client) NetworkInterfacesListAll(params *NetworkInterfacesListAllParams, authInfo runtime.ClientAuthInfoWriter) (*NetworkInterfacesListAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkInterfacesListAllParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkInterfaces_ListAll",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkInterfaces",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NetworkInterfacesListAllReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkInterfacesListAllOK), nil

}

/*
NetworkInterfacesListVirtualMachineScaleSetNetworkInterfaces The list network interface operation retrieves information about all network interfaces in a virtual machine scale set.
*/
func (a *Client) NetworkInterfacesListVirtualMachineScaleSetNetworkInterfaces(params *NetworkInterfacesListVirtualMachineScaleSetNetworkInterfacesParams, authInfo runtime.ClientAuthInfoWriter) (*NetworkInterfacesListVirtualMachineScaleSetNetworkInterfacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkInterfacesListVirtualMachineScaleSetNetworkInterfacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkInterfaces_ListVirtualMachineScaleSetNetworkInterfaces",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/networkInterfaces",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NetworkInterfacesListVirtualMachineScaleSetNetworkInterfacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkInterfacesListVirtualMachineScaleSetNetworkInterfacesOK), nil

}

/*
NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfaces The list network interface operation retrieves information about all network interfaces in a virtual machine from a virtual machine scale set.
*/
func (a *Client) NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfaces(params *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams, authInfo runtime.ClientAuthInfoWriter) (*NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkInterfaces_ListVirtualMachineScaleSetVMNetworkInterfaces",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
