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
)

// NewJobGetAllJobsLifetimeStatisticsParams creates a new JobGetAllJobsLifetimeStatisticsParams object
// with the default values initialized.
func NewJobGetAllJobsLifetimeStatisticsParams() *JobGetAllJobsLifetimeStatisticsParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobGetAllJobsLifetimeStatisticsParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewJobGetAllJobsLifetimeStatisticsParamsWithTimeout creates a new JobGetAllJobsLifetimeStatisticsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobGetAllJobsLifetimeStatisticsParamsWithTimeout(timeout time.Duration) *JobGetAllJobsLifetimeStatisticsParams {
	var (
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobGetAllJobsLifetimeStatisticsParams{
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewJobGetAllJobsLifetimeStatisticsParamsWithContext creates a new JobGetAllJobsLifetimeStatisticsParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobGetAllJobsLifetimeStatisticsParamsWithContext(ctx context.Context) *JobGetAllJobsLifetimeStatisticsParams {
	var (
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobGetAllJobsLifetimeStatisticsParams{
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*JobGetAllJobsLifetimeStatisticsParams contains all the parameters to send to the API endpoint
for the job get all jobs lifetime statistics operation typically these are written to a http.Request
*/
type JobGetAllJobsLifetimeStatisticsParams struct {

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

// WithRequestTimeout adds the timeout to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) WithRequestTimeout(timeout time.Duration) *JobGetAllJobsLifetimeStatisticsParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) WithContext(ctx context.Context) *JobGetAllJobsLifetimeStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) WithAPIVersion(aPIVersion string) *JobGetAllJobsLifetimeStatisticsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) WithClientRequestID(clientRequestID *string) *JobGetAllJobsLifetimeStatisticsParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithOcpDate adds the ocpDate to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) WithOcpDate(ocpDate *string) *JobGetAllJobsLifetimeStatisticsParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithReturnClientRequestID adds the returnClientRequestID to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) WithReturnClientRequestID(returnClientRequestID *bool) *JobGetAllJobsLifetimeStatisticsParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTimeout adds the timeout to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) WithTimeout(timeout *int32) *JobGetAllJobsLifetimeStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job get all jobs lifetime statistics params
func (o *JobGetAllJobsLifetimeStatisticsParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *JobGetAllJobsLifetimeStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
