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

// NewNetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams creates a new NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams object
// with the default values initialized.
func NewNetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams() *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	var ()
	return &NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParamsWithTimeout creates a new NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParamsWithTimeout(timeout time.Duration) *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	var ()
	return &NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams{

		timeout: timeout,
	}
}

// NewNetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParamsWithContext creates a new NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewNetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParamsWithContext(ctx context.Context) *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	var ()
	return &NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams{

		Context: ctx,
	}
}

/*NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams contains all the parameters to send to the API endpoint
for the network interfaces get virtual machine scale set network interface operation typically these are written to a http.Request
*/
type NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams struct {

	/*NrDollarExpand
	  expand references resources.

	*/
	DollarExpand *string
	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*NetworkInterfaceName
	  The name of the network interface.

	*/
	NetworkInterfaceName string
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

// WithTimeout adds the timeout to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) WithTimeout(timeout time.Duration) *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) WithContext(ctx context.Context) *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDollarExpand adds the dollarExpand to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) WithDollarExpand(dollarExpand *string) *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	o.SetDollarExpand(dollarExpand)
	return o
}

// SetDollarExpand adds the dollarExpand to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) SetDollarExpand(dollarExpand *string) {
	o.DollarExpand = dollarExpand
}

// WithAPIVersion adds the aPIVersion to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) WithAPIVersion(aPIVersion string) *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithNetworkInterfaceName adds the networkInterfaceName to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) WithNetworkInterfaceName(networkInterfaceName string) *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	o.SetNetworkInterfaceName(networkInterfaceName)
	return o
}

// SetNetworkInterfaceName adds the networkInterfaceName to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) SetNetworkInterfaceName(networkInterfaceName string) {
	o.NetworkInterfaceName = networkInterfaceName
}

// WithResourceGroupName adds the resourceGroupName to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) WithResourceGroupName(resourceGroupName string) *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) WithSubscriptionID(subscriptionID string) *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WithVirtualMachineScaleSetName adds the virtualMachineScaleSetName to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) WithVirtualMachineScaleSetName(virtualMachineScaleSetName string) *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	o.SetVirtualMachineScaleSetName(virtualMachineScaleSetName)
	return o
}

// SetVirtualMachineScaleSetName adds the virtualMachineScaleSetName to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) SetVirtualMachineScaleSetName(virtualMachineScaleSetName string) {
	o.VirtualMachineScaleSetName = virtualMachineScaleSetName
}

// WithVirtualmachineIndex adds the virtualmachineIndex to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) WithVirtualmachineIndex(virtualmachineIndex string) *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams {
	o.SetVirtualmachineIndex(virtualmachineIndex)
	return o
}

// SetVirtualmachineIndex adds the virtualmachineIndex to the network interfaces get virtual machine scale set network interface params
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) SetVirtualmachineIndex(virtualmachineIndex string) {
	o.VirtualmachineIndex = virtualmachineIndex
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.DollarExpand != nil {

		// query param $expand
		var qrNrDollarExpand string
		if o.DollarExpand != nil {
			qrNrDollarExpand = *o.DollarExpand
		}
		qNrDollarExpand := qrNrDollarExpand
		if qNrDollarExpand != "" {
			if err := r.SetQueryParam("$expand", qNrDollarExpand); err != nil {
				return err
			}
		}

	}

	// query param api-version
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("api-version", qAPIVersion); err != nil {
			return err
		}
	}

	// path param networkInterfaceName
	if err := r.SetPathParam("networkInterfaceName", o.NetworkInterfaceName); err != nil {
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
