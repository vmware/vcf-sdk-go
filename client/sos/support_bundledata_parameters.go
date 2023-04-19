// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sos

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

// NewSupportBundledataParams creates a new SupportBundledataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSupportBundledataParams() *SupportBundledataParams {
	return &SupportBundledataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSupportBundledataParamsWithTimeout creates a new SupportBundledataParams object
// with the ability to set a timeout on a request.
func NewSupportBundledataParamsWithTimeout(timeout time.Duration) *SupportBundledataParams {
	return &SupportBundledataParams{
		timeout: timeout,
	}
}

// NewSupportBundledataParamsWithContext creates a new SupportBundledataParams object
// with the ability to set a context for a request.
func NewSupportBundledataParamsWithContext(ctx context.Context) *SupportBundledataParams {
	return &SupportBundledataParams{
		Context: ctx,
	}
}

// NewSupportBundledataParamsWithHTTPClient creates a new SupportBundledataParams object
// with the ability to set a custom HTTPClient for a request.
func NewSupportBundledataParamsWithHTTPClient(client *http.Client) *SupportBundledataParams {
	return &SupportBundledataParams{
		HTTPClient: client,
	}
}

/*
SupportBundledataParams contains all the parameters to send to the API endpoint

	for the support bundledata operation.

	Typically these are written to a http.Request.
*/
type SupportBundledataParams struct {

	/* ID.

	   The Support Bundle ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the support bundledata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupportBundledataParams) WithDefaults() *SupportBundledataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the support bundledata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupportBundledataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the support bundledata params
func (o *SupportBundledataParams) WithTimeout(timeout time.Duration) *SupportBundledataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the support bundledata params
func (o *SupportBundledataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the support bundledata params
func (o *SupportBundledataParams) WithContext(ctx context.Context) *SupportBundledataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the support bundledata params
func (o *SupportBundledataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the support bundledata params
func (o *SupportBundledataParams) WithHTTPClient(client *http.Client) *SupportBundledataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the support bundledata params
func (o *SupportBundledataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the support bundledata params
func (o *SupportBundledataParams) WithID(id string) *SupportBundledataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the support bundledata params
func (o *SupportBundledataParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SupportBundledataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
