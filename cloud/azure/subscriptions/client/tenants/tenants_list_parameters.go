package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTenantsListParams creates a new TenantsListParams object
// with the default values initialized.
func NewTenantsListParams() *TenantsListParams {
	var ()
	return &TenantsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTenantsListParamsWithTimeout creates a new TenantsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTenantsListParamsWithTimeout(timeout time.Duration) *TenantsListParams {
	var ()
	return &TenantsListParams{

		timeout: timeout,
	}
}

// NewTenantsListParamsWithContext creates a new TenantsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTenantsListParamsWithContext(ctx context.Context) *TenantsListParams {
	var ()
	return &TenantsListParams{

		Context: ctx,
	}
}

/*TenantsListParams contains all the parameters to send to the API endpoint
for the tenants list operation typically these are written to a http.Request
*/
type TenantsListParams struct {

	/*APIVersion
	  The API version to use for the operation.

	*/
	APIVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tenants list params
func (o *TenantsListParams) WithTimeout(timeout time.Duration) *TenantsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenants list params
func (o *TenantsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenants list params
func (o *TenantsListParams) WithContext(ctx context.Context) *TenantsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenants list params
func (o *TenantsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the tenants list params
func (o *TenantsListParams) WithAPIVersion(aPIVersion string) *TenantsListParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the tenants list params
func (o *TenantsListParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *TenantsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// query param api-version
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("api-version", qAPIVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}