// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package credentials

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

// NewGetPasswordExpirationByTaskIDParams creates a new GetPasswordExpirationByTaskIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPasswordExpirationByTaskIDParams() *GetPasswordExpirationByTaskIDParams {
	return &GetPasswordExpirationByTaskIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPasswordExpirationByTaskIDParamsWithTimeout creates a new GetPasswordExpirationByTaskIDParams object
// with the ability to set a timeout on a request.
func NewGetPasswordExpirationByTaskIDParamsWithTimeout(timeout time.Duration) *GetPasswordExpirationByTaskIDParams {
	return &GetPasswordExpirationByTaskIDParams{
		timeout: timeout,
	}
}

// NewGetPasswordExpirationByTaskIDParamsWithContext creates a new GetPasswordExpirationByTaskIDParams object
// with the ability to set a context for a request.
func NewGetPasswordExpirationByTaskIDParamsWithContext(ctx context.Context) *GetPasswordExpirationByTaskIDParams {
	return &GetPasswordExpirationByTaskIDParams{
		Context: ctx,
	}
}

// NewGetPasswordExpirationByTaskIDParamsWithHTTPClient creates a new GetPasswordExpirationByTaskIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPasswordExpirationByTaskIDParamsWithHTTPClient(client *http.Client) *GetPasswordExpirationByTaskIDParams {
	return &GetPasswordExpirationByTaskIDParams{
		HTTPClient: client,
	}
}

/*
GetPasswordExpirationByTaskIDParams contains all the parameters to send to the API endpoint

	for the get password expiration by task ID operation.

	Typically these are written to a http.Request.
*/
type GetPasswordExpirationByTaskIDParams struct {

	/* ID.

	   The expiration fetch workflow ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get password expiration by task ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPasswordExpirationByTaskIDParams) WithDefaults() *GetPasswordExpirationByTaskIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get password expiration by task ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPasswordExpirationByTaskIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get password expiration by task ID params
func (o *GetPasswordExpirationByTaskIDParams) WithTimeout(timeout time.Duration) *GetPasswordExpirationByTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get password expiration by task ID params
func (o *GetPasswordExpirationByTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get password expiration by task ID params
func (o *GetPasswordExpirationByTaskIDParams) WithContext(ctx context.Context) *GetPasswordExpirationByTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get password expiration by task ID params
func (o *GetPasswordExpirationByTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get password expiration by task ID params
func (o *GetPasswordExpirationByTaskIDParams) WithHTTPClient(client *http.Client) *GetPasswordExpirationByTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get password expiration by task ID params
func (o *GetPasswordExpirationByTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get password expiration by task ID params
func (o *GetPasswordExpirationByTaskIDParams) WithID(id string) *GetPasswordExpirationByTaskIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get password expiration by task ID params
func (o *GetPasswordExpirationByTaskIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPasswordExpirationByTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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