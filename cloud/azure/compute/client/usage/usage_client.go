package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new usage API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for usage API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
UsageList Gets, for the specified location, the current compute resource usage information as well as the limits for compute resources under the subscription.
*/
func (a *Client) UsageList(params *UsageListParams, authInfo runtime.ClientAuthInfoWriter) (*UsageListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsageListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Usage_List",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/usages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsageListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UsageListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}