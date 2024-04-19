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

// NewGetCredentialTaskByResourceIDParams creates a new GetCredentialTaskByResourceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCredentialTaskByResourceIDParams() *GetCredentialTaskByResourceIDParams {
	return &GetCredentialTaskByResourceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCredentialTaskByResourceIDParamsWithTimeout creates a new GetCredentialTaskByResourceIDParams object
// with the ability to set a timeout on a request.
func NewGetCredentialTaskByResourceIDParamsWithTimeout(timeout time.Duration) *GetCredentialTaskByResourceIDParams {
	return &GetCredentialTaskByResourceIDParams{
		timeout: timeout,
	}
}

// NewGetCredentialTaskByResourceIDParamsWithContext creates a new GetCredentialTaskByResourceIDParams object
// with the ability to set a context for a request.
func NewGetCredentialTaskByResourceIDParamsWithContext(ctx context.Context) *GetCredentialTaskByResourceIDParams {
	return &GetCredentialTaskByResourceIDParams{
		Context: ctx,
	}
}

// NewGetCredentialTaskByResourceIDParamsWithHTTPClient creates a new GetCredentialTaskByResourceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCredentialTaskByResourceIDParamsWithHTTPClient(client *http.Client) *GetCredentialTaskByResourceIDParams {
	return &GetCredentialTaskByResourceIDParams{
		HTTPClient: client,
	}
}

/*
GetCredentialTaskByResourceIDParams contains all the parameters to send to the API endpoint

	for the get credential task by resource ID operation.

	Typically these are written to a http.Request.
*/
type GetCredentialTaskByResourceIDParams struct {

	/* ID.

	   The ID of the credentials task
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get credential task by resource ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialTaskByResourceIDParams) WithDefaults() *GetCredentialTaskByResourceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get credential task by resource ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialTaskByResourceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get credential task by resource ID params
func (o *GetCredentialTaskByResourceIDParams) WithTimeout(timeout time.Duration) *GetCredentialTaskByResourceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credential task by resource ID params
func (o *GetCredentialTaskByResourceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credential task by resource ID params
func (o *GetCredentialTaskByResourceIDParams) WithContext(ctx context.Context) *GetCredentialTaskByResourceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credential task by resource ID params
func (o *GetCredentialTaskByResourceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credential task by resource ID params
func (o *GetCredentialTaskByResourceIDParams) WithHTTPClient(client *http.Client) *GetCredentialTaskByResourceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credential task by resource ID params
func (o *GetCredentialTaskByResourceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get credential task by resource ID params
func (o *GetCredentialTaskByResourceIDParams) WithID(id string) *GetCredentialTaskByResourceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get credential task by resource ID params
func (o *GetCredentialTaskByResourceIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCredentialTaskByResourceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
