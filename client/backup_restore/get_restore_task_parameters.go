// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package backup_restore

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

// NewGetRestoreTaskParams creates a new GetRestoreTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRestoreTaskParams() *GetRestoreTaskParams {
	return &GetRestoreTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestoreTaskParamsWithTimeout creates a new GetRestoreTaskParams object
// with the ability to set a timeout on a request.
func NewGetRestoreTaskParamsWithTimeout(timeout time.Duration) *GetRestoreTaskParams {
	return &GetRestoreTaskParams{
		timeout: timeout,
	}
}

// NewGetRestoreTaskParamsWithContext creates a new GetRestoreTaskParams object
// with the ability to set a context for a request.
func NewGetRestoreTaskParamsWithContext(ctx context.Context) *GetRestoreTaskParams {
	return &GetRestoreTaskParams{
		Context: ctx,
	}
}

// NewGetRestoreTaskParamsWithHTTPClient creates a new GetRestoreTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRestoreTaskParamsWithHTTPClient(client *http.Client) *GetRestoreTaskParams {
	return &GetRestoreTaskParams{
		HTTPClient: client,
	}
}

/*
GetRestoreTaskParams contains all the parameters to send to the API endpoint

	for the get restore task operation.

	Typically these are written to a http.Request.
*/
type GetRestoreTaskParams struct {

	/* ID.

	   The restore task ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get restore task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRestoreTaskParams) WithDefaults() *GetRestoreTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get restore task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRestoreTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get restore task params
func (o *GetRestoreTaskParams) WithTimeout(timeout time.Duration) *GetRestoreTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restore task params
func (o *GetRestoreTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restore task params
func (o *GetRestoreTaskParams) WithContext(ctx context.Context) *GetRestoreTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restore task params
func (o *GetRestoreTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restore task params
func (o *GetRestoreTaskParams) WithHTTPClient(client *http.Client) *GetRestoreTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restore task params
func (o *GetRestoreTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get restore task params
func (o *GetRestoreTaskParams) WithID(id string) *GetRestoreTaskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get restore task params
func (o *GetRestoreTaskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestoreTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
