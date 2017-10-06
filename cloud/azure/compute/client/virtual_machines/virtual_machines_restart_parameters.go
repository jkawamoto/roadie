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

// NewVirtualMachinesRestartParams creates a new VirtualMachinesRestartParams object
// with the default values initialized.
func NewVirtualMachinesRestartParams() *VirtualMachinesRestartParams {
	var ()
	return &VirtualMachinesRestartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualMachinesRestartParamsWithTimeout creates a new VirtualMachinesRestartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualMachinesRestartParamsWithTimeout(timeout time.Duration) *VirtualMachinesRestartParams {
	var ()
	return &VirtualMachinesRestartParams{

		timeout: timeout,
	}
}

// NewVirtualMachinesRestartParamsWithContext creates a new VirtualMachinesRestartParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualMachinesRestartParamsWithContext(ctx context.Context) *VirtualMachinesRestartParams {
	var ()
	return &VirtualMachinesRestartParams{

		Context: ctx,
	}
}

/*VirtualMachinesRestartParams contains all the parameters to send to the API endpoint
for the virtual machines restart operation typically these are written to a http.Request
*/
type VirtualMachinesRestartParams struct {

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

// WithTimeout adds the timeout to the virtual machines restart params
func (o *VirtualMachinesRestartParams) WithTimeout(timeout time.Duration) *VirtualMachinesRestartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual machines restart params
func (o *VirtualMachinesRestartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual machines restart params
func (o *VirtualMachinesRestartParams) WithContext(ctx context.Context) *VirtualMachinesRestartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual machines restart params
func (o *VirtualMachinesRestartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the virtual machines restart params
func (o *VirtualMachinesRestartParams) WithAPIVersion(aPIVersion string) *VirtualMachinesRestartParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the virtual machines restart params
func (o *VirtualMachinesRestartParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the virtual machines restart params
func (o *VirtualMachinesRestartParams) WithResourceGroupName(resourceGroupName string) *VirtualMachinesRestartParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the virtual machines restart params
func (o *VirtualMachinesRestartParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the virtual machines restart params
func (o *VirtualMachinesRestartParams) WithSubscriptionID(subscriptionID string) *VirtualMachinesRestartParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the virtual machines restart params
func (o *VirtualMachinesRestartParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVMName adds the vMName to the virtual machines restart params
func (o *VirtualMachinesRestartParams) WithVMName(vMName string) *VirtualMachinesRestartParams {
	o.SetVMName(vMName)
	return o
}

// SetVMName adds the vmName to the virtual machines restart params
func (o *VirtualMachinesRestartParams) SetVMName(vMName string) {
	o.VMName = vMName
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualMachinesRestartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
