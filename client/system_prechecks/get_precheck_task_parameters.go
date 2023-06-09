// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package system_prechecks

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

// NewGetPrecheckTaskParams creates a new GetPrecheckTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPrecheckTaskParams() *GetPrecheckTaskParams {
	return &GetPrecheckTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrecheckTaskParamsWithTimeout creates a new GetPrecheckTaskParams object
// with the ability to set a timeout on a request.
func NewGetPrecheckTaskParamsWithTimeout(timeout time.Duration) *GetPrecheckTaskParams {
	return &GetPrecheckTaskParams{
		timeout: timeout,
	}
}

// NewGetPrecheckTaskParamsWithContext creates a new GetPrecheckTaskParams object
// with the ability to set a context for a request.
func NewGetPrecheckTaskParamsWithContext(ctx context.Context) *GetPrecheckTaskParams {
	return &GetPrecheckTaskParams{
		Context: ctx,
	}
}

// NewGetPrecheckTaskParamsWithHTTPClient creates a new GetPrecheckTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPrecheckTaskParamsWithHTTPClient(client *http.Client) *GetPrecheckTaskParams {
	return &GetPrecheckTaskParams{
		HTTPClient: client,
	}
}

/*
GetPrecheckTaskParams contains all the parameters to send to the API endpoint

	for the get precheck task operation.

	Typically these are written to a http.Request.
*/
type GetPrecheckTaskParams struct {

	/* ID.

	   Precheck Task ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get precheck task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPrecheckTaskParams) WithDefaults() *GetPrecheckTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get precheck task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPrecheckTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get precheck task params
func (o *GetPrecheckTaskParams) WithTimeout(timeout time.Duration) *GetPrecheckTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get precheck task params
func (o *GetPrecheckTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get precheck task params
func (o *GetPrecheckTaskParams) WithContext(ctx context.Context) *GetPrecheckTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get precheck task params
func (o *GetPrecheckTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get precheck task params
func (o *GetPrecheckTaskParams) WithHTTPClient(client *http.Client) *GetPrecheckTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get precheck task params
func (o *GetPrecheckTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get precheck task params
func (o *GetPrecheckTaskParams) WithID(id string) *GetPrecheckTaskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get precheck task params
func (o *GetPrecheckTaskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrecheckTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
