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

// NewDeploymentsCheckExistenceParams creates a new DeploymentsCheckExistenceParams object
// with the default values initialized.
func NewDeploymentsCheckExistenceParams() *DeploymentsCheckExistenceParams {
	var ()
	return &DeploymentsCheckExistenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeploymentsCheckExistenceParamsWithTimeout creates a new DeploymentsCheckExistenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeploymentsCheckExistenceParamsWithTimeout(timeout time.Duration) *DeploymentsCheckExistenceParams {
	var ()
	return &DeploymentsCheckExistenceParams{

		timeout: timeout,
	}
}

// NewDeploymentsCheckExistenceParamsWithContext creates a new DeploymentsCheckExistenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeploymentsCheckExistenceParamsWithContext(ctx context.Context) *DeploymentsCheckExistenceParams {
	var ()
	return &DeploymentsCheckExistenceParams{

		Context: ctx,
	}
}

// NewDeploymentsCheckExistenceParamsWithHTTPClient creates a new DeploymentsCheckExistenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeploymentsCheckExistenceParamsWithHTTPClient(client *http.Client) *DeploymentsCheckExistenceParams {
	var ()
	return &DeploymentsCheckExistenceParams{
		HTTPClient: client,
	}
}

/*DeploymentsCheckExistenceParams contains all the parameters to send to the API endpoint
for the deployments check existence operation typically these are written to a http.Request
*/
type DeploymentsCheckExistenceParams struct {

	/*APIVersion
	  The API version to use for this operation.

	*/
	APIVersion string
	/*DeploymentName
	  The name of the deployment to check.

	*/
	DeploymentName string
	/*ResourceGroupName
	  The name of the resource group with the deployment to check. The name is case insensitive.

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

// WithTimeout adds the timeout to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) WithTimeout(timeout time.Duration) *DeploymentsCheckExistenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) WithContext(ctx context.Context) *DeploymentsCheckExistenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) WithHTTPClient(client *http.Client) *DeploymentsCheckExistenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) WithAPIVersion(aPIVersion string) *DeploymentsCheckExistenceParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDeploymentName adds the deploymentName to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) WithDeploymentName(deploymentName string) *DeploymentsCheckExistenceParams {
	o.SetDeploymentName(deploymentName)
	return o
}

// SetDeploymentName adds the deploymentName to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) SetDeploymentName(deploymentName string) {
	o.DeploymentName = deploymentName
}

// WithResourceGroupName adds the resourceGroupName to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) WithResourceGroupName(resourceGroupName string) *DeploymentsCheckExistenceParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) WithSubscriptionID(subscriptionID string) *DeploymentsCheckExistenceParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the deployments check existence params
func (o *DeploymentsCheckExistenceParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeploymentsCheckExistenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
