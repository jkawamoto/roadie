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
)

// NewPoolStopResizeParams creates a new PoolStopResizeParams object
// with the default values initialized.
func NewPoolStopResizeParams() *PoolStopResizeParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolStopResizeParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewPoolStopResizeParamsWithTimeout creates a new PoolStopResizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPoolStopResizeParamsWithTimeout(timeout time.Duration) *PoolStopResizeParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolStopResizeParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewPoolStopResizeParamsWithContext creates a new PoolStopResizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPoolStopResizeParamsWithContext(ctx context.Context) *PoolStopResizeParams {
	var (
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolStopResizeParams{
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*PoolStopResizeParams contains all the parameters to send to the API endpoint
for the pool stop resize operation typically these are written to a http.Request
*/
type PoolStopResizeParams struct {

	/*IfMatch
	  An ETag is specified. Specify this header to perform the operation only if the resource's ETag is an exact match as specified.

	*/
	IfMatch *string
	/*IfModifiedSince
	  Specify this header to perform the operation only if the resource has been modified since the specified date/time.

	*/
	IfModifiedSince *string
	/*IfNoneMatch
	  An ETag is specified. Specify this header to perform the operation only if the resource's ETag does not match the specified ETag.

	*/
	IfNoneMatch *string
	/*IfUnmodifiedSince
	  Specify this header to perform the operation only if the resource has not been modified since the specified date/time.

	*/
	IfUnmodifiedSince *string
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
	/*PoolID
	  The ID of the pool whose resizing you want to stop.

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

// WithRequestTimeout adds the timeout to the pool stop resize params
func (o *PoolStopResizeParams) WithRequestTimeout(timeout time.Duration) *PoolStopResizeParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the pool stop resize params
func (o *PoolStopResizeParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the pool stop resize params
func (o *PoolStopResizeParams) WithContext(ctx context.Context) *PoolStopResizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pool stop resize params
func (o *PoolStopResizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIfMatch adds the ifMatch to the pool stop resize params
func (o *PoolStopResizeParams) WithIfMatch(ifMatch *string) *PoolStopResizeParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the pool stop resize params
func (o *PoolStopResizeParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithIfModifiedSince adds the ifModifiedSince to the pool stop resize params
func (o *PoolStopResizeParams) WithIfModifiedSince(ifModifiedSince *string) *PoolStopResizeParams {
	o.SetIfModifiedSince(ifModifiedSince)
	return o
}

// SetIfModifiedSince adds the ifModifiedSince to the pool stop resize params
func (o *PoolStopResizeParams) SetIfModifiedSince(ifModifiedSince *string) {
	o.IfModifiedSince = ifModifiedSince
}

// WithIfNoneMatch adds the ifNoneMatch to the pool stop resize params
func (o *PoolStopResizeParams) WithIfNoneMatch(ifNoneMatch *string) *PoolStopResizeParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the pool stop resize params
func (o *PoolStopResizeParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithIfUnmodifiedSince adds the ifUnmodifiedSince to the pool stop resize params
func (o *PoolStopResizeParams) WithIfUnmodifiedSince(ifUnmodifiedSince *string) *PoolStopResizeParams {
	o.SetIfUnmodifiedSince(ifUnmodifiedSince)
	return o
}

// SetIfUnmodifiedSince adds the ifUnmodifiedSince to the pool stop resize params
func (o *PoolStopResizeParams) SetIfUnmodifiedSince(ifUnmodifiedSince *string) {
	o.IfUnmodifiedSince = ifUnmodifiedSince
}

// WithAPIVersion adds the aPIVersion to the pool stop resize params
func (o *PoolStopResizeParams) WithAPIVersion(aPIVersion string) *PoolStopResizeParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the pool stop resize params
func (o *PoolStopResizeParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the pool stop resize params
func (o *PoolStopResizeParams) WithClientRequestID(clientRequestID *string) *PoolStopResizeParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the pool stop resize params
func (o *PoolStopResizeParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithOcpDate adds the ocpDate to the pool stop resize params
func (o *PoolStopResizeParams) WithOcpDate(ocpDate *string) *PoolStopResizeParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the pool stop resize params
func (o *PoolStopResizeParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithPoolID adds the poolID to the pool stop resize params
func (o *PoolStopResizeParams) WithPoolID(poolID string) *PoolStopResizeParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the pool stop resize params
func (o *PoolStopResizeParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WithReturnClientRequestID adds the returnClientRequestID to the pool stop resize params
func (o *PoolStopResizeParams) WithReturnClientRequestID(returnClientRequestID *bool) *PoolStopResizeParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the pool stop resize params
func (o *PoolStopResizeParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTimeout adds the timeout to the pool stop resize params
func (o *PoolStopResizeParams) WithTimeout(timeout *int32) *PoolStopResizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pool stop resize params
func (o *PoolStopResizeParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *PoolStopResizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.requestTimeout)
	var res []error

	if o.IfMatch != nil {

		// header param If-Match
		if err := r.SetHeaderParam("If-Match", *o.IfMatch); err != nil {
			return err
		}

	}

	if o.IfModifiedSince != nil {

		// header param If-Modified-Since
		if err := r.SetHeaderParam("If-Modified-Since", *o.IfModifiedSince); err != nil {
			return err
		}

	}

	if o.IfNoneMatch != nil {

		// header param If-None-Match
		if err := r.SetHeaderParam("If-None-Match", *o.IfNoneMatch); err != nil {
			return err
		}

	}

	if o.IfUnmodifiedSince != nil {

		// header param If-Unmodified-Since
		if err := r.SetHeaderParam("If-Unmodified-Since", *o.IfUnmodifiedSince); err != nil {
			return err
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
