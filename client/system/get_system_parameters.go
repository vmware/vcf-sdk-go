// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

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

// NewGetSystemParams creates a new GetSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemParams() *GetSystemParams {
	return &GetSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemParamsWithTimeout creates a new GetSystemParams object
// with the ability to set a timeout on a request.
func NewGetSystemParamsWithTimeout(timeout time.Duration) *GetSystemParams {
	return &GetSystemParams{
		timeout: timeout,
	}
}

// NewGetSystemParamsWithContext creates a new GetSystemParams object
// with the ability to set a context for a request.
func NewGetSystemParamsWithContext(ctx context.Context) *GetSystemParams {
	return &GetSystemParams{
		Context: ctx,
	}
}

// NewGetSystemParamsWithHTTPClient creates a new GetSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemParamsWithHTTPClient(client *http.Client) *GetSystemParams {
	return &GetSystemParams{
		HTTPClient: client,
	}
}

/*
GetSystemParams contains all the parameters to send to the API endpoint

	for the get system operation.

	Typically these are written to a http.Request.
*/
type GetSystemParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemParams) WithDefaults() *GetSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system params
func (o *GetSystemParams) WithTimeout(timeout time.Duration) *GetSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system params
func (o *GetSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system params
func (o *GetSystemParams) WithContext(ctx context.Context) *GetSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system params
func (o *GetSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system params
func (o *GetSystemParams) WithHTTPClient(client *http.Client) *GetSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system params
func (o *GetSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
