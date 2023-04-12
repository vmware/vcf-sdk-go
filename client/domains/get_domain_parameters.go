// Code generated by go-swagger; DO NOT EDIT.

package domains

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

// NewGetDomainParams creates a new GetDomainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomainParams() *GetDomainParams {
	return &GetDomainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainParamsWithTimeout creates a new GetDomainParams object
// with the ability to set a timeout on a request.
func NewGetDomainParamsWithTimeout(timeout time.Duration) *GetDomainParams {
	return &GetDomainParams{
		timeout: timeout,
	}
}

// NewGetDomainParamsWithContext creates a new GetDomainParams object
// with the ability to set a context for a request.
func NewGetDomainParamsWithContext(ctx context.Context) *GetDomainParams {
	return &GetDomainParams{
		Context: ctx,
	}
}

// NewGetDomainParamsWithHTTPClient creates a new GetDomainParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomainParamsWithHTTPClient(client *http.Client) *GetDomainParams {
	return &GetDomainParams{
		HTTPClient: client,
	}
}

/*
GetDomainParams contains all the parameters to send to the API endpoint

	for the get domain operation.

	Typically these are written to a http.Request.
*/
type GetDomainParams struct {

	/* ID.

	   Domain ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainParams) WithDefaults() *GetDomainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domain params
func (o *GetDomainParams) WithTimeout(timeout time.Duration) *GetDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain params
func (o *GetDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain params
func (o *GetDomainParams) WithContext(ctx context.Context) *GetDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain params
func (o *GetDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain params
func (o *GetDomainParams) WithHTTPClient(client *http.Client) *GetDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain params
func (o *GetDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get domain params
func (o *GetDomainParams) WithID(id string) *GetDomainParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get domain params
func (o *GetDomainParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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