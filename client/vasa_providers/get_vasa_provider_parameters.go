// Code generated by go-swagger; DO NOT EDIT.

// Copyright © 2023 VMware, Inc. All Rights Reserved.
// Code licenced under: BSD-2

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

// NewGETVasaProviderParams creates a new GETVasaProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETVasaProviderParams() *GETVasaProviderParams {
	return &GETVasaProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETVasaProviderParamsWithTimeout creates a new GETVasaProviderParams object
// with the ability to set a timeout on a request.
func NewGETVasaProviderParamsWithTimeout(timeout time.Duration) *GETVasaProviderParams {
	return &GETVasaProviderParams{
		timeout: timeout,
	}
}

// NewGETVasaProviderParamsWithContext creates a new GETVasaProviderParams object
// with the ability to set a context for a request.
func NewGETVasaProviderParamsWithContext(ctx context.Context) *GETVasaProviderParams {
	return &GETVasaProviderParams{
		Context: ctx,
	}
}

// NewGETVasaProviderParamsWithHTTPClient creates a new GETVasaProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETVasaProviderParamsWithHTTPClient(client *http.Client) *GETVasaProviderParams {
	return &GETVasaProviderParams{
		HTTPClient: client,
	}
}

/*
GETVasaProviderParams contains all the parameters to send to the API endpoint

	for the get vasa provider operation.

	Typically these are written to a http.Request.
*/
type GETVasaProviderParams struct {

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
func (o *GETVasaProviderParams) WithDefaults() *GETVasaProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETVasaProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vasa provider params
func (o *GETVasaProviderParams) WithTimeout(timeout time.Duration) *GETVasaProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vasa provider params
func (o *GETVasaProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vasa provider params
func (o *GETVasaProviderParams) WithContext(ctx context.Context) *GETVasaProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vasa provider params
func (o *GETVasaProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vasa provider params
func (o *GETVasaProviderParams) WithHTTPClient(client *http.Client) *GETVasaProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vasa provider params
func (o *GETVasaProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get vasa provider params
func (o *GETVasaProviderParams) WithID(id string) *GETVasaProviderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vasa provider params
func (o *GETVasaProviderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETVasaProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
