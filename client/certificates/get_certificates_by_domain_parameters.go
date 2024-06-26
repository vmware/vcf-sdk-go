// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package certificates

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

// NewGetCertificatesByDomainParams creates a new GetCertificatesByDomainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCertificatesByDomainParams() *GetCertificatesByDomainParams {
	return &GetCertificatesByDomainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCertificatesByDomainParamsWithTimeout creates a new GetCertificatesByDomainParams object
// with the ability to set a timeout on a request.
func NewGetCertificatesByDomainParamsWithTimeout(timeout time.Duration) *GetCertificatesByDomainParams {
	return &GetCertificatesByDomainParams{
		timeout: timeout,
	}
}

// NewGetCertificatesByDomainParamsWithContext creates a new GetCertificatesByDomainParams object
// with the ability to set a context for a request.
func NewGetCertificatesByDomainParamsWithContext(ctx context.Context) *GetCertificatesByDomainParams {
	return &GetCertificatesByDomainParams{
		Context: ctx,
	}
}

// NewGetCertificatesByDomainParamsWithHTTPClient creates a new GetCertificatesByDomainParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCertificatesByDomainParamsWithHTTPClient(client *http.Client) *GetCertificatesByDomainParams {
	return &GetCertificatesByDomainParams{
		HTTPClient: client,
	}
}

/*
GetCertificatesByDomainParams contains all the parameters to send to the API endpoint

	for the get certificates by domain operation.

	Typically these are written to a http.Request.
*/
type GetCertificatesByDomainParams struct {

	/* ID.

	   Domain ID or Name
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get certificates by domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCertificatesByDomainParams) WithDefaults() *GetCertificatesByDomainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get certificates by domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCertificatesByDomainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get certificates by domain params
func (o *GetCertificatesByDomainParams) WithTimeout(timeout time.Duration) *GetCertificatesByDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get certificates by domain params
func (o *GetCertificatesByDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get certificates by domain params
func (o *GetCertificatesByDomainParams) WithContext(ctx context.Context) *GetCertificatesByDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get certificates by domain params
func (o *GetCertificatesByDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get certificates by domain params
func (o *GetCertificatesByDomainParams) WithHTTPClient(client *http.Client) *GetCertificatesByDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get certificates by domain params
func (o *GetCertificatesByDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get certificates by domain params
func (o *GetCertificatesByDomainParams) WithID(id string) *GetCertificatesByDomainParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get certificates by domain params
func (o *GetCertificatesByDomainParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCertificatesByDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
