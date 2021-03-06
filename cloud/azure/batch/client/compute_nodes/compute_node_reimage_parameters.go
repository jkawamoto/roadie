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

	"github.com/jkawamoto/roadie/cloud/azure/batch/models"
)

// NewComputeNodeReimageParams creates a new ComputeNodeReimageParams object
// with the default values initialized.
func NewComputeNodeReimageParams() *ComputeNodeReimageParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &ComputeNodeReimageParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewComputeNodeReimageParamsWithTimeout creates a new ComputeNodeReimageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewComputeNodeReimageParamsWithTimeout(timeout time.Duration) *ComputeNodeReimageParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &ComputeNodeReimageParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewComputeNodeReimageParamsWithContext creates a new ComputeNodeReimageParams object
// with the default values initialized, and the ability to set a context for a request
func NewComputeNodeReimageParamsWithContext(ctx context.Context) *ComputeNodeReimageParams {
	var (
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &ComputeNodeReimageParams{
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*ComputeNodeReimageParams contains all the parameters to send to the API endpoint
for the compute node reimage operation typically these are written to a http.Request
*/
type ComputeNodeReimageParams struct {

	/*APIVersion
	  Client API Version.

	*/
	APIVersion string
	/*ClientRequestID
	  The caller-generated request identity, in the form of a GUID with no decoration such as curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.

	*/
	ClientRequestID *string
	/*NodeID
	  The ID of the compute node that you want to restart.

	*/
	NodeID string
	/*NodeReimageParameter
	  The parameters for the request.

	*/
	NodeReimageParameter *models.NodeReimageParameter
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

// WithRequestTimeout adds the timeout to the compute node reimage params
func (o *ComputeNodeReimageParams) WithRequestTimeout(timeout time.Duration) *ComputeNodeReimageParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the compute node reimage params
func (o *ComputeNodeReimageParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the compute node reimage params
func (o *ComputeNodeReimageParams) WithContext(ctx context.Context) *ComputeNodeReimageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the compute node reimage params
func (o *ComputeNodeReimageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the compute node reimage params
func (o *ComputeNodeReimageParams) WithAPIVersion(aPIVersion string) *ComputeNodeReimageParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the compute node reimage params
func (o *ComputeNodeReimageParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the compute node reimage params
func (o *ComputeNodeReimageParams) WithClientRequestID(clientRequestID *string) *ComputeNodeReimageParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the compute node reimage params
func (o *ComputeNodeReimageParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithNodeID adds the nodeID to the compute node reimage params
func (o *ComputeNodeReimageParams) WithNodeID(nodeID string) *ComputeNodeReimageParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the compute node reimage params
func (o *ComputeNodeReimageParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WithNodeReimageParameter adds the nodeReimageParameter to the compute node reimage params
func (o *ComputeNodeReimageParams) WithNodeReimageParameter(nodeReimageParameter *models.NodeReimageParameter) *ComputeNodeReimageParams {
	o.SetNodeReimageParameter(nodeReimageParameter)
	return o
}

// SetNodeReimageParameter adds the nodeReimageParameter to the compute node reimage params
func (o *ComputeNodeReimageParams) SetNodeReimageParameter(nodeReimageParameter *models.NodeReimageParameter) {
	o.NodeReimageParameter = nodeReimageParameter
}

// WithOcpDate adds the ocpDate to the compute node reimage params
func (o *ComputeNodeReimageParams) WithOcpDate(ocpDate *string) *ComputeNodeReimageParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the compute node reimage params
func (o *ComputeNodeReimageParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithPoolID adds the poolID to the compute node reimage params
func (o *ComputeNodeReimageParams) WithPoolID(poolID string) *ComputeNodeReimageParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the compute node reimage params
func (o *ComputeNodeReimageParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WithReturnClientRequestID adds the returnClientRequestID to the compute node reimage params
func (o *ComputeNodeReimageParams) WithReturnClientRequestID(returnClientRequestID *bool) *ComputeNodeReimageParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the compute node reimage params
func (o *ComputeNodeReimageParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTimeout adds the timeout to the compute node reimage params
func (o *ComputeNodeReimageParams) WithTimeout(timeout *int32) *ComputeNodeReimageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the compute node reimage params
func (o *ComputeNodeReimageParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *ComputeNodeReimageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.NodeReimageParameter == nil {
		o.NodeReimageParameter = new(models.NodeReimageParameter)
	}

	if err := r.SetBodyParam(o.NodeReimageParameter); err != nil {
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
