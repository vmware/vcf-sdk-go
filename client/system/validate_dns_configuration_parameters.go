// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"vcf-sdk-go/models"
)

// NewValidateDNSConfigurationParams creates a new ValidateDNSConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateDNSConfigurationParams() *ValidateDNSConfigurationParams {
	return &ValidateDNSConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateDNSConfigurationParamsWithTimeout creates a new ValidateDNSConfigurationParams object
// with the ability to set a timeout on a request.
func NewValidateDNSConfigurationParamsWithTimeout(timeout time.Duration) *ValidateDNSConfigurationParams {
	return &ValidateDNSConfigurationParams{
		timeout: timeout,
	}
}

// NewValidateDNSConfigurationParamsWithContext creates a new ValidateDNSConfigurationParams object
// with the ability to set a context for a request.
func NewValidateDNSConfigurationParamsWithContext(ctx context.Context) *ValidateDNSConfigurationParams {
	return &ValidateDNSConfigurationParams{
		Context: ctx,
	}
}

// NewValidateDNSConfigurationParamsWithHTTPClient creates a new ValidateDNSConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateDNSConfigurationParamsWithHTTPClient(client *http.Client) *ValidateDNSConfigurationParams {
	return &ValidateDNSConfigurationParams{
		HTTPClient: client,
	}
}

/*
ValidateDNSConfigurationParams contains all the parameters to send to the API endpoint

	for the validate Dns configuration operation.

	Typically these are written to a http.Request.
*/
type ValidateDNSConfigurationParams struct {

	/* DNSConfiguration.

	   dnsConfiguration
	*/
	DNSConfiguration *models.DNSConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate Dns configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateDNSConfigurationParams) WithDefaults() *ValidateDNSConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate Dns configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateDNSConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate Dns configuration params
func (o *ValidateDNSConfigurationParams) WithTimeout(timeout time.Duration) *ValidateDNSConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate Dns configuration params
func (o *ValidateDNSConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate Dns configuration params
func (o *ValidateDNSConfigurationParams) WithContext(ctx context.Context) *ValidateDNSConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate Dns configuration params
func (o *ValidateDNSConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate Dns configuration params
func (o *ValidateDNSConfigurationParams) WithHTTPClient(client *http.Client) *ValidateDNSConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate Dns configuration params
func (o *ValidateDNSConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDNSConfiguration adds the dNSConfiguration to the validate Dns configuration params
func (o *ValidateDNSConfigurationParams) WithDNSConfiguration(dNSConfiguration *models.DNSConfiguration) *ValidateDNSConfigurationParams {
	o.SetDNSConfiguration(dNSConfiguration)
	return o
}

// SetDNSConfiguration adds the dnsConfiguration to the validate Dns configuration params
func (o *ValidateDNSConfigurationParams) SetDNSConfiguration(dNSConfiguration *models.DNSConfiguration) {
	o.DNSConfiguration = dNSConfiguration
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateDNSConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.DNSConfiguration != nil {
		if err := r.SetBodyParam(o.DNSConfiguration); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
