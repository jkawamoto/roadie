package jobs

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

// NewJobPatchParams creates a new JobPatchParams object
// with the default values initialized.
func NewJobPatchParams() *JobPatchParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobPatchParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewJobPatchParamsWithTimeout creates a new JobPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobPatchParamsWithTimeout(timeout time.Duration) *JobPatchParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobPatchParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewJobPatchParamsWithContext creates a new JobPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobPatchParamsWithContext(ctx context.Context) *JobPatchParams {
	var (
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobPatchParams{
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*JobPatchParams contains all the parameters to send to the API endpoint
for the job patch operation typically these are written to a http.Request
*/
type JobPatchParams struct {

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
	/*JobID
	  The ID of the job whose properties you want to update.

	*/
	JobID string
	/*JobPatchParameter
	  The parameters for the request.

	*/
	JobPatchParameter *models.JobPatchParameter
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

// WithRequestTimeout adds the timeout to the job patch params
func (o *JobPatchParams) WithRequestTimeout(timeout time.Duration) *JobPatchParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the job patch params
func (o *JobPatchParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the job patch params
func (o *JobPatchParams) WithContext(ctx context.Context) *JobPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job patch params
func (o *JobPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIfMatch adds the ifMatch to the job patch params
func (o *JobPatchParams) WithIfMatch(ifMatch *string) *JobPatchParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the job patch params
func (o *JobPatchParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithIfModifiedSince adds the ifModifiedSince to the job patch params
func (o *JobPatchParams) WithIfModifiedSince(ifModifiedSince *string) *JobPatchParams {
	o.SetIfModifiedSince(ifModifiedSince)
	return o
}

// SetIfModifiedSince adds the ifModifiedSince to the job patch params
func (o *JobPatchParams) SetIfModifiedSince(ifModifiedSince *string) {
	o.IfModifiedSince = ifModifiedSince
}

// WithIfNoneMatch adds the ifNoneMatch to the job patch params
func (o *JobPatchParams) WithIfNoneMatch(ifNoneMatch *string) *JobPatchParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the job patch params
func (o *JobPatchParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithIfUnmodifiedSince adds the ifUnmodifiedSince to the job patch params
func (o *JobPatchParams) WithIfUnmodifiedSince(ifUnmodifiedSince *string) *JobPatchParams {
	o.SetIfUnmodifiedSince(ifUnmodifiedSince)
	return o
}

// SetIfUnmodifiedSince adds the ifUnmodifiedSince to the job patch params
func (o *JobPatchParams) SetIfUnmodifiedSince(ifUnmodifiedSince *string) {
	o.IfUnmodifiedSince = ifUnmodifiedSince
}

// WithAPIVersion adds the aPIVersion to the job patch params
func (o *JobPatchParams) WithAPIVersion(aPIVersion string) *JobPatchParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the job patch params
func (o *JobPatchParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the job patch params
func (o *JobPatchParams) WithClientRequestID(clientRequestID *string) *JobPatchParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the job patch params
func (o *JobPatchParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithJobID adds the jobID to the job patch params
func (o *JobPatchParams) WithJobID(jobID string) *JobPatchParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the job patch params
func (o *JobPatchParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WithJobPatchParameter adds the jobPatchParameter to the job patch params
func (o *JobPatchParams) WithJobPatchParameter(jobPatchParameter *models.JobPatchParameter) *JobPatchParams {
	o.SetJobPatchParameter(jobPatchParameter)
	return o
}

// SetJobPatchParameter adds the jobPatchParameter to the job patch params
func (o *JobPatchParams) SetJobPatchParameter(jobPatchParameter *models.JobPatchParameter) {
	o.JobPatchParameter = jobPatchParameter
}

// WithOcpDate adds the ocpDate to the job patch params
func (o *JobPatchParams) WithOcpDate(ocpDate *string) *JobPatchParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the job patch params
func (o *JobPatchParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithReturnClientRequestID adds the returnClientRequestID to the job patch params
func (o *JobPatchParams) WithReturnClientRequestID(returnClientRequestID *bool) *JobPatchParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the job patch params
func (o *JobPatchParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTimeout adds the timeout to the job patch params
func (o *JobPatchParams) WithTimeout(timeout *int32) *JobPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job patch params
func (o *JobPatchParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *JobPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param jobId
	if err := r.SetPathParam("jobId", o.JobID); err != nil {
		return err
	}

	if o.JobPatchParameter == nil {
		o.JobPatchParameter = new(models.JobPatchParameter)
	}

	if err := r.SetBodyParam(o.JobPatchParameter); err != nil {
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