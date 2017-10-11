package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new accounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AccountListNodeAgentSkus lists all node agent s k us supported by the azure batch service
*/
func (a *Client) AccountListNodeAgentSkus(params *AccountListNodeAgentSkusParams) (*AccountListNodeAgentSkusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountListNodeAgentSkusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Account_ListNodeAgentSkus",
		Method:             "GET",
		PathPattern:        "/nodeagentskus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json; odata=minimalmetadata"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountListNodeAgentSkusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AccountListNodeAgentSkusOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}