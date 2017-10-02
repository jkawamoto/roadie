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

// NewBatchAccountGetKeysParams creates a new BatchAccountGetKeysParams object
// with the default values initialized.
func NewBatchAccountGetKeysParams() *BatchAccountGetKeysParams {
	var ()
	return &BatchAccountGetKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBatchAccountGetKeysParamsWithTimeout creates a new BatchAccountGetKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBatchAccountGetKeysParamsWithTimeout(timeout time.Duration) *BatchAccountGetKeysParams {
	var ()
	return &BatchAccountGetKeysParams{

		timeout: timeout,
	}
}

// NewBatchAccountGetKeysParamsWithContext creates a new BatchAccountGetKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewBatchAccountGetKeysParamsWithContext(ctx context.Context) *BatchAccountGetKeysParams {
	var ()
	return &BatchAccountGetKeysParams{

		Context: ctx,
	}
}

/*BatchAccountGetKeysParams contains all the parameters to send to the API endpoint
for the batch account get keys operation typically these are written to a http.Request
*/
type BatchAccountGetKeysParams struct {

	/*AccountName
	  The name of the account.

	*/
	AccountName string
	/*APIVersion
	  The API version to be used with the HTTP request.

	*/
	APIVersion string
	/*ResourceGroupName
	  The name of the resource group that contains the Batch account.

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

// WithTimeout adds the timeout to the batch account get keys params
func (o *BatchAccountGetKeysParams) WithTimeout(timeout time.Duration) *BatchAccountGetKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch account get keys params
func (o *BatchAccountGetKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch account get keys params
func (o *BatchAccountGetKeysParams) WithContext(ctx context.Context) *BatchAccountGetKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch account get keys params
func (o *BatchAccountGetKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAccountName adds the accountName to the batch account get keys params
func (o *BatchAccountGetKeysParams) WithAccountName(accountName string) *BatchAccountGetKeysParams {
	o.SetAccountName(accountName)
	return o
}

// SetAccountName adds the accountName to the batch account get keys params
func (o *BatchAccountGetKeysParams) SetAccountName(accountName string) {
	o.AccountName = accountName
}

// WithAPIVersion adds the aPIVersion to the batch account get keys params
func (o *BatchAccountGetKeysParams) WithAPIVersion(aPIVersion string) *BatchAccountGetKeysParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the batch account get keys params
func (o *BatchAccountGetKeysParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the batch account get keys params
func (o *BatchAccountGetKeysParams) WithResourceGroupName(resourceGroupName string) *BatchAccountGetKeysParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the batch account get keys params
func (o *BatchAccountGetKeysParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the batch account get keys params
func (o *BatchAccountGetKeysParams) WithSubscriptionID(subscriptionID string) *BatchAccountGetKeysParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the batch account get keys params
func (o *BatchAccountGetKeysParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *BatchAccountGetKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param accountName
	if err := r.SetPathParam("accountName", o.AccountName); err != nil {
		return err
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

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
