package routes

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

// NewRoutesGetParams creates a new RoutesGetParams object
// with the default values initialized.
func NewRoutesGetParams() *RoutesGetParams {
	var ()
	return &RoutesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRoutesGetParamsWithTimeout creates a new RoutesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRoutesGetParamsWithTimeout(timeout time.Duration) *RoutesGetParams {
	var ()
	return &RoutesGetParams{

		timeout: timeout,
	}
}

// NewRoutesGetParamsWithContext creates a new RoutesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewRoutesGetParamsWithContext(ctx context.Context) *RoutesGetParams {
	var ()
	return &RoutesGetParams{

		Context: ctx,
	}
}

/*RoutesGetParams contains all the parameters to send to the API endpoint
for the routes get operation typically these are written to a http.Request
*/
type RoutesGetParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*RouteName
	  The name of the route.

	*/
	RouteName string
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

// WithTimeout adds the timeout to the routes get params
func (o *RoutesGetParams) WithTimeout(timeout time.Duration) *RoutesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the routes get params
func (o *RoutesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the routes get params
func (o *RoutesGetParams) WithContext(ctx context.Context) *RoutesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the routes get params
func (o *RoutesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the routes get params
func (o *RoutesGetParams) WithAPIVersion(aPIVersion string) *RoutesGetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the routes get params
func (o *RoutesGetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the routes get params
func (o *RoutesGetParams) WithResourceGroupName(resourceGroupName string) *RoutesGetParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the routes get params
func (o *RoutesGetParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithRouteName adds the routeName to the routes get params
func (o *RoutesGetParams) WithRouteName(routeName string) *RoutesGetParams {
	o.SetRouteName(routeName)
	return o
}

// SetRouteName adds the routeName to the routes get params
func (o *RoutesGetParams) SetRouteName(routeName string) {
	o.RouteName = routeName
}

// WithRouteTableName adds the routeTableName to the routes get params
func (o *RoutesGetParams) WithRouteTableName(routeTableName string) *RoutesGetParams {
	o.SetRouteTableName(routeTableName)
	return o
}

// SetRouteTableName adds the routeTableName to the routes get params
func (o *RoutesGetParams) SetRouteTableName(routeTableName string) {
	o.RouteTableName = routeTableName
}

// WithSubscriptionID adds the subscriptionID to the routes get params
func (o *RoutesGetParams) WithSubscriptionID(subscriptionID string) *RoutesGetParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the routes get params
func (o *RoutesGetParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *RoutesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param routeName
	if err := r.SetPathParam("routeName", o.RouteName); err != nil {
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
