// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vasa_providers

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

// NewGetVasaProviderParams creates a new GetVasaProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVasaProviderParams() *GetVasaProviderParams {
	return &GetVasaProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVasaProviderParamsWithTimeout creates a new GetVasaProviderParams object
// with the ability to set a timeout on a request.
func NewGetVasaProviderParamsWithTimeout(timeout time.Duration) *GetVasaProviderParams {
	return &GetVasaProviderParams{
		timeout: timeout,
	}
}

// NewGetVasaProviderParamsWithContext creates a new GetVasaProviderParams object
// with the ability to set a context for a request.
func NewGetVasaProviderParamsWithContext(ctx context.Context) *GetVasaProviderParams {
	return &GetVasaProviderParams{
		Context: ctx,
	}
}

// NewGetVasaProviderParamsWithHTTPClient creates a new GetVasaProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVasaProviderParamsWithHTTPClient(client *http.Client) *GetVasaProviderParams {
	return &GetVasaProviderParams{
		HTTPClient: client,
	}
}

/*
GetVasaProviderParams contains all the parameters to send to the API endpoint

	for the get vasa provider operation.

	Typically these are written to a http.Request.
*/
type GetVasaProviderParams struct {

	/* ID.

	   VASA Provider ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVasaProviderParams) WithDefaults() *GetVasaProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVasaProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vasa provider params
func (o *GetVasaProviderParams) WithTimeout(timeout time.Duration) *GetVasaProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vasa provider params
func (o *GetVasaProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vasa provider params
func (o *GetVasaProviderParams) WithContext(ctx context.Context) *GetVasaProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vasa provider params
func (o *GetVasaProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vasa provider params
func (o *GetVasaProviderParams) WithHTTPClient(client *http.Client) *GetVasaProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vasa provider params
func (o *GetVasaProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vasa provider params
func (o *GetVasaProviderParams) WithID(id string) *GetVasaProviderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vasa provider params
func (o *GetVasaProviderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVasaProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
