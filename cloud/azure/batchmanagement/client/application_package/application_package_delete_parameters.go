package application_package

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

// NewApplicationPackageDeleteParams creates a new ApplicationPackageDeleteParams object
// with the default values initialized.
func NewApplicationPackageDeleteParams() *ApplicationPackageDeleteParams {
	var ()
	return &ApplicationPackageDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationPackageDeleteParamsWithTimeout creates a new ApplicationPackageDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewApplicationPackageDeleteParamsWithTimeout(timeout time.Duration) *ApplicationPackageDeleteParams {
	var ()
	return &ApplicationPackageDeleteParams{

		timeout: timeout,
	}
}

// NewApplicationPackageDeleteParamsWithContext creates a new ApplicationPackageDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewApplicationPackageDeleteParamsWithContext(ctx context.Context) *ApplicationPackageDeleteParams {
	var ()
	return &ApplicationPackageDeleteParams{

		Context: ctx,
	}
}

/*ApplicationPackageDeleteParams contains all the parameters to send to the API endpoint
for the application package delete operation typically these are written to a http.Request
*/
type ApplicationPackageDeleteParams struct {

	/*AccountName
	  The name of the Batch account.

	*/
	AccountName string
	/*APIVersion
	  The API version to be used with the HTTP request.

	*/
	APIVersion string
	/*ApplicationID
	  The ID of the application.

	*/
	ApplicationID string
	/*ResourceGroupName
	  The name of the resource group that contains the Batch account.

	*/
	ResourceGroupName string
	/*SubscriptionID
	  A unique identifier of a Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string
	/*Version
	  The version of the application to delete.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the application package delete params
func (o *ApplicationPackageDeleteParams) WithTimeout(timeout time.Duration) *ApplicationPackageDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application package delete params
func (o *ApplicationPackageDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application package delete params
func (o *ApplicationPackageDeleteParams) WithContext(ctx context.Context) *ApplicationPackageDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application package delete params
func (o *ApplicationPackageDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAccountName adds the accountName to the application package delete params
func (o *ApplicationPackageDeleteParams) WithAccountName(accountName string) *ApplicationPackageDeleteParams {
	o.SetAccountName(accountName)
	return o
}

// SetAccountName adds the accountName to the application package delete params
func (o *ApplicationPackageDeleteParams) SetAccountName(accountName string) {
	o.AccountName = accountName
}

// WithAPIVersion adds the aPIVersion to the application package delete params
func (o *ApplicationPackageDeleteParams) WithAPIVersion(aPIVersion string) *ApplicationPackageDeleteParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the application package delete params
func (o *ApplicationPackageDeleteParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithApplicationID adds the applicationID to the application package delete params
func (o *ApplicationPackageDeleteParams) WithApplicationID(applicationID string) *ApplicationPackageDeleteParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the application package delete params
func (o *ApplicationPackageDeleteParams) SetApplicationID(applicationID string) {
	o.ApplicationID = applicationID
}

// WithResourceGroupName adds the resourceGroupName to the application package delete params
func (o *ApplicationPackageDeleteParams) WithResourceGroupName(resourceGroupName string) *ApplicationPackageDeleteParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the application package delete params
func (o *ApplicationPackageDeleteParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the application package delete params
func (o *ApplicationPackageDeleteParams) WithSubscriptionID(subscriptionID string) *ApplicationPackageDeleteParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the application package delete params
func (o *ApplicationPackageDeleteParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVersion adds the version to the application package delete params
func (o *ApplicationPackageDeleteParams) WithVersion(version string) *ApplicationPackageDeleteParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the application package delete params
func (o *ApplicationPackageDeleteParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationPackageDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param applicationId
	if err := r.SetPathParam("applicationId", o.ApplicationID); err != nil {
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

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
