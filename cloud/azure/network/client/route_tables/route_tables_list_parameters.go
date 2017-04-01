package route_tables

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

// NewRouteTablesListParams creates a new RouteTablesListParams object
// with the default values initialized.
func NewRouteTablesListParams() *RouteTablesListParams {
	var ()
	return &RouteTablesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRouteTablesListParamsWithTimeout creates a new RouteTablesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRouteTablesListParamsWithTimeout(timeout time.Duration) *RouteTablesListParams {
	var ()
	return &RouteTablesListParams{

		timeout: timeout,
	}
}

// NewRouteTablesListParamsWithContext creates a new RouteTablesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewRouteTablesListParamsWithContext(ctx context.Context) *RouteTablesListParams {
	var ()
	return &RouteTablesListParams{

		Context: ctx,
	}
}

/*RouteTablesListParams contains all the parameters to send to the API endpoint
for the route tables list operation typically these are written to a http.Request
*/
type RouteTablesListParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the route tables list params
func (o *RouteTablesListParams) WithTimeout(timeout time.Duration) *RouteTablesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route tables list params
func (o *RouteTablesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route tables list params
func (o *RouteTablesListParams) WithContext(ctx context.Context) *RouteTablesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route tables list params
func (o *RouteTablesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the route tables list params
func (o *RouteTablesListParams) WithAPIVersion(aPIVersion string) *RouteTablesListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the route tables list params
func (o *RouteTablesListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the route tables list params
func (o *RouteTablesListParams) WithResourceGroupName(resourceGroupName string) *RouteTablesListParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the route tables list params
func (o *RouteTablesListParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the route tables list params
func (o *RouteTablesListParams) WithSubscriptionID(subscriptionID string) *RouteTablesListParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the route tables list params
func (o *RouteTablesListParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *RouteTablesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
