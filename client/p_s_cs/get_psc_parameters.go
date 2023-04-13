// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package p_s_cs

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

// NewGETPscParams creates a new GETPscParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETPscParams() *GETPscParams {
	return &GETPscParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETPscParamsWithTimeout creates a new GETPscParams object
// with the ability to set a timeout on a request.
func NewGETPscParamsWithTimeout(timeout time.Duration) *GETPscParams {
	return &GETPscParams{
		timeout: timeout,
	}
}

// NewGETPscParamsWithContext creates a new GETPscParams object
// with the ability to set a context for a request.
func NewGETPscParamsWithContext(ctx context.Context) *GETPscParams {
	return &GETPscParams{
		Context: ctx,
	}
}

// NewGETPscParamsWithHTTPClient creates a new GETPscParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETPscParamsWithHTTPClient(client *http.Client) *GETPscParams {
	return &GETPscParams{
		HTTPClient: client,
	}
}

/*
GETPscParams contains all the parameters to send to the API endpoint

	for the get psc operation.

	Typically these are written to a http.Request.
*/
type GETPscParams struct {

	/* ID.

	   PSC ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get psc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETPscParams) WithDefaults() *GETPscParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get psc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETPscParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get psc params
func (o *GETPscParams) WithTimeout(timeout time.Duration) *GETPscParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get psc params
func (o *GETPscParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get psc params
func (o *GETPscParams) WithContext(ctx context.Context) *GETPscParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get psc params
func (o *GETPscParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get psc params
func (o *GETPscParams) WithHTTPClient(client *http.Client) *GETPscParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get psc params
func (o *GETPscParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get psc params
func (o *GETPscParams) WithID(id string) *GETPscParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get psc params
func (o *GETPscParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETPscParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
