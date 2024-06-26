// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package certificates

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

// NewDownloadCSRParams creates a new DownloadCSRParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadCSRParams() *DownloadCSRParams {
	return &DownloadCSRParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadCSRParamsWithTimeout creates a new DownloadCSRParams object
// with the ability to set a timeout on a request.
func NewDownloadCSRParamsWithTimeout(timeout time.Duration) *DownloadCSRParams {
	return &DownloadCSRParams{
		timeout: timeout,
	}
}

// NewDownloadCSRParamsWithContext creates a new DownloadCSRParams object
// with the ability to set a context for a request.
func NewDownloadCSRParamsWithContext(ctx context.Context) *DownloadCSRParams {
	return &DownloadCSRParams{
		Context: ctx,
	}
}

// NewDownloadCSRParamsWithHTTPClient creates a new DownloadCSRParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadCSRParamsWithHTTPClient(client *http.Client) *DownloadCSRParams {
	return &DownloadCSRParams{
		HTTPClient: client,
	}
}

/*
DownloadCSRParams contains all the parameters to send to the API endpoint

	for the download CSR operation.

	Typically these are written to a http.Request.
*/
type DownloadCSRParams struct {

	/* ID.

	   Domain ID or Name
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download CSR params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadCSRParams) WithDefaults() *DownloadCSRParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download CSR params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadCSRParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download CSR params
func (o *DownloadCSRParams) WithTimeout(timeout time.Duration) *DownloadCSRParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download CSR params
func (o *DownloadCSRParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download CSR params
func (o *DownloadCSRParams) WithContext(ctx context.Context) *DownloadCSRParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download CSR params
func (o *DownloadCSRParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download CSR params
func (o *DownloadCSRParams) WithHTTPClient(client *http.Client) *DownloadCSRParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download CSR params
func (o *DownloadCSRParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the download CSR params
func (o *DownloadCSRParams) WithID(id string) *DownloadCSRParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the download CSR params
func (o *DownloadCSRParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadCSRParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
