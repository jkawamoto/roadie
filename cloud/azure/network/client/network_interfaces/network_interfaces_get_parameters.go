package network_interfaces

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

// NewNetworkInterfacesGetParams creates a new NetworkInterfacesGetParams object
// with the default values initialized.
func NewNetworkInterfacesGetParams() *NetworkInterfacesGetParams {
	var ()
	return &NetworkInterfacesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkInterfacesGetParamsWithTimeout creates a new NetworkInterfacesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNetworkInterfacesGetParamsWithTimeout(timeout time.Duration) *NetworkInterfacesGetParams {
	var ()
	return &NetworkInterfacesGetParams{

		timeout: timeout,
	}
}

// NewNetworkInterfacesGetParamsWithContext creates a new NetworkInterfacesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewNetworkInterfacesGetParamsWithContext(ctx context.Context) *NetworkInterfacesGetParams {
	var ()
	return &NetworkInterfacesGetParams{

		Context: ctx,
	}
}

/*NetworkInterfacesGetParams contains all the parameters to send to the API endpoint
for the network interfaces get operation typically these are written to a http.Request
*/
type NetworkInterfacesGetParams struct {

	/*NrDollarExpand
	  expand references resources.

	*/
	DollarExpand *string
	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*NetworkInterfaceName
	  The name of the network interface.

	*/
	NetworkInterfaceName string
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*SubscriptionID
	  Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the network interfaces get params
func (o *NetworkInterfacesGetParams) WithTimeout(timeout time.Duration) *NetworkInterfacesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network interfaces get params
func (o *NetworkInterfacesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network interfaces get params
func (o *NetworkInterfacesGetParams) WithContext(ctx context.Context) *NetworkInterfacesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network interfaces get params
func (o *NetworkInterfacesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDollarExpand adds the dollarExpand to the network interfaces get params
func (o *NetworkInterfacesGetParams) WithDollarExpand(dollarExpand *string) *NetworkInterfacesGetParams {
	o.SetDollarExpand(dollarExpand)
	return o
}

// SetDollarExpand adds the dollarExpand to the network interfaces get params
func (o *NetworkInterfacesGetParams) SetDollarExpand(dollarExpand *string) {
	o.DollarExpand = dollarExpand
}

// WithAPIVersion adds the aPIVersion to the network interfaces get params
func (o *NetworkInterfacesGetParams) WithAPIVersion(aPIVersion string) *NetworkInterfacesGetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the network interfaces get params
func (o *NetworkInterfacesGetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithNetworkInterfaceName adds the networkInterfaceName to the network interfaces get params
func (o *NetworkInterfacesGetParams) WithNetworkInterfaceName(networkInterfaceName string) *NetworkInterfacesGetParams {
	o.SetNetworkInterfaceName(networkInterfaceName)
	return o
}

// SetNetworkInterfaceName adds the networkInterfaceName to the network interfaces get params
func (o *NetworkInterfacesGetParams) SetNetworkInterfaceName(networkInterfaceName string) {
	o.NetworkInterfaceName = networkInterfaceName
}

// WithResourceGroupName adds the resourceGroupName to the network interfaces get params
func (o *NetworkInterfacesGetParams) WithResourceGroupName(resourceGroupName string) *NetworkInterfacesGetParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the network interfaces get params
func (o *NetworkInterfacesGetParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the network interfaces get params
func (o *NetworkInterfacesGetParams) WithSubscriptionID(subscriptionID string) *NetworkInterfacesGetParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the network interfaces get params
func (o *NetworkInterfacesGetParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkInterfacesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.DollarExpand != nil {

		// query param $expand
		var qrNrDollarExpand string
		if o.DollarExpand != nil {
			qrNrDollarExpand = *o.DollarExpand
		}
		qNrDollarExpand := qrNrDollarExpand
		if qNrDollarExpand != "" {
			if err := r.SetQueryParam("$expand", qNrDollarExpand); err != nil {
				return err
			}
		}

	}

	// query param api-version
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("api-version", qAPIVersion); err != nil {
			return err
		}
	}

	// path param networkInterfaceName
	if err := r.SetPathParam("networkInterfaceName", o.NetworkInterfaceName); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
