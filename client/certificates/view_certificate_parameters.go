// Code generated by go-swagger; DO NOT EDIT.

package certificates

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
)

// NewViewCertificateParams creates a new ViewCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewViewCertificateParams() *ViewCertificateParams {
	return &ViewCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewViewCertificateParamsWithTimeout creates a new ViewCertificateParams object
// with the ability to set a timeout on a request.
func NewViewCertificateParamsWithTimeout(timeout time.Duration) *ViewCertificateParams {
	return &ViewCertificateParams{
		timeout: timeout,
	}
}

// NewViewCertificateParamsWithContext creates a new ViewCertificateParams object
// with the ability to set a context for a request.
func NewViewCertificateParamsWithContext(ctx context.Context) *ViewCertificateParams {
	return &ViewCertificateParams{
		Context: ctx,
	}
}

// NewViewCertificateParamsWithHTTPClient creates a new ViewCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewViewCertificateParamsWithHTTPClient(client *http.Client) *ViewCertificateParams {
	return &ViewCertificateParams{
		HTTPClient: client,
	}
}

/*
ViewCertificateParams contains all the parameters to send to the API endpoint

	for the view certificate operation.

	Typically these are written to a http.Request.
*/
type ViewCertificateParams struct {

	/* DomainName.

	   The domain name
	*/
	DomainName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the view certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ViewCertificateParams) WithDefaults() *ViewCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the view certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ViewCertificateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the view certificate params
func (o *ViewCertificateParams) WithTimeout(timeout time.Duration) *ViewCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the view certificate params
func (o *ViewCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the view certificate params
func (o *ViewCertificateParams) WithContext(ctx context.Context) *ViewCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the view certificate params
func (o *ViewCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the view certificate params
func (o *ViewCertificateParams) WithHTTPClient(client *http.Client) *ViewCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the view certificate params
func (o *ViewCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainName adds the domainName to the view certificate params
func (o *ViewCertificateParams) WithDomainName(domainName string) *ViewCertificateParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the view certificate params
func (o *ViewCertificateParams) SetDomainName(domainName string) {
	o.DomainName = domainName
}

// WriteToRequest writes these params to a swagger request
func (o *ViewCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainName
	if err := r.SetPathParam("domainName", o.DomainName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
