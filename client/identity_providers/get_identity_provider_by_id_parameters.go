// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package identity_providers

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

// NewGetIdentityProviderByIDParams creates a new GetIdentityProviderByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIdentityProviderByIDParams() *GetIdentityProviderByIDParams {
	return &GetIdentityProviderByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityProviderByIDParamsWithTimeout creates a new GetIdentityProviderByIDParams object
// with the ability to set a timeout on a request.
func NewGetIdentityProviderByIDParamsWithTimeout(timeout time.Duration) *GetIdentityProviderByIDParams {
	return &GetIdentityProviderByIDParams{
		timeout: timeout,
	}
}

// NewGetIdentityProviderByIDParamsWithContext creates a new GetIdentityProviderByIDParams object
// with the ability to set a context for a request.
func NewGetIdentityProviderByIDParamsWithContext(ctx context.Context) *GetIdentityProviderByIDParams {
	return &GetIdentityProviderByIDParams{
		Context: ctx,
	}
}

// NewGetIdentityProviderByIDParamsWithHTTPClient creates a new GetIdentityProviderByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIdentityProviderByIDParamsWithHTTPClient(client *http.Client) *GetIdentityProviderByIDParams {
	return &GetIdentityProviderByIDParams{
		HTTPClient: client,
	}
}

/*
GetIdentityProviderByIDParams contains all the parameters to send to the API endpoint

	for the get identity provider by Id operation.

	Typically these are written to a http.Request.
*/
type GetIdentityProviderByIDParams struct {

	/* ID.

	   ID of the Identity Provider
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get identity provider by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityProviderByIDParams) WithDefaults() *GetIdentityProviderByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get identity provider by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityProviderByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get identity provider by Id params
func (o *GetIdentityProviderByIDParams) WithTimeout(timeout time.Duration) *GetIdentityProviderByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identity provider by Id params
func (o *GetIdentityProviderByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identity provider by Id params
func (o *GetIdentityProviderByIDParams) WithContext(ctx context.Context) *GetIdentityProviderByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identity provider by Id params
func (o *GetIdentityProviderByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identity provider by Id params
func (o *GetIdentityProviderByIDParams) WithHTTPClient(client *http.Client) *GetIdentityProviderByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identity provider by Id params
func (o *GetIdentityProviderByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get identity provider by Id params
func (o *GetIdentityProviderByIDParams) WithID(id string) *GetIdentityProviderByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get identity provider by Id params
func (o *GetIdentityProviderByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityProviderByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
