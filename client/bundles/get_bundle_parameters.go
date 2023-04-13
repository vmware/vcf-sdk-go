// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package bundles

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

// NewGETBundleParams creates a new GETBundleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETBundleParams() *GETBundleParams {
	return &GETBundleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETBundleParamsWithTimeout creates a new GETBundleParams object
// with the ability to set a timeout on a request.
func NewGETBundleParamsWithTimeout(timeout time.Duration) *GETBundleParams {
	return &GETBundleParams{
		timeout: timeout,
	}
}

// NewGETBundleParamsWithContext creates a new GETBundleParams object
// with the ability to set a context for a request.
func NewGETBundleParamsWithContext(ctx context.Context) *GETBundleParams {
	return &GETBundleParams{
		Context: ctx,
	}
}

// NewGETBundleParamsWithHTTPClient creates a new GETBundleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETBundleParamsWithHTTPClient(client *http.Client) *GETBundleParams {
	return &GETBundleParams{
		HTTPClient: client,
	}
}

/*
GETBundleParams contains all the parameters to send to the API endpoint

	for the get bundle operation.

	Typically these are written to a http.Request.
*/
type GETBundleParams struct {

	/* ID.

	   Bundle ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETBundleParams) WithDefaults() *GETBundleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETBundleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bundle params
func (o *GETBundleParams) WithTimeout(timeout time.Duration) *GETBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundle params
func (o *GETBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundle params
func (o *GETBundleParams) WithContext(ctx context.Context) *GETBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundle params
func (o *GETBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundle params
func (o *GETBundleParams) WithHTTPClient(client *http.Client) *GETBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundle params
func (o *GETBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get bundle params
func (o *GETBundleParams) WithID(id string) *GETBundleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get bundle params
func (o *GETBundleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
