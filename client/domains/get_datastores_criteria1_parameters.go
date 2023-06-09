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

// NewGetDatastoresCriteria1Params creates a new GetDatastoresCriteria1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDatastoresCriteria1Params() *GetDatastoresCriteria1Params {
	return &GetDatastoresCriteria1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatastoresCriteria1ParamsWithTimeout creates a new GetDatastoresCriteria1Params object
// with the ability to set a timeout on a request.
func NewGetDatastoresCriteria1ParamsWithTimeout(timeout time.Duration) *GetDatastoresCriteria1Params {
	return &GetDatastoresCriteria1Params{
		timeout: timeout,
	}
}

// NewGetDatastoresCriteria1ParamsWithContext creates a new GetDatastoresCriteria1Params object
// with the ability to set a context for a request.
func NewGetDatastoresCriteria1ParamsWithContext(ctx context.Context) *GetDatastoresCriteria1Params {
	return &GetDatastoresCriteria1Params{
		Context: ctx,
	}
}

// NewGetDatastoresCriteria1ParamsWithHTTPClient creates a new GetDatastoresCriteria1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetDatastoresCriteria1ParamsWithHTTPClient(client *http.Client) *GetDatastoresCriteria1Params {
	return &GetDatastoresCriteria1Params{
		HTTPClient: client,
	}
}

/*
GetDatastoresCriteria1Params contains all the parameters to send to the API endpoint

	for the get datastores criteria 1 operation.

	Typically these are written to a http.Request.
*/
type GetDatastoresCriteria1Params struct {

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get datastores criteria 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatastoresCriteria1Params) WithDefaults() *GetDatastoresCriteria1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get datastores criteria 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatastoresCriteria1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get datastores criteria 1 params
func (o *GetDatastoresCriteria1Params) WithTimeout(timeout time.Duration) *GetDatastoresCriteria1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datastores criteria 1 params
func (o *GetDatastoresCriteria1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datastores criteria 1 params
func (o *GetDatastoresCriteria1Params) WithContext(ctx context.Context) *GetDatastoresCriteria1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datastores criteria 1 params
func (o *GetDatastoresCriteria1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datastores criteria 1 params
func (o *GetDatastoresCriteria1Params) WithHTTPClient(client *http.Client) *GetDatastoresCriteria1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datastores criteria 1 params
func (o *GetDatastoresCriteria1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get datastores criteria 1 params
func (o *GetDatastoresCriteria1Params) WithDomainID(domainID string) *GetDatastoresCriteria1Params {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get datastores criteria 1 params
func (o *GetDatastoresCriteria1Params) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatastoresCriteria1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
