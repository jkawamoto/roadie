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

// NewPoolListPoolUsageMetricsParams creates a new PoolListPoolUsageMetricsParams object
// with the default values initialized.
func NewPoolListPoolUsageMetricsParams() *PoolListPoolUsageMetricsParams {
	var (
		maxresultsDefault            = int32(1000)
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolListPoolUsageMetricsParams{
		Maxresults:            &maxresultsDefault,
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewPoolListPoolUsageMetricsParamsWithTimeout creates a new PoolListPoolUsageMetricsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPoolListPoolUsageMetricsParamsWithTimeout(timeout time.Duration) *PoolListPoolUsageMetricsParams {
	var (
		maxresultsDefault            = int32(1000)
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolListPoolUsageMetricsParams{
		Maxresults:            &maxresultsDefault,
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewPoolListPoolUsageMetricsParamsWithContext creates a new PoolListPoolUsageMetricsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPoolListPoolUsageMetricsParamsWithContext(ctx context.Context) *PoolListPoolUsageMetricsParams {
	var (
		maxresultsDefault            = int32(1000)
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &PoolListPoolUsageMetricsParams{
		Maxresults:            &maxresultsDefault,
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*PoolListPoolUsageMetricsParams contains all the parameters to send to the API endpoint
for the pool list pool usage metrics operation typically these are written to a http.Request
*/
type PoolListPoolUsageMetricsParams struct {

	/*NrDollarFilter
	  An OData $filter clause. If this is not specified the response includes all pools that existed in the account in the time range of the returned aggregation intervals.

	*/
	DollarFilter *string
	/*APIVersion
	  Client API Version.

	*/
	APIVersion string
	/*ClientRequestID
	  The caller-generated request identity, in the form of a GUID with no decoration such as curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.

	*/
	ClientRequestID *string
	/*Endtime
	  The latest time from which to include metrics. This must be at least two hours before the current time. If not specified this defaults to the end time of the last aggregation interval currently available.

	*/
	Endtime *strfmt.DateTime
	/*Maxresults
	  The maximum number of items to return in the response. A maximum of 1000 results will be returned.

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
	/*Starttime
	  The earliest time from which to include metrics. This must be at least two and a half hours before the current time. If not specified this defaults to the start time of the last aggregation interval currently available.

	*/
	Starttime *strfmt.DateTime
	/*Timeout
	  The maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.

	*/
	Timeout *int32

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) WithRequestTimeout(timeout time.Duration) *PoolListPoolUsageMetricsParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) WithContext(ctx context.Context) *PoolListPoolUsageMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDollarFilter adds the dollarFilter to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) WithDollarFilter(dollarFilter *string) *PoolListPoolUsageMetricsParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithAPIVersion adds the aPIVersion to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) WithAPIVersion(aPIVersion string) *PoolListPoolUsageMetricsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) WithClientRequestID(clientRequestID *string) *PoolListPoolUsageMetricsParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithEndtime adds the endtime to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) WithEndtime(endtime *strfmt.DateTime) *PoolListPoolUsageMetricsParams {
	o.SetEndtime(endtime)
	return o
}

// SetEndtime adds the endtime to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) SetEndtime(endtime *strfmt.DateTime) {
	o.Endtime = endtime
}

// WithMaxresults adds the maxresults to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) WithMaxresults(maxresults *int32) *PoolListPoolUsageMetricsParams {
	o.SetMaxresults(maxresults)
	return o
}

// SetMaxresults adds the maxresults to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) SetMaxresults(maxresults *int32) {
	o.Maxresults = maxresults
}

// WithOcpDate adds the ocpDate to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) WithOcpDate(ocpDate *string) *PoolListPoolUsageMetricsParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithReturnClientRequestID adds the returnClientRequestID to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) WithReturnClientRequestID(returnClientRequestID *bool) *PoolListPoolUsageMetricsParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithStarttime adds the starttime to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) WithStarttime(starttime *strfmt.DateTime) *PoolListPoolUsageMetricsParams {
	o.SetStarttime(starttime)
	return o
}

// SetStarttime adds the starttime to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) SetStarttime(starttime *strfmt.DateTime) {
	o.Starttime = starttime
}

// WithTimeout adds the timeout to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) WithTimeout(timeout *int32) *PoolListPoolUsageMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pool list pool usage metrics params
func (o *PoolListPoolUsageMetricsParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *PoolListPoolUsageMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.requestTimeout)
	var res []error

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

	if o.Endtime != nil {

		// query param endtime
		var qrEndtime strfmt.DateTime
		if o.Endtime != nil {
			qrEndtime = *o.Endtime
		}
		qEndtime := qrEndtime.String()
		if qEndtime != "" {
			if err := r.SetQueryParam("endtime", qEndtime); err != nil {
				return err
			}
		}

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

	if o.Starttime != nil {

		// query param starttime
		var qrStarttime strfmt.DateTime
		if o.Starttime != nil {
			qrStarttime = *o.Starttime
		}
		qStarttime := qrStarttime.String()
		if qStarttime != "" {
			if err := r.SetQueryParam("starttime", qStarttime); err != nil {
				return err
			}
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
