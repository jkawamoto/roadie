package virtual_machine_sizes

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

// NewVirtualMachineSizesListParams creates a new VirtualMachineSizesListParams object
// with the default values initialized.
func NewVirtualMachineSizesListParams() *VirtualMachineSizesListParams {
	var ()
	return &VirtualMachineSizesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualMachineSizesListParamsWithTimeout creates a new VirtualMachineSizesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualMachineSizesListParamsWithTimeout(timeout time.Duration) *VirtualMachineSizesListParams {
	var ()
	return &VirtualMachineSizesListParams{

		timeout: timeout,
	}
}

// NewVirtualMachineSizesListParamsWithContext creates a new VirtualMachineSizesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualMachineSizesListParamsWithContext(ctx context.Context) *VirtualMachineSizesListParams {
	var ()
	return &VirtualMachineSizesListParams{

		Context: ctx,
	}
}

/*VirtualMachineSizesListParams contains all the parameters to send to the API endpoint
for the virtual machine sizes list operation typically these are written to a http.Request
*/
type VirtualMachineSizesListParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*Location
	  The location upon which virtual-machine-sizes is queried.

	*/
	Location string
	/*SubscriptionID
	  Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtual machine sizes list params
func (o *VirtualMachineSizesListParams) WithTimeout(timeout time.Duration) *VirtualMachineSizesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual machine sizes list params
func (o *VirtualMachineSizesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual machine sizes list params
func (o *VirtualMachineSizesListParams) WithContext(ctx context.Context) *VirtualMachineSizesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual machine sizes list params
func (o *VirtualMachineSizesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the virtual machine sizes list params
func (o *VirtualMachineSizesListParams) WithAPIVersion(aPIVersion string) *VirtualMachineSizesListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the virtual machine sizes list params
func (o *VirtualMachineSizesListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithLocation adds the location to the virtual machine sizes list params
func (o *VirtualMachineSizesListParams) WithLocation(location string) *VirtualMachineSizesListParams {
	o.SetLocation(location)
	return o
}

// SetLocation adds the location to the virtual machine sizes list params
func (o *VirtualMachineSizesListParams) SetLocation(location string) {
	o.Location = location
}

// WithSubscriptionID adds the subscriptionID to the virtual machine sizes list params
func (o *VirtualMachineSizesListParams) WithSubscriptionID(subscriptionID string) *VirtualMachineSizesListParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the virtual machine sizes list params
func (o *VirtualMachineSizesListParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualMachineSizesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}