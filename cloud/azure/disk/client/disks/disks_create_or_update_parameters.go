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

	"github.com/jkawamoto/roadie/cloud/azure/disk/models"
)

// NewDisksCreateOrUpdateParams creates a new DisksCreateOrUpdateParams object
// with the default values initialized.
func NewDisksCreateOrUpdateParams() *DisksCreateOrUpdateParams {
	var ()
	return &DisksCreateOrUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDisksCreateOrUpdateParamsWithTimeout creates a new DisksCreateOrUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDisksCreateOrUpdateParamsWithTimeout(timeout time.Duration) *DisksCreateOrUpdateParams {
	var ()
	return &DisksCreateOrUpdateParams{

		timeout: timeout,
	}
}

// NewDisksCreateOrUpdateParamsWithContext creates a new DisksCreateOrUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDisksCreateOrUpdateParamsWithContext(ctx context.Context) *DisksCreateOrUpdateParams {
	var ()
	return &DisksCreateOrUpdateParams{

		Context: ctx,
	}
}

/*DisksCreateOrUpdateParams contains all the parameters to send to the API endpoint
for the disks create or update operation typically these are written to a http.Request
*/
type DisksCreateOrUpdateParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*Disk
	  Disk object supplied in the body of the Put disk operation.

	*/
	Disk *models.Disk
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

// WithTimeout adds the timeout to the disks create or update params
func (o *DisksCreateOrUpdateParams) WithTimeout(timeout time.Duration) *DisksCreateOrUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disks create or update params
func (o *DisksCreateOrUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disks create or update params
func (o *DisksCreateOrUpdateParams) WithContext(ctx context.Context) *DisksCreateOrUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disks create or update params
func (o *DisksCreateOrUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the disks create or update params
func (o *DisksCreateOrUpdateParams) WithAPIVersion(aPIVersion string) *DisksCreateOrUpdateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the disks create or update params
func (o *DisksCreateOrUpdateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDisk adds the disk to the disks create or update params
func (o *DisksCreateOrUpdateParams) WithDisk(disk *models.Disk) *DisksCreateOrUpdateParams {
	o.SetDisk(disk)
	return o
}

// SetDisk adds the disk to the disks create or update params
func (o *DisksCreateOrUpdateParams) SetDisk(disk *models.Disk) {
	o.Disk = disk
}

// WithDiskName adds the diskName to the disks create or update params
func (o *DisksCreateOrUpdateParams) WithDiskName(diskName string) *DisksCreateOrUpdateParams {
	o.SetDiskName(diskName)
	return o
}

// SetDiskName adds the diskName to the disks create or update params
func (o *DisksCreateOrUpdateParams) SetDiskName(diskName string) {
	o.DiskName = diskName
}

// WithResourceGroupName adds the resourceGroupName to the disks create or update params
func (o *DisksCreateOrUpdateParams) WithResourceGroupName(resourceGroupName string) *DisksCreateOrUpdateParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the disks create or update params
func (o *DisksCreateOrUpdateParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the disks create or update params
func (o *DisksCreateOrUpdateParams) WithSubscriptionID(subscriptionID string) *DisksCreateOrUpdateParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the disks create or update params
func (o *DisksCreateOrUpdateParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *DisksCreateOrUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Disk == nil {
		o.Disk = new(models.Disk)
	}

	if err := r.SetBodyParam(o.Disk); err != nil {
		return err
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
