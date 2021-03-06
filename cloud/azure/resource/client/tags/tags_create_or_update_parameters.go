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

// NewTagsCreateOrUpdateParams creates a new TagsCreateOrUpdateParams object
// with the default values initialized.
func NewTagsCreateOrUpdateParams() *TagsCreateOrUpdateParams {
	var ()
	return &TagsCreateOrUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTagsCreateOrUpdateParamsWithTimeout creates a new TagsCreateOrUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTagsCreateOrUpdateParamsWithTimeout(timeout time.Duration) *TagsCreateOrUpdateParams {
	var ()
	return &TagsCreateOrUpdateParams{

		timeout: timeout,
	}
}

// NewTagsCreateOrUpdateParamsWithContext creates a new TagsCreateOrUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewTagsCreateOrUpdateParamsWithContext(ctx context.Context) *TagsCreateOrUpdateParams {
	var ()
	return &TagsCreateOrUpdateParams{

		Context: ctx,
	}
}

// NewTagsCreateOrUpdateParamsWithHTTPClient creates a new TagsCreateOrUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTagsCreateOrUpdateParamsWithHTTPClient(client *http.Client) *TagsCreateOrUpdateParams {
	var ()
	return &TagsCreateOrUpdateParams{
		HTTPClient: client,
	}
}

/*TagsCreateOrUpdateParams contains all the parameters to send to the API endpoint
for the tags create or update operation typically these are written to a http.Request
*/
type TagsCreateOrUpdateParams struct {

	/*APIVersion
	  The API version to use for this operation.

	*/
	APIVersion string
	/*SubscriptionID
	  The ID of the target subscription.

	*/
	SubscriptionID string
	/*TagName
	  The name of the tag to create.

	*/
	TagName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tags create or update params
func (o *TagsCreateOrUpdateParams) WithTimeout(timeout time.Duration) *TagsCreateOrUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tags create or update params
func (o *TagsCreateOrUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tags create or update params
func (o *TagsCreateOrUpdateParams) WithContext(ctx context.Context) *TagsCreateOrUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tags create or update params
func (o *TagsCreateOrUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tags create or update params
func (o *TagsCreateOrUpdateParams) WithHTTPClient(client *http.Client) *TagsCreateOrUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tags create or update params
func (o *TagsCreateOrUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the tags create or update params
func (o *TagsCreateOrUpdateParams) WithAPIVersion(aPIVersion string) *TagsCreateOrUpdateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the tags create or update params
func (o *TagsCreateOrUpdateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithSubscriptionID adds the subscriptionID to the tags create or update params
func (o *TagsCreateOrUpdateParams) WithSubscriptionID(subscriptionID string) *TagsCreateOrUpdateParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the tags create or update params
func (o *TagsCreateOrUpdateParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithTagName adds the tagName to the tags create or update params
func (o *TagsCreateOrUpdateParams) WithTagName(tagName string) *TagsCreateOrUpdateParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the tags create or update params
func (o *TagsCreateOrUpdateParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WriteToRequest writes these params to a swagger request
func (o *TagsCreateOrUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
