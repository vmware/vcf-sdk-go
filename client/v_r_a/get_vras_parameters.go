// Code generated by go-swagger; DO NOT EDIT.

package v_r_a

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

// NewGetVrasParams creates a new GetVrasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVrasParams() *GetVrasParams {
	return &GetVrasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVrasParamsWithTimeout creates a new GetVrasParams object
// with the ability to set a timeout on a request.
func NewGetVrasParamsWithTimeout(timeout time.Duration) *GetVrasParams {
	return &GetVrasParams{
		timeout: timeout,
	}
}

// NewGetVrasParamsWithContext creates a new GetVrasParams object
// with the ability to set a context for a request.
func NewGetVrasParamsWithContext(ctx context.Context) *GetVrasParams {
	return &GetVrasParams{
		Context: ctx,
	}
}

// NewGetVrasParamsWithHTTPClient creates a new GetVrasParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVrasParamsWithHTTPClient(client *http.Client) *GetVrasParams {
	return &GetVrasParams{
		HTTPClient: client,
	}
}

/*
GetVrasParams contains all the parameters to send to the API endpoint

	for the get vras operation.

	Typically these are written to a http.Request.
*/
type GetVrasParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vras params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVrasParams) WithDefaults() *GetVrasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vras params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVrasParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vras params
func (o *GetVrasParams) WithTimeout(timeout time.Duration) *GetVrasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vras params
func (o *GetVrasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vras params
func (o *GetVrasParams) WithContext(ctx context.Context) *GetVrasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vras params
func (o *GetVrasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vras params
func (o *GetVrasParams) WithHTTPClient(client *http.Client) *GetVrasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vras params
func (o *GetVrasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVrasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
