// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package version_aliases_for_bundle_component_type

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

// NewGetVersionAliasConfigurationParams creates a new GetVersionAliasConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVersionAliasConfigurationParams() *GetVersionAliasConfigurationParams {
	return &GetVersionAliasConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionAliasConfigurationParamsWithTimeout creates a new GetVersionAliasConfigurationParams object
// with the ability to set a timeout on a request.
func NewGetVersionAliasConfigurationParamsWithTimeout(timeout time.Duration) *GetVersionAliasConfigurationParams {
	return &GetVersionAliasConfigurationParams{
		timeout: timeout,
	}
}

// NewGetVersionAliasConfigurationParamsWithContext creates a new GetVersionAliasConfigurationParams object
// with the ability to set a context for a request.
func NewGetVersionAliasConfigurationParamsWithContext(ctx context.Context) *GetVersionAliasConfigurationParams {
	return &GetVersionAliasConfigurationParams{
		Context: ctx,
	}
}

// NewGetVersionAliasConfigurationParamsWithHTTPClient creates a new GetVersionAliasConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVersionAliasConfigurationParamsWithHTTPClient(client *http.Client) *GetVersionAliasConfigurationParams {
	return &GetVersionAliasConfigurationParams{
		HTTPClient: client,
	}
}

/*
GetVersionAliasConfigurationParams contains all the parameters to send to the API endpoint

	for the get version alias configuration operation.

	Typically these are written to a http.Request.
*/
type GetVersionAliasConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get version alias configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionAliasConfigurationParams) WithDefaults() *GetVersionAliasConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get version alias configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionAliasConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get version alias configuration params
func (o *GetVersionAliasConfigurationParams) WithTimeout(timeout time.Duration) *GetVersionAliasConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get version alias configuration params
func (o *GetVersionAliasConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get version alias configuration params
func (o *GetVersionAliasConfigurationParams) WithContext(ctx context.Context) *GetVersionAliasConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get version alias configuration params
func (o *GetVersionAliasConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get version alias configuration params
func (o *GetVersionAliasConfigurationParams) WithHTTPClient(client *http.Client) *GetVersionAliasConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get version alias configuration params
func (o *GetVersionAliasConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionAliasConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
