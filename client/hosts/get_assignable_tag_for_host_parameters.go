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

// NewGetAssignableTagForHostParams creates a new GetAssignableTagForHostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAssignableTagForHostParams() *GetAssignableTagForHostParams {
	return &GetAssignableTagForHostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssignableTagForHostParamsWithTimeout creates a new GetAssignableTagForHostParams object
// with the ability to set a timeout on a request.
func NewGetAssignableTagForHostParamsWithTimeout(timeout time.Duration) *GetAssignableTagForHostParams {
	return &GetAssignableTagForHostParams{
		timeout: timeout,
	}
}

// NewGetAssignableTagForHostParamsWithContext creates a new GetAssignableTagForHostParams object
// with the ability to set a context for a request.
func NewGetAssignableTagForHostParamsWithContext(ctx context.Context) *GetAssignableTagForHostParams {
	return &GetAssignableTagForHostParams{
		Context: ctx,
	}
}

// NewGetAssignableTagForHostParamsWithHTTPClient creates a new GetAssignableTagForHostParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAssignableTagForHostParamsWithHTTPClient(client *http.Client) *GetAssignableTagForHostParams {
	return &GetAssignableTagForHostParams{
		HTTPClient: client,
	}
}

/*
GetAssignableTagForHostParams contains all the parameters to send to the API endpoint

	for the get assignable tag for host operation.

	Typically these are written to a http.Request.
*/
type GetAssignableTagForHostParams struct {

	/* ID.

	   Host ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get assignable tag for host params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAssignableTagForHostParams) WithDefaults() *GetAssignableTagForHostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get assignable tag for host params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAssignableTagForHostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get assignable tag for host params
func (o *GetAssignableTagForHostParams) WithTimeout(timeout time.Duration) *GetAssignableTagForHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get assignable tag for host params
func (o *GetAssignableTagForHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get assignable tag for host params
func (o *GetAssignableTagForHostParams) WithContext(ctx context.Context) *GetAssignableTagForHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get assignable tag for host params
func (o *GetAssignableTagForHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get assignable tag for host params
func (o *GetAssignableTagForHostParams) WithHTTPClient(client *http.Client) *GetAssignableTagForHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get assignable tag for host params
func (o *GetAssignableTagForHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get assignable tag for host params
func (o *GetAssignableTagForHostParams) WithID(id string) *GetAssignableTagForHostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get assignable tag for host params
func (o *GetAssignableTagForHostParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssignableTagForHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
