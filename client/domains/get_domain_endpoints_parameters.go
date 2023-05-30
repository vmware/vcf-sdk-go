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

// NewGetDomainEndpointsParams creates a new GetDomainEndpointsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomainEndpointsParams() *GetDomainEndpointsParams {
	return &GetDomainEndpointsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainEndpointsParamsWithTimeout creates a new GetDomainEndpointsParams object
// with the ability to set a timeout on a request.
func NewGetDomainEndpointsParamsWithTimeout(timeout time.Duration) *GetDomainEndpointsParams {
	return &GetDomainEndpointsParams{
		timeout: timeout,
	}
}

// NewGetDomainEndpointsParamsWithContext creates a new GetDomainEndpointsParams object
// with the ability to set a context for a request.
func NewGetDomainEndpointsParamsWithContext(ctx context.Context) *GetDomainEndpointsParams {
	return &GetDomainEndpointsParams{
		Context: ctx,
	}
}

// NewGetDomainEndpointsParamsWithHTTPClient creates a new GetDomainEndpointsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomainEndpointsParamsWithHTTPClient(client *http.Client) *GetDomainEndpointsParams {
	return &GetDomainEndpointsParams{
		HTTPClient: client,
	}
}

/*
GetDomainEndpointsParams contains all the parameters to send to the API endpoint

	for the get domain endpoints operation.

	Typically these are written to a http.Request.
*/
type GetDomainEndpointsParams struct {

	/* ID.

	   Domain ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get domain endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainEndpointsParams) WithDefaults() *GetDomainEndpointsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domain endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainEndpointsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domain endpoints params
func (o *GetDomainEndpointsParams) WithTimeout(timeout time.Duration) *GetDomainEndpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain endpoints params
func (o *GetDomainEndpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain endpoints params
func (o *GetDomainEndpointsParams) WithContext(ctx context.Context) *GetDomainEndpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain endpoints params
func (o *GetDomainEndpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain endpoints params
func (o *GetDomainEndpointsParams) WithHTTPClient(client *http.Client) *GetDomainEndpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain endpoints params
func (o *GetDomainEndpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get domain endpoints params
func (o *GetDomainEndpointsParams) WithID(id string) *GetDomainEndpointsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get domain endpoints params
func (o *GetDomainEndpointsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainEndpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
