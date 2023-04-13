// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package users

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

// NewGETLocalAccountParams creates a new GETLocalAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETLocalAccountParams() *GETLocalAccountParams {
	return &GETLocalAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETLocalAccountParamsWithTimeout creates a new GETLocalAccountParams object
// with the ability to set a timeout on a request.
func NewGETLocalAccountParamsWithTimeout(timeout time.Duration) *GETLocalAccountParams {
	return &GETLocalAccountParams{
		timeout: timeout,
	}
}

// NewGETLocalAccountParamsWithContext creates a new GETLocalAccountParams object
// with the ability to set a context for a request.
func NewGETLocalAccountParamsWithContext(ctx context.Context) *GETLocalAccountParams {
	return &GETLocalAccountParams{
		Context: ctx,
	}
}

// NewGETLocalAccountParamsWithHTTPClient creates a new GETLocalAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETLocalAccountParamsWithHTTPClient(client *http.Client) *GETLocalAccountParams {
	return &GETLocalAccountParams{
		HTTPClient: client,
	}
}

/*
GETLocalAccountParams contains all the parameters to send to the API endpoint

	for the get local account operation.

	Typically these are written to a http.Request.
*/
type GETLocalAccountParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get local account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETLocalAccountParams) WithDefaults() *GETLocalAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get local account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETLocalAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get local account params
func (o *GETLocalAccountParams) WithTimeout(timeout time.Duration) *GETLocalAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get local account params
func (o *GETLocalAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get local account params
func (o *GETLocalAccountParams) WithContext(ctx context.Context) *GETLocalAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get local account params
func (o *GETLocalAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get local account params
func (o *GETLocalAccountParams) WithHTTPClient(client *http.Client) *GETLocalAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get local account params
func (o *GETLocalAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GETLocalAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
