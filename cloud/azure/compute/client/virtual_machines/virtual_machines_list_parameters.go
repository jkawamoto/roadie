package virtual_machines

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

// NewVirtualMachinesListParams creates a new VirtualMachinesListParams object
// with the default values initialized.
func NewVirtualMachinesListParams() *VirtualMachinesListParams {
	var ()
	return &VirtualMachinesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualMachinesListParamsWithTimeout creates a new VirtualMachinesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualMachinesListParamsWithTimeout(timeout time.Duration) *VirtualMachinesListParams {
	var ()
	return &VirtualMachinesListParams{

		timeout: timeout,
	}
}

// NewVirtualMachinesListParamsWithContext creates a new VirtualMachinesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualMachinesListParamsWithContext(ctx context.Context) *VirtualMachinesListParams {
	var ()
	return &VirtualMachinesListParams{

		Context: ctx,
	}
}

/*VirtualMachinesListParams contains all the parameters to send to the API endpoint
for the virtual machines list operation typically these are written to a http.Request
*/
type VirtualMachinesListParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
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

// WithTimeout adds the timeout to the virtual machines list params
func (o *VirtualMachinesListParams) WithTimeout(timeout time.Duration) *VirtualMachinesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual machines list params
func (o *VirtualMachinesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual machines list params
func (o *VirtualMachinesListParams) WithContext(ctx context.Context) *VirtualMachinesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual machines list params
func (o *VirtualMachinesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the virtual machines list params
func (o *VirtualMachinesListParams) WithAPIVersion(aPIVersion string) *VirtualMachinesListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the virtual machines list params
func (o *VirtualMachinesListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the virtual machines list params
func (o *VirtualMachinesListParams) WithResourceGroupName(resourceGroupName string) *VirtualMachinesListParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the virtual machines list params
func (o *VirtualMachinesListParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the virtual machines list params
func (o *VirtualMachinesListParams) WithSubscriptionID(subscriptionID string) *VirtualMachinesListParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the virtual machines list params
func (o *VirtualMachinesListParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualMachinesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
