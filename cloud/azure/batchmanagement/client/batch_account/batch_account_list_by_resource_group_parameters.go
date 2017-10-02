package batch_account

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

// NewBatchAccountListByResourceGroupParams creates a new BatchAccountListByResourceGroupParams object
// with the default values initialized.
func NewBatchAccountListByResourceGroupParams() *BatchAccountListByResourceGroupParams {
	var ()
	return &BatchAccountListByResourceGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBatchAccountListByResourceGroupParamsWithTimeout creates a new BatchAccountListByResourceGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBatchAccountListByResourceGroupParamsWithTimeout(timeout time.Duration) *BatchAccountListByResourceGroupParams {
	var ()
	return &BatchAccountListByResourceGroupParams{

		timeout: timeout,
	}
}

// NewBatchAccountListByResourceGroupParamsWithContext creates a new BatchAccountListByResourceGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewBatchAccountListByResourceGroupParamsWithContext(ctx context.Context) *BatchAccountListByResourceGroupParams {
	var ()
	return &BatchAccountListByResourceGroupParams{

		Context: ctx,
	}
}

/*BatchAccountListByResourceGroupParams contains all the parameters to send to the API endpoint
for the batch account list by resource group operation typically these are written to a http.Request
*/
type BatchAccountListByResourceGroupParams struct {

	/*APIVersion
	  The API version to be used with the HTTP request.

	*/
	APIVersion string
	/*ResourceGroupName
	  The name of the resource group whose Batch accounts to list.

	*/
	ResourceGroupName string
	/*SubscriptionID
	  A unique identifier of a Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the batch account list by resource group params
func (o *BatchAccountListByResourceGroupParams) WithTimeout(timeout time.Duration) *BatchAccountListByResourceGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch account list by resource group params
func (o *BatchAccountListByResourceGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch account list by resource group params
func (o *BatchAccountListByResourceGroupParams) WithContext(ctx context.Context) *BatchAccountListByResourceGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch account list by resource group params
func (o *BatchAccountListByResourceGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the batch account list by resource group params
func (o *BatchAccountListByResourceGroupParams) WithAPIVersion(aPIVersion string) *BatchAccountListByResourceGroupParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the batch account list by resource group params
func (o *BatchAccountListByResourceGroupParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the batch account list by resource group params
func (o *BatchAccountListByResourceGroupParams) WithResourceGroupName(resourceGroupName string) *BatchAccountListByResourceGroupParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the batch account list by resource group params
func (o *BatchAccountListByResourceGroupParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the batch account list by resource group params
func (o *BatchAccountListByResourceGroupParams) WithSubscriptionID(subscriptionID string) *BatchAccountListByResourceGroupParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the batch account list by resource group params
func (o *BatchAccountListByResourceGroupParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *BatchAccountListByResourceGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
