package job_schedules

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

// NewJobScheduleGetParams creates a new JobScheduleGetParams object
// with the default values initialized.
func NewJobScheduleGetParams() *JobScheduleGetParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobScheduleGetParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewJobScheduleGetParamsWithTimeout creates a new JobScheduleGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobScheduleGetParamsWithTimeout(timeout time.Duration) *JobScheduleGetParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobScheduleGetParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewJobScheduleGetParamsWithContext creates a new JobScheduleGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobScheduleGetParamsWithContext(ctx context.Context) *JobScheduleGetParams {
	var (
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobScheduleGetParams{
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*JobScheduleGetParams contains all the parameters to send to the API endpoint
for the job schedule get operation typically these are written to a http.Request
*/
type JobScheduleGetParams struct {

	/*NrDollarExpand
	  An OData $expand clause.

	*/
	DollarExpand *string
	/*NrDollarSelect
	  An OData $select clause.

	*/
	DollarSelect *string
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
	/*JobScheduleID
	  The ID of the job schedule to get.

	*/
	JobScheduleID string
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

// WithRequestTimeout adds the timeout to the job schedule get params
func (o *JobScheduleGetParams) WithRequestTimeout(timeout time.Duration) *JobScheduleGetParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the job schedule get params
func (o *JobScheduleGetParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the job schedule get params
func (o *JobScheduleGetParams) WithContext(ctx context.Context) *JobScheduleGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job schedule get params
func (o *JobScheduleGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDollarExpand adds the dollarExpand to the job schedule get params
func (o *JobScheduleGetParams) WithDollarExpand(dollarExpand *string) *JobScheduleGetParams {
	o.SetDollarExpand(dollarExpand)
	return o
}

// SetDollarExpand adds the dollarExpand to the job schedule get params
func (o *JobScheduleGetParams) SetDollarExpand(dollarExpand *string) {
	o.DollarExpand = dollarExpand
}

// WithDollarSelect adds the dollarSelect to the job schedule get params
func (o *JobScheduleGetParams) WithDollarSelect(dollarSelect *string) *JobScheduleGetParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the job schedule get params
func (o *JobScheduleGetParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithIfMatch adds the ifMatch to the job schedule get params
func (o *JobScheduleGetParams) WithIfMatch(ifMatch *string) *JobScheduleGetParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the job schedule get params
func (o *JobScheduleGetParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithIfModifiedSince adds the ifModifiedSince to the job schedule get params
func (o *JobScheduleGetParams) WithIfModifiedSince(ifModifiedSince *string) *JobScheduleGetParams {
	o.SetIfModifiedSince(ifModifiedSince)
	return o
}

// SetIfModifiedSince adds the ifModifiedSince to the job schedule get params
func (o *JobScheduleGetParams) SetIfModifiedSince(ifModifiedSince *string) {
	o.IfModifiedSince = ifModifiedSince
}

// WithIfNoneMatch adds the ifNoneMatch to the job schedule get params
func (o *JobScheduleGetParams) WithIfNoneMatch(ifNoneMatch *string) *JobScheduleGetParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the job schedule get params
func (o *JobScheduleGetParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithIfUnmodifiedSince adds the ifUnmodifiedSince to the job schedule get params
func (o *JobScheduleGetParams) WithIfUnmodifiedSince(ifUnmodifiedSince *string) *JobScheduleGetParams {
	o.SetIfUnmodifiedSince(ifUnmodifiedSince)
	return o
}

// SetIfUnmodifiedSince adds the ifUnmodifiedSince to the job schedule get params
func (o *JobScheduleGetParams) SetIfUnmodifiedSince(ifUnmodifiedSince *string) {
	o.IfUnmodifiedSince = ifUnmodifiedSince
}

// WithAPIVersion adds the aPIVersion to the job schedule get params
func (o *JobScheduleGetParams) WithAPIVersion(aPIVersion string) *JobScheduleGetParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the job schedule get params
func (o *JobScheduleGetParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the job schedule get params
func (o *JobScheduleGetParams) WithClientRequestID(clientRequestID *string) *JobScheduleGetParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the job schedule get params
func (o *JobScheduleGetParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithJobScheduleID adds the jobScheduleID to the job schedule get params
func (o *JobScheduleGetParams) WithJobScheduleID(jobScheduleID string) *JobScheduleGetParams {
	o.SetJobScheduleID(jobScheduleID)
	return o
}

// SetJobScheduleID adds the jobScheduleId to the job schedule get params
func (o *JobScheduleGetParams) SetJobScheduleID(jobScheduleID string) {
	o.JobScheduleID = jobScheduleID
}

// WithOcpDate adds the ocpDate to the job schedule get params
func (o *JobScheduleGetParams) WithOcpDate(ocpDate *string) *JobScheduleGetParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the job schedule get params
func (o *JobScheduleGetParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithReturnClientRequestID adds the returnClientRequestID to the job schedule get params
func (o *JobScheduleGetParams) WithReturnClientRequestID(returnClientRequestID *bool) *JobScheduleGetParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the job schedule get params
func (o *JobScheduleGetParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTimeout adds the timeout to the job schedule get params
func (o *JobScheduleGetParams) WithTimeout(timeout *int32) *JobScheduleGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job schedule get params
func (o *JobScheduleGetParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *JobScheduleGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.requestTimeout)
	var res []error

	if o.DollarExpand != nil {

		// query param $expand
		var qrNrDollarExpand string
		if o.DollarExpand != nil {
			qrNrDollarExpand = *o.DollarExpand
		}
		qNrDollarExpand := qrNrDollarExpand
		if qNrDollarExpand != "" {
			if err := r.SetQueryParam("$expand", qNrDollarExpand); err != nil {
				return err
			}
		}

	}

	if o.DollarSelect != nil {

		// query param $select
		var qrNrDollarSelect string
		if o.DollarSelect != nil {
			qrNrDollarSelect = *o.DollarSelect
		}
		qNrDollarSelect := qrNrDollarSelect
		if qNrDollarSelect != "" {
			if err := r.SetQueryParam("$select", qNrDollarSelect); err != nil {
				return err
			}
		}

	}

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

	// path param jobScheduleId
	if err := r.SetPathParam("jobScheduleId", o.JobScheduleID); err != nil {
		return err
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
