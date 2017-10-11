package availability_sets

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

// NewAvailabilitySetsDeleteParams creates a new AvailabilitySetsDeleteParams object
// with the default values initialized.
func NewAvailabilitySetsDeleteParams() *AvailabilitySetsDeleteParams {
	var ()
	return &AvailabilitySetsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAvailabilitySetsDeleteParamsWithTimeout creates a new AvailabilitySetsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAvailabilitySetsDeleteParamsWithTimeout(timeout time.Duration) *AvailabilitySetsDeleteParams {
	var ()
	return &AvailabilitySetsDeleteParams{

		timeout: timeout,
	}
}

// NewAvailabilitySetsDeleteParamsWithContext creates a new AvailabilitySetsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewAvailabilitySetsDeleteParamsWithContext(ctx context.Context) *AvailabilitySetsDeleteParams {
	var ()
	return &AvailabilitySetsDeleteParams{

		Context: ctx,
	}
}

/*AvailabilitySetsDeleteParams contains all the parameters to send to the API endpoint
for the availability sets delete operation typically these are written to a http.Request
*/
type AvailabilitySetsDeleteParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*AvailabilitySetName
	  The name of the availability set.

	*/
	AvailabilitySetName string
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*SubscriptionID
	  Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) WithTimeout(timeout time.Duration) *AvailabilitySetsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) WithContext(ctx context.Context) *AvailabilitySetsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) WithAPIVersion(aPIVersion string) *AvailabilitySetsDeleteParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithAvailabilitySetName adds the availabilitySetName to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) WithAvailabilitySetName(availabilitySetName string) *AvailabilitySetsDeleteParams {
	o.SetAvailabilitySetName(availabilitySetName)
	return o
}

// SetAvailabilitySetName adds the availabilitySetName to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) SetAvailabilitySetName(availabilitySetName string) {
	o.AvailabilitySetName = availabilitySetName
}

// WithResourceGroupName adds the resourceGroupName to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) WithResourceGroupName(resourceGroupName string) *AvailabilitySetsDeleteParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) WithSubscriptionID(subscriptionID string) *AvailabilitySetsDeleteParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the availability sets delete params
func (o *AvailabilitySetsDeleteParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *AvailabilitySetsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param availabilitySetName
	if err := r.SetPathParam("availabilitySetName", o.AvailabilitySetName); err != nil {
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
