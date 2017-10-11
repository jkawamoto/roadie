package tags

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

// NewTagsListParams creates a new TagsListParams object
// with the default values initialized.
func NewTagsListParams() *TagsListParams {
	var ()
	return &TagsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTagsListParamsWithTimeout creates a new TagsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTagsListParamsWithTimeout(timeout time.Duration) *TagsListParams {
	var ()
	return &TagsListParams{

		timeout: timeout,
	}
}

// NewTagsListParamsWithContext creates a new TagsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTagsListParamsWithContext(ctx context.Context) *TagsListParams {
	var ()
	return &TagsListParams{

		Context: ctx,
	}
}

// NewTagsListParamsWithHTTPClient creates a new TagsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTagsListParamsWithHTTPClient(client *http.Client) *TagsListParams {
	var ()
	return &TagsListParams{
		HTTPClient: client,
	}
}

/*TagsListParams contains all the parameters to send to the API endpoint
for the tags list operation typically these are written to a http.Request
*/
type TagsListParams struct {

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

// WithTimeout adds the timeout to the tags list params
func (o *TagsListParams) WithTimeout(timeout time.Duration) *TagsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tags list params
func (o *TagsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tags list params
func (o *TagsListParams) WithContext(ctx context.Context) *TagsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tags list params
func (o *TagsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tags list params
func (o *TagsListParams) WithHTTPClient(client *http.Client) *TagsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tags list params
func (o *TagsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the tags list params
func (o *TagsListParams) WithAPIVersion(aPIVersion string) *TagsListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the tags list params
func (o *TagsListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithSubscriptionID adds the subscriptionID to the tags list params
func (o *TagsListParams) WithSubscriptionID(subscriptionID string) *TagsListParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the tags list params
func (o *TagsListParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *TagsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}