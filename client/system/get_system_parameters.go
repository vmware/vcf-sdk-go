// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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
)

// NewGETSystemParams creates a new GETSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETSystemParams() *GETSystemParams {
	return &GETSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETSystemParamsWithTimeout creates a new GETSystemParams object
// with the ability to set a timeout on a request.
func NewGETSystemParamsWithTimeout(timeout time.Duration) *GETSystemParams {
	return &GETSystemParams{
		timeout: timeout,
	}
}

// NewGETSystemParamsWithContext creates a new GETSystemParams object
// with the ability to set a context for a request.
func NewGETSystemParamsWithContext(ctx context.Context) *GETSystemParams {
	return &GETSystemParams{
		Context: ctx,
	}
}

// NewGETSystemParamsWithHTTPClient creates a new GETSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETSystemParamsWithHTTPClient(client *http.Client) *GETSystemParams {
	return &GETSystemParams{
		HTTPClient: client,
	}
}

/*
GETSystemParams contains all the parameters to send to the API endpoint

	for the get system operation.

	Typically these are written to a http.Request.
*/
type GETSystemParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETSystemParams) WithDefaults() *GETSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system params
func (o *GETSystemParams) WithTimeout(timeout time.Duration) *GETSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system params
func (o *GETSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system params
func (o *GETSystemParams) WithContext(ctx context.Context) *GETSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system params
func (o *GETSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system params
func (o *GETSystemParams) WithHTTPClient(client *http.Client) *GETSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system params
func (o *GETSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GETSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
