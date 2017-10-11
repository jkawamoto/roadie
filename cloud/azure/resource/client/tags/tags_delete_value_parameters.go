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

// NewTagsDeleteValueParams creates a new TagsDeleteValueParams object
// with the default values initialized.
func NewTagsDeleteValueParams() *TagsDeleteValueParams {
	var ()
	return &TagsDeleteValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTagsDeleteValueParamsWithTimeout creates a new TagsDeleteValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTagsDeleteValueParamsWithTimeout(timeout time.Duration) *TagsDeleteValueParams {
	var ()
	return &TagsDeleteValueParams{

		timeout: timeout,
	}
}

// NewTagsDeleteValueParamsWithContext creates a new TagsDeleteValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewTagsDeleteValueParamsWithContext(ctx context.Context) *TagsDeleteValueParams {
	var ()
	return &TagsDeleteValueParams{

		Context: ctx,
	}
}

// NewTagsDeleteValueParamsWithHTTPClient creates a new TagsDeleteValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTagsDeleteValueParamsWithHTTPClient(client *http.Client) *TagsDeleteValueParams {
	var ()
	return &TagsDeleteValueParams{
		HTTPClient: client,
	}
}

/*TagsDeleteValueParams contains all the parameters to send to the API endpoint
for the tags delete value operation typically these are written to a http.Request
*/
type TagsDeleteValueParams struct {

	/*APIVersion
	  The API version to use for this operation.

	*/
	APIVersion string
	/*SubscriptionID
	  The ID of the target subscription.

	*/
	SubscriptionID string
	/*TagName
	  The name of the tag.

	*/
	TagName string
	/*TagValue
	  The value of the tag to delete.

	*/
	TagValue string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tags delete value params
func (o *TagsDeleteValueParams) WithTimeout(timeout time.Duration) *TagsDeleteValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tags delete value params
func (o *TagsDeleteValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tags delete value params
func (o *TagsDeleteValueParams) WithContext(ctx context.Context) *TagsDeleteValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tags delete value params
func (o *TagsDeleteValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tags delete value params
func (o *TagsDeleteValueParams) WithHTTPClient(client *http.Client) *TagsDeleteValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tags delete value params
func (o *TagsDeleteValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the tags delete value params
func (o *TagsDeleteValueParams) WithAPIVersion(aPIVersion string) *TagsDeleteValueParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the tags delete value params
func (o *TagsDeleteValueParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithSubscriptionID adds the subscriptionID to the tags delete value params
func (o *TagsDeleteValueParams) WithSubscriptionID(subscriptionID string) *TagsDeleteValueParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the tags delete value params
func (o *TagsDeleteValueParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithTagName adds the tagName to the tags delete value params
func (o *TagsDeleteValueParams) WithTagName(tagName string) *TagsDeleteValueParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the tags delete value params
func (o *TagsDeleteValueParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithTagValue adds the tagValue to the tags delete value params
func (o *TagsDeleteValueParams) WithTagValue(tagValue string) *TagsDeleteValueParams {
	o.SetTagValue(tagValue)
	return o
}

// SetTagValue adds the tagValue to the tags delete value params
func (o *TagsDeleteValueParams) SetTagValue(tagValue string) {
	o.TagValue = tagValue
}

// WriteToRequest writes these params to a swagger request
func (o *TagsDeleteValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tagName
	if err := r.SetPathParam("tagName", o.TagName); err != nil {
		return err
	}

	// path param tagValue
	if err := r.SetPathParam("tagValue", o.TagValue); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}