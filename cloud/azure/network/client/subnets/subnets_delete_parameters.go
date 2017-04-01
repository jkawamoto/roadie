package subnets

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

// NewSubnetsDeleteParams creates a new SubnetsDeleteParams object
// with the default values initialized.
func NewSubnetsDeleteParams() *SubnetsDeleteParams {
	var ()
	return &SubnetsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubnetsDeleteParamsWithTimeout creates a new SubnetsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubnetsDeleteParamsWithTimeout(timeout time.Duration) *SubnetsDeleteParams {
	var ()
	return &SubnetsDeleteParams{

		timeout: timeout,
	}
}

// NewSubnetsDeleteParamsWithContext creates a new SubnetsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubnetsDeleteParamsWithContext(ctx context.Context) *SubnetsDeleteParams {
	var ()
	return &SubnetsDeleteParams{

		Context: ctx,
	}
}

/*SubnetsDeleteParams contains all the parameters to send to the API endpoint
for the subnets delete operation typically these are written to a http.Request
*/
type SubnetsDeleteParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*SubnetName
	  The name of the subnet.

	*/
	SubnetName string
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

// WithTimeout adds the timeout to the subnets delete params
func (o *SubnetsDeleteParams) WithTimeout(timeout time.Duration) *SubnetsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subnets delete params
func (o *SubnetsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subnets delete params
func (o *SubnetsDeleteParams) WithContext(ctx context.Context) *SubnetsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subnets delete params
func (o *SubnetsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the subnets delete params
func (o *SubnetsDeleteParams) WithAPIVersion(aPIVersion string) *SubnetsDeleteParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the subnets delete params
func (o *SubnetsDeleteParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the subnets delete params
func (o *SubnetsDeleteParams) WithResourceGroupName(resourceGroupName string) *SubnetsDeleteParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the subnets delete params
func (o *SubnetsDeleteParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubnetName adds the subnetName to the subnets delete params
func (o *SubnetsDeleteParams) WithSubnetName(subnetName string) *SubnetsDeleteParams {
	o.SetSubnetName(subnetName)
	return o
}

// SetSubnetName adds the subnetName to the subnets delete params
func (o *SubnetsDeleteParams) SetSubnetName(subnetName string) {
	o.SubnetName = subnetName
}

// WithSubscriptionID adds the subscriptionID to the subnets delete params
func (o *SubnetsDeleteParams) WithSubscriptionID(subscriptionID string) *SubnetsDeleteParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the subnets delete params
func (o *SubnetsDeleteParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVirtualNetworkName adds the virtualNetworkName to the subnets delete params
func (o *SubnetsDeleteParams) WithVirtualNetworkName(virtualNetworkName string) *SubnetsDeleteParams {
	o.SetVirtualNetworkName(virtualNetworkName)
	return o
}

// SetVirtualNetworkName adds the virtualNetworkName to the subnets delete params
func (o *SubnetsDeleteParams) SetVirtualNetworkName(virtualNetworkName string) {
	o.VirtualNetworkName = virtualNetworkName
}

// WriteToRequest writes these params to a swagger request
func (o *SubnetsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subnetName
	if err := r.SetPathParam("subnetName", o.SubnetName); err != nil {
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
