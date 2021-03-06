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

// NewVirtualMachinesPowerOffParams creates a new VirtualMachinesPowerOffParams object
// with the default values initialized.
func NewVirtualMachinesPowerOffParams() *VirtualMachinesPowerOffParams {
	var ()
	return &VirtualMachinesPowerOffParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualMachinesPowerOffParamsWithTimeout creates a new VirtualMachinesPowerOffParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualMachinesPowerOffParamsWithTimeout(timeout time.Duration) *VirtualMachinesPowerOffParams {
	var ()
	return &VirtualMachinesPowerOffParams{

		timeout: timeout,
	}
}

// NewVirtualMachinesPowerOffParamsWithContext creates a new VirtualMachinesPowerOffParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualMachinesPowerOffParamsWithContext(ctx context.Context) *VirtualMachinesPowerOffParams {
	var ()
	return &VirtualMachinesPowerOffParams{

		Context: ctx,
	}
}

/*VirtualMachinesPowerOffParams contains all the parameters to send to the API endpoint
for the virtual machines power off operation typically these are written to a http.Request
*/
type VirtualMachinesPowerOffParams struct {

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

// WithTimeout adds the timeout to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) WithTimeout(timeout time.Duration) *VirtualMachinesPowerOffParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) WithContext(ctx context.Context) *VirtualMachinesPowerOffParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) WithAPIVersion(aPIVersion string) *VirtualMachinesPowerOffParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) WithResourceGroupName(resourceGroupName string) *VirtualMachinesPowerOffParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) WithSubscriptionID(subscriptionID string) *VirtualMachinesPowerOffParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVMName adds the vMName to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) WithVMName(vMName string) *VirtualMachinesPowerOffParams {
	o.SetVMName(vMName)
	return o
}

// SetVMName adds the vmName to the virtual machines power off params
func (o *VirtualMachinesPowerOffParams) SetVMName(vMName string) {
	o.VMName = vMName
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualMachinesPowerOffParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
