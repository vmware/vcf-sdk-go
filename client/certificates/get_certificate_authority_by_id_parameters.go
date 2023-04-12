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

// NewGetCertificateAuthorityByIDParams creates a new GetCertificateAuthorityByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCertificateAuthorityByIDParams() *GetCertificateAuthorityByIDParams {
	return &GetCertificateAuthorityByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCertificateAuthorityByIDParamsWithTimeout creates a new GetCertificateAuthorityByIDParams object
// with the ability to set a timeout on a request.
func NewGetCertificateAuthorityByIDParamsWithTimeout(timeout time.Duration) *GetCertificateAuthorityByIDParams {
	return &GetCertificateAuthorityByIDParams{
		timeout: timeout,
	}
}

// NewGetCertificateAuthorityByIDParamsWithContext creates a new GetCertificateAuthorityByIDParams object
// with the ability to set a context for a request.
func NewGetCertificateAuthorityByIDParamsWithContext(ctx context.Context) *GetCertificateAuthorityByIDParams {
	return &GetCertificateAuthorityByIDParams{
		Context: ctx,
	}
}

// NewGetCertificateAuthorityByIDParamsWithHTTPClient creates a new GetCertificateAuthorityByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCertificateAuthorityByIDParamsWithHTTPClient(client *http.Client) *GetCertificateAuthorityByIDParams {
	return &GetCertificateAuthorityByIDParams{
		HTTPClient: client,
	}
}

/*
GetCertificateAuthorityByIDParams contains all the parameters to send to the API endpoint

	for the get certificate authority by Id operation.

	Typically these are written to a http.Request.
*/
type GetCertificateAuthorityByIDParams struct {

	/* ID.

	   CA type
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get certificate authority by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCertificateAuthorityByIDParams) WithDefaults() *GetCertificateAuthorityByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get certificate authority by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCertificateAuthorityByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get certificate authority by Id params
func (o *GetCertificateAuthorityByIDParams) WithTimeout(timeout time.Duration) *GetCertificateAuthorityByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get certificate authority by Id params
func (o *GetCertificateAuthorityByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get certificate authority by Id params
func (o *GetCertificateAuthorityByIDParams) WithContext(ctx context.Context) *GetCertificateAuthorityByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get certificate authority by Id params
func (o *GetCertificateAuthorityByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get certificate authority by Id params
func (o *GetCertificateAuthorityByIDParams) WithHTTPClient(client *http.Client) *GetCertificateAuthorityByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get certificate authority by Id params
func (o *GetCertificateAuthorityByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get certificate authority by Id params
func (o *GetCertificateAuthorityByIDParams) WithID(id string) *GetCertificateAuthorityByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get certificate authority by Id params
func (o *GetCertificateAuthorityByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCertificateAuthorityByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}