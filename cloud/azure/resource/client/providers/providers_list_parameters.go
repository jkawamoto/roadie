package providers

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

// NewProvidersListParams creates a new ProvidersListParams object
// with the default values initialized.
func NewProvidersListParams() *ProvidersListParams {
	var ()
	return &ProvidersListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProvidersListParamsWithTimeout creates a new ProvidersListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProvidersListParamsWithTimeout(timeout time.Duration) *ProvidersListParams {
	var ()
	return &ProvidersListParams{

		timeout: timeout,
	}
}

// NewProvidersListParamsWithContext creates a new ProvidersListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProvidersListParamsWithContext(ctx context.Context) *ProvidersListParams {
	var ()
	return &ProvidersListParams{

		Context: ctx,
	}
}

// NewProvidersListParamsWithHTTPClient creates a new ProvidersListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProvidersListParamsWithHTTPClient(client *http.Client) *ProvidersListParams {
	var ()
	return &ProvidersListParams{
		HTTPClient: client,
	}
}

/*ProvidersListParams contains all the parameters to send to the API endpoint
for the providers list operation typically these are written to a http.Request
*/
type ProvidersListParams struct {

	/*NrDollarExpand
	  The properties to include in the results. For example, use &$expand=metadata in the query string to retrieve resource provider metadata. To include property aliases in response, use $expand=resourceTypes/aliases.

	*/
	DollarExpand *string
	/*NrDollarTop
	  The number of results to return. If null is passed returns all deployments.

	*/
	DollarTop *int32
	/*APIVersion
	  The API version to use for this operation.

	*/
	APIVersion string
	/*SubscriptionID
	  The ID of the target subscription.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the providers list params
func (o *ProvidersListParams) WithTimeout(timeout time.Duration) *ProvidersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the providers list params
func (o *ProvidersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the providers list params
func (o *ProvidersListParams) WithContext(ctx context.Context) *ProvidersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the providers list params
func (o *ProvidersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the providers list params
func (o *ProvidersListParams) WithHTTPClient(client *http.Client) *ProvidersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the providers list params
func (o *ProvidersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarExpand adds the dollarExpand to the providers list params
func (o *ProvidersListParams) WithDollarExpand(dollarExpand *string) *ProvidersListParams {
	o.SetDollarExpand(dollarExpand)
	return o
}

// SetDollarExpand adds the dollarExpand to the providers list params
func (o *ProvidersListParams) SetDollarExpand(dollarExpand *string) {
	o.DollarExpand = dollarExpand
}

// WithDollarTop adds the dollarTop to the providers list params
func (o *ProvidersListParams) WithDollarTop(dollarTop *int32) *ProvidersListParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the providers list params
func (o *ProvidersListParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the providers list params
func (o *ProvidersListParams) WithAPIVersion(aPIVersion string) *ProvidersListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the providers list params
func (o *ProvidersListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithSubscriptionID adds the subscriptionID to the providers list params
func (o *ProvidersListParams) WithSubscriptionID(subscriptionID string) *ProvidersListParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the providers list params
func (o *ProvidersListParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ProvidersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarExpand != nil {

		// query param $expand
		var qrNrDollarExpand string
		if o.DollarExpand != nil {
			qrNrDollarExpand = *o.DollarExpand
		}
		qNrDollarExpand := qrNrDollarExpand
		if qNrDollarExpand != "" {
			if err := r.SetQueryParam("$expand", qNrDollarExpand); err != nil {
				return err
			}
		}

	}

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

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
