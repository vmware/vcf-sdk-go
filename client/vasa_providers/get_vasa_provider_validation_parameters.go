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

// NewGetVasaProviderValidationParams creates a new GetVasaProviderValidationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVasaProviderValidationParams() *GetVasaProviderValidationParams {
	return &GetVasaProviderValidationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVasaProviderValidationParamsWithTimeout creates a new GetVasaProviderValidationParams object
// with the ability to set a timeout on a request.
func NewGetVasaProviderValidationParamsWithTimeout(timeout time.Duration) *GetVasaProviderValidationParams {
	return &GetVasaProviderValidationParams{
		timeout: timeout,
	}
}

// NewGetVasaProviderValidationParamsWithContext creates a new GetVasaProviderValidationParams object
// with the ability to set a context for a request.
func NewGetVasaProviderValidationParamsWithContext(ctx context.Context) *GetVasaProviderValidationParams {
	return &GetVasaProviderValidationParams{
		Context: ctx,
	}
}

// NewGetVasaProviderValidationParamsWithHTTPClient creates a new GetVasaProviderValidationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVasaProviderValidationParamsWithHTTPClient(client *http.Client) *GetVasaProviderValidationParams {
	return &GetVasaProviderValidationParams{
		HTTPClient: client,
	}
}

/*
GetVasaProviderValidationParams contains all the parameters to send to the API endpoint

	for the get vasa provider validation operation.

	Typically these are written to a http.Request.
*/
type GetVasaProviderValidationParams struct {

	/* ID.

	   The validation ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vasa provider validation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVasaProviderValidationParams) WithDefaults() *GetVasaProviderValidationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vasa provider validation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVasaProviderValidationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vasa provider validation params
func (o *GetVasaProviderValidationParams) WithTimeout(timeout time.Duration) *GetVasaProviderValidationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vasa provider validation params
func (o *GetVasaProviderValidationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vasa provider validation params
func (o *GetVasaProviderValidationParams) WithContext(ctx context.Context) *GetVasaProviderValidationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vasa provider validation params
func (o *GetVasaProviderValidationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vasa provider validation params
func (o *GetVasaProviderValidationParams) WithHTTPClient(client *http.Client) *GetVasaProviderValidationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vasa provider validation params
func (o *GetVasaProviderValidationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vasa provider validation params
func (o *GetVasaProviderValidationParams) WithID(id string) *GetVasaProviderValidationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vasa provider validation params
func (o *GetVasaProviderValidationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVasaProviderValidationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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