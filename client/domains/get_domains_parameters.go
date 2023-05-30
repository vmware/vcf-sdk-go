// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package domains

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

// NewGetDomainsParams creates a new GetDomainsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomainsParams() *GetDomainsParams {
	return &GetDomainsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainsParamsWithTimeout creates a new GetDomainsParams object
// with the ability to set a timeout on a request.
func NewGetDomainsParamsWithTimeout(timeout time.Duration) *GetDomainsParams {
	return &GetDomainsParams{
		timeout: timeout,
	}
}

// NewGetDomainsParamsWithContext creates a new GetDomainsParams object
// with the ability to set a context for a request.
func NewGetDomainsParamsWithContext(ctx context.Context) *GetDomainsParams {
	return &GetDomainsParams{
		Context: ctx,
	}
}

// NewGetDomainsParamsWithHTTPClient creates a new GetDomainsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomainsParamsWithHTTPClient(client *http.Client) *GetDomainsParams {
	return &GetDomainsParams{
		HTTPClient: client,
	}
}

/*
GetDomainsParams contains all the parameters to send to the API endpoint

	for the get domains operation.

	Typically these are written to a http.Request.
*/
type GetDomainsParams struct {

	/* Type.

	   The type of the domain
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get domains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainsParams) WithDefaults() *GetDomainsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domains params
func (o *GetDomainsParams) WithTimeout(timeout time.Duration) *GetDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domains params
func (o *GetDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domains params
func (o *GetDomainsParams) WithContext(ctx context.Context) *GetDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domains params
func (o *GetDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domains params
func (o *GetDomainsParams) WithHTTPClient(client *http.Client) *GetDomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domains params
func (o *GetDomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithType adds the typeVar to the get domains params
func (o *GetDomainsParams) WithType(typeVar *string) *GetDomainsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get domains params
func (o *GetDomainsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
