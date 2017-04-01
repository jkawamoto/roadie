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

	"github.com/jkawamoto/roadie/cloud/azure/network/models"
)

// NewNetworkInterfacesCreateOrUpdateParams creates a new NetworkInterfacesCreateOrUpdateParams object
// with the default values initialized.
func NewNetworkInterfacesCreateOrUpdateParams() *NetworkInterfacesCreateOrUpdateParams {
	var ()
	return &NetworkInterfacesCreateOrUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkInterfacesCreateOrUpdateParamsWithTimeout creates a new NetworkInterfacesCreateOrUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNetworkInterfacesCreateOrUpdateParamsWithTimeout(timeout time.Duration) *NetworkInterfacesCreateOrUpdateParams {
	var ()
	return &NetworkInterfacesCreateOrUpdateParams{

		timeout: timeout,
	}
}

// NewNetworkInterfacesCreateOrUpdateParamsWithContext creates a new NetworkInterfacesCreateOrUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewNetworkInterfacesCreateOrUpdateParamsWithContext(ctx context.Context) *NetworkInterfacesCreateOrUpdateParams {
	var ()
	return &NetworkInterfacesCreateOrUpdateParams{

		Context: ctx,
	}
}

/*NetworkInterfacesCreateOrUpdateParams contains all the parameters to send to the API endpoint
for the network interfaces create or update operation typically these are written to a http.Request
*/
type NetworkInterfacesCreateOrUpdateParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*NetworkInterfaceName
	  The name of the network interface.

	*/
	NetworkInterfaceName string
	/*Parameters
	  Parameters supplied to the create/update NetworkInterface operation

	*/
	Parameters *models.NetworkInterface
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

// WithTimeout adds the timeout to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) WithTimeout(timeout time.Duration) *NetworkInterfacesCreateOrUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) WithContext(ctx context.Context) *NetworkInterfacesCreateOrUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) WithAPIVersion(aPIVersion string) *NetworkInterfacesCreateOrUpdateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithNetworkInterfaceName adds the networkInterfaceName to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) WithNetworkInterfaceName(networkInterfaceName string) *NetworkInterfacesCreateOrUpdateParams {
	o.SetNetworkInterfaceName(networkInterfaceName)
	return o
}

// SetNetworkInterfaceName adds the networkInterfaceName to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) SetNetworkInterfaceName(networkInterfaceName string) {
	o.NetworkInterfaceName = networkInterfaceName
}

// WithParameters adds the parameters to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) WithParameters(parameters *models.NetworkInterface) *NetworkInterfacesCreateOrUpdateParams {
	o.SetParameters(parameters)
	return o
}

// SetParameters adds the parameters to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) SetParameters(parameters *models.NetworkInterface) {
	o.Parameters = parameters
}

// WithResourceGroupName adds the resourceGroupName to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) WithResourceGroupName(resourceGroupName string) *NetworkInterfacesCreateOrUpdateParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) WithSubscriptionID(subscriptionID string) *NetworkInterfacesCreateOrUpdateParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the network interfaces create or update params
func (o *NetworkInterfacesCreateOrUpdateParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkInterfacesCreateOrUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param networkInterfaceName
	if err := r.SetPathParam("networkInterfaceName", o.NetworkInterfaceName); err != nil {
		return err
	}

	if o.Parameters == nil {
		o.Parameters = new(models.NetworkInterface)
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
