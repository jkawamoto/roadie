package virtual_machine_scale_set_vms

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

// NewVirtualMachineScaleSetVmsDeallocateParams creates a new VirtualMachineScaleSetVmsDeallocateParams object
// with the default values initialized.
func NewVirtualMachineScaleSetVmsDeallocateParams() *VirtualMachineScaleSetVmsDeallocateParams {
	var ()
	return &VirtualMachineScaleSetVmsDeallocateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualMachineScaleSetVmsDeallocateParamsWithTimeout creates a new VirtualMachineScaleSetVmsDeallocateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualMachineScaleSetVmsDeallocateParamsWithTimeout(timeout time.Duration) *VirtualMachineScaleSetVmsDeallocateParams {
	var ()
	return &VirtualMachineScaleSetVmsDeallocateParams{

		timeout: timeout,
	}
}

// NewVirtualMachineScaleSetVmsDeallocateParamsWithContext creates a new VirtualMachineScaleSetVmsDeallocateParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualMachineScaleSetVmsDeallocateParamsWithContext(ctx context.Context) *VirtualMachineScaleSetVmsDeallocateParams {
	var ()
	return &VirtualMachineScaleSetVmsDeallocateParams{

		Context: ctx,
	}
}

/*VirtualMachineScaleSetVmsDeallocateParams contains all the parameters to send to the API endpoint
for the virtual machine scale set vms deallocate operation typically these are written to a http.Request
*/
type VirtualMachineScaleSetVmsDeallocateParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*InstanceID
	  The instance ID of the virtual machine.

	*/
	InstanceID string
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*SubscriptionID
	  Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string
	/*VMScaleSetName
	  The name of the VM scale set.

	*/
	VMScaleSetName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) WithTimeout(timeout time.Duration) *VirtualMachineScaleSetVmsDeallocateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) WithContext(ctx context.Context) *VirtualMachineScaleSetVmsDeallocateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) WithAPIVersion(aPIVersion string) *VirtualMachineScaleSetVmsDeallocateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithInstanceID adds the instanceID to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) WithInstanceID(instanceID string) *VirtualMachineScaleSetVmsDeallocateParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithResourceGroupName adds the resourceGroupName to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) WithResourceGroupName(resourceGroupName string) *VirtualMachineScaleSetVmsDeallocateParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) WithSubscriptionID(subscriptionID string) *VirtualMachineScaleSetVmsDeallocateParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVMScaleSetName adds the vMScaleSetName to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) WithVMScaleSetName(vMScaleSetName string) *VirtualMachineScaleSetVmsDeallocateParams {
	o.SetVMScaleSetName(vMScaleSetName)
	return o
}

// SetVMScaleSetName adds the vmScaleSetName to the virtual machine scale set vms deallocate params
func (o *VirtualMachineScaleSetVmsDeallocateParams) SetVMScaleSetName(vMScaleSetName string) {
	o.VMScaleSetName = vMScaleSetName
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualMachineScaleSetVmsDeallocateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param instanceId
	if err := r.SetPathParam("instanceId", o.InstanceID); err != nil {
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

	// path param vmScaleSetName
	if err := r.SetPathParam("vmScaleSetName", o.VMScaleSetName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
