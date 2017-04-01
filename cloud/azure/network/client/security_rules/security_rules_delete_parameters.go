package security_rules

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

// NewSecurityRulesDeleteParams creates a new SecurityRulesDeleteParams object
// with the default values initialized.
func NewSecurityRulesDeleteParams() *SecurityRulesDeleteParams {
	var ()
	return &SecurityRulesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityRulesDeleteParamsWithTimeout creates a new SecurityRulesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecurityRulesDeleteParamsWithTimeout(timeout time.Duration) *SecurityRulesDeleteParams {
	var ()
	return &SecurityRulesDeleteParams{

		timeout: timeout,
	}
}

// NewSecurityRulesDeleteParamsWithContext creates a new SecurityRulesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecurityRulesDeleteParamsWithContext(ctx context.Context) *SecurityRulesDeleteParams {
	var ()
	return &SecurityRulesDeleteParams{

		Context: ctx,
	}
}

/*SecurityRulesDeleteParams contains all the parameters to send to the API endpoint
for the security rules delete operation typically these are written to a http.Request
*/
type SecurityRulesDeleteParams struct {

	/*APIVersion
	  Client Api Version.

	*/
	APIVersion string
	/*NetworkSecurityGroupName
	  The name of the network security group.

	*/
	NetworkSecurityGroupName string
	/*ResourceGroupName
	  The name of the resource group.

	*/
	ResourceGroupName string
	/*SecurityRuleName
	  The name of the security rule.

	*/
	SecurityRuleName string
	/*SubscriptionID
	  Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the security rules delete params
func (o *SecurityRulesDeleteParams) WithTimeout(timeout time.Duration) *SecurityRulesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security rules delete params
func (o *SecurityRulesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security rules delete params
func (o *SecurityRulesDeleteParams) WithContext(ctx context.Context) *SecurityRulesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security rules delete params
func (o *SecurityRulesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAPIVersion adds the aPIVersion to the security rules delete params
func (o *SecurityRulesDeleteParams) WithAPIVersion(aPIVersion string) *SecurityRulesDeleteParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the security rules delete params
func (o *SecurityRulesDeleteParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithNetworkSecurityGroupName adds the networkSecurityGroupName to the security rules delete params
func (o *SecurityRulesDeleteParams) WithNetworkSecurityGroupName(networkSecurityGroupName string) *SecurityRulesDeleteParams {
	o.SetNetworkSecurityGroupName(networkSecurityGroupName)
	return o
}

// SetNetworkSecurityGroupName adds the networkSecurityGroupName to the security rules delete params
func (o *SecurityRulesDeleteParams) SetNetworkSecurityGroupName(networkSecurityGroupName string) {
	o.NetworkSecurityGroupName = networkSecurityGroupName
}

// WithResourceGroupName adds the resourceGroupName to the security rules delete params
func (o *SecurityRulesDeleteParams) WithResourceGroupName(resourceGroupName string) *SecurityRulesDeleteParams {
	o.SetResourceGroupName(resourceGroupName)
	return o
}

// SetResourceGroupName adds the resourceGroupName to the security rules delete params
func (o *SecurityRulesDeleteParams) SetResourceGroupName(resourceGroupName string) {
	o.ResourceGroupName = resourceGroupName
}

// WithSecurityRuleName adds the securityRuleName to the security rules delete params
func (o *SecurityRulesDeleteParams) WithSecurityRuleName(securityRuleName string) *SecurityRulesDeleteParams {
	o.SetSecurityRuleName(securityRuleName)
	return o
}

// SetSecurityRuleName adds the securityRuleName to the security rules delete params
func (o *SecurityRulesDeleteParams) SetSecurityRuleName(securityRuleName string) {
	o.SecurityRuleName = securityRuleName
}

// WithSubscriptionID adds the subscriptionID to the security rules delete params
func (o *SecurityRulesDeleteParams) WithSubscriptionID(subscriptionID string) *SecurityRulesDeleteParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the security rules delete params
func (o *SecurityRulesDeleteParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityRulesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param networkSecurityGroupName
	if err := r.SetPathParam("networkSecurityGroupName", o.NetworkSecurityGroupName); err != nil {
		return err
	}

	// path param resourceGroupName
	if err := r.SetPathParam("resourceGroupName", o.ResourceGroupName); err != nil {
		return err
	}

	// path param securityRuleName
	if err := r.SetPathParam("securityRuleName", o.SecurityRuleName); err != nil {
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
