package accounts

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

// NewAccountListNodeAgentSkusParams creates a new AccountListNodeAgentSkusParams object
// with the default values initialized.
func NewAccountListNodeAgentSkusParams() *AccountListNodeAgentSkusParams {
	var (
		maxresultsDefault            = int32(1000)
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &AccountListNodeAgentSkusParams{
		Maxresults:            &maxresultsDefault,
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewAccountListNodeAgentSkusParamsWithTimeout creates a new AccountListNodeAgentSkusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAccountListNodeAgentSkusParamsWithTimeout(timeout time.Duration) *AccountListNodeAgentSkusParams {
	var (
		maxresultsDefault            = int32(1000)
		returnClientRequestIDDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &AccountListNodeAgentSkusParams{
		Maxresults:            &maxresultsDefault,
		ReturnClientRequestID: &returnClientRequestIDDefault,
		Timeout:               &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewAccountListNodeAgentSkusParamsWithContext creates a new AccountListNodeAgentSkusParams object
// with the default values initialized, and the ability to set a context for a request
func NewAccountListNodeAgentSkusParamsWithContext(ctx context.Context) *AccountListNodeAgentSkusParams {
	var (
		maxresultsDefault            = int32(1000)
		returnClientRequestIdDefault = bool(false)
		timeoutDefault               = int32(30)
	)
	return &AccountListNodeAgentSkusParams{
		Maxresults:            &maxresultsDefault,
		ReturnClientRequestID: &returnClientRequestIdDefault,
		Timeout:               &timeoutDefault,

		Context: ctx,
	}
}

/*AccountListNodeAgentSkusParams contains all the parameters to send to the API endpoint
for the account list node agent skus operation typically these are written to a http.Request
*/
type AccountListNodeAgentSkusParams struct {

	/*NrDollarFilter
	  An OData $filter clause.

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
	/*Timeout
	  The maximum time that the server can spend processing the request, in seconds. The default is 30 seconds.

	*/
	Timeout *int32

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) WithRequestTimeout(timeout time.Duration) *AccountListNodeAgentSkusParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) WithContext(ctx context.Context) *AccountListNodeAgentSkusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDollarFilter adds the dollarFilter to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) WithDollarFilter(dollarFilter *string) *AccountListNodeAgentSkusParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithAPIVersion adds the aPIVersion to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) WithAPIVersion(aPIVersion string) *AccountListNodeAgentSkusParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithClientRequestID adds the clientRequestID to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) WithClientRequestID(clientRequestID *string) *AccountListNodeAgentSkusParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithMaxresults adds the maxresults to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) WithMaxresults(maxresults *int32) *AccountListNodeAgentSkusParams {
	o.SetMaxresults(maxresults)
	return o
}

// SetMaxresults adds the maxresults to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) SetMaxresults(maxresults *int32) {
	o.Maxresults = maxresults
}

// WithOcpDate adds the ocpDate to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) WithOcpDate(ocpDate *string) *AccountListNodeAgentSkusParams {
	o.SetOcpDate(ocpDate)
	return o
}

// SetOcpDate adds the ocpDate to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) SetOcpDate(ocpDate *string) {
	o.OcpDate = ocpDate
}

// WithReturnClientRequestID adds the returnClientRequestID to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) WithReturnClientRequestID(returnClientRequestID *bool) *AccountListNodeAgentSkusParams {
	o.SetReturnClientRequestID(returnClientRequestID)
	return o
}

// SetReturnClientRequestID adds the returnClientRequestId to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) SetReturnClientRequestID(returnClientRequestID *bool) {
	o.ReturnClientRequestID = returnClientRequestID
}

// WithTimeout adds the timeout to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) WithTimeout(timeout *int32) *AccountListNodeAgentSkusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account list node agent skus params
func (o *AccountListNodeAgentSkusParams) SetTimeout(timeout *int32) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *AccountListNodeAgentSkusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
