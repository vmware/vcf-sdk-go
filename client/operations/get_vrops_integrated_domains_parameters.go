// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package operations

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

// NewGetVROPSIntegratedDomainsParams creates a new GetVROPSIntegratedDomainsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVROPSIntegratedDomainsParams() *GetVROPSIntegratedDomainsParams {
	return &GetVROPSIntegratedDomainsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVROPSIntegratedDomainsParamsWithTimeout creates a new GetVROPSIntegratedDomainsParams object
// with the ability to set a timeout on a request.
func NewGetVROPSIntegratedDomainsParamsWithTimeout(timeout time.Duration) *GetVROPSIntegratedDomainsParams {
	return &GetVROPSIntegratedDomainsParams{
		timeout: timeout,
	}
}

// NewGetVROPSIntegratedDomainsParamsWithContext creates a new GetVROPSIntegratedDomainsParams object
// with the ability to set a context for a request.
func NewGetVROPSIntegratedDomainsParamsWithContext(ctx context.Context) *GetVROPSIntegratedDomainsParams {
	return &GetVROPSIntegratedDomainsParams{
		Context: ctx,
	}
}

// NewGetVROPSIntegratedDomainsParamsWithHTTPClient creates a new GetVROPSIntegratedDomainsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVROPSIntegratedDomainsParamsWithHTTPClient(client *http.Client) *GetVROPSIntegratedDomainsParams {
	return &GetVROPSIntegratedDomainsParams{
		HTTPClient: client,
	}
}

/*
GetVROPSIntegratedDomainsParams contains all the parameters to send to the API endpoint

	for the get Vrops integrated domains operation.

	Typically these are written to a http.Request.
*/
type GetVROPSIntegratedDomainsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vrops integrated domains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVROPSIntegratedDomainsParams) WithDefaults() *GetVROPSIntegratedDomainsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vrops integrated domains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVROPSIntegratedDomainsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Vrops integrated domains params
func (o *GetVROPSIntegratedDomainsParams) WithTimeout(timeout time.Duration) *GetVROPSIntegratedDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vrops integrated domains params
func (o *GetVROPSIntegratedDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vrops integrated domains params
func (o *GetVROPSIntegratedDomainsParams) WithContext(ctx context.Context) *GetVROPSIntegratedDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vrops integrated domains params
func (o *GetVROPSIntegratedDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vrops integrated domains params
func (o *GetVROPSIntegratedDomainsParams) WithHTTPClient(client *http.Client) *GetVROPSIntegratedDomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vrops integrated domains params
func (o *GetVROPSIntegratedDomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVROPSIntegratedDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
