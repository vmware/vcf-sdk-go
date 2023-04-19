// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

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

// NewGETSDDCManagerInfoParams creates a new GETSDDCManagerInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETSDDCManagerInfoParams() *GETSDDCManagerInfoParams {
	return &GETSDDCManagerInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETSDDCManagerInfoParamsWithTimeout creates a new GETSDDCManagerInfoParams object
// with the ability to set a timeout on a request.
func NewGETSDDCManagerInfoParamsWithTimeout(timeout time.Duration) *GETSDDCManagerInfoParams {
	return &GETSDDCManagerInfoParams{
		timeout: timeout,
	}
}

// NewGETSDDCManagerInfoParamsWithContext creates a new GETSDDCManagerInfoParams object
// with the ability to set a context for a request.
func NewGETSDDCManagerInfoParamsWithContext(ctx context.Context) *GETSDDCManagerInfoParams {
	return &GETSDDCManagerInfoParams{
		Context: ctx,
	}
}

// NewGETSDDCManagerInfoParamsWithHTTPClient creates a new GETSDDCManagerInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETSDDCManagerInfoParamsWithHTTPClient(client *http.Client) *GETSDDCManagerInfoParams {
	return &GETSDDCManagerInfoParams{
		HTTPClient: client,
	}
}

/*
GETSDDCManagerInfoParams contains all the parameters to send to the API endpoint

	for the get Sddc manager info operation.

	Typically these are written to a http.Request.
*/
type GETSDDCManagerInfoParams struct {

	/* ID.

	   SDDC ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Sddc manager info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETSDDCManagerInfoParams) WithDefaults() *GETSDDCManagerInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Sddc manager info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETSDDCManagerInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Sddc manager info params
func (o *GETSDDCManagerInfoParams) WithTimeout(timeout time.Duration) *GETSDDCManagerInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Sddc manager info params
func (o *GETSDDCManagerInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Sddc manager info params
func (o *GETSDDCManagerInfoParams) WithContext(ctx context.Context) *GETSDDCManagerInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Sddc manager info params
func (o *GETSDDCManagerInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Sddc manager info params
func (o *GETSDDCManagerInfoParams) WithHTTPClient(client *http.Client) *GETSDDCManagerInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Sddc manager info params
func (o *GETSDDCManagerInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get Sddc manager info params
func (o *GETSDDCManagerInfoParams) WithID(id string) *GETSDDCManagerInfoParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get Sddc manager info params
func (o *GETSDDCManagerInfoParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETSDDCManagerInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
