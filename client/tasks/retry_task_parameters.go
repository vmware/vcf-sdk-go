// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package tasks

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

// NewRetryTaskParams creates a new RetryTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetryTaskParams() *RetryTaskParams {
	return &RetryTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetryTaskParamsWithTimeout creates a new RetryTaskParams object
// with the ability to set a timeout on a request.
func NewRetryTaskParamsWithTimeout(timeout time.Duration) *RetryTaskParams {
	return &RetryTaskParams{
		timeout: timeout,
	}
}

// NewRetryTaskParamsWithContext creates a new RetryTaskParams object
// with the ability to set a context for a request.
func NewRetryTaskParamsWithContext(ctx context.Context) *RetryTaskParams {
	return &RetryTaskParams{
		Context: ctx,
	}
}

// NewRetryTaskParamsWithHTTPClient creates a new RetryTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetryTaskParamsWithHTTPClient(client *http.Client) *RetryTaskParams {
	return &RetryTaskParams{
		HTTPClient: client,
	}
}

/*
RetryTaskParams contains all the parameters to send to the API endpoint

	for the retry task operation.

	Typically these are written to a http.Request.
*/
type RetryTaskParams struct {

	/* ID.

	   Task id retry
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retry task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetryTaskParams) WithDefaults() *RetryTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retry task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetryTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retry task params
func (o *RetryTaskParams) WithTimeout(timeout time.Duration) *RetryTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retry task params
func (o *RetryTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retry task params
func (o *RetryTaskParams) WithContext(ctx context.Context) *RetryTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retry task params
func (o *RetryTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retry task params
func (o *RetryTaskParams) WithHTTPClient(client *http.Client) *RetryTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retry task params
func (o *RetryTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the retry task params
func (o *RetryTaskParams) WithID(id string) *RetryTaskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the retry task params
func (o *RetryTaskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RetryTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
