package deployment_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeploymentOperationsListParams creates a new DeploymentOperationsListParams object
// with the default values initialized.
func NewDeploymentOperationsListParams() *DeploymentOperationsListParams {
	var ()
	return &DeploymentOperationsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeploymentOperationsListParamsWithTimeout creates a new DeploymentOperationsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeploymentOperationsListParamsWithTimeout(timeout time.Duration) *DeploymentOperationsListParams {
	var ()
	return &DeploymentOperationsListParams{

		timeout: timeout,
	}
}

// NewDeploymentOperationsListParamsWithContext creates a new DeploymentOperationsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeploymentOperationsListParamsWithContext(ctx context.Context) *DeploymentOperationsListParams {
	var ()
	return &DeploymentOperationsListParams{

		Context: ctx,
	}
}

// NewDeploymentOperationsListParamsWithHTTPClient creates a new DeploymentOperationsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeploymentOperationsListParamsWithHTTPClient(client *http.Client) *DeploymentOperationsListParams {
	var ()
	return &DeploymentOperationsListParams{
		HTTPClient: client,
	}
}

/*DeploymentOperationsListParams contains all the parameters to send to the API endpoint
for the deployment operations list operation typically these are written to a http.Request
*/
type DeploymentOperationsListParams struct {

	/*NrDollarTop
	  The number of results to return.

	*/
	DollarTop *int32
	/*APIVersion
	  The API version to use for this operation.

	*/
	APIVersion string
	/*DeploymentName
	  The name of the deployment with the operation to get.

	*/
	DeploymentName string
	/*ResourceGroupName
	  The name of the resource group. The name is case insensitive.

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

// WithTimeout adds the timeout to the deployment operations list params
func (o *DeploymentOperationsListParams) WithTimeout(timeout time.Duration) *DeploymentOperationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deployment operations list params
func (o *DeploymentOperationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deployment operations list params
func (o *DeploymentOperationsListParams) WithContext(ctx context.Context) *DeploymentOperationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deployment operations list params
func (o *DeploymentOperationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deployment operations list params
func (o *DeploymentOperationsListParams) WithHTTPClient(client *http.Client) *DeploymentOperationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deployment operations list params
func (o *DeploymentOperationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarTop adds the dollarTop to the deployment operations list params
func (o *DeploymentOperationsListParams) WithDollarTop(dollarTop *int32) *DeploymentOperationsListParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the deployment operations list params
func (o *DeploymentOperationsListParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the deployment operations list params
func (o *DeploymentOperationsListParams) WithAPIVersion(aPIVersion string) *DeploymentOperationsListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the deployment operations list params
func (o *DeploymentOperationsListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDeploymentName adds the deploymentName to the deployment operations list params
func (o *DeploymentOperationsListParams) WithDeploymentName(deploymentName string) *DeploymentOperationsListParams {
	o.SetDeploymentName(deploymentName)
	return o
}

// SetDeploymentName adds the deploymentName to the deployment operations list params
func (o *DeploymentOperationsListParams) SetDeploymentName(deploymentName string) {
	o.DeploymentName = deploymentName
}

// WithResourceGroupName adds the resourceGroupName to the deployment operations list params
func (o *DeploymentOperationsListParams) WithResourceGroupName(resourceGroupName string) *DeploymentOperationsListParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the deployment operations list params
func (o *DeploymentOperationsListParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the deployment operations list params
func (o *DeploymentOperationsListParams) WithSubscriptionID(subscriptionID string) *DeploymentOperationsListParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the deployment operations list params
func (o *DeploymentOperationsListParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeploymentOperationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarTop != nil {

		// query param $top
		var qrNrDollarTop int32
		if o.DollarTop != nil {
			qrNrDollarTop = *o.DollarTop
		}
		qNrDollarTop := swag.FormatInt32(qrNrDollarTop)
		if qNrDollarTop != "" {
			if err := r.SetQueryParam("$top", qNrDollarTop); err != nil {
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

	// path param deploymentName
	if err := r.SetPathParam("deploymentName", o.DeploymentName); err != nil {
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
