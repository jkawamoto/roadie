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

// NewComputeNodeAddUserParams creates a new ComputeNodeAddUserParams object
// with the default values initialized.
func NewComputeNodeAddUserParams() *ComputeNodeAddUserParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &ComputeNodeAddUserParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewComputeNodeAddUserParamsWithTimeout creates a new ComputeNodeAddUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewComputeNodeAddUserParamsWithTimeout(timeout time.Duration) *ComputeNodeAddUserParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &ComputeNodeAddUserParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewComputeNodeAddUserParamsWithContext creates a new ComputeNodeAddUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewComputeNodeAddUserParamsWithContext(ctx context.Context) *ComputeNodeAddUserParams {
	var (
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &ComputeNodeAddUserParams{
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*ComputeNodeAddUserParams contains all the parameters to send to the API endpoint
for the compute node add user operation typically these are written to a http.Request
*/
type ComputeNodeAddUserParams struct {

	/*User
	  The user account to be created.

	*/
	User *models.ComputeNodeUser
	/*APIVersion
	  Client API Version.

	*/
	APIVersion string
	/*ClientRequestID
	  The caller-generated request identity, in the form of a GUID with no decoration such as curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.

	*/
	ClientRequestID *string
	/*NodeID
	  The ID of the machine on which you want to create a user account.

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

// WithRequestTimeout adds the timeout to the compute node add user params
func (o *ComputeNodeAddUserParams) WithRequestTimeout(timeout time.Duration) *ComputeNodeAddUserParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the compute node add user params
func (o *ComputeNodeAddUserParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the compute node add user params
func (o *ComputeNodeAddUserParams) WithContext(ctx context.Context) *ComputeNodeAddUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the compute node add user params
func (o *ComputeNodeAddUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithUser adds the user to the compute node add user params
func (o *ComputeNodeAddUserParams) WithUser(user *models.ComputeNodeUser) *ComputeNodeAddUserParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the compute node add user params
func (o *ComputeNodeAddUserParams) SetUser(user *models.ComputeNodeUser) {
	o.User = user
}

// WithAPIVersion adds the aPIVersion to the compute node add user params
func (o *ComputeNodeAddUserParams) WithAPIVersion(aPIVersion string) *ComputeNodeAddUserParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the compute node add user params
func (o *ComputeNodeAddUserParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the compute node add user params
func (o *ComputeNodeAddUserParams) WithClientRequestID(clientRequestID *string) *ComputeNodeAddUserParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the compute node add user params
func (o *ComputeNodeAddUserParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithNodeID adds the nodeID to the compute node add user params
func (o *ComputeNodeAddUserParams) WithNodeID(nodeID string) *ComputeNodeAddUserParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the compute node add user params
func (o *ComputeNodeAddUserParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WithOcpDate adds the ocpDate to the compute node add user params
func (o *ComputeNodeAddUserParams) WithOcpDate(ocpDate *string) *ComputeNodeAddUserParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the compute node add user params
func (o *ComputeNodeAddUserParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithPoolID adds the poolID to the compute node add user params
func (o *ComputeNodeAddUserParams) WithPoolID(poolID string) *ComputeNodeAddUserParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the compute node add user params
func (o *ComputeNodeAddUserParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WithReturnClientRequestID adds the returnClientRequestID to the compute node add user params
func (o *ComputeNodeAddUserParams) WithReturnClientRequestID(returnClientRequestID *bool) *ComputeNodeAddUserParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the compute node add user params
func (o *ComputeNodeAddUserParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTimeout adds the timeout to the compute node add user params
func (o *ComputeNodeAddUserParams) WithTimeout(timeout *int32) *ComputeNodeAddUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the compute node add user params
func (o *ComputeNodeAddUserParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *ComputeNodeAddUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.requestTimeout)
	var res []error

	if o.User == nil {
		o.User = new(models.ComputeNodeUser)
	}

	if err := r.SetBodyParam(o.User); err != nil {
		return err
	}

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