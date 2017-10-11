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

// NewImagesDeleteParams creates a new ImagesDeleteParams object
// with the default values initialized.
func NewImagesDeleteParams() *ImagesDeleteParams {
	var ()
	return &ImagesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImagesDeleteParamsWithTimeout creates a new ImagesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImagesDeleteParamsWithTimeout(timeout time.Duration) *ImagesDeleteParams {
	var ()
	return &ImagesDeleteParams{

		timeout: timeout,
	}
}

// NewImagesDeleteParamsWithContext creates a new ImagesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewImagesDeleteParamsWithContext(ctx context.Context) *ImagesDeleteParams {
	var ()
	return &ImagesDeleteParams{

		Context: ctx,
	}
}

/*ImagesDeleteParams contains all the parameters to send to the API endpoint
for the images delete operation typically these are written to a http.Request
*/
type ImagesDeleteParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*ImageName
	  The name of the image.

	*/
	ImageName string
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*SubscriptionID
	  Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the images delete params
func (o *ImagesDeleteParams) WithTimeout(timeout time.Duration) *ImagesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the images delete params
func (o *ImagesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the images delete params
func (o *ImagesDeleteParams) WithContext(ctx context.Context) *ImagesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the images delete params
func (o *ImagesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the images delete params
func (o *ImagesDeleteParams) WithAPIVersion(aPIVersion string) *ImagesDeleteParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the images delete params
func (o *ImagesDeleteParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithImageName adds the imageName to the images delete params
func (o *ImagesDeleteParams) WithImageName(imageName string) *ImagesDeleteParams {
	o.SetImageName(imageName)
	return o
}

// SetImageName adds the imageName to the images delete params
func (o *ImagesDeleteParams) SetImageName(imageName string) {
	o.ImageName = imageName
}

// WithResourceGroupName adds the resourceGroupName to the images delete params
func (o *ImagesDeleteParams) WithResourceGroupName(resourceGroupName string) *ImagesDeleteParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the images delete params
func (o *ImagesDeleteParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the images delete params
func (o *ImagesDeleteParams) WithSubscriptionID(subscriptionID string) *ImagesDeleteParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the images delete params
func (o *ImagesDeleteParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ImagesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param imageName
	if err := r.SetPathParam("imageName", o.ImageName); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}