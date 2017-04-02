package disks

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

// NewDisksRevokeAccessParams creates a new DisksRevokeAccessParams object
// with the default values initialized.
func NewDisksRevokeAccessParams() *DisksRevokeAccessParams {
	var ()
	return &DisksRevokeAccessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDisksRevokeAccessParamsWithTimeout creates a new DisksRevokeAccessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDisksRevokeAccessParamsWithTimeout(timeout time.Duration) *DisksRevokeAccessParams {
	var ()
	return &DisksRevokeAccessParams{

		timeout: timeout,
	}
}

// NewDisksRevokeAccessParamsWithContext creates a new DisksRevokeAccessParams object
// with the default values initialized, and the ability to set a context for a request
func NewDisksRevokeAccessParamsWithContext(ctx context.Context) *DisksRevokeAccessParams {
	var ()
	return &DisksRevokeAccessParams{

		Context: ctx,
	}
}

/*DisksRevokeAccessParams contains all the parameters to send to the API endpoint
for the disks revoke access operation typically these are written to a http.Request
*/
type DisksRevokeAccessParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*DiskName
	  The name of the disk within the given subscription and resource group.

	*/
	DiskName string
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

// WithTimeout adds the timeout to the disks revoke access params
func (o *DisksRevokeAccessParams) WithTimeout(timeout time.Duration) *DisksRevokeAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disks revoke access params
func (o *DisksRevokeAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disks revoke access params
func (o *DisksRevokeAccessParams) WithContext(ctx context.Context) *DisksRevokeAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disks revoke access params
func (o *DisksRevokeAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the disks revoke access params
func (o *DisksRevokeAccessParams) WithAPIVersion(aPIVersion string) *DisksRevokeAccessParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the disks revoke access params
func (o *DisksRevokeAccessParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDiskName adds the diskName to the disks revoke access params
func (o *DisksRevokeAccessParams) WithDiskName(diskName string) *DisksRevokeAccessParams {
	o.SetDiskName(diskName)
	return o
}

// SetDiskName adds the diskName to the disks revoke access params
func (o *DisksRevokeAccessParams) SetDiskName(diskName string) {
	o.DiskName = diskName
}

// WithResourceGroupName adds the resourceGroupName to the disks revoke access params
func (o *DisksRevokeAccessParams) WithResourceGroupName(resourceGroupName string) *DisksRevokeAccessParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the disks revoke access params
func (o *DisksRevokeAccessParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the disks revoke access params
func (o *DisksRevokeAccessParams) WithSubscriptionID(subscriptionID string) *DisksRevokeAccessParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the disks revoke access params
func (o *DisksRevokeAccessParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *DisksRevokeAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param diskName
	if err := r.SetPathParam("diskName", o.DiskName); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
