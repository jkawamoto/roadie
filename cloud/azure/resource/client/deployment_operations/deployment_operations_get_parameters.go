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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeploymentOperationsGetParams creates a new DeploymentOperationsGetParams object
// with the default values initialized.
func NewDeploymentOperationsGetParams() *DeploymentOperationsGetParams {
	var ()
	return &DeploymentOperationsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeploymentOperationsGetParamsWithTimeout creates a new DeploymentOperationsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeploymentOperationsGetParamsWithTimeout(timeout time.Duration) *DeploymentOperationsGetParams {
	var ()
	return &DeploymentOperationsGetParams{

		timeout: timeout,
	}
}

// NewDeploymentOperationsGetParamsWithContext creates a new DeploymentOperationsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeploymentOperationsGetParamsWithContext(ctx context.Context) *DeploymentOperationsGetParams {
	var ()
	return &DeploymentOperationsGetParams{

		Context: ctx,
	}
}

// NewDeploymentOperationsGetParamsWithHTTPClient creates a new DeploymentOperationsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeploymentOperationsGetParamsWithHTTPClient(client *http.Client) *DeploymentOperationsGetParams {
	var ()
	return &DeploymentOperationsGetParams{
		HTTPClient: client,
	}
}

/*DeploymentOperationsGetParams contains all the parameters to send to the API endpoint
for the deployment operations get operation typically these are written to a http.Request
*/
type DeploymentOperationsGetParams struct {

	/*APIVersion
	  The API version to use for this operation.

	*/
	APIVersion string
	/*DeploymentName
	  The name of the deployment.

	*/
	DeploymentName string
	/*OperationID
	  The ID of the operation to get.

	*/
	OperationID string
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

// WithTimeout adds the timeout to the deployment operations get params
func (o *DeploymentOperationsGetParams) WithTimeout(timeout time.Duration) *DeploymentOperationsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deployment operations get params
func (o *DeploymentOperationsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deployment operations get params
func (o *DeploymentOperationsGetParams) WithContext(ctx context.Context) *DeploymentOperationsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deployment operations get params
func (o *DeploymentOperationsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deployment operations get params
func (o *DeploymentOperationsGetParams) WithHTTPClient(client *http.Client) *DeploymentOperationsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deployment operations get params
func (o *DeploymentOperationsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the deployment operations get params
func (o *DeploymentOperationsGetParams) WithAPIVersion(aPIVersion string) *DeploymentOperationsGetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the deployment operations get params
func (o *DeploymentOperationsGetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDeploymentName adds the deploymentName to the deployment operations get params
func (o *DeploymentOperationsGetParams) WithDeploymentName(deploymentName string) *DeploymentOperationsGetParams {
	o.SetDeploymentName(deploymentName)
	return o
}

// SetDeploymentName adds the deploymentName to the deployment operations get params
func (o *DeploymentOperationsGetParams) SetDeploymentName(deploymentName string) {
	o.DeploymentName = deploymentName
}

// WithOperationID adds the operationID to the deployment operations get params
func (o *DeploymentOperationsGetParams) WithOperationID(operationID string) *DeploymentOperationsGetParams {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the deployment operations get params
func (o *DeploymentOperationsGetParams) SetOperationID(operationID string) {
	o.OperationID = operationID
}

// WithResourceGroupName adds the resourceGroupName to the deployment operations get params
func (o *DeploymentOperationsGetParams) WithResourceGroupName(resourceGroupName string) *DeploymentOperationsGetParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the deployment operations get params
func (o *DeploymentOperationsGetParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the deployment operations get params
func (o *DeploymentOperationsGetParams) WithSubscriptionID(subscriptionID string) *DeploymentOperationsGetParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the deployment operations get params
func (o *DeploymentOperationsGetParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeploymentOperationsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param operationId
	if err := r.SetPathParam("operationId", o.OperationID); err != nil {
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
