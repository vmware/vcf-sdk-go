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

// NewGetUsersOfVasaProviderParams creates a new GetUsersOfVasaProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsersOfVasaProviderParams() *GetUsersOfVasaProviderParams {
	return &GetUsersOfVasaProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersOfVasaProviderParamsWithTimeout creates a new GetUsersOfVasaProviderParams object
// with the ability to set a timeout on a request.
func NewGetUsersOfVasaProviderParamsWithTimeout(timeout time.Duration) *GetUsersOfVasaProviderParams {
	return &GetUsersOfVasaProviderParams{
		timeout: timeout,
	}
}

// NewGetUsersOfVasaProviderParamsWithContext creates a new GetUsersOfVasaProviderParams object
// with the ability to set a context for a request.
func NewGetUsersOfVasaProviderParamsWithContext(ctx context.Context) *GetUsersOfVasaProviderParams {
	return &GetUsersOfVasaProviderParams{
		Context: ctx,
	}
}

// NewGetUsersOfVasaProviderParamsWithHTTPClient creates a new GetUsersOfVasaProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsersOfVasaProviderParamsWithHTTPClient(client *http.Client) *GetUsersOfVasaProviderParams {
	return &GetUsersOfVasaProviderParams{
		HTTPClient: client,
	}
}

/*
GetUsersOfVasaProviderParams contains all the parameters to send to the API endpoint

	for the get users of vasa provider operation.

	Typically these are written to a http.Request.
*/
type GetUsersOfVasaProviderParams struct {

	/* ID.

	   VASA Provider ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get users of vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersOfVasaProviderParams) WithDefaults() *GetUsersOfVasaProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get users of vasa provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersOfVasaProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get users of vasa provider params
func (o *GetUsersOfVasaProviderParams) WithTimeout(timeout time.Duration) *GetUsersOfVasaProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users of vasa provider params
func (o *GetUsersOfVasaProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users of vasa provider params
func (o *GetUsersOfVasaProviderParams) WithContext(ctx context.Context) *GetUsersOfVasaProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users of vasa provider params
func (o *GetUsersOfVasaProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users of vasa provider params
func (o *GetUsersOfVasaProviderParams) WithHTTPClient(client *http.Client) *GetUsersOfVasaProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users of vasa provider params
func (o *GetUsersOfVasaProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get users of vasa provider params
func (o *GetUsersOfVasaProviderParams) WithID(id string) *GetUsersOfVasaProviderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get users of vasa provider params
func (o *GetUsersOfVasaProviderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersOfVasaProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
