package virtual_machine_scale_sets

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

	"github.com/jkawamoto/roadie/cloud/azure/compute/models"
)

// NewVirtualMachineScaleSetsStartParams creates a new VirtualMachineScaleSetsStartParams object
// with the default values initialized.
func NewVirtualMachineScaleSetsStartParams() *VirtualMachineScaleSetsStartParams {
	var ()
	return &VirtualMachineScaleSetsStartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualMachineScaleSetsStartParamsWithTimeout creates a new VirtualMachineScaleSetsStartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualMachineScaleSetsStartParamsWithTimeout(timeout time.Duration) *VirtualMachineScaleSetsStartParams {
	var ()
	return &VirtualMachineScaleSetsStartParams{

		timeout: timeout,
	}
}

// NewVirtualMachineScaleSetsStartParamsWithContext creates a new VirtualMachineScaleSetsStartParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualMachineScaleSetsStartParamsWithContext(ctx context.Context) *VirtualMachineScaleSetsStartParams {
	var ()
	return &VirtualMachineScaleSetsStartParams{

		Context: ctx,
	}
}

/*VirtualMachineScaleSetsStartParams contains all the parameters to send to the API endpoint
for the virtual machine scale sets start operation typically these are written to a http.Request
*/
type VirtualMachineScaleSetsStartParams struct {

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
	/*VMInstanceIds
	  A list of virtual machine instance IDs from the VM scale set.

	*/
	VMInstanceIds *models.VirtualMachineScaleSetVMInstanceIds
	/*VMScaleSetName
	  The name of the VM scale set.

	*/
	VMScaleSetName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) WithTimeout(timeout time.Duration) *VirtualMachineScaleSetsStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) WithContext(ctx context.Context) *VirtualMachineScaleSetsStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) WithAPIVersion(aPIVersion string) *VirtualMachineScaleSetsStartParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) WithResourceGroupName(resourceGroupName string) *VirtualMachineScaleSetsStartParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) WithSubscriptionID(subscriptionID string) *VirtualMachineScaleSetsStartParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVMInstanceIds adds the vMInstanceIds to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) WithVMInstanceIds(vMInstanceIds *models.VirtualMachineScaleSetVMInstanceIds) *VirtualMachineScaleSetsStartParams {
	o.SetVMInstanceIds(vMInstanceIds)
	return o
}

// SetVMInstanceIds adds the vmInstanceIds to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) SetVMInstanceIds(vMInstanceIds *models.VirtualMachineScaleSetVMInstanceIds) {
	o.VMInstanceIds = vMInstanceIds
}

// WithVMScaleSetName adds the vMScaleSetName to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) WithVMScaleSetName(vMScaleSetName string) *VirtualMachineScaleSetsStartParams {
	o.SetVMScaleSetName(vMScaleSetName)
	return o
}

// SetVMScaleSetName adds the vmScaleSetName to the virtual machine scale sets start params
func (o *VirtualMachineScaleSetsStartParams) SetVMScaleSetName(vMScaleSetName string) {
	o.VMScaleSetName = vMScaleSetName
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualMachineScaleSetsStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.VMInstanceIds == nil {
		o.VMInstanceIds = new(models.VirtualMachineScaleSetVMInstanceIds)
	}

	if err := r.SetBodyParam(o.VMInstanceIds); err != nil {
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
