package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new images API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for images API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ImagesCreateOrUpdate Create or update an image.
*/
func (a *Client) ImagesCreateOrUpdate(params *ImagesCreateOrUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*ImagesCreateOrUpdateOK, *ImagesCreateOrUpdateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesCreateOrUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Images_CreateOrUpdate",
		Method:             "PUT",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesCreateOrUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ImagesCreateOrUpdateOK:
		return value, nil, nil
	case *ImagesCreateOrUpdateCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
ImagesDelete Deletes an Image.
*/
func (a *Client) ImagesDelete(params *ImagesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*ImagesDeleteOK, *ImagesDeleteAccepted, *ImagesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Images_Delete",
		Method:             "DELETE",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ImagesDeleteOK:
		return value, nil, nil, nil
	case *ImagesDeleteAccepted:
		return nil, value, nil, nil
	case *ImagesDeleteNoContent:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
ImagesGet Gets an image.
*/
func (a *Client) ImagesGet(params *ImagesGetParams, authInfo runtime.ClientAuthInfoWriter) (*ImagesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Images_Get",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ImagesGetOK), nil

}

/*
ImagesList Gets the list of Images in the subscription. Use nextLink property in the response to get the next page of Images. Do this till nextLink is not null to fetch all the Images.
*/
func (a *Client) ImagesList(params *ImagesListParams, authInfo runtime.ClientAuthInfoWriter) (*ImagesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Images_List",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ImagesListOK), nil

}

/*
ImagesListByResourceGroup Gets the list of images under a resource group.
*/
func (a *Client) ImagesListByResourceGroup(params *ImagesListByResourceGroupParams, authInfo runtime.ClientAuthInfoWriter) (*ImagesListByResourceGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesListByResourceGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Images_ListByResourceGroup",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesListByResourceGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ImagesListByResourceGroupOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
