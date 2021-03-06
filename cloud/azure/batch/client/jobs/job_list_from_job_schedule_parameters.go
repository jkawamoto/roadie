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

// NewJobListFromJobScheduleParams creates a new JobListFromJobScheduleParams object
// with the default values initialized.
func NewJobListFromJobScheduleParams() *JobListFromJobScheduleParams {
	var (
		maxresultsDefault            = int32(1000)
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobListFromJobScheduleParams{
		Maxresults:            &maxresultsDefault,
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewJobListFromJobScheduleParamsWithTimeout creates a new JobListFromJobScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobListFromJobScheduleParamsWithTimeout(timeout time.Duration) *JobListFromJobScheduleParams {
	var (
		maxresultsDefault            = int32(1000)
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobListFromJobScheduleParams{
		Maxresults:            &maxresultsDefault,
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewJobListFromJobScheduleParamsWithContext creates a new JobListFromJobScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobListFromJobScheduleParamsWithContext(ctx context.Context) *JobListFromJobScheduleParams {
	var (
		maxresultsDefault            = int32(1000)
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &JobListFromJobScheduleParams{
		Maxresults:            &maxresultsDefault,
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*JobListFromJobScheduleParams contains all the parameters to send to the API endpoint
for the job list from job schedule operation typically these are written to a http.Request
*/
type JobListFromJobScheduleParams struct {

	/*NrDollarExpand
	  An OData $expand clause.

	*/
	DollarExpand *string
	/*NrDollarFilter
	  An OData $filter clause.

	*/
	DollarFilter *string
	/*NrDollarSelect
	  An OData $select clause.

	*/
	DollarSelect *string
	/*APIVersion
	  Client API Version.

	*/
	APIVersion string
	/*ClientRequestID
	  The caller-generated request identity, in the form of a GUID with no decoration such as curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.

	*/
	ClientRequestID *string
	/*JobScheduleID
	  The ID of the job schedule from which you want to get a list of jobs.

	*/
	JobScheduleID string
	/*Maxresults
	  The maximum number of items to return in the response. A maximum of 1000 jobs can be returned.

	*/
	Maxresults *int32
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

// WithRequestTimeout adds the timeout to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithRequestTimeout(timeout time.Duration) *JobListFromJobScheduleParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithContext(ctx context.Context) *JobListFromJobScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDollarExpand adds the dollarExpand to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithDollarExpand(dollarExpand *string) *JobListFromJobScheduleParams {
	o.SetDollarExpand(dollarExpand)
	return o
}

// SetDollarExpand adds the dollarExpand to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetDollarExpand(dollarExpand *string) {
	o.DollarExpand = dollarExpand
}

// WithDollarFilter adds the dollarFilter to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithDollarFilter(dollarFilter *string) *JobListFromJobScheduleParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarSelect adds the dollarSelect to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithDollarSelect(dollarSelect *string) *JobListFromJobScheduleParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithAPIVersion adds the aPIVersion to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithAPIVersion(aPIVersion string) *JobListFromJobScheduleParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithClientRequestID(clientRequestID *string) *JobListFromJobScheduleParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithJobScheduleID adds the jobScheduleID to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithJobScheduleID(jobScheduleID string) *JobListFromJobScheduleParams {
	o.SetJobScheduleID(jobScheduleID)
	return o
}

// SetJobScheduleID adds the jobScheduleId to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetJobScheduleID(jobScheduleID string) {
	o.JobScheduleID = jobScheduleID
}

// WithMaxresults adds the maxresults to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithMaxresults(maxresults *int32) *JobListFromJobScheduleParams {
	o.SetMaxresults(maxresults)
	return o
}

// SetMaxresults adds the maxresults to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetMaxresults(maxresults *int32) {
	o.Maxresults = maxresults
}

// WithOcpDate adds the ocpDate to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithOcpDate(ocpDate *string) *JobListFromJobScheduleParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithReturnClientRequestID adds the returnClientRequestID to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithReturnClientRequestID(returnClientRequestID *bool) *JobListFromJobScheduleParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTimeout adds the timeout to the job list from job schedule params
func (o *JobListFromJobScheduleParams) WithTimeout(timeout *int32) *JobListFromJobScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job list from job schedule params
func (o *JobListFromJobScheduleParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *JobListFromJobScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DollarFilter != nil {

		// query param $filter
		var qrNrDollarFilter string
		if o.DollarFilter != nil {
			qrNrDollarFilter = *o.DollarFilter
		}
		qNrDollarFilter := qrNrDollarFilter
		if qNrDollarFilter != "" {
			if err := r.SetQueryParam("$filter", qNrDollarFilter); err != nil {
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

	if o.Maxresults != nil {

		// query param maxresults
		var qrMaxresults int32
		if o.Maxresults != nil {
			qrMaxresults = *o.Maxresults
		}
		qMaxresults := swag.FormatInt32(qrMaxresults)
		if qMaxresults != "" {
			if err := r.SetQueryParam("maxresults", qMaxresults); err != nil {
				return err
			}
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
