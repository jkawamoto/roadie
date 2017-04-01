package virtual_networks

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
)

// NewVirtualNetworksDeleteParams creates a new VirtualNetworksDeleteParams object
// with the default values initialized.
func NewVirtualNetworksDeleteParams() *VirtualNetworksDeleteParams {
	var ()
	return &VirtualNetworksDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualNetworksDeleteParamsWithTimeout creates a new VirtualNetworksDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualNetworksDeleteParamsWithTimeout(timeout time.Duration) *VirtualNetworksDeleteParams {
	var ()
	return &VirtualNetworksDeleteParams{

		timeout: timeout,
	}
}

// NewVirtualNetworksDeleteParamsWithContext creates a new VirtualNetworksDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualNetworksDeleteParamsWithContext(ctx context.Context) *VirtualNetworksDeleteParams {
	var ()
	return &VirtualNetworksDeleteParams{

		Context: ctx,
	}
}

/*VirtualNetworksDeleteParams contains all the parameters to send to the API endpoint
for the virtual networks delete operation typically these are written to a http.Request
*/
type VirtualNetworksDeleteParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*SubscriptionID
	  Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string
	/*VirtualNetworkName
	  The name of the virtual network.

	*/
	VirtualNetworkName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) WithTimeout(timeout time.Duration) *VirtualNetworksDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) WithContext(ctx context.Context) *VirtualNetworksDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) WithAPIVersion(aPIVersion string) *VirtualNetworksDeleteParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) WithResourceGroupName(resourceGroupName string) *VirtualNetworksDeleteParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) WithSubscriptionID(subscriptionID string) *VirtualNetworksDeleteParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVirtualNetworkName adds the virtualNetworkName to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) WithVirtualNetworkName(virtualNetworkName string) *VirtualNetworksDeleteParams {
	o.SetVirtualNetworkName(virtualNetworkName)
	return o
}

// SetVirtualNetworkName adds the virtualNetworkName to the virtual networks delete params
func (o *VirtualNetworksDeleteParams) SetVirtualNetworkName(virtualNetworkName string) {
	o.VirtualNetworkName = virtualNetworkName
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualNetworksDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resourceGroupName
	if err := r.SetPathParam("resourceGroupName", o.ResourceGroupName); err != nil {
		return err
	}

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID); err != nil {
		return err
	}

	// path param virtualNetworkName
	if err := r.SetPathParam("virtualNetworkName", o.VirtualNetworkName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
