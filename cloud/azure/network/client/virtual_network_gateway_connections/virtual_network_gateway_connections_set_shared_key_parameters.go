package virtual_network_gateway_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/roadie/cloud/azure/network/models"
)

// NewVirtualNetworkGatewayConnectionsSetSharedKeyParams creates a new VirtualNetworkGatewayConnectionsSetSharedKeyParams object
// with the default values initialized.
func NewVirtualNetworkGatewayConnectionsSetSharedKeyParams() *VirtualNetworkGatewayConnectionsSetSharedKeyParams {
	var ()
	return &VirtualNetworkGatewayConnectionsSetSharedKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualNetworkGatewayConnectionsSetSharedKeyParamsWithTimeout creates a new VirtualNetworkGatewayConnectionsSetSharedKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualNetworkGatewayConnectionsSetSharedKeyParamsWithTimeout(timeout time.Duration) *VirtualNetworkGatewayConnectionsSetSharedKeyParams {
	var ()
	return &VirtualNetworkGatewayConnectionsSetSharedKeyParams{

		timeout: timeout,
	}
}

// NewVirtualNetworkGatewayConnectionsSetSharedKeyParamsWithContext creates a new VirtualNetworkGatewayConnectionsSetSharedKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualNetworkGatewayConnectionsSetSharedKeyParamsWithContext(ctx context.Context) *VirtualNetworkGatewayConnectionsSetSharedKeyParams {
	var ()
	return &VirtualNetworkGatewayConnectionsSetSharedKeyParams{

		Context: ctx,
	}
}

/*VirtualNetworkGatewayConnectionsSetSharedKeyParams contains all the parameters to send to the API endpoint
for the virtual network gateway connections set shared key operation typically these are written to a http.Request
*/
type VirtualNetworkGatewayConnectionsSetSharedKeyParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*Parameters
	  Parameters supplied to the Begin Set Virtual Network Gateway conection Shared key operation throughNetwork resource provider.

	*/
	Parameters *models.ConnectionSharedKey
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*SubscriptionID
	  Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string
	/*VirtualNetworkGatewayConnectionName
	  The virtual network gateway connection name.

	*/
	VirtualNetworkGatewayConnectionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) WithTimeout(timeout time.Duration) *VirtualNetworkGatewayConnectionsSetSharedKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) WithContext(ctx context.Context) *VirtualNetworkGatewayConnectionsSetSharedKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) WithAPIVersion(aPIVersion string) *VirtualNetworkGatewayConnectionsSetSharedKeyParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithParameters adds the parameters to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) WithParameters(parameters *models.ConnectionSharedKey) *VirtualNetworkGatewayConnectionsSetSharedKeyParams {
	o.SetParameters(parameters)
	return o
}

// SetParameters adds the parameters to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) SetParameters(parameters *models.ConnectionSharedKey) {
	o.Parameters = parameters
}

// WithResourceGroupName adds the resourceGroupName to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) WithResourceGroupName(resourceGroupName string) *VirtualNetworkGatewayConnectionsSetSharedKeyParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) WithSubscriptionID(subscriptionID string) *VirtualNetworkGatewayConnectionsSetSharedKeyParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVirtualNetworkGatewayConnectionName adds the virtualNetworkGatewayConnectionName to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) WithVirtualNetworkGatewayConnectionName(virtualNetworkGatewayConnectionName string) *VirtualNetworkGatewayConnectionsSetSharedKeyParams {
	o.SetVirtualNetworkGatewayConnectionName(virtualNetworkGatewayConnectionName)
	return o
}

// SetVirtualNetworkGatewayConnectionName adds the virtualNetworkGatewayConnectionName to the virtual network gateway connections set shared key params
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) SetVirtualNetworkGatewayConnectionName(virtualNetworkGatewayConnectionName string) {
	o.VirtualNetworkGatewayConnectionName = virtualNetworkGatewayConnectionName
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualNetworkGatewayConnectionsSetSharedKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// query param api-version
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("api-version", qAPIVersion); err != nil {
			return err
		}
	}

	if o.Parameters == nil {
		o.Parameters = new(models.ConnectionSharedKey)
	}

	if err := r.SetBodyParam(o.Parameters); err != nil {
		return err
	}

	// path param resourceGroupName
	if err := r.SetPathParam("resourceGroupName", o.ResourceGroupName); err != nil {
		return err
	}

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID); err != nil {
		return err
	}

	// path param virtualNetworkGatewayConnectionName
	if err := r.SetPathParam("virtualNetworkGatewayConnectionName", o.VirtualNetworkGatewayConnectionName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
