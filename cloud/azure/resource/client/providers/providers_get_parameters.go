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

	strfmt "github.com/go-openapi/strfmt"
)

// NewProvidersGetParams creates a new ProvidersGetParams object
// with the default values initialized.
func NewProvidersGetParams() *ProvidersGetParams {
	var ()
	return &ProvidersGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProvidersGetParamsWithTimeout creates a new ProvidersGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProvidersGetParamsWithTimeout(timeout time.Duration) *ProvidersGetParams {
	var ()
	return &ProvidersGetParams{

		timeout: timeout,
	}
}

// NewProvidersGetParamsWithContext creates a new ProvidersGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewProvidersGetParamsWithContext(ctx context.Context) *ProvidersGetParams {
	var ()
	return &ProvidersGetParams{

		Context: ctx,
	}
}

// NewProvidersGetParamsWithHTTPClient creates a new ProvidersGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProvidersGetParamsWithHTTPClient(client *http.Client) *ProvidersGetParams {
	var ()
	return &ProvidersGetParams{
		HTTPClient: client,
	}
}

/*ProvidersGetParams contains all the parameters to send to the API endpoint
for the providers get operation typically these are written to a http.Request
*/
type ProvidersGetParams struct {

	/*NrDollarExpand
	  The $expand query parameter. For example, to include property aliases in response, use $expand=resourceTypes/aliases.

	*/
	DollarExpand *string
	/*APIVersion
	  The API version to use for this operation.

	*/
	APIVersion string
	/*ResourceProviderNamespace
	  The namespace of the resource provider.

	*/
	ResourceProviderNamespace string
	/*SubscriptionID
	  The ID of the target subscription.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the providers get params
func (o *ProvidersGetParams) WithTimeout(timeout time.Duration) *ProvidersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the providers get params
func (o *ProvidersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the providers get params
func (o *ProvidersGetParams) WithContext(ctx context.Context) *ProvidersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the providers get params
func (o *ProvidersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the providers get params
func (o *ProvidersGetParams) WithHTTPClient(client *http.Client) *ProvidersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the providers get params
func (o *ProvidersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarExpand adds the dollarExpand to the providers get params
func (o *ProvidersGetParams) WithDollarExpand(dollarExpand *string) *ProvidersGetParams {
	o.SetDollarExpand(dollarExpand)
	return o
}

// SetDollarExpand adds the dollarExpand to the providers get params
func (o *ProvidersGetParams) SetDollarExpand(dollarExpand *string) {
	o.DollarExpand = dollarExpand
}

// WithAPIVersion adds the aPIVersion to the providers get params
func (o *ProvidersGetParams) WithAPIVersion(aPIVersion string) *ProvidersGetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the providers get params
func (o *ProvidersGetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceProviderNamespace adds the resourceProviderNamespace to the providers get params
func (o *ProvidersGetParams) WithResourceProviderNamespace(resourceProviderNamespace string) *ProvidersGetParams {
	o.SetResourceProviderNamespace(resourceProviderNamespace)
	return o
}

// SetResourceProviderNamespace adds the resourceProviderNamespace to the providers get params
func (o *ProvidersGetParams) SetResourceProviderNamespace(resourceProviderNamespace string) {
	o.ResourceProviderNamespace = resourceProviderNamespace
}

// WithSubscriptionID adds the subscriptionID to the providers get params
func (o *ProvidersGetParams) WithSubscriptionID(subscriptionID string) *ProvidersGetParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the providers get params
func (o *ProvidersGetParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ProvidersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param api-version
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("api-version", qAPIVersion); err != nil {
			return err
		}
	}

	// path param resourceProviderNamespace
	if err := r.SetPathParam("resourceProviderNamespace", o.ResourceProviderNamespace); err != nil {
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
