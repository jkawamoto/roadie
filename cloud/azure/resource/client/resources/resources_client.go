package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new resources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for resources API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ResourcesCheckExistence Checks whether a resource exists.
*/
func (a *Client) ResourcesCheckExistence(params *ResourcesCheckExistenceParams, authInfo runtime.ClientAuthInfoWriter) (*ResourcesCheckExistenceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourcesCheckExistenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Resources_CheckExistence",
		Method:             "HEAD",
		PathPattern:        "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourcesCheckExistenceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResourcesCheckExistenceNoContent), nil

}

/*
ResourcesCheckExistenceByID Checks by ID whether a resource exists.
*/
func (a *Client) ResourcesCheckExistenceByID(params *ResourcesCheckExistenceByIDParams, authInfo runtime.ClientAuthInfoWriter) (*ResourcesCheckExistenceByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourcesCheckExistenceByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Resources_CheckExistenceById",
		Method:             "HEAD",
		PathPattern:        "/{resourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourcesCheckExistenceByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResourcesCheckExistenceByIDNoContent), nil

}

/*
ResourcesCreateOrUpdate Creates a resource.
*/
func (a *Client) ResourcesCreateOrUpdate(params *ResourcesCreateOrUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*ResourcesCreateOrUpdateOK, *ResourcesCreateOrUpdateCreated, *ResourcesCreateOrUpdateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourcesCreateOrUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Resources_CreateOrUpdate",
		Method:             "PUT",
		PathPattern:        "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourcesCreateOrUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ResourcesCreateOrUpdateOK:
		return value, nil, nil, nil
	case *ResourcesCreateOrUpdateCreated:
		return nil, value, nil, nil
	case *ResourcesCreateOrUpdateAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
ResourcesCreateOrUpdateByID Create a resource by ID.
*/
func (a *Client) ResourcesCreateOrUpdateByID(params *ResourcesCreateOrUpdateByIDParams, authInfo runtime.ClientAuthInfoWriter) (*ResourcesCreateOrUpdateByIDOK, *ResourcesCreateOrUpdateByIDCreated, *ResourcesCreateOrUpdateByIDAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourcesCreateOrUpdateByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Resources_CreateOrUpdateById",
		Method:             "PUT",
		PathPattern:        "/{resourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourcesCreateOrUpdateByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ResourcesCreateOrUpdateByIDOK:
		return value, nil, nil, nil
	case *ResourcesCreateOrUpdateByIDCreated:
		return nil, value, nil, nil
	case *ResourcesCreateOrUpdateByIDAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
ResourcesDelete Deletes a resource.
*/
func (a *Client) ResourcesDelete(params *ResourcesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*ResourcesDeleteOK, *ResourcesDeleteAccepted, *ResourcesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourcesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Resources_Delete",
		Method:             "DELETE",
		PathPattern:        "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourcesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ResourcesDeleteOK:
		return value, nil, nil, nil
	case *ResourcesDeleteAccepted:
		return nil, value, nil, nil
	case *ResourcesDeleteNoContent:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
ResourcesDeleteByID Deletes a resource by ID.
*/
func (a *Client) ResourcesDeleteByID(params *ResourcesDeleteByIDParams, authInfo runtime.ClientAuthInfoWriter) (*ResourcesDeleteByIDOK, *ResourcesDeleteByIDAccepted, *ResourcesDeleteByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourcesDeleteByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Resources_DeleteById",
		Method:             "DELETE",
		PathPattern:        "/{resourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourcesDeleteByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ResourcesDeleteByIDOK:
		return value, nil, nil, nil
	case *ResourcesDeleteByIDAccepted:
		return nil, value, nil, nil
	case *ResourcesDeleteByIDNoContent:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
ResourcesGet Gets a resource.
*/
func (a *Client) ResourcesGet(params *ResourcesGetParams, authInfo runtime.ClientAuthInfoWriter) (*ResourcesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourcesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Resources_Get",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourcesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResourcesGetOK), nil

}

/*
ResourcesGetByID Gets a resource by ID.
*/
func (a *Client) ResourcesGetByID(params *ResourcesGetByIDParams, authInfo runtime.ClientAuthInfoWriter) (*ResourcesGetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourcesGetByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Resources_GetById",
		Method:             "GET",
		PathPattern:        "/{resourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourcesGetByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResourcesGetByIDOK), nil

}

/*
ResourcesList Get all the resources in a subscription.
*/
func (a *Client) ResourcesList(params *ResourcesListParams, authInfo runtime.ClientAuthInfoWriter) (*ResourcesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourcesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Resources_List",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourcesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResourcesListOK), nil

}

/*
ResourcesMoveResources moves resources from one resource group to another resource group

The resources to move must be in the same source resource group. The target resource group may be in a different subscription. When moving resources, both the source group and the target group are locked for the duration of the operation. Write and delete operations are blocked on the groups until the move completes.
*/
func (a *Client) ResourcesMoveResources(params *ResourcesMoveResourcesParams, authInfo runtime.ClientAuthInfoWriter) (*ResourcesMoveResourcesAccepted, *ResourcesMoveResourcesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourcesMoveResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Resources_MoveResources",
		Method:             "POST",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{sourceResourceGroupName}/moveResources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourcesMoveResourcesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ResourcesMoveResourcesAccepted:
		return value, nil, nil
	case *ResourcesMoveResourcesNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
