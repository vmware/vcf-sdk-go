// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package nsxt_clusters

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

// NewGetNsxCriteriaParams creates a new GetNsxCriteriaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNsxCriteriaParams() *GetNsxCriteriaParams {
	return &GetNsxCriteriaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNsxCriteriaParamsWithTimeout creates a new GetNsxCriteriaParams object
// with the ability to set a timeout on a request.
func NewGetNsxCriteriaParamsWithTimeout(timeout time.Duration) *GetNsxCriteriaParams {
	return &GetNsxCriteriaParams{
		timeout: timeout,
	}
}

// NewGetNsxCriteriaParamsWithContext creates a new GetNsxCriteriaParams object
// with the ability to set a context for a request.
func NewGetNsxCriteriaParamsWithContext(ctx context.Context) *GetNsxCriteriaParams {
	return &GetNsxCriteriaParams{
		Context: ctx,
	}
}

// NewGetNsxCriteriaParamsWithHTTPClient creates a new GetNsxCriteriaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNsxCriteriaParamsWithHTTPClient(client *http.Client) *GetNsxCriteriaParams {
	return &GetNsxCriteriaParams{
		HTTPClient: client,
	}
}

/*
GetNsxCriteriaParams contains all the parameters to send to the API endpoint

	for the get nsx criteria operation.

	Typically these are written to a http.Request.
*/
type GetNsxCriteriaParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nsx criteria params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNsxCriteriaParams) WithDefaults() *GetNsxCriteriaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nsx criteria params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNsxCriteriaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nsx criteria params
func (o *GetNsxCriteriaParams) WithTimeout(timeout time.Duration) *GetNsxCriteriaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nsx criteria params
func (o *GetNsxCriteriaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nsx criteria params
func (o *GetNsxCriteriaParams) WithContext(ctx context.Context) *GetNsxCriteriaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nsx criteria params
func (o *GetNsxCriteriaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nsx criteria params
func (o *GetNsxCriteriaParams) WithHTTPClient(client *http.Client) *GetNsxCriteriaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nsx criteria params
func (o *GetNsxCriteriaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNsxCriteriaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
