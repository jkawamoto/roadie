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

// NewPoolUpgradeOSParams creates a new PoolUpgradeOSParams object
// with the default values initialized.
func NewPoolUpgradeOSParams() *PoolUpgradeOSParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolUpgradeOSParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewPoolUpgradeOSParamsWithTimeout creates a new PoolUpgradeOSParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPoolUpgradeOSParamsWithTimeout(timeout time.Duration) *PoolUpgradeOSParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolUpgradeOSParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewPoolUpgradeOSParamsWithContext creates a new PoolUpgradeOSParams object
// with the default values initialized, and the ability to set a context for a request
func NewPoolUpgradeOSParamsWithContext(ctx context.Context) *PoolUpgradeOSParams {
	var (
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolUpgradeOSParams{
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*PoolUpgradeOSParams contains all the parameters to send to the API endpoint
for the pool upgrade o s operation typically these are written to a http.Request
*/
type PoolUpgradeOSParams struct {

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
	  The ID of the pool to upgrade.

	*/
	PoolID string
	/*PoolUpgradeOSParameter
	  The parameters for the request.

	*/
	PoolUpgradeOSParameter *models.PoolUpgradeOSParameter
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

// WithRequestTimeout adds the timeout to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithRequestTimeout(timeout time.Duration) *PoolUpgradeOSParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithContext(ctx context.Context) *PoolUpgradeOSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIfMatch adds the ifMatch to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithIfMatch(ifMatch *string) *PoolUpgradeOSParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithIfModifiedSince adds the ifModifiedSince to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithIfModifiedSince(ifModifiedSince *string) *PoolUpgradeOSParams {
	o.SetIfModifiedSince(ifModifiedSince)
	return o
}

// SetIfModifiedSince adds the ifModifiedSince to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetIfModifiedSince(ifModifiedSince *string) {
	o.IfModifiedSince = ifModifiedSince
}

// WithIfNoneMatch adds the ifNoneMatch to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithIfNoneMatch(ifNoneMatch *string) *PoolUpgradeOSParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithIfUnmodifiedSince adds the ifUnmodifiedSince to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithIfUnmodifiedSince(ifUnmodifiedSince *string) *PoolUpgradeOSParams {
	o.SetIfUnmodifiedSince(ifUnmodifiedSince)
	return o
}

// SetIfUnmodifiedSince adds the ifUnmodifiedSince to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetIfUnmodifiedSince(ifUnmodifiedSince *string) {
	o.IfUnmodifiedSince = ifUnmodifiedSince
}

// WithAPIVersion adds the aPIVersion to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithAPIVersion(aPIVersion string) *PoolUpgradeOSParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithClientRequestID(clientRequestID *string) *PoolUpgradeOSParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithOcpDate adds the ocpDate to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithOcpDate(ocpDate *string) *PoolUpgradeOSParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithPoolID adds the poolID to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithPoolID(poolID string) *PoolUpgradeOSParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WithPoolUpgradeOSParameter adds the poolUpgradeOSParameter to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithPoolUpgradeOSParameter(poolUpgradeOSParameter *models.PoolUpgradeOSParameter) *PoolUpgradeOSParams {
	o.SetPoolUpgradeOSParameter(poolUpgradeOSParameter)
	return o
}

// SetPoolUpgradeOSParameter adds the poolUpgradeOSParameter to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetPoolUpgradeOSParameter(poolUpgradeOSParameter *models.PoolUpgradeOSParameter) {
	o.PoolUpgradeOSParameter = poolUpgradeOSParameter
}

// WithReturnClientRequestID adds the returnClientRequestID to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithReturnClientRequestID(returnClientRequestID *bool) *PoolUpgradeOSParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTimeout adds the timeout to the pool upgrade o s params
func (o *PoolUpgradeOSParams) WithTimeout(timeout *int32) *PoolUpgradeOSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pool upgrade o s params
func (o *PoolUpgradeOSParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *PoolUpgradeOSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PoolUpgradeOSParameter == nil {
		o.PoolUpgradeOSParameter = new(models.PoolUpgradeOSParameter)
	}

	if err := r.SetBodyParam(o.PoolUpgradeOSParameter); err != nil {
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
