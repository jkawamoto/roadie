package deployments

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

// NewDeploymentsGetParams creates a new DeploymentsGetParams object
// with the default values initialized.
func NewDeploymentsGetParams() *DeploymentsGetParams {
	var ()
	return &DeploymentsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeploymentsGetParamsWithTimeout creates a new DeploymentsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeploymentsGetParamsWithTimeout(timeout time.Duration) *DeploymentsGetParams {
	var ()
	return &DeploymentsGetParams{

		timeout: timeout,
	}
}

// NewDeploymentsGetParamsWithContext creates a new DeploymentsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeploymentsGetParamsWithContext(ctx context.Context) *DeploymentsGetParams {
	var ()
	return &DeploymentsGetParams{

		Context: ctx,
	}
}

// NewDeploymentsGetParamsWithHTTPClient creates a new DeploymentsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeploymentsGetParamsWithHTTPClient(client *http.Client) *DeploymentsGetParams {
	var ()
	return &DeploymentsGetParams{
		HTTPClient: client,
	}
}

/*DeploymentsGetParams contains all the parameters to send to the API endpoint
for the deployments get operation typically these are written to a http.Request
*/
type DeploymentsGetParams struct {

	/*APIVersion
	  The API version to use for this operation.

	*/
	APIVersion string
	/*DeploymentName
	  The name of the deployment to get.

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

// WithTimeout adds the timeout to the deployments get params
func (o *DeploymentsGetParams) WithTimeout(timeout time.Duration) *DeploymentsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deployments get params
func (o *DeploymentsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deployments get params
func (o *DeploymentsGetParams) WithContext(ctx context.Context) *DeploymentsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deployments get params
func (o *DeploymentsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deployments get params
func (o *DeploymentsGetParams) WithHTTPClient(client *http.Client) *DeploymentsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deployments get params
func (o *DeploymentsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the deployments get params
func (o *DeploymentsGetParams) WithAPIVersion(aPIVersion string) *DeploymentsGetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the deployments get params
func (o *DeploymentsGetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDeploymentName adds the deploymentName to the deployments get params
func (o *DeploymentsGetParams) WithDeploymentName(deploymentName string) *DeploymentsGetParams {
	o.SetDeploymentName(deploymentName)
	return o
}

// SetDeploymentName adds the deploymentName to the deployments get params
func (o *DeploymentsGetParams) SetDeploymentName(deploymentName string) {
	o.DeploymentName = deploymentName
}

// WithResourceGroupName adds the resourceGroupName to the deployments get params
func (o *DeploymentsGetParams) WithResourceGroupName(resourceGroupName string) *DeploymentsGetParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the deployments get params
func (o *DeploymentsGetParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the deployments get params
func (o *DeploymentsGetParams) WithSubscriptionID(subscriptionID string) *DeploymentsGetParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the deployments get params
func (o *DeploymentsGetParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeploymentsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
