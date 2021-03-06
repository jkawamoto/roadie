package pools

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

// NewPoolAddParams creates a new PoolAddParams object
// with the default values initialized.
func NewPoolAddParams() *PoolAddParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolAddParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewPoolAddParamsWithTimeout creates a new PoolAddParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPoolAddParamsWithTimeout(timeout time.Duration) *PoolAddParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolAddParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewPoolAddParamsWithContext creates a new PoolAddParams object
// with the default values initialized, and the ability to set a context for a request
func NewPoolAddParamsWithContext(ctx context.Context) *PoolAddParams {
	var (
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolAddParams{
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*PoolAddParams contains all the parameters to send to the API endpoint
for the pool add operation typically these are written to a http.Request
*/
type PoolAddParams struct {

	/*Pool
	  The pool to be added.

	*/
	Pool *models.PoolAddParameter
	/*APIVersion
	  Client API Version.

	*/
	APIVersion string
	/*ClientRequestID
	  The caller-generated request identity, in the form of a GUID with no decoration such as curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.

	*/
	ClientRequestID *string
	/*OcpDate
	  The time the request was issued. If not specified, this header will be automatically populated with the current system clock time.

	*/
	OcpDate *string
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

// WithRequestTimeout adds the timeout to the pool add params
func (o *PoolAddParams) WithRequestTimeout(timeout time.Duration) *PoolAddParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the pool add params
func (o *PoolAddParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the pool add params
func (o *PoolAddParams) WithContext(ctx context.Context) *PoolAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pool add params
func (o *PoolAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithPool adds the pool to the pool add params
func (o *PoolAddParams) WithPool(pool *models.PoolAddParameter) *PoolAddParams {
	o.SetPool(pool)
	return o
}

// SetPool adds the pool to the pool add params
func (o *PoolAddParams) SetPool(pool *models.PoolAddParameter) {
	o.Pool = pool
}

// WithAPIVersion adds the aPIVersion to the pool add params
func (o *PoolAddParams) WithAPIVersion(aPIVersion string) *PoolAddParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the pool add params
func (o *PoolAddParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the pool add params
func (o *PoolAddParams) WithClientRequestID(clientRequestID *string) *PoolAddParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the pool add params
func (o *PoolAddParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithOcpDate adds the ocpDate to the pool add params
func (o *PoolAddParams) WithOcpDate(ocpDate *string) *PoolAddParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the pool add params
func (o *PoolAddParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithReturnClientRequestID adds the returnClientRequestID to the pool add params
func (o *PoolAddParams) WithReturnClientRequestID(returnClientRequestID *bool) *PoolAddParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the pool add params
func (o *PoolAddParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTimeout adds the timeout to the pool add params
func (o *PoolAddParams) WithTimeout(timeout *int32) *PoolAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pool add params
func (o *PoolAddParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *PoolAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.requestTimeout)
	var res []error

	if o.Pool == nil {
		o.Pool = new(models.PoolAddParameter)
	}

	if err := r.SetBodyParam(o.Pool); err != nil {
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

	if o.OcpDate != nil {

		// header param ocp-date
		if err := r.SetHeaderParam("ocp-date", *o.OcpDate); err != nil {
			return err
		}

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
