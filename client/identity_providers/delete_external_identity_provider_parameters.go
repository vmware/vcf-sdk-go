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

// NewDeleteExternalIdentityProviderParams creates a new DeleteExternalIdentityProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteExternalIdentityProviderParams() *DeleteExternalIdentityProviderParams {
	return &DeleteExternalIdentityProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteExternalIdentityProviderParamsWithTimeout creates a new DeleteExternalIdentityProviderParams object
// with the ability to set a timeout on a request.
func NewDeleteExternalIdentityProviderParamsWithTimeout(timeout time.Duration) *DeleteExternalIdentityProviderParams {
	return &DeleteExternalIdentityProviderParams{
		timeout: timeout,
	}
}

// NewDeleteExternalIdentityProviderParamsWithContext creates a new DeleteExternalIdentityProviderParams object
// with the ability to set a context for a request.
func NewDeleteExternalIdentityProviderParamsWithContext(ctx context.Context) *DeleteExternalIdentityProviderParams {
	return &DeleteExternalIdentityProviderParams{
		Context: ctx,
	}
}

// NewDeleteExternalIdentityProviderParamsWithHTTPClient creates a new DeleteExternalIdentityProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteExternalIdentityProviderParamsWithHTTPClient(client *http.Client) *DeleteExternalIdentityProviderParams {
	return &DeleteExternalIdentityProviderParams{
		HTTPClient: client,
	}
}

/*
DeleteExternalIdentityProviderParams contains all the parameters to send to the API endpoint

	for the delete external identity provider operation.

	Typically these are written to a http.Request.
*/
type DeleteExternalIdentityProviderParams struct {

	/* ID.

	   ID of Identity Provider
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete external identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExternalIdentityProviderParams) WithDefaults() *DeleteExternalIdentityProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete external identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExternalIdentityProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete external identity provider params
func (o *DeleteExternalIdentityProviderParams) WithTimeout(timeout time.Duration) *DeleteExternalIdentityProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete external identity provider params
func (o *DeleteExternalIdentityProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete external identity provider params
func (o *DeleteExternalIdentityProviderParams) WithContext(ctx context.Context) *DeleteExternalIdentityProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete external identity provider params
func (o *DeleteExternalIdentityProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete external identity provider params
func (o *DeleteExternalIdentityProviderParams) WithHTTPClient(client *http.Client) *DeleteExternalIdentityProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete external identity provider params
func (o *DeleteExternalIdentityProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete external identity provider params
func (o *DeleteExternalIdentityProviderParams) WithID(id string) *DeleteExternalIdentityProviderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete external identity provider params
func (o *DeleteExternalIdentityProviderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteExternalIdentityProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
