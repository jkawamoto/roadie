package express_route_circuit_arp_table

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

// NewExpressRouteCircuitsListArpTableParams creates a new ExpressRouteCircuitsListArpTableParams object
// with the default values initialized.
func NewExpressRouteCircuitsListArpTableParams() *ExpressRouteCircuitsListArpTableParams {
	var ()
	return &ExpressRouteCircuitsListArpTableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExpressRouteCircuitsListArpTableParamsWithTimeout creates a new ExpressRouteCircuitsListArpTableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExpressRouteCircuitsListArpTableParamsWithTimeout(timeout time.Duration) *ExpressRouteCircuitsListArpTableParams {
	var ()
	return &ExpressRouteCircuitsListArpTableParams{

		timeout: timeout,
	}
}

// NewExpressRouteCircuitsListArpTableParamsWithContext creates a new ExpressRouteCircuitsListArpTableParams object
// with the default values initialized, and the ability to set a context for a request
func NewExpressRouteCircuitsListArpTableParamsWithContext(ctx context.Context) *ExpressRouteCircuitsListArpTableParams {
	var ()
	return &ExpressRouteCircuitsListArpTableParams{

		Context: ctx,
	}
}

/*ExpressRouteCircuitsListArpTableParams contains all the parameters to send to the API endpoint
for the express route circuits list arp table operation typically these are written to a http.Request
*/
type ExpressRouteCircuitsListArpTableParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*CircuitName
	  The name of the circuit.

	*/
	CircuitName string
	/*DevicePath
	  The path of the device.

	*/
	DevicePath string
	/*PeeringName
	  The name of the peering.

	*/
	PeeringName string
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

// WithTimeout adds the timeout to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) WithTimeout(timeout time.Duration) *ExpressRouteCircuitsListArpTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) WithContext(ctx context.Context) *ExpressRouteCircuitsListArpTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) WithAPIVersion(aPIVersion string) *ExpressRouteCircuitsListArpTableParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithCircuitName adds the circuitName to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) WithCircuitName(circuitName string) *ExpressRouteCircuitsListArpTableParams {
	o.SetCircuitName(circuitName)
	return o
}

// SetCircuitName adds the circuitName to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) SetCircuitName(circuitName string) {
	o.CircuitName = circuitName
}

// WithDevicePath adds the devicePath to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) WithDevicePath(devicePath string) *ExpressRouteCircuitsListArpTableParams {
	o.SetDevicePath(devicePath)
	return o
}

// SetDevicePath adds the devicePath to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) SetDevicePath(devicePath string) {
	o.DevicePath = devicePath
}

// WithPeeringName adds the peeringName to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) WithPeeringName(peeringName string) *ExpressRouteCircuitsListArpTableParams {
	o.SetPeeringName(peeringName)
	return o
}

// SetPeeringName adds the peeringName to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) SetPeeringName(peeringName string) {
	o.PeeringName = peeringName
}

// WithResourceGroupName adds the resourceGroupName to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) WithResourceGroupName(resourceGroupName string) *ExpressRouteCircuitsListArpTableParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) WithSubscriptionID(subscriptionID string) *ExpressRouteCircuitsListArpTableParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the express route circuits list arp table params
func (o *ExpressRouteCircuitsListArpTableParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ExpressRouteCircuitsListArpTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param circuitName
	if err := r.SetPathParam("circuitName", o.CircuitName); err != nil {
		return err
	}

	// path param devicePath
	if err := r.SetPathParam("devicePath", o.DevicePath); err != nil {
		return err
	}

	// path param peeringName
	if err := r.SetPathParam("peeringName", o.PeeringName); err != nil {
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
