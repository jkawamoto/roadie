package resource_groups

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

// NewResourceGroupsGetParams creates a new ResourceGroupsGetParams object
// with the default values initialized.
func NewResourceGroupsGetParams() *ResourceGroupsGetParams {
	var ()
	return &ResourceGroupsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResourceGroupsGetParamsWithTimeout creates a new ResourceGroupsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResourceGroupsGetParamsWithTimeout(timeout time.Duration) *ResourceGroupsGetParams {
	var ()
	return &ResourceGroupsGetParams{

		timeout: timeout,
	}
}

// NewResourceGroupsGetParamsWithContext creates a new ResourceGroupsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewResourceGroupsGetParamsWithContext(ctx context.Context) *ResourceGroupsGetParams {
	var ()
	return &ResourceGroupsGetParams{

		Context: ctx,
	}
}

// NewResourceGroupsGetParamsWithHTTPClient creates a new ResourceGroupsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResourceGroupsGetParamsWithHTTPClient(client *http.Client) *ResourceGroupsGetParams {
	var ()
	return &ResourceGroupsGetParams{
		HTTPClient: client,
	}
}

/*ResourceGroupsGetParams contains all the parameters to send to the API endpoint
for the resource groups get operation typically these are written to a http.Request
*/
type ResourceGroupsGetParams struct {

	/*APIVersion
	  The API version to use for this operation.

	*/
	APIVersion string
	/*ResourceGroupName
	  The name of the resource group to get. The name is case insensitive.

	*/
	ResourceGroupName string
	/*SubscriptionID
	  The ID of the target subscription.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resource groups get params
func (o *ResourceGroupsGetParams) WithTimeout(timeout time.Duration) *ResourceGroupsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource groups get params
func (o *ResourceGroupsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource groups get params
func (o *ResourceGroupsGetParams) WithContext(ctx context.Context) *ResourceGroupsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource groups get params
func (o *ResourceGroupsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource groups get params
func (o *ResourceGroupsGetParams) WithHTTPClient(client *http.Client) *ResourceGroupsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource groups get params
func (o *ResourceGroupsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the resource groups get params
func (o *ResourceGroupsGetParams) WithAPIVersion(aPIVersion string) *ResourceGroupsGetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the resource groups get params
func (o *ResourceGroupsGetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the resource groups get params
func (o *ResourceGroupsGetParams) WithResourceGroupName(resourceGroupName string) *ResourceGroupsGetParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the resource groups get params
func (o *ResourceGroupsGetParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the resource groups get params
func (o *ResourceGroupsGetParams) WithSubscriptionID(subscriptionID string) *ResourceGroupsGetParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the resource groups get params
func (o *ResourceGroupsGetParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceGroupsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
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
