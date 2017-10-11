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

// NewVirtualMachinesStartParams creates a new VirtualMachinesStartParams object
// with the default values initialized.
func NewVirtualMachinesStartParams() *VirtualMachinesStartParams {
	var ()
	return &VirtualMachinesStartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualMachinesStartParamsWithTimeout creates a new VirtualMachinesStartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualMachinesStartParamsWithTimeout(timeout time.Duration) *VirtualMachinesStartParams {
	var ()
	return &VirtualMachinesStartParams{

		timeout: timeout,
	}
}

// NewVirtualMachinesStartParamsWithContext creates a new VirtualMachinesStartParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualMachinesStartParamsWithContext(ctx context.Context) *VirtualMachinesStartParams {
	var ()
	return &VirtualMachinesStartParams{

		Context: ctx,
	}
}

/*VirtualMachinesStartParams contains all the parameters to send to the API endpoint
for the virtual machines start operation typically these are written to a http.Request
*/
type VirtualMachinesStartParams struct {

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
	/*VMName
	  The name of the virtual machine.

	*/
	VMName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtual machines start params
func (o *VirtualMachinesStartParams) WithTimeout(timeout time.Duration) *VirtualMachinesStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual machines start params
func (o *VirtualMachinesStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual machines start params
func (o *VirtualMachinesStartParams) WithContext(ctx context.Context) *VirtualMachinesStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual machines start params
func (o *VirtualMachinesStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the virtual machines start params
func (o *VirtualMachinesStartParams) WithAPIVersion(aPIVersion string) *VirtualMachinesStartParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the virtual machines start params
func (o *VirtualMachinesStartParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the virtual machines start params
func (o *VirtualMachinesStartParams) WithResourceGroupName(resourceGroupName string) *VirtualMachinesStartParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the virtual machines start params
func (o *VirtualMachinesStartParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the virtual machines start params
func (o *VirtualMachinesStartParams) WithSubscriptionID(subscriptionID string) *VirtualMachinesStartParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the virtual machines start params
func (o *VirtualMachinesStartParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVMName adds the vMName to the virtual machines start params
func (o *VirtualMachinesStartParams) WithVMName(vMName string) *VirtualMachinesStartParams {
	o.SetVMName(vMName)
	return o
}

// SetVMName adds the vmName to the virtual machines start params
func (o *VirtualMachinesStartParams) SetVMName(vMName string) {
	o.VMName = vMName
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualMachinesStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param vmName
	if err := r.SetPathParam("vmName", o.VMName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}