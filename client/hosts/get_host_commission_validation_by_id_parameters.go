// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package hosts

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

// NewGetHostCommissionValidationByIDParams creates a new GetHostCommissionValidationByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHostCommissionValidationByIDParams() *GetHostCommissionValidationByIDParams {
	return &GetHostCommissionValidationByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostCommissionValidationByIDParamsWithTimeout creates a new GetHostCommissionValidationByIDParams object
// with the ability to set a timeout on a request.
func NewGetHostCommissionValidationByIDParamsWithTimeout(timeout time.Duration) *GetHostCommissionValidationByIDParams {
	return &GetHostCommissionValidationByIDParams{
		timeout: timeout,
	}
}

// NewGetHostCommissionValidationByIDParamsWithContext creates a new GetHostCommissionValidationByIDParams object
// with the ability to set a context for a request.
func NewGetHostCommissionValidationByIDParamsWithContext(ctx context.Context) *GetHostCommissionValidationByIDParams {
	return &GetHostCommissionValidationByIDParams{
		Context: ctx,
	}
}

// NewGetHostCommissionValidationByIDParamsWithHTTPClient creates a new GetHostCommissionValidationByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHostCommissionValidationByIDParamsWithHTTPClient(client *http.Client) *GetHostCommissionValidationByIDParams {
	return &GetHostCommissionValidationByIDParams{
		HTTPClient: client,
	}
}

/*
GetHostCommissionValidationByIDParams contains all the parameters to send to the API endpoint

	for the get host commission validation by ID operation.

	Typically these are written to a http.Request.
*/
type GetHostCommissionValidationByIDParams struct {

	/* ID.

	   The validation ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get host commission validation by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostCommissionValidationByIDParams) WithDefaults() *GetHostCommissionValidationByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get host commission validation by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostCommissionValidationByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get host commission validation by ID params
func (o *GetHostCommissionValidationByIDParams) WithTimeout(timeout time.Duration) *GetHostCommissionValidationByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get host commission validation by ID params
func (o *GetHostCommissionValidationByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get host commission validation by ID params
func (o *GetHostCommissionValidationByIDParams) WithContext(ctx context.Context) *GetHostCommissionValidationByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get host commission validation by ID params
func (o *GetHostCommissionValidationByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get host commission validation by ID params
func (o *GetHostCommissionValidationByIDParams) WithHTTPClient(client *http.Client) *GetHostCommissionValidationByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get host commission validation by ID params
func (o *GetHostCommissionValidationByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get host commission validation by ID params
func (o *GetHostCommissionValidationByIDParams) WithID(id string) *GetHostCommissionValidationByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get host commission validation by ID params
func (o *GetHostCommissionValidationByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostCommissionValidationByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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