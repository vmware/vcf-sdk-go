// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

package vasa_providers

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

// NewGETVasaProvidersParams creates a new GETVasaProvidersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETVasaProvidersParams() *GETVasaProvidersParams {
	return &GETVasaProvidersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETVasaProvidersParamsWithTimeout creates a new GETVasaProvidersParams object
// with the ability to set a timeout on a request.
func NewGETVasaProvidersParamsWithTimeout(timeout time.Duration) *GETVasaProvidersParams {
	return &GETVasaProvidersParams{
		timeout: timeout,
	}
}

// NewGETVasaProvidersParamsWithContext creates a new GETVasaProvidersParams object
// with the ability to set a context for a request.
func NewGETVasaProvidersParamsWithContext(ctx context.Context) *GETVasaProvidersParams {
	return &GETVasaProvidersParams{
		Context: ctx,
	}
}

// NewGETVasaProvidersParamsWithHTTPClient creates a new GETVasaProvidersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETVasaProvidersParamsWithHTTPClient(client *http.Client) *GETVasaProvidersParams {
	return &GETVasaProvidersParams{
		HTTPClient: client,
	}
}

/*
GETVasaProvidersParams contains all the parameters to send to the API endpoint

	for the get vasa providers operation.

	Typically these are written to a http.Request.
*/
type GETVasaProvidersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vasa providers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETVasaProvidersParams) WithDefaults() *GETVasaProvidersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vasa providers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETVasaProvidersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vasa providers params
func (o *GETVasaProvidersParams) WithTimeout(timeout time.Duration) *GETVasaProvidersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vasa providers params
func (o *GETVasaProvidersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vasa providers params
func (o *GETVasaProvidersParams) WithContext(ctx context.Context) *GETVasaProvidersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vasa providers params
func (o *GETVasaProvidersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vasa providers params
func (o *GETVasaProvidersParams) WithHTTPClient(client *http.Client) *GETVasaProvidersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vasa providers params
func (o *GETVasaProvidersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GETVasaProvidersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
