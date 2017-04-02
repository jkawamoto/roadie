package snapshots

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

// NewSnapshotsCreateOrUpdateParams creates a new SnapshotsCreateOrUpdateParams object
// with the default values initialized.
func NewSnapshotsCreateOrUpdateParams() *SnapshotsCreateOrUpdateParams {
	var ()
	return &SnapshotsCreateOrUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotsCreateOrUpdateParamsWithTimeout creates a new SnapshotsCreateOrUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSnapshotsCreateOrUpdateParamsWithTimeout(timeout time.Duration) *SnapshotsCreateOrUpdateParams {
	var ()
	return &SnapshotsCreateOrUpdateParams{

		timeout: timeout,
	}
}

// NewSnapshotsCreateOrUpdateParamsWithContext creates a new SnapshotsCreateOrUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSnapshotsCreateOrUpdateParamsWithContext(ctx context.Context) *SnapshotsCreateOrUpdateParams {
	var ()
	return &SnapshotsCreateOrUpdateParams{

		Context: ctx,
	}
}

/*SnapshotsCreateOrUpdateParams contains all the parameters to send to the API endpoint
for the snapshots create or update operation typically these are written to a http.Request
*/
type SnapshotsCreateOrUpdateParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*Snapshot
	  Snapshot object supplied in the body of the Put disk operation.

	*/
	Snapshot *models.Snapshot
	/*SnapshotName
	  The name of the snapshot within the given subscription and resource group.

	*/
	SnapshotName string
	/*SubscriptionID
	  Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) WithTimeout(timeout time.Duration) *SnapshotsCreateOrUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) WithContext(ctx context.Context) *SnapshotsCreateOrUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) WithAPIVersion(aPIVersion string) *SnapshotsCreateOrUpdateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithResourceGroupName adds the resourceGroupName to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) WithResourceGroupName(resourceGroupName string) *SnapshotsCreateOrUpdateParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSnapshot adds the snapshot to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) WithSnapshot(snapshot *models.Snapshot) *SnapshotsCreateOrUpdateParams {
	o.SetSnapshot(snapshot)
	return o
}

// SetSnapshot adds the snapshot to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) SetSnapshot(snapshot *models.Snapshot) {
	o.Snapshot = snapshot
}

// WithSnapshotName adds the snapshotName to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) WithSnapshotName(snapshotName string) *SnapshotsCreateOrUpdateParams {
	o.SetSnapshotName(snapshotName)
	return o
}

// SetSnapshotName adds the snapshotName to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) SetSnapshotName(snapshotName string) {
	o.SnapshotName = snapshotName
}

// WithSubscriptionID adds the subscriptionID to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) WithSubscriptionID(subscriptionID string) *SnapshotsCreateOrUpdateParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the snapshots create or update params
func (o *SnapshotsCreateOrUpdateParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotsCreateOrUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Snapshot == nil {
		o.Snapshot = new(models.Snapshot)
	}

	if err := r.SetBodyParam(o.Snapshot); err != nil {
		return err
	}

	// path param snapshotName
	if err := r.SetPathParam("snapshotName", o.SnapshotName); err != nil {
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
