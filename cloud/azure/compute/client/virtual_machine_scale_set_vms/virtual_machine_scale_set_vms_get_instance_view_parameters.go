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

// NewVirtualMachineScaleSetVmsGetInstanceViewParams creates a new VirtualMachineScaleSetVmsGetInstanceViewParams object
// with the default values initialized.
func NewVirtualMachineScaleSetVmsGetInstanceViewParams() *VirtualMachineScaleSetVmsGetInstanceViewParams {
	var ()
	return &VirtualMachineScaleSetVmsGetInstanceViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualMachineScaleSetVmsGetInstanceViewParamsWithTimeout creates a new VirtualMachineScaleSetVmsGetInstanceViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualMachineScaleSetVmsGetInstanceViewParamsWithTimeout(timeout time.Duration) *VirtualMachineScaleSetVmsGetInstanceViewParams {
	var ()
	return &VirtualMachineScaleSetVmsGetInstanceViewParams{

		timeout: timeout,
	}
}

// NewVirtualMachineScaleSetVmsGetInstanceViewParamsWithContext creates a new VirtualMachineScaleSetVmsGetInstanceViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualMachineScaleSetVmsGetInstanceViewParamsWithContext(ctx context.Context) *VirtualMachineScaleSetVmsGetInstanceViewParams {
	var ()
	return &VirtualMachineScaleSetVmsGetInstanceViewParams{

		Context: ctx,
	}
}

/*VirtualMachineScaleSetVmsGetInstanceViewParams contains all the parameters to send to the API endpoint
for the virtual machine scale set vms get instance view operation typically these are written to a http.Request
*/
type VirtualMachineScaleSetVmsGetInstanceViewParams struct {

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

// WithTimeout adds the timeout to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) WithTimeout(timeout time.Duration) *VirtualMachineScaleSetVmsGetInstanceViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) WithContext(ctx context.Context) *VirtualMachineScaleSetVmsGetInstanceViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) WithAPIVersion(aPIVersion string) *VirtualMachineScaleSetVmsGetInstanceViewParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithInstanceID adds the instanceID to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) WithInstanceID(instanceID string) *VirtualMachineScaleSetVmsGetInstanceViewParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithResourceGroupName adds the resourceGroupName to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) WithResourceGroupName(resourceGroupName string) *VirtualMachineScaleSetVmsGetInstanceViewParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) WithSubscriptionID(subscriptionID string) *VirtualMachineScaleSetVmsGetInstanceViewParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVMScaleSetName adds the vMScaleSetName to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) WithVMScaleSetName(vMScaleSetName string) *VirtualMachineScaleSetVmsGetInstanceViewParams {
	o.SetVMScaleSetName(vMScaleSetName)
	return o
}

// SetVMScaleSetName adds the vmScaleSetName to the virtual machine scale set vms get instance view params
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) SetVMScaleSetName(vMScaleSetName string) {
	o.VMScaleSetName = vMScaleSetName
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualMachineScaleSetVmsGetInstanceViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
