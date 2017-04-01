package load_balancers

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

	"github.com/jkawamoto/roadie/cloud/azure/network/models"
)

// NewLoadBalancersCreateOrUpdateParams creates a new LoadBalancersCreateOrUpdateParams object
// with the default values initialized.
func NewLoadBalancersCreateOrUpdateParams() *LoadBalancersCreateOrUpdateParams {
	var ()
	return &LoadBalancersCreateOrUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadBalancersCreateOrUpdateParamsWithTimeout creates a new LoadBalancersCreateOrUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadBalancersCreateOrUpdateParamsWithTimeout(timeout time.Duration) *LoadBalancersCreateOrUpdateParams {
	var ()
	return &LoadBalancersCreateOrUpdateParams{

		timeout: timeout,
	}
}

// NewLoadBalancersCreateOrUpdateParamsWithContext creates a new LoadBalancersCreateOrUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadBalancersCreateOrUpdateParamsWithContext(ctx context.Context) *LoadBalancersCreateOrUpdateParams {
	var ()
	return &LoadBalancersCreateOrUpdateParams{

		Context: ctx,
	}
}

/*LoadBalancersCreateOrUpdateParams contains all the parameters to send to the API endpoint
for the load balancers create or update operation typically these are written to a http.Request
*/
type LoadBalancersCreateOrUpdateParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*LoadBalancerName
	  The name of the loadBalancer.

	*/
	LoadBalancerName string
	/*Parameters
	  Parameters supplied to the create/delete LoadBalancer operation

	*/
	Parameters *models.LoadBalancer
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*SubscriptionID
	  Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) WithTimeout(timeout time.Duration) *LoadBalancersCreateOrUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) WithContext(ctx context.Context) *LoadBalancersCreateOrUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) WithAPIVersion(aPIVersion string) *LoadBalancersCreateOrUpdateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithLoadBalancerName adds the loadBalancerName to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) WithLoadBalancerName(loadBalancerName string) *LoadBalancersCreateOrUpdateParams {
	o.SetLoadBalancerName(loadBalancerName)
	return o
}

// SetLoadBalancerName adds the loadBalancerName to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) SetLoadBalancerName(loadBalancerName string) {
	o.LoadBalancerName = loadBalancerName
}

// WithParameters adds the parameters to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) WithParameters(parameters *models.LoadBalancer) *LoadBalancersCreateOrUpdateParams {
	o.SetParameters(parameters)
	return o
}

// SetParameters adds the parameters to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) SetParameters(parameters *models.LoadBalancer) {
	o.Parameters = parameters
}

// WithResourceGroupName adds the resourceGroupName to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) WithResourceGroupName(resourceGroupName string) *LoadBalancersCreateOrUpdateParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSubscriptionID adds the subscriptionID to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) WithSubscriptionID(subscriptionID string) *LoadBalancersCreateOrUpdateParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the load balancers create or update params
func (o *LoadBalancersCreateOrUpdateParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *LoadBalancersCreateOrUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param loadBalancerName
	if err := r.SetPathParam("loadBalancerName", o.LoadBalancerName); err != nil {
		return err
	}

	if o.Parameters == nil {
		o.Parameters = new(models.LoadBalancer)
	}

	if err := r.SetBodyParam(o.Parameters); err != nil {
		return err
	}

	// path param resourceGroupName
	if err := r.SetPathParam("resourceGroupName", o.ResourceGroupName); err != nil {
		return err
	}

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
