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

// NewRemoveVasaProviderParams creates a new RemoveVasaProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveVasaProviderParams() *RemoveVasaProviderParams {
	return &RemoveVasaProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveVasaProviderParamsWithTimeout creates a new RemoveVasaProviderParams object
// with the ability to set a timeout on a request.
func NewRemoveVasaProviderParamsWithTimeout(timeout time.Duration) *RemoveVasaProviderParams {
	return &RemoveVasaProviderParams{
		timeout: timeout,
	}
}

// NewRemoveVasaProviderParamsWithContext creates a new RemoveVasaProviderParams object
// with the ability to set a context for a request.
func NewRemoveVasaProviderParamsWithContext(ctx context.Context) *RemoveVasaProviderParams {
	return &RemoveVasaProviderParams{
		Context: ctx,
	}
}

// NewRemoveVasaProviderParamsWithHTTPClient creates a new RemoveVasaProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveVasaProviderParamsWithHTTPClient(client *http.Client) *RemoveVasaProviderParams {
	return &RemoveVasaProviderParams{
		HTTPClient: client,
	}
}

/*
RemoveVasaProviderParams contains all the parameters to send to the API endpoint

	for the remove vasa provider operation.

	Typically these are written to a http.Request.
*/
type RemoveVasaProviderParams struct {

	/* ID.

	   VASA Provider ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveVasaProviderParams) WithDefaults() *RemoveVasaProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveVasaProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove vasa provider params
func (o *RemoveVasaProviderParams) WithTimeout(timeout time.Duration) *RemoveVasaProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove vasa provider params
func (o *RemoveVasaProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove vasa provider params
func (o *RemoveVasaProviderParams) WithContext(ctx context.Context) *RemoveVasaProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove vasa provider params
func (o *RemoveVasaProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove vasa provider params
func (o *RemoveVasaProviderParams) WithHTTPClient(client *http.Client) *RemoveVasaProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove vasa provider params
func (o *RemoveVasaProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the remove vasa provider params
func (o *RemoveVasaProviderParams) WithID(id string) *RemoveVasaProviderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove vasa provider params
func (o *RemoveVasaProviderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveVasaProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
