// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc_managers

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

// NewGETSDDCManagerParams creates a new GETSDDCManagerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETSDDCManagerParams() *GETSDDCManagerParams {
	return &GETSDDCManagerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETSDDCManagerParamsWithTimeout creates a new GETSDDCManagerParams object
// with the ability to set a timeout on a request.
func NewGETSDDCManagerParamsWithTimeout(timeout time.Duration) *GETSDDCManagerParams {
	return &GETSDDCManagerParams{
		timeout: timeout,
	}
}

// NewGETSDDCManagerParamsWithContext creates a new GETSDDCManagerParams object
// with the ability to set a context for a request.
func NewGETSDDCManagerParamsWithContext(ctx context.Context) *GETSDDCManagerParams {
	return &GETSDDCManagerParams{
		Context: ctx,
	}
}

// NewGETSDDCManagerParamsWithHTTPClient creates a new GETSDDCManagerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETSDDCManagerParamsWithHTTPClient(client *http.Client) *GETSDDCManagerParams {
	return &GETSDDCManagerParams{
		HTTPClient: client,
	}
}

/*
GETSDDCManagerParams contains all the parameters to send to the API endpoint

	for the get Sddc manager operation.

	Typically these are written to a http.Request.
*/
type GETSDDCManagerParams struct {

	/* ID.

	   Sddc Manager ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Sddc manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETSDDCManagerParams) WithDefaults() *GETSDDCManagerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Sddc manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETSDDCManagerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Sddc manager params
func (o *GETSDDCManagerParams) WithTimeout(timeout time.Duration) *GETSDDCManagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Sddc manager params
func (o *GETSDDCManagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Sddc manager params
func (o *GETSDDCManagerParams) WithContext(ctx context.Context) *GETSDDCManagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Sddc manager params
func (o *GETSDDCManagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Sddc manager params
func (o *GETSDDCManagerParams) WithHTTPClient(client *http.Client) *GETSDDCManagerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Sddc manager params
func (o *GETSDDCManagerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get Sddc manager params
func (o *GETSDDCManagerParams) WithID(id string) *GETSDDCManagerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get Sddc manager params
func (o *GETSDDCManagerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETSDDCManagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
