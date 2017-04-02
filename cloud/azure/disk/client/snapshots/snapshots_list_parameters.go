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
)

// NewSnapshotsListParams creates a new SnapshotsListParams object
// with the default values initialized.
func NewSnapshotsListParams() *SnapshotsListParams {
	var ()
	return &SnapshotsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotsListParamsWithTimeout creates a new SnapshotsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSnapshotsListParamsWithTimeout(timeout time.Duration) *SnapshotsListParams {
	var ()
	return &SnapshotsListParams{

		timeout: timeout,
	}
}

// NewSnapshotsListParamsWithContext creates a new SnapshotsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSnapshotsListParamsWithContext(ctx context.Context) *SnapshotsListParams {
	var ()
	return &SnapshotsListParams{

		Context: ctx,
	}
}

/*SnapshotsListParams contains all the parameters to send to the API endpoint
for the snapshots list operation typically these are written to a http.Request
*/
type SnapshotsListParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*SubscriptionID
	  Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the snapshots list params
func (o *SnapshotsListParams) WithTimeout(timeout time.Duration) *SnapshotsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshots list params
func (o *SnapshotsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshots list params
func (o *SnapshotsListParams) WithContext(ctx context.Context) *SnapshotsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshots list params
func (o *SnapshotsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the snapshots list params
func (o *SnapshotsListParams) WithAPIVersion(aPIVersion string) *SnapshotsListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the snapshots list params
func (o *SnapshotsListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithSubscriptionID adds the subscriptionID to the snapshots list params
func (o *SnapshotsListParams) WithSubscriptionID(subscriptionID string) *SnapshotsListParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the snapshots list params
func (o *SnapshotsListParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
