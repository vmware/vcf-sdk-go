// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package upgradables

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

// NewGETUpgradablesParams creates a new GETUpgradablesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETUpgradablesParams() *GETUpgradablesParams {
	return &GETUpgradablesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETUpgradablesParamsWithTimeout creates a new GETUpgradablesParams object
// with the ability to set a timeout on a request.
func NewGETUpgradablesParamsWithTimeout(timeout time.Duration) *GETUpgradablesParams {
	return &GETUpgradablesParams{
		timeout: timeout,
	}
}

// NewGETUpgradablesParamsWithContext creates a new GETUpgradablesParams object
// with the ability to set a context for a request.
func NewGETUpgradablesParamsWithContext(ctx context.Context) *GETUpgradablesParams {
	return &GETUpgradablesParams{
		Context: ctx,
	}
}

// NewGETUpgradablesParamsWithHTTPClient creates a new GETUpgradablesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETUpgradablesParamsWithHTTPClient(client *http.Client) *GETUpgradablesParams {
	return &GETUpgradablesParams{
		HTTPClient: client,
	}
}

/*
GETUpgradablesParams contains all the parameters to send to the API endpoint

	for the get upgradables operation.

	Typically these are written to a http.Request.
*/
type GETUpgradablesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get upgradables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETUpgradablesParams) WithDefaults() *GETUpgradablesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get upgradables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETUpgradablesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get upgradables params
func (o *GETUpgradablesParams) WithTimeout(timeout time.Duration) *GETUpgradablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get upgradables params
func (o *GETUpgradablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get upgradables params
func (o *GETUpgradablesParams) WithContext(ctx context.Context) *GETUpgradablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get upgradables params
func (o *GETUpgradablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get upgradables params
func (o *GETUpgradablesParams) WithHTTPClient(client *http.Client) *GETUpgradablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get upgradables params
func (o *GETUpgradablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GETUpgradablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
