package resources

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

// NewResourcesListParams creates a new ResourcesListParams object
// with the default values initialized.
func NewResourcesListParams() *ResourcesListParams {
	var ()
	return &ResourcesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResourcesListParamsWithTimeout creates a new ResourcesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResourcesListParamsWithTimeout(timeout time.Duration) *ResourcesListParams {
	var ()
	return &ResourcesListParams{

		timeout: timeout,
	}
}

// NewResourcesListParamsWithContext creates a new ResourcesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewResourcesListParamsWithContext(ctx context.Context) *ResourcesListParams {
	var ()
	return &ResourcesListParams{

		Context: ctx,
	}
}

// NewResourcesListParamsWithHTTPClient creates a new ResourcesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResourcesListParamsWithHTTPClient(client *http.Client) *ResourcesListParams {
	var ()
	return &ResourcesListParams{
		HTTPClient: client,
	}
}

/*ResourcesListParams contains all the parameters to send to the API endpoint
for the resources list operation typically these are written to a http.Request
*/
type ResourcesListParams struct {

	/*NrDollarExpand
	  The $expand query parameter.

	*/
	DollarExpand *string
	/*NrDollarFilter
	  The filter to apply on the operation.

	*/
	DollarFilter *string
	/*NrDollarTop
	  The number of results to return. If null is passed, returns all resource groups.

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

// WithTimeout adds the timeout to the resources list params
func (o *ResourcesListParams) WithTimeout(timeout time.Duration) *ResourcesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resources list params
func (o *ResourcesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resources list params
func (o *ResourcesListParams) WithContext(ctx context.Context) *ResourcesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resources list params
func (o *ResourcesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resources list params
func (o *ResourcesListParams) WithHTTPClient(client *http.Client) *ResourcesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resources list params
func (o *ResourcesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarExpand adds the dollarExpand to the resources list params
func (o *ResourcesListParams) WithDollarExpand(dollarExpand *string) *ResourcesListParams {
	o.SetDollarExpand(dollarExpand)
	return o
}

// SetDollarExpand adds the dollarExpand to the resources list params
func (o *ResourcesListParams) SetDollarExpand(dollarExpand *string) {
	o.DollarExpand = dollarExpand
}

// WithDollarFilter adds the dollarFilter to the resources list params
func (o *ResourcesListParams) WithDollarFilter(dollarFilter *string) *ResourcesListParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the resources list params
func (o *ResourcesListParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarTop adds the dollarTop to the resources list params
func (o *ResourcesListParams) WithDollarTop(dollarTop *int32) *ResourcesListParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the resources list params
func (o *ResourcesListParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the resources list params
func (o *ResourcesListParams) WithAPIVersion(aPIVersion string) *ResourcesListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the resources list params
func (o *ResourcesListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithSubscriptionID adds the subscriptionID to the resources list params
func (o *ResourcesListParams) WithSubscriptionID(subscriptionID string) *ResourcesListParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the resources list params
func (o *ResourcesListParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ResourcesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DollarFilter != nil {

		// query param $filter
		var qrNrDollarFilter string
		if o.DollarFilter != nil {
			qrNrDollarFilter = *o.DollarFilter
		}
		qNrDollarFilter := qrNrDollarFilter
		if qNrDollarFilter != "" {
			if err := r.SetQueryParam("$filter", qNrDollarFilter); err != nil {
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