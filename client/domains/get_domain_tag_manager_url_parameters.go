// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package domains

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

// NewGetDomainTagManagerURLParams creates a new GetDomainTagManagerURLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomainTagManagerURLParams() *GetDomainTagManagerURLParams {
	return &GetDomainTagManagerURLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainTagManagerURLParamsWithTimeout creates a new GetDomainTagManagerURLParams object
// with the ability to set a timeout on a request.
func NewGetDomainTagManagerURLParamsWithTimeout(timeout time.Duration) *GetDomainTagManagerURLParams {
	return &GetDomainTagManagerURLParams{
		timeout: timeout,
	}
}

// NewGetDomainTagManagerURLParamsWithContext creates a new GetDomainTagManagerURLParams object
// with the ability to set a context for a request.
func NewGetDomainTagManagerURLParamsWithContext(ctx context.Context) *GetDomainTagManagerURLParams {
	return &GetDomainTagManagerURLParams{
		Context: ctx,
	}
}

// NewGetDomainTagManagerURLParamsWithHTTPClient creates a new GetDomainTagManagerURLParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomainTagManagerURLParamsWithHTTPClient(client *http.Client) *GetDomainTagManagerURLParams {
	return &GetDomainTagManagerURLParams{
		HTTPClient: client,
	}
}

/*
GetDomainTagManagerURLParams contains all the parameters to send to the API endpoint

	for the get domain tag manager Url operation.

	Typically these are written to a http.Request.
*/
type GetDomainTagManagerURLParams struct {

	/* ID.

	   Domain ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get domain tag manager Url params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainTagManagerURLParams) WithDefaults() *GetDomainTagManagerURLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domain tag manager Url params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainTagManagerURLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domain tag manager Url params
func (o *GetDomainTagManagerURLParams) WithTimeout(timeout time.Duration) *GetDomainTagManagerURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain tag manager Url params
func (o *GetDomainTagManagerURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain tag manager Url params
func (o *GetDomainTagManagerURLParams) WithContext(ctx context.Context) *GetDomainTagManagerURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain tag manager Url params
func (o *GetDomainTagManagerURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain tag manager Url params
func (o *GetDomainTagManagerURLParams) WithHTTPClient(client *http.Client) *GetDomainTagManagerURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain tag manager Url params
func (o *GetDomainTagManagerURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get domain tag manager Url params
func (o *GetDomainTagManagerURLParams) WithID(id string) *GetDomainTagManagerURLParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get domain tag manager Url params
func (o *GetDomainTagManagerURLParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainTagManagerURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
