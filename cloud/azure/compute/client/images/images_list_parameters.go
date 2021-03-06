package images

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

// NewImagesListParams creates a new ImagesListParams object
// with the default values initialized.
func NewImagesListParams() *ImagesListParams {
	var ()
	return &ImagesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImagesListParamsWithTimeout creates a new ImagesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImagesListParamsWithTimeout(timeout time.Duration) *ImagesListParams {
	var ()
	return &ImagesListParams{

		timeout: timeout,
	}
}

// NewImagesListParamsWithContext creates a new ImagesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewImagesListParamsWithContext(ctx context.Context) *ImagesListParams {
	var ()
	return &ImagesListParams{

		Context: ctx,
	}
}

/*ImagesListParams contains all the parameters to send to the API endpoint
for the images list operation typically these are written to a http.Request
*/
type ImagesListParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*SubscriptionID
	  Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the images list params
func (o *ImagesListParams) WithTimeout(timeout time.Duration) *ImagesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the images list params
func (o *ImagesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the images list params
func (o *ImagesListParams) WithContext(ctx context.Context) *ImagesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the images list params
func (o *ImagesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the images list params
func (o *ImagesListParams) WithAPIVersion(aPIVersion string) *ImagesListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the images list params
func (o *ImagesListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithSubscriptionID adds the subscriptionID to the images list params
func (o *ImagesListParams) WithSubscriptionID(subscriptionID string) *ImagesListParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the images list params
func (o *ImagesListParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ImagesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
