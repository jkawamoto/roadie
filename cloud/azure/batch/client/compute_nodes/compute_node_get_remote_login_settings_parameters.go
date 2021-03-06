package compute_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewComputeNodeGetRemoteLoginSettingsParams creates a new ComputeNodeGetRemoteLoginSettingsParams object
// with the default values initialized.
func NewComputeNodeGetRemoteLoginSettingsParams() *ComputeNodeGetRemoteLoginSettingsParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &ComputeNodeGetRemoteLoginSettingsParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewComputeNodeGetRemoteLoginSettingsParamsWithTimeout creates a new ComputeNodeGetRemoteLoginSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewComputeNodeGetRemoteLoginSettingsParamsWithTimeout(timeout time.Duration) *ComputeNodeGetRemoteLoginSettingsParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &ComputeNodeGetRemoteLoginSettingsParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewComputeNodeGetRemoteLoginSettingsParamsWithContext creates a new ComputeNodeGetRemoteLoginSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewComputeNodeGetRemoteLoginSettingsParamsWithContext(ctx context.Context) *ComputeNodeGetRemoteLoginSettingsParams {
	var (
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &ComputeNodeGetRemoteLoginSettingsParams{
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*ComputeNodeGetRemoteLoginSettingsParams contains all the parameters to send to the API endpoint
for the compute node get remote login settings operation typically these are written to a http.Request
*/
type ComputeNodeGetRemoteLoginSettingsParams struct {

	/*APIVersion
	  Client API Version.

	*/
	APIVersion string
	/*ClientRequestID
	  The caller-generated request identity, in the form of a GUID with no decoration such as curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.

	*/
	ClientRequestID *string
	/*NodeID
	  The ID of the compute node for which to obtain the remote login settings.

	*/
	NodeID string
	/*OcpDate
	  The time the request was issued. If not specified, this header will be automatically populated with the current system clock time.

	*/
	OcpDate *string
	/*PoolID
	  The ID of the pool that contains the compute node.

	*/
	PoolID string
	/*ReturnClientRequestID
	  Whether the server should return the client-request-id in the response.

	*/
	ReturnClientRequestID *bool
	/*Timeout
	  The maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.

	*/
	Timeout *int32

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) WithRequestTimeout(timeout time.Duration) *ComputeNodeGetRemoteLoginSettingsParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) WithContext(ctx context.Context) *ComputeNodeGetRemoteLoginSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) WithAPIVersion(aPIVersion string) *ComputeNodeGetRemoteLoginSettingsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) WithClientRequestID(clientRequestID *string) *ComputeNodeGetRemoteLoginSettingsParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithNodeID adds the nodeID to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) WithNodeID(nodeID string) *ComputeNodeGetRemoteLoginSettingsParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WithOcpDate adds the ocpDate to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) WithOcpDate(ocpDate *string) *ComputeNodeGetRemoteLoginSettingsParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithPoolID adds the poolID to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) WithPoolID(poolID string) *ComputeNodeGetRemoteLoginSettingsParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WithReturnClientRequestID adds the returnClientRequestID to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) WithReturnClientRequestID(returnClientRequestID *bool) *ComputeNodeGetRemoteLoginSettingsParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTimeout adds the timeout to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) WithTimeout(timeout *int32) *ComputeNodeGetRemoteLoginSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the compute node get remote login settings params
func (o *ComputeNodeGetRemoteLoginSettingsParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *ComputeNodeGetRemoteLoginSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.requestTimeout)
	var res []error

	// query param api-version
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("api-version", qAPIVersion); err != nil {
			return err
		}
	}

	if o.ClientRequestID != nil {

		// header param client-request-id
		if err := r.SetHeaderParam("client-request-id", *o.ClientRequestID); err != nil {
			return err
		}

	}

	// path param nodeId
	if err := r.SetPathParam("nodeId", o.NodeID); err != nil {
		return err
	}

	if o.OcpDate != nil {

		// header param ocp-date
		if err := r.SetHeaderParam("ocp-date", *o.OcpDate); err != nil {
			return err
		}

	}

	// path param poolId
	if err := r.SetPathParam("poolId", o.PoolID); err != nil {
		return err
	}

	if o.ReturnClientRequestID != nil {

		// header param return-client-request-id
		if err := r.SetHeaderParam("return-client-request-id", swag.FormatBool(*o.ReturnClientRequestID)); err != nil {
			return err
		}

	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int32
		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt32(qrTimeout)
		if qTimeout != "" {
			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
