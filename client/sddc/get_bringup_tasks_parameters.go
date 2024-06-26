// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

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

// NewGetBringupTasksParams creates a new GetBringupTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBringupTasksParams() *GetBringupTasksParams {
	return &GetBringupTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBringupTasksParamsWithTimeout creates a new GetBringupTasksParams object
// with the ability to set a timeout on a request.
func NewGetBringupTasksParamsWithTimeout(timeout time.Duration) *GetBringupTasksParams {
	return &GetBringupTasksParams{
		timeout: timeout,
	}
}

// NewGetBringupTasksParamsWithContext creates a new GetBringupTasksParams object
// with the ability to set a context for a request.
func NewGetBringupTasksParamsWithContext(ctx context.Context) *GetBringupTasksParams {
	return &GetBringupTasksParams{
		Context: ctx,
	}
}

// NewGetBringupTasksParamsWithHTTPClient creates a new GetBringupTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBringupTasksParamsWithHTTPClient(client *http.Client) *GetBringupTasksParams {
	return &GetBringupTasksParams{
		HTTPClient: client,
	}
}

/*
GetBringupTasksParams contains all the parameters to send to the API endpoint

	for the get bringup tasks operation.

	Typically these are written to a http.Request.
*/
type GetBringupTasksParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bringup tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBringupTasksParams) WithDefaults() *GetBringupTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bringup tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBringupTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bringup tasks params
func (o *GetBringupTasksParams) WithTimeout(timeout time.Duration) *GetBringupTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bringup tasks params
func (o *GetBringupTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bringup tasks params
func (o *GetBringupTasksParams) WithContext(ctx context.Context) *GetBringupTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bringup tasks params
func (o *GetBringupTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bringup tasks params
func (o *GetBringupTasksParams) WithHTTPClient(client *http.Client) *GetBringupTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bringup tasks params
func (o *GetBringupTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetBringupTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
