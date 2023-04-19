// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vrli

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

// NewGETIntegratedDomains1Params creates a new GETIntegratedDomains1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETIntegratedDomains1Params() *GETIntegratedDomains1Params {
	return &GETIntegratedDomains1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETIntegratedDomains1ParamsWithTimeout creates a new GETIntegratedDomains1Params object
// with the ability to set a timeout on a request.
func NewGETIntegratedDomains1ParamsWithTimeout(timeout time.Duration) *GETIntegratedDomains1Params {
	return &GETIntegratedDomains1Params{
		timeout: timeout,
	}
}

// NewGETIntegratedDomains1ParamsWithContext creates a new GETIntegratedDomains1Params object
// with the ability to set a context for a request.
func NewGETIntegratedDomains1ParamsWithContext(ctx context.Context) *GETIntegratedDomains1Params {
	return &GETIntegratedDomains1Params{
		Context: ctx,
	}
}

// NewGETIntegratedDomains1ParamsWithHTTPClient creates a new GETIntegratedDomains1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGETIntegratedDomains1ParamsWithHTTPClient(client *http.Client) *GETIntegratedDomains1Params {
	return &GETIntegratedDomains1Params{
		HTTPClient: client,
	}
}

/*
GETIntegratedDomains1Params contains all the parameters to send to the API endpoint

	for the get integrated domains 1 operation.

	Typically these are written to a http.Request.
*/
type GETIntegratedDomains1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get integrated domains 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETIntegratedDomains1Params) WithDefaults() *GETIntegratedDomains1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get integrated domains 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETIntegratedDomains1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get integrated domains 1 params
func (o *GETIntegratedDomains1Params) WithTimeout(timeout time.Duration) *GETIntegratedDomains1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrated domains 1 params
func (o *GETIntegratedDomains1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrated domains 1 params
func (o *GETIntegratedDomains1Params) WithContext(ctx context.Context) *GETIntegratedDomains1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrated domains 1 params
func (o *GETIntegratedDomains1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrated domains 1 params
func (o *GETIntegratedDomains1Params) WithHTTPClient(client *http.Client) *GETIntegratedDomains1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrated domains 1 params
func (o *GETIntegratedDomains1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GETIntegratedDomains1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
