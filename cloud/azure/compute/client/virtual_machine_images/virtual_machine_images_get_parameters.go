package virtual_machine_images

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

// NewVirtualMachineImagesGetParams creates a new VirtualMachineImagesGetParams object
// with the default values initialized.
func NewVirtualMachineImagesGetParams() *VirtualMachineImagesGetParams {
	var ()
	return &VirtualMachineImagesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualMachineImagesGetParamsWithTimeout creates a new VirtualMachineImagesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualMachineImagesGetParamsWithTimeout(timeout time.Duration) *VirtualMachineImagesGetParams {
	var ()
	return &VirtualMachineImagesGetParams{

		timeout: timeout,
	}
}

// NewVirtualMachineImagesGetParamsWithContext creates a new VirtualMachineImagesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualMachineImagesGetParamsWithContext(ctx context.Context) *VirtualMachineImagesGetParams {
	var ()
	return &VirtualMachineImagesGetParams{

		Context: ctx,
	}
}

/*VirtualMachineImagesGetParams contains all the parameters to send to the API endpoint
for the virtual machine images get operation typically these are written to a http.Request
*/
type VirtualMachineImagesGetParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*Location
	  The name of a supported Azure region.

	*/
	Location string
	/*Offer
	  A valid image publisher offer.

	*/
	Offer string
	/*PublisherName
	  A valid image publisher.

	*/
	PublisherName string
	/*Skus
	  A valid image SKU.

	*/
	Skus string
	/*SubscriptionID
	  Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string
	/*Version
	  A valid image SKU version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) WithTimeout(timeout time.Duration) *VirtualMachineImagesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) WithContext(ctx context.Context) *VirtualMachineImagesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) WithAPIVersion(aPIVersion string) *VirtualMachineImagesGetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithLocation adds the location to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) WithLocation(location string) *VirtualMachineImagesGetParams {
	o.SetLocation(location)
	return o
}

// SetLocation adds the location to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) SetLocation(location string) {
	o.Location = location
}

// WithOffer adds the offer to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) WithOffer(offer string) *VirtualMachineImagesGetParams {
	o.SetOffer(offer)
	return o
}

// SetOffer adds the offer to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) SetOffer(offer string) {
	o.Offer = offer
}

// WithPublisherName adds the publisherName to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) WithPublisherName(publisherName string) *VirtualMachineImagesGetParams {
	o.SetPublisherName(publisherName)
	return o
}

// SetPublisherName adds the publisherName to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) SetPublisherName(publisherName string) {
	o.PublisherName = publisherName
}

// WithSkus adds the skus to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) WithSkus(skus string) *VirtualMachineImagesGetParams {
	o.SetSkus(skus)
	return o
}

// SetSkus adds the skus to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) SetSkus(skus string) {
	o.Skus = skus
}

// WithSubscriptionID adds the subscriptionID to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) WithSubscriptionID(subscriptionID string) *VirtualMachineImagesGetParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVersion adds the version to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) WithVersion(version string) *VirtualMachineImagesGetParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the virtual machine images get params
func (o *VirtualMachineImagesGetParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualMachineImagesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param location
	if err := r.SetPathParam("location", o.Location); err != nil {
		return err
	}

	// path param offer
	if err := r.SetPathParam("offer", o.Offer); err != nil {
		return err
	}

	// path param publisherName
	if err := r.SetPathParam("publisherName", o.PublisherName); err != nil {
		return err
	}

	// path param skus
	if err := r.SetPathParam("skus", o.Skus); err != nil {
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