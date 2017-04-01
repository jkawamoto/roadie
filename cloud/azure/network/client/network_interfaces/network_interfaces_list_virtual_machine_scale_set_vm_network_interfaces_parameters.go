package network_interfaces

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

// NewNetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams creates a new NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams object
// with the default values initialized.
func NewNetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams() *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams {
	var ()
	return &NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParamsWithTimeout creates a new NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParamsWithTimeout(timeout time.Duration) *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams {
	var ()
	return &NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams{

		timeout: timeout,
	}
}

// NewNetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParamsWithContext creates a new NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewNetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParamsWithContext(ctx context.Context) *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams {
	var ()
	return &NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams{

		Context: ctx,
	}
}

/*NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams contains all the parameters to send to the API endpoint
for the network interfaces list virtual machine scale set VM network interfaces operation typically these are written to a http.Request
*/
type NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*SubscriptionID
	  Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string
	/*VirtualMachineScaleSetName
	  The name of the virtual machine scale set.

	*/
	VirtualMachineScaleSetName string
	/*VirtualmachineIndex
	  The virtual machine index.

	*/
	VirtualmachineIndex string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) WithTimeout(timeout time.Duration) *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) WithContext(ctx context.Context) *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) WithAPIVersion(aPIVersion string) *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) WithResourceGroupName(resourceGroupName string) *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) WithSubscriptionID(subscriptionID string) *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVirtualMachineScaleSetName adds the virtualMachineScaleSetName to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) WithVirtualMachineScaleSetName(virtualMachineScaleSetName string) *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams {
	o.SetVirtualMachineScaleSetName(virtualMachineScaleSetName)
	return o
}

// SetVirtualMachineScaleSetName adds the virtualMachineScaleSetName to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) SetVirtualMachineScaleSetName(virtualMachineScaleSetName string) {
	o.VirtualMachineScaleSetName = virtualMachineScaleSetName
}

// WithVirtualmachineIndex adds the virtualmachineIndex to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) WithVirtualmachineIndex(virtualmachineIndex string) *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams {
	o.SetVirtualmachineIndex(virtualmachineIndex)
	return o
}

// SetVirtualmachineIndex adds the virtualmachineIndex to the network interfaces list virtual machine scale set VM network interfaces params
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) SetVirtualmachineIndex(virtualmachineIndex string) {
	o.VirtualmachineIndex = virtualmachineIndex
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param virtualMachineScaleSetName
	if err := r.SetPathParam("virtualMachineScaleSetName", o.VirtualMachineScaleSetName); err != nil {
		return err
	}

	// path param virtualmachineIndex
	if err := r.SetPathParam("virtualmachineIndex", o.VirtualmachineIndex); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
