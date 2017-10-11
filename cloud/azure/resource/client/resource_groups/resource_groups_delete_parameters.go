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

// NewResourceGroupsDeleteParams creates a new ResourceGroupsDeleteParams object
// with the default values initialized.
func NewResourceGroupsDeleteParams() *ResourceGroupsDeleteParams {
	var ()
	return &ResourceGroupsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResourceGroupsDeleteParamsWithTimeout creates a new ResourceGroupsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResourceGroupsDeleteParamsWithTimeout(timeout time.Duration) *ResourceGroupsDeleteParams {
	var ()
	return &ResourceGroupsDeleteParams{

		timeout: timeout,
	}
}

// NewResourceGroupsDeleteParamsWithContext creates a new ResourceGroupsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewResourceGroupsDeleteParamsWithContext(ctx context.Context) *ResourceGroupsDeleteParams {
	var ()
	return &ResourceGroupsDeleteParams{

		Context: ctx,
	}
}

// NewResourceGroupsDeleteParamsWithHTTPClient creates a new ResourceGroupsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResourceGroupsDeleteParamsWithHTTPClient(client *http.Client) *ResourceGroupsDeleteParams {
	var ()
	return &ResourceGroupsDeleteParams{
		HTTPClient: client,
	}
}

/*ResourceGroupsDeleteParams contains all the parameters to send to the API endpoint
for the resource groups delete operation typically these are written to a http.Request
*/
type ResourceGroupsDeleteParams struct {

	/*APIVersion
	  The API version to use for this operation.

	*/
	APIVersion string
	/*ResourceGroupName
	  The name of the resource group to delete. The name is case insensitive.

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

// WithTimeout adds the timeout to the resource groups delete params
func (o *ResourceGroupsDeleteParams) WithTimeout(timeout time.Duration) *ResourceGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource groups delete params
func (o *ResourceGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource groups delete params
func (o *ResourceGroupsDeleteParams) WithContext(ctx context.Context) *ResourceGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource groups delete params
func (o *ResourceGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource groups delete params
func (o *ResourceGroupsDeleteParams) WithHTTPClient(client *http.Client) *ResourceGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource groups delete params
func (o *ResourceGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the resource groups delete params
func (o *ResourceGroupsDeleteParams) WithAPIVersion(aPIVersion string) *ResourceGroupsDeleteParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the resource groups delete params
func (o *ResourceGroupsDeleteParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the resource groups delete params
func (o *ResourceGroupsDeleteParams) WithResourceGroupName(resourceGroupName string) *ResourceGroupsDeleteParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the resource groups delete params
func (o *ResourceGroupsDeleteParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the resource groups delete params
func (o *ResourceGroupsDeleteParams) WithSubscriptionID(subscriptionID string) *ResourceGroupsDeleteParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the resource groups delete params
func (o *ResourceGroupsDeleteParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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