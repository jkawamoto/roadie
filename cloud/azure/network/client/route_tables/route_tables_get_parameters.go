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

// NewRouteTablesGetParams creates a new RouteTablesGetParams object
// with the default values initialized.
func NewRouteTablesGetParams() *RouteTablesGetParams {
	var ()
	return &RouteTablesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRouteTablesGetParamsWithTimeout creates a new RouteTablesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRouteTablesGetParamsWithTimeout(timeout time.Duration) *RouteTablesGetParams {
	var ()
	return &RouteTablesGetParams{

		timeout: timeout,
	}
}

// NewRouteTablesGetParamsWithContext creates a new RouteTablesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewRouteTablesGetParamsWithContext(ctx context.Context) *RouteTablesGetParams {
	var ()
	return &RouteTablesGetParams{

		Context: ctx,
	}
}

/*RouteTablesGetParams contains all the parameters to send to the API endpoint
for the route tables get operation typically these are written to a http.Request
*/
type RouteTablesGetParams struct {

	/*NrDollarExpand
	  expand references resources.

	*/
	DollarExpand *string
	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*RouteTableName
	  The name of the route table.

	*/
	RouteTableName string
	/*SubscriptionID
	  Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the route tables get params
func (o *RouteTablesGetParams) WithTimeout(timeout time.Duration) *RouteTablesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route tables get params
func (o *RouteTablesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route tables get params
func (o *RouteTablesGetParams) WithContext(ctx context.Context) *RouteTablesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route tables get params
func (o *RouteTablesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDollarExpand adds the dollarExpand to the route tables get params
func (o *RouteTablesGetParams) WithDollarExpand(dollarExpand *string) *RouteTablesGetParams {
	o.SetDollarExpand(dollarExpand)
	return o
}

// SetDollarExpand adds the dollarExpand to the route tables get params
func (o *RouteTablesGetParams) SetDollarExpand(dollarExpand *string) {
	o.DollarExpand = dollarExpand
}

// WithAPIVersion adds the aPIVersion to the route tables get params
func (o *RouteTablesGetParams) WithAPIVersion(aPIVersion string) *RouteTablesGetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the route tables get params
func (o *RouteTablesGetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the route tables get params
func (o *RouteTablesGetParams) WithResourceGroupName(resourceGroupName string) *RouteTablesGetParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the route tables get params
func (o *RouteTablesGetParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithRouteTableName adds the routeTableName to the route tables get params
func (o *RouteTablesGetParams) WithRouteTableName(routeTableName string) *RouteTablesGetParams {
	o.SetRouteTableName(routeTableName)
	return o
}

// SetRouteTableName adds the routeTableName to the route tables get params
func (o *RouteTablesGetParams) SetRouteTableName(routeTableName string) {
	o.RouteTableName = routeTableName
}

// WithSubscriptionID adds the subscriptionID to the route tables get params
func (o *RouteTablesGetParams) WithSubscriptionID(subscriptionID string) *RouteTablesGetParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the route tables get params
func (o *RouteTablesGetParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *RouteTablesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resourceGroupName
	if err := r.SetPathParam("resourceGroupName", o.ResourceGroupName); err != nil {
		return err
	}

	// path param routeTableName
	if err := r.SetPathParam("routeTableName", o.RouteTableName); err != nil {
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
